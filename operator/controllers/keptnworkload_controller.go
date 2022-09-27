/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	"github.com/keptn-sandbox/lifecycle-controller/operator/api/v1alpha1/common"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	"reflect"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"time"

	klcv1alpha1 "github.com/keptn-sandbox/lifecycle-controller/operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// KeptnWorkloadReconciler reconciles a KeptnWorkload object
type KeptnWorkloadReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
	Log      logr.Logger
}

//+kubebuilder:rbac:groups=lifecycle.keptn.sh,resources=keptnworkloads,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=lifecycle.keptn.sh,resources=keptnworkloads/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=lifecycle.keptn.sh,resources=keptnworkloads/finalizers,verbs=update
//+kubebuilder:rbac:groups=lifecycle.keptn.sh,resources=keptnworkloadinstances,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=lifecycle.keptn.sh,resources=keptnworkloadinstances/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=lifecycle.keptn.sh,resources=keptnworkloadinstances/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the KeptnWorkload object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.2/pkg/reconcile
func (r *KeptnWorkloadReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	r.Log.Info("Searching for workload")

	workload := &klcv1alpha1.KeptnWorkload{}
	err := r.Get(ctx, req.NamespacedName, workload)
	if errors.IsNotFound(err) {
		return reconcile.Result{}, nil
	}

	if err != nil {
		return reconcile.Result{}, fmt.Errorf("could not fetch Workload: %+v", err)
	}

	r.Log.Info("Reconciling Keptn Workload", "workload", workload.Name)

	workloadInstance := &klcv1alpha1.KeptnWorkloadInstance{}
	newWorkloadInstance, err := r.createWorkloadInstance(workload)
	if err != nil {
		return reconcile.Result{}, err
	}

	// Try to find the workload instance
	err = r.Get(ctx, types.NamespacedName{Namespace: workload.Namespace, Name: workload.GetWorkloadInstanceName()}, workloadInstance)
	// If the workload instance does not exist, create it
	if err != nil && errors.IsNotFound(err) {
		err = r.Client.Create(ctx, newWorkloadInstance)
		if err != nil {
			r.Log.Error(err, "could not create Workload Instance")
			r.Recorder.Event(workload, "Warning", "WorkloadInstanceNotCreated", fmt.Sprintf("Could not create KeptnWorkloadInstance / Namespace: %s, Name: %s ", workloadInstance.Namespace, workloadInstance.Name))
			return ctrl.Result{}, err
		}
		r.Recorder.Event(workload, "Normal", "WorkloadInstanceCreated", fmt.Sprintf("Created KeptnWorkloadInstance / Namespace: %s, Name: %s ", workloadInstance.Namespace, workloadInstance.Name))
	} else if err != nil {
		r.Log.Error(err, "could not get Workload Instance")
		return ctrl.Result{}, err
	}

	// If the workload instance exists, check if it is up to date, update it if it is not
	if !reflect.DeepEqual(workloadInstance.Spec, newWorkloadInstance.Spec) {
		workloadInstance.Spec = newWorkloadInstance.Spec
		err = r.Client.Update(ctx, workloadInstance)
		if err != nil {
			r.Log.Error(err, "could not update Workload Instance")
			r.Recorder.Event(workload, "Warning", "WorkloadInstanceNotUpdated", fmt.Sprintf("Could not update KeptnWorkloadInstance / Namespace: %s, Name: %s ", workloadInstance.Namespace, workloadInstance.Name))
			return ctrl.Result{}, err
		}
		r.Recorder.Event(workload, "Normal", "WorkloadInstanceUpdated", fmt.Sprintf("Updated KeptnWorkloadInstance / Namespace: %s, Name: %s ", workloadInstance.Namespace, workloadInstance.Name))
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *KeptnWorkloadReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&klcv1alpha1.KeptnWorkload{}).
		Complete(r)
}

func (r *KeptnWorkloadReconciler) createWorkloadInstance(workload *klcv1alpha1.KeptnWorkload) (*klcv1alpha1.KeptnWorkloadInstance, error) {
	workloadInstance := &klcv1alpha1.KeptnWorkloadInstance{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: workload.Annotations,
			Name:        workload.GetWorkloadInstanceName(),
			Namespace:   workload.Namespace,
		},
		Spec: klcv1alpha1.KeptnWorkloadInstanceSpec{
			KeptnWorkloadSpec: workload.Spec,
			WorkloadName:      workload.Name,
		},
	}
	err := controllerutil.SetControllerReference(workload, workloadInstance, r.Scheme)
	if err != nil {
		r.Log.Error(err, "could not set controller reference for WorkloadInstance: "+workloadInstance.Name)
	}
	return workloadInstance, err
}

func (r *KeptnWorkloadReconciler) generateK8sEvent(workload *klcv1alpha1.KeptnWorkload, workloadInstance *klcv1alpha1.KeptnWorkloadInstance) *corev1.Event {
	return &corev1.Event{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName:    workloadInstance.Name + "-created-",
			Namespace:       workloadInstance.Namespace,
			ResourceVersion: "v1alpha1",
			Labels: map[string]string{
				common.AppAnnotation:      workload.Spec.AppName,
				common.WorkloadAnnotation: workloadInstance.Name,
			},
		},
		InvolvedObject: corev1.ObjectReference{
			Kind:      workloadInstance.Kind,
			Namespace: workloadInstance.Namespace,
			Name:      workloadInstance.Name,
		},
		Reason:  "created",
		Message: "workloadInstance " + workloadInstance.Name + " was created",
		Source: corev1.EventSource{
			Component: workloadInstance.Kind,
		},
		Type: "Normal",
		EventTime: metav1.MicroTime{
			Time: time.Now().UTC(),
		},
		FirstTimestamp: metav1.Time{
			Time: time.Now().UTC(),
		},
		LastTimestamp: metav1.Time{
			Time: time.Now().UTC(),
		},
		Action:              "created",
		ReportingController: "workload-controller",
		ReportingInstance:   "workload-controller",
	}
}
