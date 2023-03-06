// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package fake

import (
	klcv1alpha3 "github.com/keptn/lifecycle-toolkit/operator/apis/lifecycle/v1alpha3"
	apicommon "github.com/keptn/lifecycle-toolkit/operator/apis/lifecycle/v1alpha3/common"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"sync"
	"time"
)

// PhaseItemMock is a mock implementation of interfaces.PhaseItem.
//
// 	func TestSomethingThatUsesPhaseItem(t *testing.T) {
//
// 		// make and configure a mocked interfaces.PhaseItem
// 		mockedPhaseItem := &PhaseItemMock{
// 			CompleteFunc: func()  {
// 				panic("mock out the Complete method")
// 			},
// 			DeprecateRemainingPhasesFunc: func(phase apicommon.KeptnPhaseType)  {
// 				panic("mock out the DeprecateRemainingPhases method")
// 			},
// 			GenerateEvaluationFunc: func(evaluationDefinition string, checkType apicommon.CheckType) klcv1alpha3.KeptnEvaluation {
// 				panic("mock out the GenerateEvaluation method")
// 			},
// 			GenerateTaskFunc: func(taskDefinition string, checkType apicommon.CheckType) klcv1alpha3.KeptnTask {
// 				panic("mock out the GenerateTask method")
// 			},
// 			GetAppNameFunc: func() string {
// 				panic("mock out the GetAppName method")
// 			},
// 			GetCurrentPhaseFunc: func() string {
// 				panic("mock out the GetCurrentPhase method")
// 			},
// 			GetEndTimeFunc: func() time.Time {
// 				panic("mock out the GetEndTime method")
// 			},
// 			GetNamespaceFunc: func() string {
// 				panic("mock out the GetNamespace method")
// 			},
// 			GetParentNameFunc: func() string {
// 				panic("mock out the GetParentName method")
// 			},
// 			GetPostDeploymentEvaluationTaskStatusFunc: func() []klcv1alpha3.ItemStatus {
// 				panic("mock out the GetPostDeploymentEvaluationTaskStatus method")
// 			},
// 			GetPostDeploymentEvaluationsFunc: func() []string {
// 				panic("mock out the GetPostDeploymentEvaluations method")
// 			},
// 			GetPostDeploymentTaskStatusFunc: func() []klcv1alpha3.ItemStatus {
// 				panic("mock out the GetPostDeploymentTaskStatus method")
// 			},
// 			GetPostDeploymentTasksFunc: func() []string {
// 				panic("mock out the GetPostDeploymentTasks method")
// 			},
// 			GetPreDeploymentEvaluationTaskStatusFunc: func() []klcv1alpha3.ItemStatus {
// 				panic("mock out the GetPreDeploymentEvaluationTaskStatus method")
// 			},
// 			GetPreDeploymentEvaluationsFunc: func() []string {
// 				panic("mock out the GetPreDeploymentEvaluations method")
// 			},
// 			GetPreDeploymentTaskStatusFunc: func() []klcv1alpha3.ItemStatus {
// 				panic("mock out the GetPreDeploymentTaskStatus method")
// 			},
// 			GetPreDeploymentTasksFunc: func() []string {
// 				panic("mock out the GetPreDeploymentTasks method")
// 			},
// 			GetPreviousVersionFunc: func() string {
// 				panic("mock out the GetPreviousVersion method")
// 			},
// 			GetSpanAttributesFunc: func() []attribute.KeyValue {
// 				panic("mock out the GetSpanAttributes method")
// 			},
// 			GetStartTimeFunc: func() time.Time {
// 				panic("mock out the GetStartTime method")
// 			},
// 			GetStateFunc: func() apicommon.KeptnState {
// 				panic("mock out the GetState method")
// 			},
// 			GetVersionFunc: func() string {
// 				panic("mock out the GetVersion method")
// 			},
// 			IsEndTimeSetFunc: func() bool {
// 				panic("mock out the IsEndTimeSet method")
// 			},
// 			SetCurrentPhaseFunc: func(s string)  {
// 				panic("mock out the SetCurrentPhase method")
// 			},
// 			SetSpanAttributesFunc: func(span trace.Span)  {
// 				panic("mock out the SetSpanAttributes method")
// 			},
// 			SetStateFunc: func(keptnState apicommon.KeptnState)  {
// 				panic("mock out the SetState method")
// 			},
// 		}
//
// 		// use mockedPhaseItem in code that requires interfaces.PhaseItem
// 		// and then make assertions.
//
// 	}
type PhaseItemMock struct {
	// CompleteFunc mocks the Complete method.
	CompleteFunc func()

	// DeprecateRemainingPhasesFunc mocks the DeprecateRemainingPhases method.
	DeprecateRemainingPhasesFunc func(phase apicommon.KeptnPhaseType)

	// GenerateEvaluationFunc mocks the GenerateEvaluation method.
	GenerateEvaluationFunc func(evaluationDefinition string, checkType apicommon.CheckType) klcv1alpha3.KeptnEvaluation

	// GenerateTaskFunc mocks the GenerateTask method.
	GenerateTaskFunc func(taskDefinition string, checkType apicommon.CheckType) klcv1alpha3.KeptnTask

	// GetAppNameFunc mocks the GetAppName method.
	GetAppNameFunc func() string

	// GetCurrentPhaseFunc mocks the GetCurrentPhase method.
	GetCurrentPhaseFunc func() string

	// GetEndTimeFunc mocks the GetEndTime method.
	GetEndTimeFunc func() time.Time

	// GetNamespaceFunc mocks the GetNamespace method.
	GetNamespaceFunc func() string

	// GetParentNameFunc mocks the GetParentName method.
	GetParentNameFunc func() string

	// GetPostDeploymentEvaluationTaskStatusFunc mocks the GetPostDeploymentEvaluationTaskStatus method.
	GetPostDeploymentEvaluationTaskStatusFunc func() []klcv1alpha3.ItemStatus

	// GetPostDeploymentEvaluationsFunc mocks the GetPostDeploymentEvaluations method.
	GetPostDeploymentEvaluationsFunc func() []string

	// GetPostDeploymentTaskStatusFunc mocks the GetPostDeploymentTaskStatus method.
	GetPostDeploymentTaskStatusFunc func() []klcv1alpha3.ItemStatus

	// GetPostDeploymentTasksFunc mocks the GetPostDeploymentTasks method.
	GetPostDeploymentTasksFunc func() []string

	// GetPreDeploymentEvaluationTaskStatusFunc mocks the GetPreDeploymentEvaluationTaskStatus method.
	GetPreDeploymentEvaluationTaskStatusFunc func() []klcv1alpha3.ItemStatus

	// GetPreDeploymentEvaluationsFunc mocks the GetPreDeploymentEvaluations method.
	GetPreDeploymentEvaluationsFunc func() []string

	// GetPreDeploymentTaskStatusFunc mocks the GetPreDeploymentTaskStatus method.
	GetPreDeploymentTaskStatusFunc func() []klcv1alpha3.ItemStatus

	// GetPreDeploymentTasksFunc mocks the GetPreDeploymentTasks method.
	GetPreDeploymentTasksFunc func() []string

	// GetPreviousVersionFunc mocks the GetPreviousVersion method.
	GetPreviousVersionFunc func() string

	// GetSpanAttributesFunc mocks the GetSpanAttributes method.
	GetSpanAttributesFunc func() []attribute.KeyValue

	// GetStartTimeFunc mocks the GetStartTime method.
	GetStartTimeFunc func() time.Time

	// GetStateFunc mocks the GetState method.
	GetStateFunc func() apicommon.KeptnState

	// GetVersionFunc mocks the GetVersion method.
	GetVersionFunc func() string

	// IsEndTimeSetFunc mocks the IsEndTimeSet method.
	IsEndTimeSetFunc func() bool

	// SetCurrentPhaseFunc mocks the SetCurrentPhase method.
	SetCurrentPhaseFunc func(s string)

	// SetSpanAttributesFunc mocks the SetSpanAttributes method.
	SetSpanAttributesFunc func(span trace.Span)

	// SetStateFunc mocks the SetState method.
	SetStateFunc func(keptnState apicommon.KeptnState)

	// calls tracks calls to the methods.
	calls struct {
		// Complete holds details about calls to the Complete method.
		Complete []struct {
		}
		// DeprecateRemainingPhases holds details about calls to the DeprecateRemainingPhases method.
		DeprecateRemainingPhases []struct {
			// Phase is the phase argument value.
			Phase apicommon.KeptnPhaseType
		}
		// GenerateEvaluation holds details about calls to the GenerateEvaluation method.
		GenerateEvaluation []struct {
			// EvaluationDefinition is the evaluationDefinition argument value.
			EvaluationDefinition string
			// CheckType is the checkType argument value.
			CheckType apicommon.CheckType
		}
		// GenerateTask holds details about calls to the GenerateTask method.
		GenerateTask []struct {
			// TaskDefinition is the taskDefinition argument value.
			TaskDefinition string
			// CheckType is the checkType argument value.
			CheckType apicommon.CheckType
		}
		// GetAppName holds details about calls to the GetAppName method.
		GetAppName []struct {
		}
		// GetCurrentPhase holds details about calls to the GetCurrentPhase method.
		GetCurrentPhase []struct {
		}
		// GetEndTime holds details about calls to the GetEndTime method.
		GetEndTime []struct {
		}
		// GetNamespace holds details about calls to the GetNamespace method.
		GetNamespace []struct {
		}
		// GetParentName holds details about calls to the GetParentName method.
		GetParentName []struct {
		}
		// GetPostDeploymentEvaluationTaskStatus holds details about calls to the GetPostDeploymentEvaluationTaskStatus method.
		GetPostDeploymentEvaluationTaskStatus []struct {
		}
		// GetPostDeploymentEvaluations holds details about calls to the GetPostDeploymentEvaluations method.
		GetPostDeploymentEvaluations []struct {
		}
		// GetPostDeploymentTaskStatus holds details about calls to the GetPostDeploymentTaskStatus method.
		GetPostDeploymentTaskStatus []struct {
		}
		// GetPostDeploymentTasks holds details about calls to the GetPostDeploymentTasks method.
		GetPostDeploymentTasks []struct {
		}
		// GetPreDeploymentEvaluationTaskStatus holds details about calls to the GetPreDeploymentEvaluationTaskStatus method.
		GetPreDeploymentEvaluationTaskStatus []struct {
		}
		// GetPreDeploymentEvaluations holds details about calls to the GetPreDeploymentEvaluations method.
		GetPreDeploymentEvaluations []struct {
		}
		// GetPreDeploymentTaskStatus holds details about calls to the GetPreDeploymentTaskStatus method.
		GetPreDeploymentTaskStatus []struct {
		}
		// GetPreDeploymentTasks holds details about calls to the GetPreDeploymentTasks method.
		GetPreDeploymentTasks []struct {
		}
		// GetPreviousVersion holds details about calls to the GetPreviousVersion method.
		GetPreviousVersion []struct {
		}
		// GetSpanAttributes holds details about calls to the GetSpanAttributes method.
		GetSpanAttributes []struct {
		}
		// GetStartTime holds details about calls to the GetStartTime method.
		GetStartTime []struct {
		}
		// GetState holds details about calls to the GetState method.
		GetState []struct {
		}
		// GetVersion holds details about calls to the GetVersion method.
		GetVersion []struct {
		}
		// IsEndTimeSet holds details about calls to the IsEndTimeSet method.
		IsEndTimeSet []struct {
		}
		// SetCurrentPhase holds details about calls to the SetCurrentPhase method.
		SetCurrentPhase []struct {
			// S is the s argument value.
			S string
		}
		// SetSpanAttributes holds details about calls to the SetSpanAttributes method.
		SetSpanAttributes []struct {
			// Span is the span argument value.
			Span trace.Span
		}
		// SetState holds details about calls to the SetState method.
		SetState []struct {
			// KeptnState is the keptnState argument value.
			KeptnState apicommon.KeptnState
		}
	}
	lockComplete                              sync.RWMutex
	lockDeprecateRemainingPhases              sync.RWMutex
	lockGenerateEvaluation                    sync.RWMutex
	lockGenerateTask                          sync.RWMutex
	lockGetAppName                            sync.RWMutex
	lockGetCurrentPhase                       sync.RWMutex
	lockGetEndTime                            sync.RWMutex
	lockGetNamespace                          sync.RWMutex
	lockGetParentName                         sync.RWMutex
	lockGetPostDeploymentEvaluationTaskStatus sync.RWMutex
	lockGetPostDeploymentEvaluations          sync.RWMutex
	lockGetPostDeploymentTaskStatus           sync.RWMutex
	lockGetPostDeploymentTasks                sync.RWMutex
	lockGetPreDeploymentEvaluationTaskStatus  sync.RWMutex
	lockGetPreDeploymentEvaluations           sync.RWMutex
	lockGetPreDeploymentTaskStatus            sync.RWMutex
	lockGetPreDeploymentTasks                 sync.RWMutex
	lockGetPreviousVersion                    sync.RWMutex
	lockGetSpanAttributes                     sync.RWMutex
	lockGetStartTime                          sync.RWMutex
	lockGetState                              sync.RWMutex
	lockGetVersion                            sync.RWMutex
	lockIsEndTimeSet                          sync.RWMutex
	lockSetCurrentPhase                       sync.RWMutex
	lockSetSpanAttributes                     sync.RWMutex
	lockSetState                              sync.RWMutex
}

