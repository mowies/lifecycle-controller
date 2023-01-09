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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KeptnConfigSpec defines the desired state of KeptnConfig
type KeptnConfigSpec struct {

	//+kubebuilder:default=""

	// OTelCollectorUrl can be used to set the Open Telemetry collector that the operator should use
	OTelCollectorUrl string `json:"OTelCollectorUrl,omitempty"`
}

// KeptnConfigStatus defines the observed state of KeptnConfig
type KeptnConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// KeptnConfig is the Schema for the keptnconfigs API
type KeptnConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KeptnConfigSpec   `json:"spec,omitempty"`
	Status KeptnConfigStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// KeptnConfigList contains a list of KeptnConfig
type KeptnConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeptnConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KeptnConfig{}, &KeptnConfigList{})
}
