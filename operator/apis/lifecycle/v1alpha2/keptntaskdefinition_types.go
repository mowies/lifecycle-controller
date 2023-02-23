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

package v1alpha2

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// KeptnTaskDefinitionSpec defines tasks to be executed.
// This can be expressed in the FunctionSpec.
type KeptnTaskDefinitionSpec struct {
	// Function defines what function should be executed in this task
	Function FunctionSpec `json:"function,omitempty"`
}

// FunctionSpec defines code that is executed as part of a [KeptnTaskDefinition](#keptntaskdefinition).
type FunctionSpec struct {
	// FunctionReference can be used to execute a function that was already defined before
	FunctionReference FunctionReference `json:"functionRef,omitempty"`
	// Inline can be used to provide code to be executed inline in the manifest
	Inline Inline `json:"inline,omitempty"`
	// HttpReference can be used to reference a Deno script from an external source
	HttpReference      HttpReference      `json:"httpRef,omitempty"`
	ConfigMapReference ConfigMapReference `json:"configMapRef,omitempty"`
	Parameters         TaskParameters     `json:"parameters,omitempty"`
	SecureParameters   SecureParameters   `json:"secureParameters,omitempty"`
}

type ConfigMapReference struct {
	Name string `json:"name,omitempty"`
}

// FunctionReference Execute another `KeptnTaskDefinition` that has been defined.
// Populate this field with the value of the `name` field
// for the `KeptnTaskDefinition` to be called.
// This is commonly used to call a general function
// that is used in multiple place with different parameters.
// An example is:
//
//  ```yaml
//  spec:
//    function:
//      functionRef:
//        name: slack-notification
//  ```
type FunctionReference struct {
	// Name defines the name of the function to be referenced
	Name string `json:"name,omitempty"`
}

// Inline Include the actual executable code to execute.
// This can be written as a full-fledged Deno script.
// For example:
//
// ```yaml
// function:
//   inline:
//     code: |
//       console.log("Deployment Task has been executed");
// ```
type Inline struct {
	// Code defines the code that should be run
	Code string `json:"code,omitempty"`
}

// HttpReference Specify a Deno script to be executed at runtime
// from the remote webserver that is specified.
// For example:
//
// ```yaml
// name: hello-keptn-http
//   spec:
//     function:
//       httpRef:
//         url: <url>
// ```
type HttpReference struct {
	// Url defines the Deno script that should be executed from a remote server
	Url string `json:"url,omitempty"`
}

type ContainerSpec struct {
}

// KeptnTaskDefinitionStatus defines the observed state of KeptnTaskDefinition
type KeptnTaskDefinitionStatus struct {
	Function FunctionStatus `json:"function,omitempty"`
}

type FunctionStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	ConfigMap string `json:"configMap,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status

// A KeptnTaskDefinition is a CRD used to define tasks that can be run by the Keptn Lifecycle Toolkit as part of
// pre- and post-deployment phases of a deployment. The task definition is a [Deno](https://deno.land/) script.
// Please, refer to the [function runtime](https://github.com/keptn/lifecycle-toolkit/tree/main/functions-runtime)
// for more information about the runtime.
// In the future, we also intend to support other runtimes, especially running a container image directly.
type KeptnTaskDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// KeptnTaskDefinitionSpec defines tasks to be executed
	Spec   KeptnTaskDefinitionSpec   `json:"spec,omitempty"`
	Status KeptnTaskDefinitionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KeptnTaskDefinitionList contains a list of KeptnTaskDefinition
type KeptnTaskDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeptnTaskDefinition `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KeptnTaskDefinition{}, &KeptnTaskDefinitionList{})
}
