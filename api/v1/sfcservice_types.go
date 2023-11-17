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

// SFCServiceSpec defines the desired state of SFCService

// If any error, refer: https://github.com/kubernetes-sigs/controller-tools/issues/772#issuecomment-1416027564
// just set CONTROLLER_TOOLS_VERSION ?= v0.11.2 in Makefile and k8s.io/api v0.26.1 in go.mod, and it's working now! thanks!
type SFCServiceSpec corev1.ServiceSpec

// type SFCServiceSpec struct {
// 	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
// 	// Important: Run "make" to regenerate code after modifying this file

// 	// Foo is an example field of SFCService. Edit sfcservice_types.go to remove/update
// 	Foo string `json:"foo,omitempty"`
// }

// SFCServiceStatus defines the observed state of SFCService
type SFCServiceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	ReadyForDeploy bool   `json:"readyfordeploy,omitempty"`
	Location       string `json:"location,omitempty"`
	GithubLink     string `json:"githublink,omitempty"`
	Deployed       bool   `json:"deployed,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// SFCService is the Schema for the sfcservices API
type SFCService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SFCServiceSpec   `json:"spec,omitempty"`
	Status SFCServiceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SFCServiceList contains a list of SFCService
type SFCServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SFCService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SFCService{}, &SFCServiceList{})
}
