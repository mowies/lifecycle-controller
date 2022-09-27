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
	"k8s.io/apimachinery/pkg/types"
	"strings"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KeptnWorkloadSpec defines the desired state of KeptnWorkload
type KeptnWorkloadSpec struct {
	AppName                string            `json:"app"`
	Version                string            `json:"version"`
	PreDeploymentTasks     []string          `json:"preDeploymentTasks,omitempty"`
	PostDeploymentTasks    []string          `json:"postDeploymentTasks,omitempty"`
	PreDeploymentAnalysis  []string          `json:"preDeploymentAnalysis,omitempty"`
	PostDeploymentAnalysis []string          `json:"postDeploymentAnalysis,omitempty"`
	ResourceReference      ResourceReference `json:"resourceReference"`
}

// KeptnWorkloadStatus defines the observed state of KeptnWorkload
type KeptnWorkloadStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// KeptnWorkload is the Schema for the keptnworkloads API
type KeptnWorkload struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KeptnWorkloadSpec   `json:"spec,omitempty"`
	Status KeptnWorkloadStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// KeptnWorkloadList contains a list of KeptnWorkload
type KeptnWorkloadList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeptnWorkload `json:"items"`
}

type ResourceReference struct {
	UID  types.UID `json:"uid"`
	Kind string    `json:"kind"`
}

func init() {
	SchemeBuilder.Register(&KeptnWorkload{}, &KeptnWorkloadList{})
}

func (w KeptnWorkload) GetWorkloadInstanceName() string {
	return strings.ToLower(w.Name + "-" + w.Spec.Version)
}
