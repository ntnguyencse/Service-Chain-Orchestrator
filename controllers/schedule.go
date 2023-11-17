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

package controllers

import (
	"fmt"

	sfcv1 "github.com/ntnguyencse/Service-Chain-Orchestrator/api/v1"
)

type SystemStatus struct {
	CPULoad           float32
	MemoryAvailable   int32
	MemoryTotal       int32
	CPUUtilization    float32
	MemoryUtilization float32
}
type NetworkMeshInformation struct {
	Latency float32
}

func Schedule(sfc *sfcv1.ServiceFunctionChain) {
	// Schedule
	// Step 1:
	GetRequestFromQueue(sfc)
	// Step 2:
	graph, err := GetSystemStatus(sfc)
	if err != nil {
		fmt.Println("Error when get system status")
	}
	//Step 3:
	PreFilterClusterPlacement(sfc, &graph)
	// Step 4:
	FilteringClusterPlacement(sfc)
	// Step 5:
	PostFilterClusterPlacement(sfc)
	// Step 6:
	ScoringClusterPlacement(sfc)
	// Step 7:
	NormalizeScore(sfc)
	// Step 8:
	Reverse(sfc)
	// Step 9:
	PreBindClusterPlacement(sfc)
	// Step 10:
	BindWorkloadsToCluster(sfc)
	// Step 11:
	PostBindWorkloads(sfc)

}

func GetRequestFromQueue(sfc *sfcv1.ServiceFunctionChain) {
	// Queue of request deploy workload
	loggerSFCS.Info("Get request from Queue to Schedule")
	loggerSFCS.Info("Request Schedule: ")
}

func GetSystemStatus(sfc *sfcv1.ServiceFunctionChain) (Graph, error) {
	// Get current status of system to determine state of system, resources, ultization
	return GetSystemTopologySimlated(3), nil
}

func PreFilterClusterPlacement(sfc *sfcv1.ServiceFunctionChain, graph *Graph) {
	// Pre-process cluster palcement before filtering
	// Remove some palcement not suitable for workload
}

func FilteringClusterPlacement(sfc *sfcv1.ServiceFunctionChain) {
	// Filter cluster placement that could handle workloads
}

func PostFilterClusterPlacement(sfc *sfcv1.ServiceFunctionChain) {
	// Post Process Cluster Placement before Scoring function
}

func ScoringClusterPlacement(sfc *sfcv1.ServiceFunctionChain) {
	// Evaluate and score the cluster palcement based on score rules:
	// - Based on Computational Resources: Ranks filtered clusters to choose the most suitable placement based on score rules (the most suitable is the least score)
	// - Based on SLA: between service to service or service to end-user if service running on this placement.
	// Scheduler assigns service function to the cluster with the lowest ranking

}

func NormalizeScore(sfc *sfcv1.ServiceFunctionChain) {

	// Normalize Score marks of cluster placements
}

func Reverse(sfc *sfcv1.ServiceFunctionChain) {
	// Depend on
}
func PreBindClusterPlacement(sfc *sfcv1.ServiceFunctionChain) {
	// Prepare before bind workload to cluster

}

func BindWorkloadsToCluster(sfc *sfcv1.ServiceFunctionChain) {
	// Assigning workload to Kubernetes cluster

}

func PostBindWorkloads(sfc *sfcv1.ServiceFunctionChain) {
	// Update or do the job after bind
}
