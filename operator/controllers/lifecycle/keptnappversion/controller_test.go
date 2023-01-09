package keptnappversion

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"testing"

	lfcv1alpha2 "github.com/keptn/lifecycle-toolkit/operator/apis/lifecycle/v1alpha2"
	apicommon "github.com/keptn/lifecycle-toolkit/operator/apis/lifecycle/v1alpha2/common"
	controllercommon "github.com/keptn/lifecycle-toolkit/operator/controllers/common"
	"github.com/keptn/lifecycle-toolkit/operator/controllers/common/fake"
	interfacesfake "github.com/keptn/lifecycle-toolkit/operator/controllers/lifecycle/interfaces/fake"
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/trace"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

type contextID string

const CONTEXTID contextID = "start"

// this test checks if the chain of reconcile events is correct
func TestKeptnAppVersionReconciler_reconcile(t *testing.T) {

	r, eventChannel, tracer, _ := setupReconciler(t)

	tests := []struct {
		name       string
		req        ctrl.Request
		wantErr    error
		events     []string //check correct events are generated
		startTrace bool
	}{
		{
			name: "new appVersion with no workload nor evaluation should finish",
			req: ctrl.Request{
				NamespacedName: types.NamespacedName{
					Namespace: "default",
					Name:      "myappversion-1.0.0",
				},
			},
			wantErr: nil,
			events: []string{
				`AppPreDeployTasksStarted`,
				`AppPreDeployTasksSucceeded`,
				`AppPreDeployEvaluationsSucceeded`,
				`AppDeploySucceeded`,
				`AppPostDeployTasksSucceeded`,
				`AppPostDeployEvaluationsSucceeded`,
				`AppPostDeployEvaluationsFinished`,
			},
			startTrace: true,
		},
		{
			name: "notfound should not return error nor event",
			req: ctrl.Request{
				NamespacedName: types.NamespacedName{
					Namespace: "default",
					Name:      "mynotthereapp",
				},
			},
			wantErr: nil,
		},
		{
			name: "existing appVersion has finished",
			req: ctrl.Request{
				NamespacedName: types.NamespacedName{
					Namespace: "default",
					Name:      "myfinishedapp-1.0.0",
				},
			},
			wantErr:    nil,
			events:     []string{`AppPostDeployEvaluationsFinished`},
			startTrace: true,
		},
	}

	//setting up fakeclient CRD data

	err := controllercommon.AddAppVersion(r.Client, "default", "myappversion", "1.0.0", nil, lfcv1alpha2.KeptnAppVersionStatus{Status: apicommon.StatePending})
	require.Nil(t, err)
	err = controllercommon.AddAppVersion(r.Client, "default", "myfinishedapp", "1.0.0", nil, createFinishedAppVersionStatus())
	require.Nil(t, err)

	traces := 0
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := r.Reconcile(context.WithValue(context.TODO(), CONTEXTID, tt.req.Name), tt.req)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Reconcile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.events != nil {
				for _, e := range tt.events {
					event := <-eventChannel
					assert.Equal(t, strings.Contains(event, tt.req.Name), true, "wrong appversion")
					assert.Equal(t, strings.Contains(event, tt.req.Namespace), true, "wrong namespace")
					assert.Equal(t, strings.Contains(event, e), true, fmt.Sprintf("no %s found in %s", e, event))
				}

			}
			if tt.startTrace {
				//A different trace for each app-version
				assert.Equal(t, tracer.StartCalls()[traces].SpanName, "reconcile_app_version")
				assert.Equal(t, tracer.StartCalls()[traces].Ctx.Value(CONTEXTID), tt.req.Name)
				traces++
			}
		})

	}

}

