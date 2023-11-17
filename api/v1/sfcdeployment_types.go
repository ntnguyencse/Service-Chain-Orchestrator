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
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SFCDeploymentSpec defines the desired state of SFCDeployment

// If any error, refer: https://github.com/kubernetes-sigs/controller-tools/issues/772#issuecomment-1416027564
// just set CONTROLLER_TOOLS_VERSION ?= v0.11.2 in Makefile and k8s.io/api v0.26.1 in go.mod, and it's working now! thanks!

type SFCDeploymentSpec appsv1.DeploymentSpec

// SFCDeploymentStatus defines the observed state of SFCDeployment
type SFCDeploymentStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	ReadyForDeploy bool   `json:"readyfordeploy,omitempty"`
	Location       string `json:"location,omitempty"`
	GithubLink     string `json:"githublink,omitempty"`
	Deployed       bool   `json:"deployed,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// SFCDeployment is the Schema for the sfcdeployments API
type SFCDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SFCDeploymentSpec   `json:"spec,omitempty"`
	Status SFCDeploymentStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SFCDeploymentList contains a list of SFCDeployment
type SFCDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SFCDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SFCDeployment{}, &SFCDeploymentList{})
}