// Complete calls CompleteFunc.
func (mock *PhaseItemMock) Complete() {
	if mock.CompleteFunc == nil {
		panic("PhaseItemMock.CompleteFunc: method is nil but PhaseItem.Complete was just called")
	}
	callInfo := struct {
	}{}
	mock.lockComplete.Lock()
	mock.calls.Complete = append(mock.calls.Complete, callInfo)
	mock.lockComplete.Unlock()
	mock.CompleteFunc()
}

// CompleteCalls gets all the calls that were made to Complete.
// Check the length with:
//     len(mockedPhaseItem.CompleteCalls())
func (mock *PhaseItemMock) CompleteCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockComplete.RLock()
	calls = mock.calls.Complete
	mock.lockComplete.RUnlock()
	return calls
}

// DeprecateRemainingPhases calls DeprecateRemainingPhasesFunc.
func (mock *PhaseItemMock) DeprecateRemainingPhases(phase apicommon.KeptnPhaseType) {
	if mock.DeprecateRemainingPhasesFunc == nil {
		panic("PhaseItemMock.DeprecateRemainingPhasesFunc: method is nil but PhaseItem.DeprecateRemainingPhases was just called")
	}
	callInfo := struct {
		Phase apicommon.KeptnPhaseType
	}{
		Phase: phase,
	}
	mock.lockDeprecateRemainingPhases.Lock()
	mock.calls.DeprecateRemainingPhases = append(mock.calls.DeprecateRemainingPhases, callInfo)
	mock.lockDeprecateRemainingPhases.Unlock()
	mock.DeprecateRemainingPhasesFunc(phase)
}

