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

package keptntask

import (
	"context"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	klcv1alpha1 "github.com/keptn-sandbox/lifecycle-controller/operator/api/v1alpha1"
	"github.com/keptn-sandbox/lifecycle-controller/operator/api/v1alpha1/common"
	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// KeptnTaskReconciler reconciles a KeptnTask object
type KeptnTaskReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
	Log      logr.Logger
	Meters   common.KeptnMeters
}

//+kubebuilder:rbac:groups=lifecycle.keptn.sh,resources=keptntasks,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=lifecycle.keptn.sh,resources=keptntasks/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=lifecycle.keptn.sh,resources=keptntasks/finalizers,verbs=update
//+kubebuilder:rbac:groups=batch,resources=jobs,verbs=create;get;update;list;watch
//+kubebuilder:rbac:groups=batch,resources=jobs/status,verbs=get;list

func (r *KeptnTaskReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	r.Log.Info("Reconciling KeptnTask")
	task := &klcv1alpha1.KeptnTask{}

	if err := r.Client.Get(ctx, req.NamespacedName, task); err != nil {
		if errors.IsNotFound(err) {
			// taking down all associated K8s resources is handled by K8s
			r.Log.Info("KeptnTask resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		r.Log.Error(err, "Failed to get the KeptnTask")
		return ctrl.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
	}

	isRunning, err := r.IsJobRunning(ctx, *task, req.Namespace)
	if err != nil {
		r.Log.Error(err, "Could not check if job is running")
		return ctrl.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
	}

	if !isRunning {
		err = r.createJob(ctx, req, task)
		if err != nil {
			return ctrl.Result{Requeue: true}, err
		}
	}

	if task.Status.Status != common.StateSucceeded {
		err := r.updateJob(ctx, req, task)
		if err != nil {
			return ctrl.Result{Requeue: true, RequeueAfter: 10 * time.Second}, err
		}
	}

	r.Log.Info("Finished Reconciling KeptnTask")
	return ctrl.Result{Requeue: true, RequeueAfter: 10 * time.Second}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *KeptnTaskReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&klcv1alpha1.KeptnTask{}).
		Owns(&batchv1.Job{}).
		Complete(r)
}

func (r *KeptnTaskReconciler) IsJobRunning(ctx context.Context, task klcv1alpha1.KeptnTask, namespace string) (bool, error) {
	jobList := &batchv1.JobList{}

	jobLabels := client.MatchingLabels{}
	for k, v := range createKeptnLabels(task) {
		jobLabels[k] = v
	}

	if len(jobLabels) == 0 {
		return false, fmt.Errorf("no labels found for task: %s", task.Name)
	}

	if err := r.Client.List(ctx, jobList, client.InNamespace(namespace), jobLabels); err != nil {
		return false, err
	}

	if len(jobList.Items) > 0 {
		return true, nil
	}

	return false, nil
}
