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
	"time"

	"github.com/go-logr/logr"
	version "github.com/hashicorp/go-version"
	klcv1alpha2 "github.com/keptn/lifecycle-toolkit/operator/apis/lifecycle/v1alpha2"
	apicommon "github.com/keptn/lifecycle-toolkit/operator/apis/lifecycle/v1alpha2/common"
	controllercommon "github.com/keptn/lifecycle-toolkit/operator/controllers/common"
	controllererrors "github.com/keptn/lifecycle-toolkit/operator/controllers/errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// KeptnWorkloadInstanceReconciler reconciles a KeptnWorkloadInstance object
type KeptnWorkloadInstanceReconciler struct {
	client.Client
	Scheme      *runtime.Scheme
	Recorder    record.EventRecorder
	Log         logr.Logger
	Meters      apicommon.KeptnMeters
	Tracer      trace.Tracer
	SpanHandler *controllercommon.SpanHandler
}

//+kubebuilder:rbac:groups=lifecycle.keptn.sh,resources=keptnworkloadinstances,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=lifecycle.keptn.sh,resources=keptnworkloadinstances/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=lifecycle.keptn.sh,resources=keptnworkloadinstances/finalizers,verbs=update
//+kubebuilder:rbac:groups=lifecycle.keptn.sh,resources=keptntasks,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=lifecycle.keptn.sh,resources=keptntasks/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=lifecycle.keptn.sh,resources=keptntasks/finalizers,verbs=update
//+kubebuilder:rbac:groups=core,resources=events,verbs=create;watch;patch
//+kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch
//+kubebuilder:rbac:groups=apps,resources=replicasets;deployments;statefulsets;daemonsets,verbs=get;list;watch

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

	//retrieve workload instance
	workloadInstance := &klcv1alpha2.KeptnWorkloadInstance{}
	err := r.Get(ctx, req.NamespacedName, workloadInstance)
	if errors.IsNotFound(err) {
		return reconcile.Result{}, nil
	}

	if err != nil {
		r.Log.Error(err, "Workload Instance not found")
		return reconcile.Result{}, fmt.Errorf(controllererrors.ErrCannotRetrieveWorkloadInstancesMsg, err)
	}

	//setup otel
	traceContextCarrier := propagation.MapCarrier(workloadInstance.Annotations)
	ctx = otel.GetTextMapPropagator().Extract(ctx, traceContextCarrier)

	ctx, span := r.Tracer.Start(ctx, "reconcile_workload_instance", trace.WithSpanKind(trace.SpanKindConsumer))
	defer span.End()

	workloadInstance.SetSpanAttributes(span)

	workloadInstance.SetStartTime()

	defer func(span trace.Span, workloadInstance *klcv1alpha2.KeptnWorkloadInstance) {
		if workloadInstance.IsEndTimeSet() {
			r.Log.Info("Increasing deployment count")
			attrs := workloadInstance.GetMetricsAttributes()
			r.Meters.DeploymentCount.Add(ctx, 1, attrs...)
		}
		span.End()
	}(span, workloadInstance)

	//Wait for pre-evaluation checks of App
	phase := apicommon.PhaseAppPreEvaluation

	found, appVersion, err := r.getAppVersionForWorkloadInstance(ctx, workloadInstance)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		controllercommon.RecordEvent(r.Recorder, phase, "Warning", workloadInstance, "GetAppVersionFailed", "has failed since app could not be retrieved", workloadInstance.GetVersion())
		return reconcile.Result{Requeue: true, RequeueAfter: 10 * time.Second}, fmt.Errorf(controllererrors.ErrCannotFetchAppVersionForWorkloadInstanceMsg + err.Error())
	} else if !found {
		span.SetStatus(codes.Error, "app could not be found")
		controllercommon.RecordEvent(r.Recorder, phase, "Warning", workloadInstance, "AppVersionNotFound", "has failed since app could not be found", workloadInstance.GetVersion())
		return reconcile.Result{Requeue: true, RequeueAfter: 10 * time.Second}, fmt.Errorf(controllererrors.ErrCannotFetchAppVersionForWorkloadInstanceMsg)
	}

	appPreEvalStatus := appVersion.Status.PreDeploymentEvaluationStatus
	if !appPreEvalStatus.IsSucceeded() {
		if appPreEvalStatus.IsFailed() {
			controllercommon.RecordEvent(r.Recorder, phase, "Warning", workloadInstance, "Failed", "has failed since app has failed", workloadInstance.GetVersion())
			return ctrl.Result{Requeue: true, RequeueAfter: 20 * time.Second}, nil
		}
		controllercommon.RecordEvent(r.Recorder, phase, "Normal", workloadInstance, "NotFinished", "Pre evaluations tasks for app not finished", workloadInstance.GetVersion())
		return ctrl.Result{Requeue: true, RequeueAfter: 20 * time.Second}, nil
	}

	controllercommon.RecordEvent(r.Recorder, phase, "Normal", workloadInstance, "FinishedSuccess", "Pre evaluations tasks for app have finished successfully", workloadInstance.GetVersion())

	//Wait for pre-deployment checks of Workload
	phase = apicommon.PhaseWorkloadPreDeployment
	phaseHandler := controllercommon.PhaseHandler{
		Client:      r.Client,
		Recorder:    r.Recorder,
		Log:         r.Log,
		SpanHandler: r.SpanHandler,
	}

	// set the App trace id if not already set
	if len(workloadInstance.Spec.TraceId) < 1 {
		appDeploymentTraceID := appVersion.Status.PhaseTraceIDs[apicommon.PhaseAppDeployment.ShortName]
		if appDeploymentTraceID != nil {
			workloadInstance.Spec.TraceId = appDeploymentTraceID
		} else {
			workloadInstance.Spec.TraceId = appVersion.Spec.TraceId
		}
		if err := r.Update(ctx, workloadInstance); err != nil {
			return ctrl.Result{}, err
		}
	}

	appTraceContextCarrier := propagation.MapCarrier(workloadInstance.Spec.TraceId)
	ctxAppTrace := otel.GetTextMapPropagator().Extract(context.TODO(), appTraceContextCarrier)

	// this will be the parent span for all phases of the WorkloadInstance
	ctxWorkloadTrace, spanWorkloadTrace, err := r.SpanHandler.GetSpan(ctxAppTrace, r.Tracer, workloadInstance, "")
	if err != nil {
		r.Log.Error(err, "could not get span")
	}

	if workloadInstance.Status.CurrentPhase == "" {
		spanWorkloadTrace.AddEvent("WorkloadInstance Pre-Deployment Tasks started", trace.WithTimestamp(time.Now()))
		controllercommon.RecordEvent(r.Recorder, phase, "Normal", workloadInstance, "Started", "have started", workloadInstance.GetVersion())
	}

	if !workloadInstance.IsPreDeploymentSucceeded() {
		reconcilePre := func(phaseCtx context.Context) (apicommon.KeptnState, error) {
			return r.reconcilePrePostDeployment(ctx, phaseCtx, workloadInstance, apicommon.PreDeploymentCheckType)
		}
		result, err := phaseHandler.HandlePhase(ctx, ctxWorkloadTrace, r.Tracer, workloadInstance, phase, span, reconcilePre)
		if !result.Continue {
			return result.Result, err
		}
	}

	//Wait for pre-evaluation checks of Workload
	phase = apicommon.PhaseWorkloadPreEvaluation
	if !workloadInstance.IsPreDeploymentEvaluationSucceeded() {
		reconcilePreEval := func(phaseCtx context.Context) (apicommon.KeptnState, error) {
			return r.reconcilePrePostEvaluation(ctx, phaseCtx, workloadInstance, apicommon.PreDeploymentEvaluationCheckType)
		}
		result, err := phaseHandler.HandlePhase(ctx, ctxWorkloadTrace, r.Tracer, workloadInstance, phase, span, reconcilePreEval)
		if !result.Continue {
			return result.Result, err
		}
	}

	//Wait for deployment of Workload
	phase = apicommon.PhaseWorkloadDeployment
	if !workloadInstance.IsDeploymentSucceeded() {
		reconcileWorkloadInstance := func(phaseCtx context.Context) (apicommon.KeptnState, error) {
			return r.reconcileDeployment(ctx, workloadInstance)
		}
		result, err := phaseHandler.HandlePhase(ctx, ctxWorkloadTrace, r.Tracer, workloadInstance, phase, span, reconcileWorkloadInstance)
		if !result.Continue {
			return result.Result, err
		}
	}

	//Wait for post-deployment checks of Workload
	phase = apicommon.PhaseWorkloadPostDeployment
	if !workloadInstance.IsPostDeploymentSucceeded() {
		reconcilePostDeployment := func(phaseCtx context.Context) (apicommon.KeptnState, error) {
			return r.reconcilePrePostDeployment(ctx, phaseCtx, workloadInstance, apicommon.PostDeploymentCheckType)
		}
		result, err := phaseHandler.HandlePhase(ctx, ctxWorkloadTrace, r.Tracer, workloadInstance, phase, span, reconcilePostDeployment)
		if !result.Continue {
			return result.Result, err
		}
	}

	//Wait for post-evaluation checks of Workload
	phase = apicommon.PhaseWorkloadPostEvaluation
	if !workloadInstance.IsPostDeploymentEvaluationSucceeded() {
		reconcilePostEval := func(phaseCtx context.Context) (apicommon.KeptnState, error) {
			return r.reconcilePrePostEvaluation(ctx, phaseCtx, workloadInstance, apicommon.PostDeploymentEvaluationCheckType)
		}
		result, err := phaseHandler.HandlePhase(ctx, ctxWorkloadTrace, r.Tracer, workloadInstance, phase, span, reconcilePostEval)
		if !result.Continue {
			return result.Result, err
		}
	}

	// WorkloadInstance is completed at this place
	if !workloadInstance.IsEndTimeSet() {
		workloadInstance.Status.CurrentPhase = apicommon.PhaseCompleted.ShortName
		workloadInstance.Status.Status = apicommon.StateSucceeded
		workloadInstance.SetEndTime()
	}

	err = r.Client.Status().Update(ctx, workloadInstance)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return ctrl.Result{Requeue: true}, err
	}

	attrs := workloadInstance.GetMetricsAttributes()

	// metrics: add deployment duration
	duration := workloadInstance.Status.EndTime.Time.Sub(workloadInstance.Status.StartTime.Time)
	r.Meters.DeploymentDuration.Record(ctx, duration.Seconds(), attrs...)

	spanWorkloadTrace.AddEvent(workloadInstance.Name + " has finished")
	spanWorkloadTrace.SetStatus(codes.Ok, "Finished")
	spanWorkloadTrace.End()
	if err := r.SpanHandler.UnbindSpan(workloadInstance, ""); err != nil {
		r.Log.Error(err, controllererrors.ErrCouldNotUnbindSpan, workloadInstance.Name)
	}

	controllercommon.RecordEvent(r.Recorder, phase, "Normal", workloadInstance, "Finished", "is finished", workloadInstance.GetVersion())

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *KeptnWorkloadInstanceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		// predicate disabling the auto reconciliation after updating the object status
		For(&klcv1alpha2.KeptnWorkloadInstance{}, builder.WithPredicates(predicate.GenerationChangedPredicate{})).
		Complete(r)
}