// DeprecateRemainingPhasesCalls gets all the calls that were made to DeprecateRemainingPhases.
// Check the length with:
//     len(mockedPhaseItem.DeprecateRemainingPhasesCalls())
func (mock *PhaseItemMock) DeprecateRemainingPhasesCalls() []struct {
	Phase apicommon.KeptnPhaseType
} {
	var calls []struct {
		Phase apicommon.KeptnPhaseType
	}
	mock.lockDeprecateRemainingPhases.RLock()
	calls = mock.calls.DeprecateRemainingPhases
	mock.lockDeprecateRemainingPhases.RUnlock()
	return calls
}

// GenerateEvaluation calls GenerateEvaluationFunc.
func (mock *PhaseItemMock) GenerateEvaluation(evaluationDefinition string, checkType apicommon.CheckType) klcv1alpha3.KeptnEvaluation {
	if mock.GenerateEvaluationFunc == nil {
		panic("PhaseItemMock.GenerateEvaluationFunc: method is nil but PhaseItem.GenerateEvaluation was just called")
	}
	callInfo := struct {
		EvaluationDefinition string
		CheckType            apicommon.CheckType
	}{
		EvaluationDefinition: evaluationDefinition,
		CheckType:            checkType,
	}
	mock.lockGenerateEvaluation.Lock()
	mock.calls.GenerateEvaluation = append(mock.calls.GenerateEvaluation, callInfo)
	mock.lockGenerateEvaluation.Unlock()
	return mock.GenerateEvaluationFunc(evaluationDefinition, checkType)
}

