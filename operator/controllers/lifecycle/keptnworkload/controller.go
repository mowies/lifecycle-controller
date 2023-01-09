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

package keptnworkload

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	klcv1alpha2 "github.com/keptn/lifecycle-toolkit/operator/apis/lifecycle/v1alpha2"
	controllererrors "github.com/keptn/lifecycle-toolkit/operator/controllers/errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// KeptnWorkloadReconciler reconciles a KeptnWorkload object
type KeptnWorkloadReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
	Log      logr.Logger
	Tracer   trace.Tracer
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

	workload := &klcv1alpha2.KeptnWorkload{}
	err := r.Get(ctx, req.NamespacedName, workload)
	if errors.IsNotFound(err) {
		return reconcile.Result{}, nil
	}
	if err != nil {
		return reconcile.Result{}, fmt.Errorf(controllererrors.ErrCannotRetrieveWorkloadMsg, err)
	}

	traceContextCarrier := propagation.MapCarrier(workload.Annotations)
	ctx = otel.GetTextMapPropagator().Extract(ctx, traceContextCarrier)

	ctx, span := r.Tracer.Start(ctx, "reconcile_workload", trace.WithSpanKind(trace.SpanKindConsumer))
	defer span.End()

	workload.SetSpanAttributes(span)

	r.Log.Info("Reconciling Keptn Workload", "workload", workload.Name)

	workloadInstance := &klcv1alpha2.KeptnWorkloadInstance{}

	// Try to find the workload instance
	err = r.Get(ctx, types.NamespacedName{Namespace: workload.Namespace, Name: workload.GetWorkloadInstanceName()}, workloadInstance)
	// If the workload instance does not exist, create it
	if errors.IsNotFound(err) {
		workloadInstance, err := r.createWorkloadInstance(ctx, workload)
		if err != nil {
			span.SetStatus(codes.Error, err.Error())
			return reconcile.Result{}, err
		}
		err = r.Client.Create(ctx, workloadInstance)
		if err != nil {
			r.Log.Error(err, "could not create Workload Instance")
			span.SetStatus(codes.Error, err.Error())
			r.Recorder.Event(workload, "Warning", "WorkloadInstanceNotCreated", fmt.Sprintf("Could not create KeptnWorkloadInstance / Namespace: %s, Name: %s ", workloadInstance.Namespace, workloadInstance.Name))
			return ctrl.Result{}, err
		}
		r.Recorder.Event(workload, "Normal", "WorkloadInstanceCreated", fmt.Sprintf("Created KeptnWorkloadInstance / Namespace: %s, Name: %s ", workloadInstance.Namespace, workloadInstance.Name))
		workload.Status.CurrentVersion = workload.Spec.Version
		if err := r.Client.Status().Update(ctx, workload); err != nil {
			r.Log.Error(err, "could not update Current Version of Workload")
			return ctrl.Result{}, err
		}
		return ctrl.Result{}, nil
	}
	if err != nil {
		r.Log.Error(err, "could not get Workload Instance")
		span.SetStatus(codes.Error, err.Error())
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *KeptnWorkloadReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&klcv1alpha2.KeptnWorkload{}, builder.WithPredicates(predicate.GenerationChangedPredicate{})).
		Complete(r)
}

func (r *KeptnWorkloadReconciler) createWorkloadInstance(ctx context.Context, workload *klcv1alpha2.KeptnWorkload) (*klcv1alpha2.KeptnWorkloadInstance, error) {
	ctx, span := r.Tracer.Start(ctx, "create_workload_instance", trace.WithSpanKind(trace.SpanKindProducer))
	defer span.End()

	workload.SetSpanAttributes(span)

	// create TraceContext
	// follow up with a Keptn propagator that JSON-encoded the OTel map into our own key
	traceContextCarrier := propagation.MapCarrier{}
	otel.GetTextMapPropagator().Inject(ctx, traceContextCarrier)

	previousVersion := ""
	if workload.Spec.Version != workload.Status.CurrentVersion {
		previousVersion = workload.Status.CurrentVersion
	}

	workloadInstance := workload.GenerateWorkloadInstance(previousVersion, traceContextCarrier)
	err := controllerutil.SetControllerReference(workload, &workloadInstance, r.Scheme)
	if err != nil {
		r.Log.Error(err, "could not set controller reference for WorkloadInstance: "+workloadInstance.Name)
	}

	return &workloadInstance, err
}
