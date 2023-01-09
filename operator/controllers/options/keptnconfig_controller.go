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

package options

import (
	"context"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	controllercommon "github.com/keptn/lifecycle-toolkit/operator/controllers/common"
	controllererrors "github.com/keptn/lifecycle-toolkit/operator/controllers/errors"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	kltoptionsv1alpha1 "github.com/keptn/lifecycle-toolkit/operator/apis/options/v1alpha1"
)

// KeptnConfigReconciler reconciles a KeptnConfig object
type KeptnConfigReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Log    logr.Logger
}

//+kubebuilder:rbac:groups=options.keptn.sh,resources=keptnconfigs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=options.keptn.sh,resources=keptnconfigs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=options.keptn.sh,resources=keptnconfigs/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *KeptnConfigReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	config := &kltoptionsv1alpha1.KeptnConfig{}
	err := r.Get(ctx, req.NamespacedName, config)
	if errors.IsNotFound(err) {
		return reconcile.Result{}, nil
	}
	if err != nil {
		return reconcile.Result{}, fmt.Errorf(controllererrors.ErrCannotRetrieveConfigMsg, err)
	}

	r.Log.Info("Reconciling Keptn Config", "config", config.Name)

	if err := controllercommon.GetOtelInstance().InitOtelCollector(config.Spec.OTelCollectorUrl); err != nil {
		r.Log.Error(err, "unable to initialize OTel tracer options")
		return ctrl.Result{Requeue: true, RequeueAfter: 10 * time.Second}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *KeptnConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&kltoptionsv1alpha1.KeptnConfig{}).
		Complete(r)
}