func (r *KeptnWorkloadInstanceReconciler) getAppVersionForWorkloadInstance(ctx context.Context, wli *klcv1alpha2.KeptnWorkloadInstance) (bool, klcv1alpha2.KeptnAppVersion, error) {
	apps := &klcv1alpha2.KeptnAppVersionList{}

	// TODO add label selector for looking up by name?
	if err := r.Client.List(ctx, apps, client.InNamespace(wli.Namespace)); err != nil {
		return false, klcv1alpha2.KeptnAppVersion{}, err
	}

	// due to effectivity reasons deprecated KeptnAppVersions are removed from the list, as there is
	// no point in iterating through them in the next steps
	apps.RemoveDeprecated()

	workloadFound, latestVersion, err := getLatestAppVersion(apps, wli)
	if err != nil {
		r.Log.Error(err, "could not look up KeptnAppVersion for WorkloadInstance")
		return false, latestVersion, err
	}

	if latestVersion.Spec.Version == "" || !workloadFound {
		return false, klcv1alpha2.KeptnAppVersion{}, nil
	}
	return true, latestVersion, nil
}

func getLatestAppVersion(apps *klcv1alpha2.KeptnAppVersionList, wli *klcv1alpha2.KeptnWorkloadInstance) (bool, klcv1alpha2.KeptnAppVersion, error) {
	latestVersion := klcv1alpha2.KeptnAppVersion{}
	// ignore the potential error since this can not return an error with 0.0.0
	oldVersion, _ := version.NewVersion("0.0.0")

	workloadFound := false
	for _, app := range apps.Items {
		if app.Spec.AppName == wli.Spec.AppName {
			for _, appWorkload := range app.Spec.Workloads {
				if appWorkload.Version == wli.Spec.Version && app.GetWorkloadNameOfApp(appWorkload.Name) == wli.Spec.WorkloadName {
					workloadFound = true
					newVersion, err := version.NewVersion(app.Spec.Version)
					if err != nil {
						return false, klcv1alpha2.KeptnAppVersion{}, err
					}
					if newVersion.GreaterThan(oldVersion) {
						latestVersion = app
						oldVersion = newVersion
					}
				}
			}
		}
	}
	return workloadFound, latestVersion, nil
}
