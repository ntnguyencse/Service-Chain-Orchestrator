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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SchedulerSpec defines the desired state of Scheduler
type SchedulerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Scheduler. Edit scheduler_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// SchedulerStatus defines the observed state of Scheduler
type SchedulerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Scheduler is the Schema for the schedulers API
type Scheduler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SchedulerSpec   `json:"spec,omitempty"`
	Status SchedulerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SchedulerList contains a list of Scheduler
type SchedulerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Scheduler `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Scheduler{}, &SchedulerList{})
}

// Graph
// Source: https://gist.github.com/snassr/e79f4953eeb8d813be82eda00adeef57
// Graph represents a set of vertices connected by edges.
type Graph struct {
	Vertices map[int]*Vertex
}

// Vertex is a node in the graph that stores the int value at that node
// along with a map to the vertices it is connected to via edges.
type Vertex struct {
	NetworkData NetworkDataInfo
	Edges       map[int]*Edge
}
type NetworkDataInfo struct {
	// Latency in mili seconds
	// End-to-end delay: time taken for a packet to be transmitted across a network from user to service or reverse.
	Latency int
	// Round trip delay (milisecond)
	// Round-trip Delay: the amount of time it takes for a data packet to be sent plus the amount of time it takes to acknowledge that signal to be received.
	RoundTripDelay int
	// Packet Loss Rate (percent)
	// Packet Loss Rate: the number of packets not received divided by the total number of packets sent
	PacketLossRate int
	// Hop Count
	// Hop Count: the number of network devices through which data passes from source to destination.
	HopCount int
	// Throught put ( bytes/second)
	// Throughput: the rate of message delivery over a communication channel.
	Throughput int
	// AvailableBandwidth (bytes/second)
	// Available bandwidth: The maximum amount of data transmitted over an connection in a given amount of time
	AvailableBandwidth int
	// DelayVariation percent
	// Delay Variation:  the difference in end-to-end one-way delay between selected packets in a flow with any lost packets being ignored
	DelayVariation int
}

// Edge represents an edge in the graph and the destination vertex.
type Edge struct {
	EdgeCluster EdgeClusterInfo
	Vertex      *Vertex
}
type EdgeClusterInfo struct {
	Location  string
	Longitude string
	Latitude  string
}

// AddVertex adds a vertex to the graph with no edges.
func (this *Graph) AddVertex(key int, val NetworkDataInfo) {
	this.Vertices[key] = &Vertex{NetworkData: val, Edges: map[int]*Edge{}}
}

// AddEdge adds an edge between existing source and existing destination vertex.
func (this *Graph) AddEdge(srcKey, destKey int, edgeinfo EdgeClusterInfo) {
	// check if src & dest exist
	if _, ok := this.Vertices[srcKey]; !ok {
		return
	}
	if _, ok := this.Vertices[destKey]; !ok {
		return
	}

	// add edge src --> dest
	this.Vertices[srcKey].Edges[destKey] = &Edge{EdgeCluster: edgeinfo, Vertex: this.Vertices[destKey]}
}

// Neighbors returns all vertex values that have an edge from
// the provided src vertex.
func (this *Graph) Neighbors(srcKey int) []NetworkDataInfo {
	result := []NetworkDataInfo{}

	if _, ok := this.Vertices[srcKey]; !ok {
		return result
	}

	for _, edge := range this.Vertices[srcKey].Edges {
		result = append(result, edge.Vertex.NetworkData)
	}

	return result
}
