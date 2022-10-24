package keptntask

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"

	klcv1alpha1 "github.com/keptn/lifecycle-controller/operator/api/v1alpha1"
	"github.com/keptn/lifecycle-controller/operator/api/v1alpha1/common"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

type FunctionExecutionParams struct {
	ConfigMap        string
	Parameters       map[string]string
	SecureParameters string
	URL              string
	Context          klcv1alpha1.TaskContext
}

func (r *KeptnTaskReconciler) generateFunctionJob(task *klcv1alpha1.KeptnTask, params FunctionExecutionParams) (*batchv1.Job, error) {
	randomId := rand.Intn(99999-10000) + 10000
	jobId := fmt.Sprintf("klc-%s-%d", common.TruncateString(task.Name, common.MaxTaskNameLength), randomId)
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      jobId,
			Namespace: task.Namespace,
			Labels:    createKeptnLabels(*task),
		},
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					RestartPolicy: "OnFailure",
				},
			},
		},
	}
	err := controllerutil.SetControllerReference(task, job, r.Scheme)
	if err != nil {
		r.Log.Error(err, "could not set controller reference:")
	}

	container := corev1.Container{
		Name:  "keptn-function-runner",
		Image: os.Getenv("FUNCTION_RUNNER_IMAGE"),
	}

	var envVars []corev1.EnvVar

	if len(params.Parameters) > 0 {
		jsonParams, err := json.Marshal(params.Parameters)
		if err != nil {
			return job, fmt.Errorf("could not marshal parameters")
		}
		envVars = append(envVars, corev1.EnvVar{Name: "DATA", Value: string(jsonParams)})
	}

	jsonParams, err := json.Marshal(params.Context)
	if err != nil {
		return job, fmt.Errorf("could not marshal parameters")
	}
	envVars = append(envVars, corev1.EnvVar{Name: "CONTEXT", Value: string(jsonParams)})

	if params.SecureParameters != "" {
		envVars = append(envVars, corev1.EnvVar{
			Name: "SECURE_DATA",
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{Name: params.SecureParameters},
					Key:                  "SECURE_DATA",
				},
			},
		})
	}

	// Mount the function code if a ConfigMap is provided
	// The ConfigMap might be provided manually or created by the TaskDefinition controller
	if params.ConfigMap != "" {
		envVars = append(envVars, corev1.EnvVar{Name: "SCRIPT", Value: "/var/data/function.ts"})

		job.Spec.Template.Spec.Volumes = []corev1.Volume{
			{
				Name: "function-mount",
				VolumeSource: corev1.VolumeSource{
					ConfigMap: &corev1.ConfigMapVolumeSource{
						LocalObjectReference: corev1.LocalObjectReference{
							Name: params.ConfigMap,
						},
					},
				},
			},
		}
		container.VolumeMounts = []corev1.VolumeMount{
			{
				Name:      "function-mount",
				ReadOnly:  true,
				MountPath: "/var/data/function.ts",
				SubPath:   "code",
			},
		}
	} else {
		envVars = append(envVars, corev1.EnvVar{Name: "SCRIPT", Value: params.URL})
	}

	container.Env = envVars
	job.Spec.Template.Spec.Containers = []corev1.Container{
		container,
	}
	return job, nil
}

func (r *KeptnTaskReconciler) parseFunctionTaskDefinition(definition *klcv1alpha1.KeptnTaskDefinition) (FunctionExecutionParams, bool, error) {
	params := FunctionExecutionParams{}

	// Firstly check if this task definition has a parent object
	hasParent := false
	if definition.Spec.Function.FunctionReference != (klcv1alpha1.FunctionReference{}) {
		hasParent = true
	}

	if definition.Status.Function.ConfigMap != "" && definition.Spec.Function.HttpReference.Url != "" {
		r.Log.Info(fmt.Sprintf("The JobDefinition contains a ConfigMap and a HTTP Reference, ConfigMap is used / Namespace: %s, Name: %s  ", definition.Namespace, definition.Name))
	}

	// Check if there is a ConfigMap with the function for this object
	if definition.Status.Function.ConfigMap != "" {
		params.ConfigMap = definition.Status.Function.ConfigMap
	} else {
		// If not, check if it has an HTTP reference. If this is also not the case and the object has no parent, something is wrong
		if definition.Spec.Function.HttpReference.Url == "" && !hasParent {
			return params, false, fmt.Errorf("No ConfigMap specified or HTTP source specified in TaskDefinition) / Namespace: %s, Name: %s ", definition.Namespace, definition.Name)
		}
		params.URL = definition.Spec.Function.HttpReference.Url
	}

	// Check if there are parameters provided
	if len(definition.Spec.Function.Parameters.Inline) > 0 {
		params.Parameters = definition.Spec.Function.Parameters.Inline
	}

	// Check if there is a secret for secret params provided
	if definition.Spec.Function.SecureParameters.Secret != "" {
		params.SecureParameters = definition.Spec.Function.SecureParameters.Secret
	}
	return params, hasParent, nil
}
