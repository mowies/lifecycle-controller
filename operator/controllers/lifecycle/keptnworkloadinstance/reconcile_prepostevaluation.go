package keptnworkloadinstance

import (
	"context"
	"fmt"

	klcv1alpha2 "github.com/keptn/lifecycle-toolkit/operator/apis/lifecycle/v1alpha2"
	apicommon "github.com/keptn/lifecycle-toolkit/operator/apis/lifecycle/v1alpha2/common"
	controllercommon "github.com/keptn/lifecycle-toolkit/operator/controllers/common"
)

func (r *KeptnWorkloadInstanceReconciler) reconcilePrePostEvaluation(ctx context.Context, phaseCtx context.Context, workloadInstance *klcv1alpha2.KeptnWorkloadInstance, checkType apicommon.CheckType) (apicommon.KeptnState, error) {
	evaluationHandler := controllercommon.EvaluationHandler{
		Client:      r.Client,
		Recorder:    r.Recorder,
		Log:         r.Log,
		Tracer:      r.Tracer,
		Scheme:      r.Scheme,
		SpanHandler: r.SpanHandler,
	}

	evaluationCreateAttributes := controllercommon.EvaluationCreateAttributes{
		SpanName:  fmt.Sprintf(apicommon.CreateWorkloadEvalSpanName, checkType),
		CheckType: checkType,
	}

	newStatus, state, err := evaluationHandler.ReconcileEvaluations(ctx, phaseCtx, workloadInstance, evaluationCreateAttributes)
	if err != nil {
		return apicommon.StateUnknown, err
	}

	overallState := apicommon.GetOverallState(state)

	switch checkType {
	case apicommon.PreDeploymentEvaluationCheckType:
		workloadInstance.Status.PreDeploymentEvaluationStatus = overallState
		workloadInstance.Status.PreDeploymentEvaluationTaskStatus = newStatus
	case apicommon.PostDeploymentEvaluationCheckType:
		workloadInstance.Status.PostDeploymentEvaluationStatus = overallState
		workloadInstance.Status.PostDeploymentEvaluationTaskStatus = newStatus
	}

	// Write Status Field
	err = r.Client.Status().Update(ctx, workloadInstance)
	if err != nil {
		return apicommon.StateUnknown, err
	}
	return overallState, nil
}