// GenerateEvaluationCalls gets all the calls that were made to GenerateEvaluation.
// Check the length with:
//     len(mockedPhaseItem.GenerateEvaluationCalls())
func (mock *PhaseItemMock) GenerateEvaluationCalls() []struct {
	EvaluationDefinition string
	CheckType            apicommon.CheckType
} {
	var calls []struct {
		EvaluationDefinition string
		CheckType            apicommon.CheckType
	}
	mock.lockGenerateEvaluation.RLock()
	calls = mock.calls.GenerateEvaluation
	mock.lockGenerateEvaluation.RUnlock()
	return calls
}

// GenerateTask calls GenerateTaskFunc.
func (mock *PhaseItemMock) GenerateTask(taskDefinition string, checkType apicommon.CheckType) klcv1alpha3.KeptnTask {
	if mock.GenerateTaskFunc == nil {
		panic("PhaseItemMock.GenerateTaskFunc: method is nil but PhaseItem.GenerateTask was just called")
	}
	callInfo := struct {
		TaskDefinition string
		CheckType      apicommon.CheckType
	}{
		TaskDefinition: taskDefinition,
		CheckType:      checkType,
	}
	mock.lockGenerateTask.Lock()
	mock.calls.GenerateTask = append(mock.calls.GenerateTask, callInfo)
	mock.lockGenerateTask.Unlock()
	return mock.GenerateTaskFunc(taskDefinition, checkType)
}

// GenerateTaskCalls gets all the calls that were made to GenerateTask.
// Check the length with:
//     len(mockedPhaseItem.GenerateTaskCalls())
func (mock *PhaseItemMock) GenerateTaskCalls() []struct {
	TaskDefinition string
	CheckType      apicommon.CheckType
} {
	var calls []struct {
		TaskDefinition string
		CheckType      apicommon.CheckType
	}
	mock.lockGenerateTask.RLock()
	calls = mock.calls.GenerateTask
	mock.lockGenerateTask.RUnlock()
	return calls
}

