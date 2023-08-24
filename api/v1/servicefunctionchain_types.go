/*
Copyright 2023 Nguyen Thanh Nguyen.

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

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ServiceFunctionChainSpec defines the desired state of ServiceFunctionChain
type ServiceFunctionChainSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	NumberOfNode                 int    `json:"numberofnode,omitempty"`
	DefaultServiceLevelAgreement string `json:"defaultSLA,omitempty"`
}
type LinkService struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Service           ServiceDefinition       `json:"service,omitempty"`
	Deployment        *corev1.ObjectReference `json:"deployment,omitempty"`
}
type ServiceDefinition struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Connectivity      Connectivity            `json:"connectivity,omitempty"`
	ServiceRef        *corev1.ObjectReference `json:"ServiceRef,omitempty"`
	TargetServiceRef  *corev1.ObjectReference `json:"TargetServiceRef,omitempty"`
}
type Connectivity map[string]string

// ServiceFunctionChainStatus defines the observed state of ServiceFunctionChain
type ServiceFunctionChainStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Placement    string `json:"placement,omitempty"`
	OldPlacement string `json:"oldplacement,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ServiceFunctionChain is the Schema for the servicefunctionchains API
type ServiceFunctionChain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceFunctionChainSpec   `json:"spec,omitempty"`
	Status ServiceFunctionChainStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ServiceFunctionChainList contains a list of ServiceFunctionChain
type ServiceFunctionChainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceFunctionChain `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ServiceFunctionChain{}, &ServiceFunctionChainList{})
}