func createFinishedAppVersionStatus() lfcv1alpha2.KeptnAppVersionStatus {
	return lfcv1alpha2.KeptnAppVersionStatus{
		CurrentPhase:                       apicommon.PhaseCompleted.ShortName,
		PreDeploymentStatus:                apicommon.StateSucceeded,
		PostDeploymentStatus:               apicommon.StateSucceeded,
		PreDeploymentEvaluationStatus:      apicommon.StateSucceeded,
		PostDeploymentEvaluationStatus:     apicommon.StateSucceeded,
		PreDeploymentTaskStatus:            []lfcv1alpha2.TaskStatus{{Status: apicommon.StateSucceeded}},
		PostDeploymentTaskStatus:           []lfcv1alpha2.TaskStatus{{Status: apicommon.StateSucceeded}},
		PreDeploymentEvaluationTaskStatus:  []lfcv1alpha2.EvaluationStatus{{Status: apicommon.StateSucceeded}},
		PostDeploymentEvaluationTaskStatus: []lfcv1alpha2.EvaluationStatus{{Status: apicommon.StateSucceeded}},
		WorkloadOverallStatus:              apicommon.StateSucceeded,
		WorkloadStatus:                     []lfcv1alpha2.WorkloadStatus{{Status: apicommon.StateSucceeded}},
		Status:                             apicommon.StateSucceeded,
	}
}

func setupReconcilerWithMeters(t *testing.T) *KeptnAppVersionReconciler {
	//setup logger
	opts := zap.Options{
		Development: true,
	}
	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))

	//fake a tracer
	tr := &interfacesfake.ITracerMock{StartFunc: func(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
		return ctx, trace.SpanFromContext(ctx)
	}}

	r := &KeptnAppVersionReconciler{
		Log:    ctrl.Log.WithName("test-appVersionController"),
		Tracer: tr,
		Meters: controllercommon.InitAppMeters(),
	}
	return r
}

func setupReconciler(t *testing.T) (*KeptnAppVersionReconciler, chan string, *interfacesfake.ITracerMock, *fake.ISpanHandlerMock) {
	//setup logger
	opts := zap.Options{
		Development: true,
	}
	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))

	//fake a tracer
	tr := &interfacesfake.ITracerMock{StartFunc: func(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
		return ctx, trace.SpanFromContext(ctx)
	}}

	//fake span handler

	spanRecorder := &fake.ISpanHandlerMock{
		GetSpanFunc: func(ctx context.Context, tracer trace.Tracer, reconcileObject client.Object, phase string) (context.Context, trace.Span, error) {
			return ctx, trace.SpanFromContext(ctx), nil
		},
		UnbindSpanFunc: func(reconcileObject client.Object, phase string) error { return nil },
	}

	fakeClient, err := fake.NewClient()
	if err != nil {
		t.Errorf("Reconcile() error when setting up fake client %v", err)
	}
	recorder := record.NewFakeRecorder(100)
	r := &KeptnAppVersionReconciler{
		Client:      fakeClient,
		Scheme:      scheme.Scheme,
		Recorder:    recorder,
		Log:         ctrl.Log.WithName("test-appVersionController"),
		Tracer:      tr,
		SpanHandler: spanRecorder,
		Meters:      controllercommon.InitAppMeters(),
	}
	return r, recorder.Events, tr, spanRecorder
}

func TestKeptnApVersionReconciler_setupSpansContexts(t *testing.T) {

	r := setupReconcilerWithMeters(t)
	type args struct {
		ctx        context.Context
		appVersion *lfcv1alpha2.KeptnAppVersion
	}
	tests := []struct {
		name    string
		args    args
		baseCtx context.Context
		appCtx  context.Context
	}{
		{name: "Current trace ctx should be != than app trace context",
			args: args{
				ctx:        context.WithValue(context.TODO(), CONTEXTID, 1),
				appVersion: &lfcv1alpha2.KeptnAppVersion{},
			},
			baseCtx: context.WithValue(context.TODO(), CONTEXTID, 1),
			appCtx:  context.TODO(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, ctxAppTrace, _, _ := setupSpansContexts(tt.args.ctx, tt.args.appVersion, r)
			if !reflect.DeepEqual(ctx, tt.baseCtx) {
				t.Errorf("setupSpansContexts() got: %v as baseCtx, wanted: %v", ctx, tt.baseCtx)
			}
			if !reflect.DeepEqual(ctxAppTrace, tt.appCtx) {
				t.Errorf("setupSpansContexts() got: %v as appCtx, wanted: %v", ctxAppTrace, tt.appCtx)
			}
		})
	}
}