// GetAppName calls GetAppNameFunc.
func (mock *PhaseItemMock) GetAppName() string {
	if mock.GetAppNameFunc == nil {
		panic("PhaseItemMock.GetAppNameFunc: method is nil but PhaseItem.GetAppName was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetAppName.Lock()
	mock.calls.GetAppName = append(mock.calls.GetAppName, callInfo)
	mock.lockGetAppName.Unlock()
	return mock.GetAppNameFunc()
}

// GetAppNameCalls gets all the calls that were made to GetAppName.
// Check the length with:
//     len(mockedPhaseItem.GetAppNameCalls())
func (mock *PhaseItemMock) GetAppNameCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetAppName.RLock()
	calls = mock.calls.GetAppName
	mock.lockGetAppName.RUnlock()
	return calls
}

// GetCurrentPhase calls GetCurrentPhaseFunc.
func (mock *PhaseItemMock) GetCurrentPhase() string {
	if mock.GetCurrentPhaseFunc == nil {
		panic("PhaseItemMock.GetCurrentPhaseFunc: method is nil but PhaseItem.GetCurrentPhase was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetCurrentPhase.Lock()
	mock.calls.GetCurrentPhase = append(mock.calls.GetCurrentPhase, callInfo)
	mock.lockGetCurrentPhase.Unlock()
	return mock.GetCurrentPhaseFunc()
}

// GetCurrentPhaseCalls gets all the calls that were made to GetCurrentPhase.
// Check the length with:
//     len(mockedPhaseItem.GetCurrentPhaseCalls())
func (mock *PhaseItemMock) GetCurrentPhaseCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetCurrentPhase.RLock()
	calls = mock.calls.GetCurrentPhase
	mock.lockGetCurrentPhase.RUnlock()
	return calls
}

// GetEndTime calls GetEndTimeFunc.
func (mock *PhaseItemMock) GetEndTime() time.Time {
	if mock.GetEndTimeFunc == nil {
		panic("PhaseItemMock.GetEndTimeFunc: method is nil but PhaseItem.GetEndTime was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetEndTime.Lock()
	mock.calls.GetEndTime = append(mock.calls.GetEndTime, callInfo)
	mock.lockGetEndTime.Unlock()
	return mock.GetEndTimeFunc()
}

// GetEndTimeCalls gets all the calls that were made to GetEndTime.
// Check the length with:
//     len(mockedPhaseItem.GetEndTimeCalls())
func (mock *PhaseItemMock) GetEndTimeCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetEndTime.RLock()
	calls = mock.calls.GetEndTime
	mock.lockGetEndTime.RUnlock()
	return calls
}

// GetNamespace calls GetNamespaceFunc.
func (mock *PhaseItemMock) GetNamespace() string {
	if mock.GetNamespaceFunc == nil {
		panic("PhaseItemMock.GetNamespaceFunc: method is nil but PhaseItem.GetNamespace was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetNamespace.Lock()
	mock.calls.GetNamespace = append(mock.calls.GetNamespace, callInfo)
	mock.lockGetNamespace.Unlock()
	return mock.GetNamespaceFunc()
}

// GetNamespaceCalls gets all the calls that were made to GetNamespace.
// Check the length with:
//     len(mockedPhaseItem.GetNamespaceCalls())
func (mock *PhaseItemMock) GetNamespaceCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetNamespace.RLock()
	calls = mock.calls.GetNamespace
	mock.lockGetNamespace.RUnlock()
	return calls
}

// GetParentName calls GetParentNameFunc.
func (mock *PhaseItemMock) GetParentName() string {
	if mock.GetParentNameFunc == nil {
		panic("PhaseItemMock.GetParentNameFunc: method is nil but PhaseItem.GetParentName was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetParentName.Lock()
	mock.calls.GetParentName = append(mock.calls.GetParentName, callInfo)
	mock.lockGetParentName.Unlock()
	return mock.GetParentNameFunc()
}

// GetParentNameCalls gets all the calls that were made to GetParentName.
// Check the length with:
//     len(mockedPhaseItem.GetParentNameCalls())
func (mock *PhaseItemMock) GetParentNameCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetParentName.RLock()
	calls = mock.calls.GetParentName
	mock.lockGetParentName.RUnlock()
	return calls
}

// GetPostDeploymentEvaluationTaskStatus calls GetPostDeploymentEvaluationTaskStatusFunc.
func (mock *PhaseItemMock) GetPostDeploymentEvaluationTaskStatus() []klcv1alpha3.ItemStatus {
	if mock.GetPostDeploymentEvaluationTaskStatusFunc == nil {
		panic("PhaseItemMock.GetPostDeploymentEvaluationTaskStatusFunc: method is nil but PhaseItem.GetPostDeploymentEvaluationTaskStatus was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetPostDeploymentEvaluationTaskStatus.Lock()
	mock.calls.GetPostDeploymentEvaluationTaskStatus = append(mock.calls.GetPostDeploymentEvaluationTaskStatus, callInfo)
	mock.lockGetPostDeploymentEvaluationTaskStatus.Unlock()
	return mock.GetPostDeploymentEvaluationTaskStatusFunc()
}

// GetPostDeploymentEvaluationTaskStatusCalls gets all the calls that were made to GetPostDeploymentEvaluationTaskStatus.
// Check the length with:
//     len(mockedPhaseItem.GetPostDeploymentEvaluationTaskStatusCalls())
func (mock *PhaseItemMock) GetPostDeploymentEvaluationTaskStatusCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetPostDeploymentEvaluationTaskStatus.RLock()
	calls = mock.calls.GetPostDeploymentEvaluationTaskStatus
	mock.lockGetPostDeploymentEvaluationTaskStatus.RUnlock()
	return calls
}

// GetPostDeploymentEvaluations calls GetPostDeploymentEvaluationsFunc.
func (mock *PhaseItemMock) GetPostDeploymentEvaluations() []string {
	if mock.GetPostDeploymentEvaluationsFunc == nil {
		panic("PhaseItemMock.GetPostDeploymentEvaluationsFunc: method is nil but PhaseItem.GetPostDeploymentEvaluations was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetPostDeploymentEvaluations.Lock()
	mock.calls.GetPostDeploymentEvaluations = append(mock.calls.GetPostDeploymentEvaluations, callInfo)
	mock.lockGetPostDeploymentEvaluations.Unlock()
	return mock.GetPostDeploymentEvaluationsFunc()
}

// GetPostDeploymentEvaluationsCalls gets all the calls that were made to GetPostDeploymentEvaluations.
// Check the length with:
//     len(mockedPhaseItem.GetPostDeploymentEvaluationsCalls())
func (mock *PhaseItemMock) GetPostDeploymentEvaluationsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetPostDeploymentEvaluations.RLock()
	calls = mock.calls.GetPostDeploymentEvaluations
	mock.lockGetPostDeploymentEvaluations.RUnlock()
	return calls
}

// GetPostDeploymentTaskStatus calls GetPostDeploymentTaskStatusFunc.
func (mock *PhaseItemMock) GetPostDeploymentTaskStatus() []klcv1alpha3.ItemStatus {
	if mock.GetPostDeploymentTaskStatusFunc == nil {
		panic("PhaseItemMock.GetPostDeploymentTaskStatusFunc: method is nil but PhaseItem.GetPostDeploymentTaskStatus was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetPostDeploymentTaskStatus.Lock()
	mock.calls.GetPostDeploymentTaskStatus = append(mock.calls.GetPostDeploymentTaskStatus, callInfo)
	mock.lockGetPostDeploymentTaskStatus.Unlock()
	return mock.GetPostDeploymentTaskStatusFunc()
}

// GetPostDeploymentTaskStatusCalls gets all the calls that were made to GetPostDeploymentTaskStatus.
// Check the length with:
//     len(mockedPhaseItem.GetPostDeploymentTaskStatusCalls())
func (mock *PhaseItemMock) GetPostDeploymentTaskStatusCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetPostDeploymentTaskStatus.RLock()
	calls = mock.calls.GetPostDeploymentTaskStatus
	mock.lockGetPostDeploymentTaskStatus.RUnlock()
	return calls
}

// GetPostDeploymentTasks calls GetPostDeploymentTasksFunc.
func (mock *PhaseItemMock) GetPostDeploymentTasks() []string {
	if mock.GetPostDeploymentTasksFunc == nil {
		panic("PhaseItemMock.GetPostDeploymentTasksFunc: method is nil but PhaseItem.GetPostDeploymentTasks was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetPostDeploymentTasks.Lock()
	mock.calls.GetPostDeploymentTasks = append(mock.calls.GetPostDeploymentTasks, callInfo)
	mock.lockGetPostDeploymentTasks.Unlock()
	return mock.GetPostDeploymentTasksFunc()
}

// GetPostDeploymentTasksCalls gets all the calls that were made to GetPostDeploymentTasks.
// Check the length with:
//     len(mockedPhaseItem.GetPostDeploymentTasksCalls())
func (mock *PhaseItemMock) GetPostDeploymentTasksCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetPostDeploymentTasks.RLock()
	calls = mock.calls.GetPostDeploymentTasks
	mock.lockGetPostDeploymentTasks.RUnlock()
	return calls
}

// GetPreDeploymentEvaluationTaskStatus calls GetPreDeploymentEvaluationTaskStatusFunc.
func (mock *PhaseItemMock) GetPreDeploymentEvaluationTaskStatus() []klcv1alpha3.ItemStatus {
	if mock.GetPreDeploymentEvaluationTaskStatusFunc == nil {
		panic("PhaseItemMock.GetPreDeploymentEvaluationTaskStatusFunc: method is nil but PhaseItem.GetPreDeploymentEvaluationTaskStatus was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetPreDeploymentEvaluationTaskStatus.Lock()
	mock.calls.GetPreDeploymentEvaluationTaskStatus = append(mock.calls.GetPreDeploymentEvaluationTaskStatus, callInfo)
	mock.lockGetPreDeploymentEvaluationTaskStatus.Unlock()
	return mock.GetPreDeploymentEvaluationTaskStatusFunc()
}

// GetPreDeploymentEvaluationTaskStatusCalls gets all the calls that were made to GetPreDeploymentEvaluationTaskStatus.
// Check the length with:
//     len(mockedPhaseItem.GetPreDeploymentEvaluationTaskStatusCalls())
func (mock *PhaseItemMock) GetPreDeploymentEvaluationTaskStatusCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetPreDeploymentEvaluationTaskStatus.RLock()
	calls = mock.calls.GetPreDeploymentEvaluationTaskStatus
	mock.lockGetPreDeploymentEvaluationTaskStatus.RUnlock()
	return calls
}

// GetPreDeploymentEvaluations calls GetPreDeploymentEvaluationsFunc.
func (mock *PhaseItemMock) GetPreDeploymentEvaluations() []string {
	if mock.GetPreDeploymentEvaluationsFunc == nil {
		panic("PhaseItemMock.GetPreDeploymentEvaluationsFunc: method is nil but PhaseItem.GetPreDeploymentEvaluations was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetPreDeploymentEvaluations.Lock()
	mock.calls.GetPreDeploymentEvaluations = append(mock.calls.GetPreDeploymentEvaluations, callInfo)
	mock.lockGetPreDeploymentEvaluations.Unlock()
	return mock.GetPreDeploymentEvaluationsFunc()
}

// GetPreDeploymentEvaluationsCalls gets all the calls that were made to GetPreDeploymentEvaluations.
// Check the length with:
//     len(mockedPhaseItem.GetPreDeploymentEvaluationsCalls())
func (mock *PhaseItemMock) GetPreDeploymentEvaluationsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetPreDeploymentEvaluations.RLock()
	calls = mock.calls.GetPreDeploymentEvaluations
	mock.lockGetPreDeploymentEvaluations.RUnlock()
	return calls
}

// GetPreDeploymentTaskStatus calls GetPreDeploymentTaskStatusFunc.
func (mock *PhaseItemMock) GetPreDeploymentTaskStatus() []klcv1alpha3.ItemStatus {
	if mock.GetPreDeploymentTaskStatusFunc == nil {
		panic("PhaseItemMock.GetPreDeploymentTaskStatusFunc: method is nil but PhaseItem.GetPreDeploymentTaskStatus was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetPreDeploymentTaskStatus.Lock()
	mock.calls.GetPreDeploymentTaskStatus = append(mock.calls.GetPreDeploymentTaskStatus, callInfo)
	mock.lockGetPreDeploymentTaskStatus.Unlock()
	return mock.GetPreDeploymentTaskStatusFunc()
}

// GetPreDeploymentTaskStatusCalls gets all the calls that were made to GetPreDeploymentTaskStatus.
// Check the length with:
//     len(mockedPhaseItem.GetPreDeploymentTaskStatusCalls())
func (mock *PhaseItemMock) GetPreDeploymentTaskStatusCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetPreDeploymentTaskStatus.RLock()
	calls = mock.calls.GetPreDeploymentTaskStatus
	mock.lockGetPreDeploymentTaskStatus.RUnlock()
	return calls
}

// GetPreDeploymentTasks calls GetPreDeploymentTasksFunc.
func (mock *PhaseItemMock) GetPreDeploymentTasks() []string {
	if mock.GetPreDeploymentTasksFunc == nil {
		panic("PhaseItemMock.GetPreDeploymentTasksFunc: method is nil but PhaseItem.GetPreDeploymentTasks was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetPreDeploymentTasks.Lock()
	mock.calls.GetPreDeploymentTasks = append(mock.calls.GetPreDeploymentTasks, callInfo)
	mock.lockGetPreDeploymentTasks.Unlock()
	return mock.GetPreDeploymentTasksFunc()
}

// GetPreDeploymentTasksCalls gets all the calls that were made to GetPreDeploymentTasks.
// Check the length with:
//     len(mockedPhaseItem.GetPreDeploymentTasksCalls())
func (mock *PhaseItemMock) GetPreDeploymentTasksCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetPreDeploymentTasks.RLock()
	calls = mock.calls.GetPreDeploymentTasks
	mock.lockGetPreDeploymentTasks.RUnlock()
	return calls
}

// GetPreviousVersion calls GetPreviousVersionFunc.
func (mock *PhaseItemMock) GetPreviousVersion() string {
	if mock.GetPreviousVersionFunc == nil {
		panic("PhaseItemMock.GetPreviousVersionFunc: method is nil but PhaseItem.GetPreviousVersion was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetPreviousVersion.Lock()
	mock.calls.GetPreviousVersion = append(mock.calls.GetPreviousVersion, callInfo)
	mock.lockGetPreviousVersion.Unlock()
	return mock.GetPreviousVersionFunc()
}

// GetPreviousVersionCalls gets all the calls that were made to GetPreviousVersion.
// Check the length with:
//     len(mockedPhaseItem.GetPreviousVersionCalls())
func (mock *PhaseItemMock) GetPreviousVersionCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetPreviousVersion.RLock()
	calls = mock.calls.GetPreviousVersion
	mock.lockGetPreviousVersion.RUnlock()
	return calls
}

// GetSpanAttributes calls GetSpanAttributesFunc.
func (mock *PhaseItemMock) GetSpanAttributes() []attribute.KeyValue {
	if mock.GetSpanAttributesFunc == nil {
		panic("PhaseItemMock.GetSpanAttributesFunc: method is nil but PhaseItem.GetSpanAttributes was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetSpanAttributes.Lock()
	mock.calls.GetSpanAttributes = append(mock.calls.GetSpanAttributes, callInfo)
	mock.lockGetSpanAttributes.Unlock()
	return mock.GetSpanAttributesFunc()
}

// GetSpanAttributesCalls gets all the calls that were made to GetSpanAttributes.
// Check the length with:
//     len(mockedPhaseItem.GetSpanAttributesCalls())
func (mock *PhaseItemMock) GetSpanAttributesCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetSpanAttributes.RLock()
	calls = mock.calls.GetSpanAttributes
	mock.lockGetSpanAttributes.RUnlock()
	return calls
}

// GetStartTime calls GetStartTimeFunc.
func (mock *PhaseItemMock) GetStartTime() time.Time {
	if mock.GetStartTimeFunc == nil {
		panic("PhaseItemMock.GetStartTimeFunc: method is nil but PhaseItem.GetStartTime was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetStartTime.Lock()
	mock.calls.GetStartTime = append(mock.calls.GetStartTime, callInfo)
	mock.lockGetStartTime.Unlock()
	return mock.GetStartTimeFunc()
}

// GetStartTimeCalls gets all the calls that were made to GetStartTime.
// Check the length with:
//     len(mockedPhaseItem.GetStartTimeCalls())
func (mock *PhaseItemMock) GetStartTimeCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetStartTime.RLock()
	calls = mock.calls.GetStartTime
	mock.lockGetStartTime.RUnlock()
	return calls
}

// GetState calls GetStateFunc.
func (mock *PhaseItemMock) GetState() apicommon.KeptnState {
	if mock.GetStateFunc == nil {
		panic("PhaseItemMock.GetStateFunc: method is nil but PhaseItem.GetState was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetState.Lock()
	mock.calls.GetState = append(mock.calls.GetState, callInfo)
	mock.lockGetState.Unlock()
	return mock.GetStateFunc()
}

// GetStateCalls gets all the calls that were made to GetState.
// Check the length with:
//     len(mockedPhaseItem.GetStateCalls())
func (mock *PhaseItemMock) GetStateCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetState.RLock()
	calls = mock.calls.GetState
	mock.lockGetState.RUnlock()
	return calls
}

// GetVersion calls GetVersionFunc.
func (mock *PhaseItemMock) GetVersion() string {
	if mock.GetVersionFunc == nil {
		panic("PhaseItemMock.GetVersionFunc: method is nil but PhaseItem.GetVersion was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetVersion.Lock()
	mock.calls.GetVersion = append(mock.calls.GetVersion, callInfo)
	mock.lockGetVersion.Unlock()
	return mock.GetVersionFunc()
}

// GetVersionCalls gets all the calls that were made to GetVersion.
// Check the length with:
//     len(mockedPhaseItem.GetVersionCalls())
func (mock *PhaseItemMock) GetVersionCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetVersion.RLock()
	calls = mock.calls.GetVersion
	mock.lockGetVersion.RUnlock()
	return calls
}

// IsEndTimeSet calls IsEndTimeSetFunc.
func (mock *PhaseItemMock) IsEndTimeSet() bool {
	if mock.IsEndTimeSetFunc == nil {
		panic("PhaseItemMock.IsEndTimeSetFunc: method is nil but PhaseItem.IsEndTimeSet was just called")
	}
	callInfo := struct {
	}{}
	mock.lockIsEndTimeSet.Lock()
	mock.calls.IsEndTimeSet = append(mock.calls.IsEndTimeSet, callInfo)
	mock.lockIsEndTimeSet.Unlock()
	return mock.IsEndTimeSetFunc()
}

// IsEndTimeSetCalls gets all the calls that were made to IsEndTimeSet.
// Check the length with:
//     len(mockedPhaseItem.IsEndTimeSetCalls())
func (mock *PhaseItemMock) IsEndTimeSetCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockIsEndTimeSet.RLock()
	calls = mock.calls.IsEndTimeSet
	mock.lockIsEndTimeSet.RUnlock()
	return calls
}

// SetCurrentPhase calls SetCurrentPhaseFunc.
func (mock *PhaseItemMock) SetCurrentPhase(s string) {
	if mock.SetCurrentPhaseFunc == nil {
		panic("PhaseItemMock.SetCurrentPhaseFunc: method is nil but PhaseItem.SetCurrentPhase was just called")
	}
	callInfo := struct {
		S string
	}{
		S: s,
	}
	mock.lockSetCurrentPhase.Lock()
	mock.calls.SetCurrentPhase = append(mock.calls.SetCurrentPhase, callInfo)
	mock.lockSetCurrentPhase.Unlock()
	mock.SetCurrentPhaseFunc(s)
}

// SetCurrentPhaseCalls gets all the calls that were made to SetCurrentPhase.
// Check the length with:
//     len(mockedPhaseItem.SetCurrentPhaseCalls())
func (mock *PhaseItemMock) SetCurrentPhaseCalls() []struct {
	S string
} {
	var calls []struct {
		S string
	}
	mock.lockSetCurrentPhase.RLock()
	calls = mock.calls.SetCurrentPhase
	mock.lockSetCurrentPhase.RUnlock()
	return calls
}

// SetSpanAttributes calls SetSpanAttributesFunc.
func (mock *PhaseItemMock) SetSpanAttributes(span trace.Span) {
	if mock.SetSpanAttributesFunc == nil {
		panic("PhaseItemMock.SetSpanAttributesFunc: method is nil but PhaseItem.SetSpanAttributes was just called")
	}
	callInfo := struct {
		Span trace.Span
	}{
		Span: span,
	}
	mock.lockSetSpanAttributes.Lock()
	mock.calls.SetSpanAttributes = append(mock.calls.SetSpanAttributes, callInfo)
	mock.lockSetSpanAttributes.Unlock()
	mock.SetSpanAttributesFunc(span)
}

// SetSpanAttributesCalls gets all the calls that were made to SetSpanAttributes.
// Check the length with:
//     len(mockedPhaseItem.SetSpanAttributesCalls())
func (mock *PhaseItemMock) SetSpanAttributesCalls() []struct {
	Span trace.Span
} {
	var calls []struct {
		Span trace.Span
	}
	mock.lockSetSpanAttributes.RLock()
	calls = mock.calls.SetSpanAttributes
	mock.lockSetSpanAttributes.RUnlock()
	return calls
}

// SetState calls SetStateFunc.
func (mock *PhaseItemMock) SetState(keptnState apicommon.KeptnState) {
	if mock.SetStateFunc == nil {
		panic("PhaseItemMock.SetStateFunc: method is nil but PhaseItem.SetState was just called")
	}
	callInfo := struct {
		KeptnState apicommon.KeptnState
	}{
		KeptnState: keptnState,
	}
	mock.lockSetState.Lock()
	mock.calls.SetState = append(mock.calls.SetState, callInfo)
	mock.lockSetState.Unlock()
	mock.SetStateFunc(keptnState)
}

// SetStateCalls gets all the calls that were made to SetState.
// Check the length with:
//     len(mockedPhaseItem.SetStateCalls())
func (mock *PhaseItemMock) SetStateCalls() []struct {
	KeptnState apicommon.KeptnState
} {
	var calls []struct {
		KeptnState apicommon.KeptnState
	}
	mock.lockSetState.RLock()
	calls = mock.calls.SetState
	mock.lockSetState.RUnlock()
	return calls
}