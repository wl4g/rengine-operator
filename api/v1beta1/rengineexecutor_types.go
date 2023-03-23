/*
Copyright 2023 James Wong<jameswong1376@gmail.com>.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RengineExecutorSpec defines the desired state of RengineExecutor
type RengineExecutorSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Image string `json:"image,omitempty"`
}

// RengineExecutorStatus defines the observed state of RengineExecutor
type RengineExecutorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	MinReplicas int `json:"minReplicas,omitempty"`
	MaxReplicas int `json:"maxReplicas,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// RengineExecutor is the Schema for the rengineexecutors API
type RengineExecutor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RengineExecutorSpec   `json:"spec,omitempty"`
	Status RengineExecutorStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RengineExecutorList contains a list of RengineExecutor
type RengineExecutorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RengineExecutor `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RengineExecutor{}, &RengineExecutorList{})
}
