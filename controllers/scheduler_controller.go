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
	"context"
	"errors"

	"github.com/go-logr/logr"
	sfcv1 "github.com/ntnguyencse/Service-Chain-Orchestrator/api/v1"
	intentv1 "github.com/ntnguyencse/l-kaas/api/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

type ScheduleResult struct {
	// Name of the selected cluster.
	SuggestedCluster string
	// The number of nodes the scheduler evaluated the pod against in the filtering
	// phase and beyond.
	EvaluatedCluster int
	// The number of nodes out of the evaluated ones that fit the pod.
	FeasibleCluster int
	// contains filtered or unexported fields
}

// SchedulerReconciler reconciles a Scheduler object
type SchedulerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	l      logr.Logger
	s      *json.Serializer
}

var (
	loggerSD = ctrl.Log.WithName("Scheduler Controller")
)

//+kubebuilder:rbac:groups=sfc.automation.dcn.ssu.ac.kr,resources=schedulers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=sfc.automation.dcn.ssu.ac.kr,resources=schedulers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=sfc.automation.dcn.ssu.ac.kr,resources=schedulers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Scheduler object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *SchedulerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	loggerSD.Info("Start Scheduler COntroller")
	SFC := &sfcv1.ServiceFunctionChain{}
	if err := r.Client.Get(ctx, req.NamespacedName, SFC); err != nil {
		if apierrors.IsNotFound(err) {
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			return ctrl.Result{}, nil
		}

		// Error reading the object - requeue the request.
		return ctrl.Result{}, err
	}
	listLocation := []string{"Location 1", "Location 2", "Location 3"}
	loggerSD.Info("Start scheduling Service Function Chain")
	listSF := SFC.Spec.Links
	for id, SF := range listSF {
		loggerSD.Info("Print Service Function: ", SF.Metadata.Name, SF.Deployment)
		loggerSD.Info("Finding placement for ", SF.Metadata.Name, SF.Deployment)
		loggerSD.Info("Placement for ", SF.Metadata.Name, listLocation[id])

	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *SchedulerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&sfcv1.ServiceFunctionChain{}).
		Complete(r)
}

func (r *SchedulerReconciler) GetLogicalCluster(ctx context.Context, req ctrl.Request) (*intentv1.ListLogicalCluster, error) {
	err := errors.New("Error when get logical cluster")
	listLogicalCluster := &intentv1.ListLogicalCluster{}
	if err := r.Client.List(ctx, listLogicalCluster); err != nil {

		loggerSD.Error(err, "Error when get Logical Cluster", "Error when list logical cluster")
		return nil, err

	}
	return listLogicalCluster, err
}
func FindLogicalCluster(name string, list *intentv1.ListLogicalCluster) (intentv1.LogicalCluster, error) {
	err := errors.New("Error when get logical cluster")
	var logicalCluster intentv1.logicalCluster
	for _, item := range list {
		if item.name == name {
			logicalCluster = item
		}
	}

	return logicalCluster, err
}
func (r *SchedulerReconciler) MakeLogicalTopology(ctx context.Context, rep ctrl.Request) error {
	err := errors.New("Error when get logical cluster")

	// var numberOfNodes int = 4
	// var listNodes []Node

	// for id, i := range listNodes {

	// }

	return err

}
func (r *SchedulerReconciler) ChooseLocationForDeployment(ctx context.Context, deployment sfcv1.SFCDeployment) error {

	return nil
}

// Node represents a node in the graph.
type Node struct {
	Name  string
	Edges []*Node
}

// AddEdge adds a new edge to the node.
func (n *Node) Add_Edge(node *Node) {

	n.Edges = append(n.Edges, node)
}
func ScheduleServiceDeployment() (ScheduleResult, error) {
	var result ScheduleResult

	return result, nil
}
