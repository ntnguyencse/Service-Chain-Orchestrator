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

func Schedule() {
	// Schedule
	// Step 1:
	GetRequestFromQueue()
	// Step 2:
	GetSystemStatus()
	//Step 3:
	PreFilterClusterPlacement()
	// Step 4:
	FilteringClusterPlacement()
	// Step 5:
	PostFilterClusterPlacement()
	// Step 6:
	ScoringClusterPlacement()
	// Step 7:
	NormalizeScore()
	// Step 8:
	Reverse()
	// Step 9:
	PreBindClusterPlacement()
	// Step 10:
	BindWorkloadsToCluster()
	// Step 11:
	PostBindWorkloads()

}

func GetRequestFromQueue() {
	// Queue of request deploy workload
	
}

func GetSystemStatus() {
	// Get current status of system to determine state of system, resources, ultization

}

func PreFilterClusterPlacement() {
	// Pre-process cluster palcement before filtering

}

func FilteringClusterPlacement() {
	// Filter cluster placement that could handle workloads
}

func PostFilterClusterPlacement() {
	// Post Process Cluster Placement before Scoring function
}

func ScoringClusterPlacement() {
	// Evaluate and score the cluster palcement based on score rules:
	// - Based on Computational Resources: Ranks filtered clusters to choose the most suitable placement based on score rules (the most suitable is the least score)
	// - Based on SLA: between service to service or service to end-user if service running on this placement.
	// Scheduler assigns service function to the cluster with the lowest ranking

}

func NormalizeScore() {

	// Normalize Score marks of cluster placements
}

func Reverse() {
	// Depend on
}
func PreBindClusterPlacement() {
	// Prepare before bind workload to cluster

}

func BindWorkloadsToCluster() {
	// Assigning workload to Kubernetes cluster

}

func PostBindWorkloads() {
	// Update or do the job after bind
}
