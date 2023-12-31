//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Connectivity) DeepCopyInto(out *Connectivity) {
	{
		in := &in
		*out = make(Connectivity, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Connectivity.
func (in Connectivity) DeepCopy() Connectivity {
	if in == nil {
		return nil
	}
	out := new(Connectivity)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectivitySpec) DeepCopyInto(out *ConnectivitySpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectivitySpec.
func (in *ConnectivitySpec) DeepCopy() *ConnectivitySpec {
	if in == nil {
		return nil
	}
	out := new(ConnectivitySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Edge) DeepCopyInto(out *Edge) {
	*out = *in
	out.EdgeCluster = in.EdgeCluster
	if in.Vertex != nil {
		in, out := &in.Vertex, &out.Vertex
		*out = new(Vertex)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Edge.
func (in *Edge) DeepCopy() *Edge {
	if in == nil {
		return nil
	}
	out := new(Edge)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeClusterInfo) DeepCopyInto(out *EdgeClusterInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeClusterInfo.
func (in *EdgeClusterInfo) DeepCopy() *EdgeClusterInfo {
	if in == nil {
		return nil
	}
	out := new(EdgeClusterInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Graph) DeepCopyInto(out *Graph) {
	*out = *in
	if in.Vertices != nil {
		in, out := &in.Vertices, &out.Vertices
		*out = make(map[int]*Vertex, len(*in))
		for key, val := range *in {
			var outVal *Vertex
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(Vertex)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Graph.
func (in *Graph) DeepCopy() *Graph {
	if in == nil {
		return nil
	}
	out := new(Graph)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LatencyPercentileScoreSpec) DeepCopyInto(out *LatencyPercentileScoreSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LatencyPercentileScoreSpec.
func (in *LatencyPercentileScoreSpec) DeepCopy() *LatencyPercentileScoreSpec {
	if in == nil {
		return nil
	}
	out := new(LatencyPercentileScoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinkService) DeepCopyInto(out *LinkService) {
	*out = *in
	out.Metadata = in.Metadata
	in.Service.DeepCopyInto(&out.Service)
	if in.Deployment != nil {
		in, out := &in.Deployment, &out.Deployment
		*out = new(corev1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinkService.
func (in *LinkService) DeepCopy() *LinkService {
	if in == nil {
		return nil
	}
	out := new(LinkService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetaData) DeepCopyInto(out *MetaData) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetaData.
func (in *MetaData) DeepCopy() *MetaData {
	if in == nil {
		return nil
	}
	out := new(MetaData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkDataInfo) DeepCopyInto(out *NetworkDataInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkDataInfo.
func (in *NetworkDataInfo) DeepCopy() *NetworkDataInfo {
	if in == nil {
		return nil
	}
	out := new(NetworkDataInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SFCDeployment) DeepCopyInto(out *SFCDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SFCDeployment.
func (in *SFCDeployment) DeepCopy() *SFCDeployment {
	if in == nil {
		return nil
	}
	out := new(SFCDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SFCDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SFCDeploymentList) DeepCopyInto(out *SFCDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SFCDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SFCDeploymentList.
func (in *SFCDeploymentList) DeepCopy() *SFCDeploymentList {
	if in == nil {
		return nil
	}
	out := new(SFCDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SFCDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SFCDeploymentSpec) DeepCopyInto(out *SFCDeploymentSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	in.Template.DeepCopyInto(&out.Template)
	in.Strategy.DeepCopyInto(&out.Strategy)
	if in.RevisionHistoryLimit != nil {
		in, out := &in.RevisionHistoryLimit, &out.RevisionHistoryLimit
		*out = new(int32)
		**out = **in
	}
	if in.ProgressDeadlineSeconds != nil {
		in, out := &in.ProgressDeadlineSeconds, &out.ProgressDeadlineSeconds
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SFCDeploymentSpec.
func (in *SFCDeploymentSpec) DeepCopy() *SFCDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(SFCDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SFCDeploymentStatus) DeepCopyInto(out *SFCDeploymentStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SFCDeploymentStatus.
func (in *SFCDeploymentStatus) DeepCopy() *SFCDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(SFCDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SFCService) DeepCopyInto(out *SFCService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SFCService.
func (in *SFCService) DeepCopy() *SFCService {
	if in == nil {
		return nil
	}
	out := new(SFCService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SFCService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SFCServiceList) DeepCopyInto(out *SFCServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SFCService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SFCServiceList.
func (in *SFCServiceList) DeepCopy() *SFCServiceList {
	if in == nil {
		return nil
	}
	out := new(SFCServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SFCServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SFCServiceSpec) DeepCopyInto(out *SFCServiceSpec) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]corev1.ServicePort, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ClusterIPs != nil {
		in, out := &in.ClusterIPs, &out.ClusterIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExternalIPs != nil {
		in, out := &in.ExternalIPs, &out.ExternalIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LoadBalancerSourceRanges != nil {
		in, out := &in.LoadBalancerSourceRanges, &out.LoadBalancerSourceRanges
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SessionAffinityConfig != nil {
		in, out := &in.SessionAffinityConfig, &out.SessionAffinityConfig
		*out = new(corev1.SessionAffinityConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.IPFamilies != nil {
		in, out := &in.IPFamilies, &out.IPFamilies
		*out = make([]corev1.IPFamily, len(*in))
		copy(*out, *in)
	}
	if in.IPFamilyPolicy != nil {
		in, out := &in.IPFamilyPolicy, &out.IPFamilyPolicy
		*out = new(corev1.IPFamilyPolicy)
		**out = **in
	}
	if in.AllocateLoadBalancerNodePorts != nil {
		in, out := &in.AllocateLoadBalancerNodePorts, &out.AllocateLoadBalancerNodePorts
		*out = new(bool)
		**out = **in
	}
	if in.LoadBalancerClass != nil {
		in, out := &in.LoadBalancerClass, &out.LoadBalancerClass
		*out = new(string)
		**out = **in
	}
	if in.InternalTrafficPolicy != nil {
		in, out := &in.InternalTrafficPolicy, &out.InternalTrafficPolicy
		*out = new(corev1.ServiceInternalTrafficPolicyType)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SFCServiceSpec.
func (in *SFCServiceSpec) DeepCopy() *SFCServiceSpec {
	if in == nil {
		return nil
	}
	out := new(SFCServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SFCServiceStatus) DeepCopyInto(out *SFCServiceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SFCServiceStatus.
func (in *SFCServiceStatus) DeepCopy() *SFCServiceStatus {
	if in == nil {
		return nil
	}
	out := new(SFCServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Scheduler) DeepCopyInto(out *Scheduler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Scheduler.
func (in *Scheduler) DeepCopy() *Scheduler {
	if in == nil {
		return nil
	}
	out := new(Scheduler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Scheduler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerList) DeepCopyInto(out *SchedulerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Scheduler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerList.
func (in *SchedulerList) DeepCopy() *SchedulerList {
	if in == nil {
		return nil
	}
	out := new(SchedulerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SchedulerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerSpec) DeepCopyInto(out *SchedulerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerSpec.
func (in *SchedulerSpec) DeepCopy() *SchedulerSpec {
	if in == nil {
		return nil
	}
	out := new(SchedulerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerStatus) DeepCopyInto(out *SchedulerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerStatus.
func (in *SchedulerStatus) DeepCopy() *SchedulerStatus {
	if in == nil {
		return nil
	}
	out := new(SchedulerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceDefinition) DeepCopyInto(out *ServiceDefinition) {
	*out = *in
	out.Metadata = in.Metadata
	if in.Connectivity != nil {
		in, out := &in.Connectivity, &out.Connectivity
		*out = make(Connectivity, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ServiceRef != nil {
		in, out := &in.ServiceRef, &out.ServiceRef
		*out = new(corev1.ObjectReference)
		**out = **in
	}
	if in.TargetServiceRef != nil {
		in, out := &in.TargetServiceRef, &out.TargetServiceRef
		*out = new(corev1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceDefinition.
func (in *ServiceDefinition) DeepCopy() *ServiceDefinition {
	if in == nil {
		return nil
	}
	out := new(ServiceDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceFunctionChain) DeepCopyInto(out *ServiceFunctionChain) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceFunctionChain.
func (in *ServiceFunctionChain) DeepCopy() *ServiceFunctionChain {
	if in == nil {
		return nil
	}
	out := new(ServiceFunctionChain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceFunctionChain) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceFunctionChainList) DeepCopyInto(out *ServiceFunctionChainList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceFunctionChain, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceFunctionChainList.
func (in *ServiceFunctionChainList) DeepCopy() *ServiceFunctionChainList {
	if in == nil {
		return nil
	}
	out := new(ServiceFunctionChainList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceFunctionChainList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceFunctionChainSpec) DeepCopyInto(out *ServiceFunctionChainSpec) {
	*out = *in
	if in.Links != nil {
		in, out := &in.Links, &out.Links
		*out = make([]LinkService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceFunctionChainSpec.
func (in *ServiceFunctionChainSpec) DeepCopy() *ServiceFunctionChainSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceFunctionChainSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceFunctionChainStatus) DeepCopyInto(out *ServiceFunctionChainStatus) {
	*out = *in
	if in.ServiceFunctions != nil {
		in, out := &in.ServiceFunctions, &out.ServiceFunctions
		*out = make([]ServiceFunctionInfo, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceFunctionChainStatus.
func (in *ServiceFunctionChainStatus) DeepCopy() *ServiceFunctionChainStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceFunctionChainStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceFunctionInfo) DeepCopyInto(out *ServiceFunctionInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceFunctionInfo.
func (in *ServiceFunctionInfo) DeepCopy() *ServiceFunctionInfo {
	if in == nil {
		return nil
	}
	out := new(ServiceFunctionInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceLevelAgreement) DeepCopyInto(out *ServiceLevelAgreement) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceLevelAgreement.
func (in *ServiceLevelAgreement) DeepCopy() *ServiceLevelAgreement {
	if in == nil {
		return nil
	}
	out := new(ServiceLevelAgreement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceLevelAgreement) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceLevelAgreementList) DeepCopyInto(out *ServiceLevelAgreementList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceLevelAgreement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceLevelAgreementList.
func (in *ServiceLevelAgreementList) DeepCopy() *ServiceLevelAgreementList {
	if in == nil {
		return nil
	}
	out := new(ServiceLevelAgreementList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceLevelAgreementList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceLevelAgreementSpec) DeepCopyInto(out *ServiceLevelAgreementSpec) {
	*out = *in
	out.Connectivity = in.Connectivity
	out.LatencyPercentileScore = in.LatencyPercentileScore
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceLevelAgreementSpec.
func (in *ServiceLevelAgreementSpec) DeepCopy() *ServiceLevelAgreementSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceLevelAgreementSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceLevelAgreementStatus) DeepCopyInto(out *ServiceLevelAgreementStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceLevelAgreementStatus.
func (in *ServiceLevelAgreementStatus) DeepCopy() *ServiceLevelAgreementStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceLevelAgreementStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Vertex) DeepCopyInto(out *Vertex) {
	*out = *in
	out.NetworkData = in.NetworkData
	if in.Edges != nil {
		in, out := &in.Edges, &out.Edges
		*out = make(map[int]*Edge, len(*in))
		for key, val := range *in {
			var outVal *Edge
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(Edge)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Vertex.
func (in *Vertex) DeepCopy() *Vertex {
	if in == nil {
		return nil
	}
	out := new(Vertex)
	in.DeepCopyInto(out)
	return out
}
