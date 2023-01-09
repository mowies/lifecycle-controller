package common

import (
	"context"
	"strings"
	"testing"

	"github.com/keptn/lifecycle-toolkit/operator/apis/lifecycle/v1alpha2"
	apicommon "github.com/keptn/lifecycle-toolkit/operator/apis/lifecycle/v1alpha2/common"
	kltfake "github.com/keptn/lifecycle-toolkit/operator/controllers/common/fake"
	controllererrors "github.com/keptn/lifecycle-toolkit/operator/controllers/errors"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/trace"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestTaskHandler(t *testing.T) {
	tests := []struct {
		name            string
		object          client.Object
		createAttr      TaskCreateAttributes
		wantStatus      []v1alpha2.TaskStatus
		wantSummary     apicommon.StatusSummary
		taskObj         v1alpha2.KeptnTask
		wantErr         error
		getSpanCalls    int
		unbindSpanCalls int
	}{
		{
			name:            "cannot unwrap object",
			object:          &v1alpha2.KeptnTask{},
			taskObj:         v1alpha2.KeptnTask{},
			createAttr:      TaskCreateAttributes{},
			wantStatus:      nil,
			wantSummary:     apicommon.StatusSummary{},
			wantErr:         controllererrors.ErrCannotWrapToPhaseItem,
			getSpanCalls:    0,
			unbindSpanCalls: 0,
		},
		{
			name:    "no tasks",
			object:  &v1alpha2.KeptnAppVersion{},
			taskObj: v1alpha2.KeptnTask{},
			createAttr: TaskCreateAttributes{
				SpanName:       "",
				TaskDefinition: "",
				CheckType:      apicommon.PreDeploymentCheckType,
			},
			wantStatus:      []v1alpha2.TaskStatus(nil),
			wantSummary:     apicommon.StatusSummary{},
			wantErr:         nil,
			getSpanCalls:    0,
			unbindSpanCalls: 0,
		},
		{
			name: "task not started",
			object: &v1alpha2.KeptnAppVersion{
				Spec: v1alpha2.KeptnAppVersionSpec{
					KeptnAppSpec: v1alpha2.KeptnAppSpec{
						PreDeploymentTasks: []string{"task-def"},
					},
				},
			},
			taskObj: v1alpha2.KeptnTask{},
			createAttr: TaskCreateAttributes{
				SpanName:       "",
				TaskDefinition: "task-def",
				CheckType:      apicommon.PreDeploymentCheckType,
			},
			wantStatus: []v1alpha2.TaskStatus{
				{
					TaskDefinitionName: "task-def",
					Status:             apicommon.StatePending,
					TaskName:           "pre-task-def-",
				},
			},
			wantSummary:     apicommon.StatusSummary{Total: 1, Pending: 1},
			wantErr:         nil,
			getSpanCalls:    1,
			unbindSpanCalls: 0,
		},
		{
			name: "already done task",
			object: &v1alpha2.KeptnAppVersion{
				Spec: v1alpha2.KeptnAppVersionSpec{
					KeptnAppSpec: v1alpha2.KeptnAppSpec{
						PreDeploymentTasks: []string{"task-def"},
					},
				},
				Status: v1alpha2.KeptnAppVersionStatus{
					PreDeploymentStatus: apicommon.StateSucceeded,
					PreDeploymentTaskStatus: []v1alpha2.TaskStatus{
						{
							TaskDefinitionName: "task-def",
							Status:             apicommon.StateSucceeded,
							TaskName:           "pre-task-def-",
						},
					},
				},
			},
			taskObj: v1alpha2.KeptnTask{},
			createAttr: TaskCreateAttributes{
				SpanName:       "",
				TaskDefinition: "task-def",
				CheckType:      apicommon.PreDeploymentCheckType,
			},
			wantStatus: []v1alpha2.TaskStatus{
				{
					TaskDefinitionName: "task-def",
					Status:             apicommon.StateSucceeded,
					TaskName:           "pre-task-def-",
				},
			},
			wantSummary:     apicommon.StatusSummary{Total: 1, Succeeded: 1},
			wantErr:         nil,
			getSpanCalls:    0,
			unbindSpanCalls: 0,
		},
		{
			name: "failed task",
			object: &v1alpha2.KeptnAppVersion{
				ObjectMeta: v1.ObjectMeta{
					Namespace: "namespace",
				},
				Spec: v1alpha2.KeptnAppVersionSpec{
					KeptnAppSpec: v1alpha2.KeptnAppSpec{
						PreDeploymentTasks: []string{"task-def"},
					},
				},
				Status: v1alpha2.KeptnAppVersionStatus{
					PreDeploymentStatus: apicommon.StateSucceeded,
					PreDeploymentTaskStatus: []v1alpha2.TaskStatus{
						{
							TaskDefinitionName: "task-def",
							Status:             apicommon.StateProgressing,
							TaskName:           "pre-task-def-",
						},
					},
				},
			},
			taskObj: v1alpha2.KeptnTask{
				ObjectMeta: v1.ObjectMeta{
					Namespace: "namespace",
					Name:      "pre-task-def-",
				},
				Status: v1alpha2.KeptnTaskStatus{
					Status: apicommon.StateFailed,
				},
			},
			createAttr: TaskCreateAttributes{
				SpanName:       "",
				TaskDefinition: "task-def",
				CheckType:      apicommon.PreDeploymentCheckType,
			},
			wantStatus: []v1alpha2.TaskStatus{
				{
					TaskDefinitionName: "task-def",
					Status:             apicommon.StateFailed,
					TaskName:           "pre-task-def-",
				},
			},
			wantSummary:     apicommon.StatusSummary{Total: 1, Failed: 1},
			wantErr:         nil,
			getSpanCalls:    1,
			unbindSpanCalls: 1,
		},
		{
			name: "succeeded task",
			object: &v1alpha2.KeptnAppVersion{
				ObjectMeta: v1.ObjectMeta{
					Namespace: "namespace",
				},
				Spec: v1alpha2.KeptnAppVersionSpec{
					KeptnAppSpec: v1alpha2.KeptnAppSpec{
						PreDeploymentTasks: []string{"task-def"},
					},
				},
				Status: v1alpha2.KeptnAppVersionStatus{
					PreDeploymentStatus: apicommon.StateSucceeded,
					PreDeploymentTaskStatus: []v1alpha2.TaskStatus{
						{
							TaskDefinitionName: "task-def",
							Status:             apicommon.StateProgressing,
							TaskName:           "pre-task-def-",
						},
					},
				},
			},
			taskObj: v1alpha2.KeptnTask{
				ObjectMeta: v1.ObjectMeta{
					Namespace: "namespace",
					Name:      "pre-task-def-",
				},
				Status: v1alpha2.KeptnTaskStatus{
					Status: apicommon.StateSucceeded,
				},
			},
			createAttr: TaskCreateAttributes{
				SpanName:       "",
				TaskDefinition: "task-def",
				CheckType:      apicommon.PreDeploymentCheckType,
			},
			wantStatus: []v1alpha2.TaskStatus{
				{
					TaskDefinitionName: "task-def",
					Status:             apicommon.StateSucceeded,
					TaskName:           "pre-task-def-",
				},
			},
			wantSummary:     apicommon.StatusSummary{Total: 1, Succeeded: 1},
			wantErr:         nil,
			getSpanCalls:    1,
			unbindSpanCalls: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := v1alpha2.AddToScheme(scheme.Scheme)
			require.Nil(t, err)
			spanHandlerMock := kltfake.ISpanHandlerMock{
				GetSpanFunc: func(ctx context.Context, tracer trace.Tracer, reconcileObject client.Object, phase string) (context.Context, trace.Span, error) {
					return context.TODO(), trace.SpanFromContext(context.TODO()), nil
				},
				UnbindSpanFunc: func(reconcileObject client.Object, phase string) error {
					return nil
				},
			}
			handler := TaskHandler{
				SpanHandler: &spanHandlerMock,
				Log:         ctrl.Log.WithName("controller"),
				Recorder:    record.NewFakeRecorder(100),
				Client:      fake.NewClientBuilder().WithObjects(&tt.taskObj).Build(),
				Tracer:      trace.NewNoopTracerProvider().Tracer("tracer"),
				Scheme:      scheme.Scheme,
			}
			status, summary, err := handler.ReconcileTasks(context.TODO(), context.TODO(), tt.object, tt.createAttr)
			if len(tt.wantStatus) == len(status) {
				for j, item := range status {
					require.Equal(t, tt.wantStatus[j].TaskDefinitionName, item.TaskDefinitionName)
					require.True(t, strings.Contains(item.TaskName, tt.wantStatus[j].TaskName))
					require.Equal(t, tt.wantStatus[j].Status, item.Status)
				}
			} else {
				t.Errorf("unexpected result, want %+v, got %+v", tt.wantStatus, status)
			}
			require.Equal(t, tt.wantSummary, summary)
			require.Equal(t, tt.wantErr, err)
			require.Equal(t, tt.getSpanCalls, len(spanHandlerMock.GetSpanCalls()))
			require.Equal(t, tt.unbindSpanCalls, len(spanHandlerMock.UnbindSpanCalls()))
		})
	}
}

