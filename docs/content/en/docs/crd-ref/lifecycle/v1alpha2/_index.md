---
title: v1alpha2
description: Reference information for lifecycle.keptn.sh/v1alpha2
---
<!-- markdownlint-disable -->

## Packages
- [lifecycle.keptn.sh/v1alpha2](#lifecyclekeptnshv1alpha2)


## lifecycle.keptn.sh/v1alpha2

Package v1alpha2 contains API Schema definitions for the lifecycle v1alpha2 API group

### Resource Types
- [KeptnApp](#keptnapp)
- [KeptnAppList](#keptnapplist)
- [KeptnAppVersion](#keptnappversion)
- [KeptnAppVersionList](#keptnappversionlist)
- [KeptnEvaluation](#keptnevaluation)
- [KeptnEvaluationDefinition](#keptnevaluationdefinition)
- [KeptnEvaluationDefinitionList](#keptnevaluationdefinitionlist)
- [KeptnEvaluationList](#keptnevaluationlist)
- [KeptnEvaluationProvider](#keptnevaluationprovider)
- [KeptnEvaluationProviderList](#keptnevaluationproviderlist)
- [KeptnTask](#keptntask)
- [KeptnTaskDefinition](#keptntaskdefinition)
- [KeptnTaskDefinitionList](#keptntaskdefinitionlist)
- [KeptnTaskList](#keptntasklist)
- [KeptnWorkload](#keptnworkload)
- [KeptnWorkloadInstance](#keptnworkloadinstance)
- [KeptnWorkloadInstanceList](#keptnworkloadinstancelist)
- [KeptnWorkloadList](#keptnworkloadlist)



#### ConfigMapReference





_Appears in:_
- [FunctionSpec](#functionspec)

| Field | Description |
| --- | --- |
| `name` _string_ |  |




#### EvaluationStatusItem





_Appears in:_
- [KeptnEvaluationStatus](#keptnevaluationstatus)

| Field | Description |
| --- | --- |
| `value` _string_ |  |
| `message` _string_ |  |


#### FunctionReference



FunctionReference Execute another `KeptnTaskDefinition` that has been defined.
Populate this field with the value of the `name` field
for the `KeptnTaskDefinition` to be called.
This is commonly used to call a general function
that is used in multiple place with different parameters.
An example is:

 ```yaml
 spec:
   function:
     functionRef:
       name: slack-notification
 ```

_Appears in:_
- [FunctionSpec](#functionspec)

| Field | Description |
| --- | --- |
| `name` _string_ | Name defines the name of the function to be referenced |


#### FunctionSpec



FunctionSpec defines code that is executed as part of a [KeptnTaskDefinition](#keptntaskdefinition).

_Appears in:_
- [KeptnTaskDefinitionSpec](#keptntaskdefinitionspec)

| Field | Description |
| --- | --- |
| `functionRef` _[FunctionReference](#functionreference)_ | FunctionReference can be used to execute a function that was already defined before |
| `inline` _[Inline](#inline)_ | Inline can be used to provide code to be executed inline in the manifest |
| `httpRef` _[HttpReference](#httpreference)_ | HttpReference can be used to reference a Deno script from an external source |
| `configMapRef` _[ConfigMapReference](#configmapreference)_ |  |
| `parameters` _[TaskParameters](#taskparameters)_ |  |
| `secureParameters` _[SecureParameters](#secureparameters)_ |  |


#### FunctionStatus





_Appears in:_
- [KeptnTaskDefinitionStatus](#keptntaskdefinitionstatus)

| Field | Description |
| --- | --- |
| `configMap` _string_ | INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run "make" to regenerate code after modifying this file |


#### HttpReference



HttpReference Specify a Deno script to be executed at runtime
from the remote webserver that is specified.
For example:

```yaml
name: hello-keptn-http
  spec:
    function:
      httpRef:
        url: <url>
```

_Appears in:_
- [FunctionSpec](#functionspec)

| Field | Description |
| --- | --- |
| `url` _string_ | Url defines the Deno script that should be executed from a remote server |


#### Inline



Inline Include the actual executable code to execute.
This can be written as a full-fledged Deno script.
For example:

```yaml
function:
  inline:
    code: |
      console.log("Deployment Task has been executed");
```

_Appears in:_
- [FunctionSpec](#functionspec)

| Field | Description |
| --- | --- |
| `code` _string_ | Code defines the code that should be run |


#### ItemStatus





_Appears in:_
- [KeptnAppVersionStatus](#keptnappversionstatus)
- [KeptnWorkloadInstanceStatus](#keptnworkloadinstancestatus)

| Field | Description |
| --- | --- |
| `definitionName` _string_ | DefinitionName is the name of the EvaluationDefinition/TaskDefiniton |
| `name` _string_ | Name is the name of the Evaluation/Task |
| `startTime` _[Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#time-v1-meta)_ |  |
| `endTime` _[Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#time-v1-meta)_ |  |


#### KeptnApp



KeptnApp is the Schema for the keptnapps API

_Appears in:_
- [KeptnAppList](#keptnapplist)

| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha2`
| `kind` _string_ | `KeptnApp`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[KeptnAppSpec](#keptnappspec)_ |  |


#### KeptnAppList



KeptnAppList contains a list of KeptnApp



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha2`
| `kind` _string_ | `KeptnAppList`
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `items` _[KeptnApp](#keptnapp) array_ |  |


#### KeptnAppSpec



KeptnAppSpec defines the desired state of KeptnApp

_Appears in:_
- [KeptnApp](#keptnapp)
- [KeptnAppVersionSpec](#keptnappversionspec)

| Field | Description |
| --- | --- |
| `version` _string_ |  |
| `revision` _integer_ |  |
| `workloads` _[KeptnWorkloadRef](#keptnworkloadref) array_ |  |
| `preDeploymentTasks` _string array_ |  |
| `postDeploymentTasks` _string array_ |  |
| `preDeploymentEvaluations` _string array_ |  |
| `postDeploymentEvaluations` _string array_ |  |




#### KeptnAppVersion



KeptnAppVersion is the Schema for the keptnappversions API

_Appears in:_
- [KeptnAppVersionList](#keptnappversionlist)

| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha2`
| `kind` _string_ | `KeptnAppVersion`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[KeptnAppVersionSpec](#keptnappversionspec)_ |  |


#### KeptnAppVersionList



KeptnAppVersionList contains a list of KeptnAppVersion



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha2`
| `kind` _string_ | `KeptnAppVersionList`
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `items` _[KeptnAppVersion](#keptnappversion) array_ |  |


#### KeptnAppVersionSpec



KeptnAppVersionSpec defines the desired state of KeptnAppVersion

_Appears in:_
- [KeptnAppVersion](#keptnappversion)

| Field | Description |
| --- | --- |
| `version` _string_ |  |
| `revision` _integer_ |  |
| `workloads` _[KeptnWorkloadRef](#keptnworkloadref) array_ |  |
| `preDeploymentTasks` _string array_ |  |
| `postDeploymentTasks` _string array_ |  |
| `preDeploymentEvaluations` _string array_ |  |
| `postDeploymentEvaluations` _string array_ |  |
| `appName` _string_ |  |
| `previousVersion` _string_ |  |
| `traceId` _object (keys:string, values:string)_ |  |




#### KeptnEvaluation



KeptnEvaluation is the Schema for the keptnevaluations API

_Appears in:_
- [KeptnEvaluationList](#keptnevaluationlist)

| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha2`
| `kind` _string_ | `KeptnEvaluation`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[KeptnEvaluationSpec](#keptnevaluationspec)_ |  |


#### KeptnEvaluationDefinition



KeptnEvaluationDefinition is the Schema for the keptnevaluationdefinitions API

_Appears in:_
- [KeptnEvaluationDefinitionList](#keptnevaluationdefinitionlist)

| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha2`
| `kind` _string_ | `KeptnEvaluationDefinition`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[KeptnEvaluationDefinitionSpec](#keptnevaluationdefinitionspec)_ |  |


#### KeptnEvaluationDefinitionList



KeptnEvaluationDefinitionList contains a list of KeptnEvaluationDefinition



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha2`
| `kind` _string_ | `KeptnEvaluationDefinitionList`
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `items` _[KeptnEvaluationDefinition](#keptnevaluationdefinition) array_ |  |


#### KeptnEvaluationDefinitionSpec



KeptnEvaluationDefinitionSpec defines the desired state of KeptnEvaluationDefinition

_Appears in:_
- [KeptnEvaluationDefinition](#keptnevaluationdefinition)

| Field | Description |
| --- | --- |
| `source` _string_ |  |
| `objectives` _[Objective](#objective) array_ |  |




#### KeptnEvaluationList



KeptnEvaluationList contains a list of KeptnEvaluation



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha2`
| `kind` _string_ | `KeptnEvaluationList`
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `items` _[KeptnEvaluation](#keptnevaluation) array_ |  |


#### KeptnEvaluationProvider



KeptnEvaluationProvider is the Schema for the keptnevaluationproviders API

_Appears in:_
- [KeptnEvaluationProviderList](#keptnevaluationproviderlist)

| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha2`
| `kind` _string_ | `KeptnEvaluationProvider`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[KeptnEvaluationProviderSpec](#keptnevaluationproviderspec)_ |  |


#### KeptnEvaluationProviderList



KeptnEvaluationProviderList contains a list of KeptnEvaluationProvider



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha2`
| `kind` _string_ | `KeptnEvaluationProviderList`
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `items` _[KeptnEvaluationProvider](#keptnevaluationprovider) array_ |  |


#### KeptnEvaluationProviderSpec



KeptnEvaluationProviderSpec defines the desired state of KeptnEvaluationProvider

_Appears in:_
- [KeptnEvaluationProvider](#keptnevaluationprovider)

| Field | Description |
| --- | --- |
| `targetServer` _string_ |  |
| `secretKeyRef` _[SecretKeySelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#secretkeyselector-v1-core)_ |  |




#### KeptnEvaluationSpec



KeptnEvaluationSpec defines the desired state of KeptnEvaluation

_Appears in:_
- [KeptnEvaluation](#keptnevaluation)

| Field | Description |
| --- | --- |
| `workload` _string_ |  |
| `workloadVersion` _string_ |  |
| `appName` _string_ |  |
| `appVersion` _string_ |  |
| `evaluationDefinition` _string_ |  |
| `retries` _integer_ |  |
| `retryInterval` _[Duration](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#duration-v1-meta)_ |  |
| `failAction` _string_ |  |
| `checkType` _CheckType_ |  |




#### KeptnTask



KeptnTask is the Schema for the keptntasks API

_Appears in:_
- [KeptnTaskList](#keptntasklist)

| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha2`
| `kind` _string_ | `KeptnTask`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[KeptnTaskSpec](#keptntaskspec)_ |  |


#### KeptnTaskDefinition



A KeptnTaskDefinition is a CRD used to define tasks that can be run by the Keptn Lifecycle Toolkit as part of
pre- and post-deployment phases of a deployment. The task definition is a [Deno](https://deno.land/) script.
Please, refer to the [function runtime](https://github.com/keptn/lifecycle-toolkit/tree/main/functions-runtime)
for more information about the runtime.
In the future, we also intend to support other runtimes, especially running a container image directly.

_Appears in:_
- [KeptnTaskDefinitionList](#keptntaskdefinitionlist)

| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha2`
| `kind` _string_ | `KeptnTaskDefinition`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[KeptnTaskDefinitionSpec](#keptntaskdefinitionspec)_ | KeptnTaskDefinitionSpec defines tasks to be executed |


#### KeptnTaskDefinitionList



KeptnTaskDefinitionList contains a list of KeptnTaskDefinition



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha2`
| `kind` _string_ | `KeptnTaskDefinitionList`
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `items` _[KeptnTaskDefinition](#keptntaskdefinition) array_ |  |


#### KeptnTaskDefinitionSpec



KeptnTaskDefinitionSpec defines tasks to be executed.
This can be expressed in the FunctionSpec.

_Appears in:_
- [KeptnTaskDefinition](#keptntaskdefinition)

| Field | Description |
| --- | --- |
| `function` _[FunctionSpec](#functionspec)_ | Function defines what function should be executed in this task |




#### KeptnTaskList



KeptnTaskList contains a list of KeptnTask



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha2`
| `kind` _string_ | `KeptnTaskList`
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `items` _[KeptnTask](#keptntask) array_ |  |


#### KeptnTaskSpec



KeptnTaskSpec defines the desired state of KeptnTask

_Appears in:_
- [KeptnTask](#keptntask)

| Field | Description |
| --- | --- |
| `workload` _string_ |  |
| `workloadVersion` _string_ |  |
| `app` _string_ |  |
| `appVersion` _string_ |  |
| `taskDefinition` _string_ |  |
| `context` _[TaskContext](#taskcontext)_ |  |
| `parameters` _[TaskParameters](#taskparameters)_ |  |
| `secureParameters` _[SecureParameters](#secureparameters)_ |  |
| `checkType` _CheckType_ |  |




#### KeptnWorkload



KeptnWorkload is the Schema for the keptnworkloads API

_Appears in:_
- [KeptnWorkloadList](#keptnworkloadlist)

| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha2`
| `kind` _string_ | `KeptnWorkload`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[KeptnWorkloadSpec](#keptnworkloadspec)_ |  |


#### KeptnWorkloadInstance



KeptnWorkloadInstance is the Schema for the keptnworkloadinstances API

_Appears in:_
- [KeptnWorkloadInstanceList](#keptnworkloadinstancelist)

| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha2`
| `kind` _string_ | `KeptnWorkloadInstance`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[KeptnWorkloadInstanceSpec](#keptnworkloadinstancespec)_ |  |


#### KeptnWorkloadInstanceList



KeptnWorkloadInstanceList contains a list of KeptnWorkloadInstance



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha2`
| `kind` _string_ | `KeptnWorkloadInstanceList`
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `items` _[KeptnWorkloadInstance](#keptnworkloadinstance) array_ |  |


#### KeptnWorkloadInstanceSpec



KeptnWorkloadInstanceSpec defines the desired state of KeptnWorkloadInstance

_Appears in:_
- [KeptnWorkloadInstance](#keptnworkloadinstance)

| Field | Description |
| --- | --- |
| `app` _string_ |  |
| `version` _string_ |  |
| `preDeploymentTasks` _string array_ |  |
| `postDeploymentTasks` _string array_ |  |
| `preDeploymentEvaluations` _string array_ |  |
| `postDeploymentEvaluations` _string array_ |  |
| `resourceReference` _[ResourceReference](#resourcereference)_ |  |
| `workloadName` _string_ |  |
| `previousVersion` _string_ |  |
| `traceId` _object (keys:string, values:string)_ |  |




#### KeptnWorkloadList



KeptnWorkloadList contains a list of KeptnWorkload



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `lifecycle.keptn.sh/v1alpha2`
| `kind` _string_ | `KeptnWorkloadList`
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `items` _[KeptnWorkload](#keptnworkload) array_ |  |


#### KeptnWorkloadRef





_Appears in:_
- [KeptnAppSpec](#keptnappspec)
- [WorkloadStatus](#workloadstatus)

| Field | Description |
| --- | --- |
| `name` _string_ |  |
| `version` _string_ |  |


#### KeptnWorkloadSpec



KeptnWorkloadSpec defines the desired state of KeptnWorkload

_Appears in:_
- [KeptnWorkload](#keptnworkload)
- [KeptnWorkloadInstanceSpec](#keptnworkloadinstancespec)

| Field | Description |
| --- | --- |
| `app` _string_ |  |
| `version` _string_ |  |
| `preDeploymentTasks` _string array_ |  |
| `postDeploymentTasks` _string array_ |  |
| `preDeploymentEvaluations` _string array_ |  |
| `postDeploymentEvaluations` _string array_ |  |
| `resourceReference` _[ResourceReference](#resourcereference)_ |  |




#### Objective





_Appears in:_
- [KeptnEvaluationDefinitionSpec](#keptnevaluationdefinitionspec)

| Field | Description |
| --- | --- |
| `name` _string_ |  |
| `query` _string_ |  |
| `evaluationTarget` _string_ |  |


#### ResourceReference





_Appears in:_
- [KeptnWorkloadSpec](#keptnworkloadspec)

| Field | Description |
| --- | --- |
| `uid` _UID_ |  |
| `kind` _string_ |  |
| `name` _string_ |  |


#### SecureParameters





_Appears in:_
- [FunctionSpec](#functionspec)
- [KeptnTaskSpec](#keptntaskspec)

| Field | Description |
| --- | --- |
| `secret` _string_ |  |


#### TaskContext





_Appears in:_
- [KeptnTaskSpec](#keptntaskspec)

| Field | Description |
| --- | --- |
| `workloadName` _string_ |  |
| `appName` _string_ |  |
| `appVersion` _string_ |  |
| `workloadVersion` _string_ |  |
| `taskType` _string_ |  |
| `objectType` _string_ |  |


#### TaskParameters





_Appears in:_
- [FunctionSpec](#functionspec)
- [KeptnTaskSpec](#keptntaskspec)

| Field | Description |
| --- | --- |
| `map` _object (keys:string, values:string)_ |  |


#### WorkloadStatus





_Appears in:_
- [KeptnAppVersionStatus](#keptnappversionstatus)

| Field | Description |
| --- | --- |
| `workload` _[KeptnWorkloadRef](#keptnworkloadref)_ |  |


