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

package keptnworkloadinstance

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	"github.com/google/uuid"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"time"

	klcv1alpha1 "github.com/keptn-sandbox/lifecycle-controller/operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// KeptnWorkloadInstanceReconciler reconciles a KeptnWorkloadInstance object
type KeptnWorkloadInstanceReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
	Log      logr.Logger
}

//+kubebuilder:rbac:groups=lifecycle.keptn.sh,resources=keptnworkloadinstances,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=lifecycle.keptn.sh,resources=keptnworkloadinstances/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=lifecycle.keptn.sh,resources=keptnworkloadinstances/finalizers,verbs=update
//+kubebuilder:rbac:groups=lifecycle.keptn.sh,resources=events,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=lifecycle.keptn.sh,resources=events/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=lifecycle.keptn.sh,resources=events/finalizers,verbs=update
//+kubebuilder:rbac:groups=core,resources=events,verbs=create;watch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the KeptnWorkloadInstance object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.2/pkg/reconcile
func (r *KeptnWorkloadInstanceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	r.Log.Info("Searching for Keptn Workload Instance")

	workloadInstance := &klcv1alpha1.KeptnWorkloadInstance{}
	err := r.Get(ctx, req.NamespacedName, workloadInstance)
	if errors.IsNotFound(err) {
		return reconcile.Result{}, nil
	}

	if err != nil {
		return reconcile.Result{}, fmt.Errorf("could not fetch KeptnWorkloadInstance: %+v", err)
	}

	// check if the workloadInstance is completed (scheduled checks are finished)
	if workloadInstance.IsCompleted() {
		return reconcile.Result{}, nil
	}

	r.Log.Info("Reconciling KeptnWorkloadInstance", "workloadInstance", workloadInstance.Name)

	if workloadInstance.IsDeploymentCheckNotCreated() {
		r.Log.Info("Deployment checks do not exist, creating")

		preDeploymentCheckName, err := r.startPreDeploymentChecks(ctx, workloadInstance)
		if err != nil {
			r.Log.Error(err, "Could not start pre-deployment checks")
			return reconcile.Result{}, err
		}

		workloadInstance.Status.PreDeploymentTaskName = preDeploymentCheckName
		workloadInstance.Status.PreDeploymentPhase = klcv1alpha1.WorkloadInstancePhaseRunning

		r.Recorder.Event(workloadInstance, "Normal", "PreDeploymentChecksStarted", fmt.Sprintf("Started Pre-Deployment Checks / Namespace: %s, Name: %s ", workloadInstance.Namespace, workloadInstance.Name))

		if err := r.Status().Update(ctx, workloadInstance); err != nil {
			r.Log.Error(err, "Could not update KeptnWorkloadInstance")
			return reconcile.Result{}, err
		}
		return ctrl.Result{Requeue: true, RequeueAfter: 5 * time.Second}, nil
	}

	preDeploymentChecksEvent, err := r.getPreDeploymentChecksEvent(ctx, workloadInstance)
	if err != nil {
		r.Log.Error(err, "Could not retrieve pre-deployment checks Event")
		return reconcile.Result{}, err
	}

	r.Log.Info("Checking status")

	if preDeploymentChecksEvent.IsCompleted() {
		if preDeploymentChecksEvent.Status.Phase == klcv1alpha1.EventFailed {
			workloadInstance.Status.PreDeploymentPhase = klcv1alpha1.WorkloadInstancePhaseFailed
		} else {
			workloadInstance.Status.PreDeploymentPhase = klcv1alpha1.WorkloadInstancePhaseSucceeded
		}

		if err := r.Delete(ctx, preDeploymentChecksEvent); err != nil {
			r.Log.Error(err, "Could not delete Event")
			return reconcile.Result{}, err
		}

		if err := r.Status().Update(ctx, workloadInstance); err != nil {
			r.Log.Error(err, "Could not update KeptnWorkloadInstance")
			return reconcile.Result{}, err
		}

		r.Recorder.Event(workloadInstance, "Normal", "PreDeploymentChecksFinished", fmt.Sprintf("Finished Pre-Deployment Checks / Namespace: %s, Name: %s ", workloadInstance.Namespace, workloadInstance.Name))

		return reconcile.Result{}, nil
	}

	return ctrl.Result{Requeue: true, RequeueAfter: 5 * time.Second}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *KeptnWorkloadInstanceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&klcv1alpha1.KeptnWorkloadInstance{}).
		Complete(r)
}

func (r *KeptnWorkloadInstanceReconciler) generateSuffix() string {
	uid := uuid.New().String()
	return uid[:10]
}

func (r *KeptnWorkloadInstanceReconciler) startPreDeploymentChecks(ctx context.Context, workloadInstance *klcv1alpha1.KeptnWorkloadInstance) (string, error) {
	event := &klcv1alpha1.Event{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: workloadInstance.Annotations,
			Name:        workloadInstance.Name + "-" + r.generateSuffix(),
			Namespace:   workloadInstance.Namespace,
		},
		Spec: klcv1alpha1.EventSpec{
			Service:     workloadInstance.Name,
			Application: workloadInstance.Spec.AppName,
			JobSpec:     workloadInstance.Spec.PreDeploymentCheck.JobSpec,
		},
	}
	for i := 0; i < 5; i++ {
		if err := r.Create(ctx, event); err != nil {
			if errors.IsAlreadyExists(err) {
				event.Name = workloadInstance.Name + "-" + r.generateSuffix()
				continue
			}
			return "", err
		}
		break
	}
	return event.Name, nil
}

func (r *KeptnWorkloadInstanceReconciler) getPreDeploymentChecksEvent(ctx context.Context, workloadInstance *klcv1alpha1.KeptnWorkloadInstance) (*klcv1alpha1.Event, error) {
	event := &klcv1alpha1.Event{}
	err := r.Get(ctx, types.NamespacedName{Name: workloadInstance.Status.PreDeploymentTaskName, Namespace: workloadInstance.Namespace}, event)
	if errors.IsNotFound(err) {
		return nil, err
	}

	return event, nil
}