func TestTaskHandler_createTask(t *testing.T) {
	tests := []struct {
		name       string
		object     client.Object
		createAttr TaskCreateAttributes
		wantName   string
		wantErr    error
	}{
		{
			name:       "cannot unwrap object",
			object:     &v1alpha2.KeptnEvaluation{},
			createAttr: TaskCreateAttributes{},
			wantName:   "",
			wantErr:    controllererrors.ErrCannotWrapToPhaseItem,
		},
		{
			name: "created task",
			object: &v1alpha2.KeptnAppVersion{
				ObjectMeta: v1.ObjectMeta{
					Namespace: "namespace",
				},
				Spec: v1alpha2.KeptnAppVersionSpec{
					KeptnAppSpec: v1alpha2.KeptnAppSpec{
						PreDeploymentTasks: []string{"task-def"},
					},
				},
			},
			createAttr: TaskCreateAttributes{
				SpanName:       "",
				CheckType:      apicommon.PreDeploymentCheckType,
				TaskDefinition: "task-def",
			},
			wantName: "pre-task-def-",
			wantErr:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := v1alpha2.AddToScheme(scheme.Scheme)
			require.Nil(t, err)
			handler := TaskHandler{
				SpanHandler: &kltfake.ISpanHandlerMock{},
				Log:         ctrl.Log.WithName("controller"),
				Recorder:    record.NewFakeRecorder(100),
				Client:      fake.NewClientBuilder().Build(),
				Tracer:      trace.NewNoopTracerProvider().Tracer("tracer"),
				Scheme:      scheme.Scheme,
			}
			name, err := handler.CreateKeptnTask(context.TODO(), "namespace", tt.object, tt.createAttr)
			require.True(t, strings.Contains(name, tt.wantName))
			require.Equal(t, tt.wantErr, err)
		})
	}
}
