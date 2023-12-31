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

	github "github.com/google/go-github/v56/github"
	intentv1 "github.com/ntnguyencse/L-KaaS/api/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	// capiulti "sigs.k8s.io/cluster-api/util"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	"github.com/go-logr/logr"
	sfcv1 "github.com/ntnguyencse/Service-Chain-Orchestrator/api/v1"
)

// ServiceFunctionChainReconciler reconciles a ServiceFunctionChain object
type ServiceFunctionChainReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	l      logr.Logger
	s      *json.Serializer
}

const (
	HEADSFC      int    = 1
	ENDSFC              = 2
	BETWEENSFC          = 3
	SFCFinalizer string = "sfc.automation.dcn.ssu.ac.kr"
)

var (
	loggerSFC = ctrl.Log.WithName("SFC Main Controller")
)

type ClusterResource struct {
	Name            string
	CPUAvailable    int64
	MemoryAvailable int64
	GPUAvailable    int64
	VPUAvailable    int64
}

//+kubebuilder:rbac:groups=sfc.automation.dcn.ssu.ac.kr,resources=servicefunctionchains,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=sfc.automation.dcn.ssu.ac.kr,resources=servicefunctionchains/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=sfc.automation.dcn.ssu.ac.kr,resources=servicefunctionchains/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the ServiceFunctionChain object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *ServiceFunctionChainReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
	token, err := GetGithubToken(GithubTokenFilePath)
	if err != nil {
		loggerSFC.Error(err, "Error when get github token", GithubTokenFilePath, err)
		return ctrl.Result{}, nil
	}
	client := github.NewClient(nil).WithAuthToken(token)

	repo := GitRepository{
		Owner:    "SFC-Demo",
		RepoName: "edge",
	}
	// TODO(user): your logic here
	loggerSFC.Info("Start SFC Main controller")

	SFCObjject := &sfcv1.ServiceFunctionChain{}
	if err := r.Client.Get(ctx, req.NamespacedName, SFCObjject); err != nil {
		if apierrors.IsNotFound(err) {
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return ctrl.Result{}, err
	}
	// Check if object is deleted
	if !SFCObjject.ObjectMeta.DeletionTimestamp.IsZero() {
		loggerSFC.Info("SFCObjject is Deleted")

		return r.ReconcilerDelete(ctx, SFCObjject)
	}
	for _, link := range SFCObjject.Spec.Links {
		// Translate
		translatedString, err1 := r.TranslateSFCtoServiceMeshDeployment(ctx, req, link)
		if err1 != nil {
			loggerSFC.Error(err1, "Error when translate SFC to Service Mesh")
		} else {
			loggerSFC.Info(translatedString)
		}

	}
	// Translate SFC to deployment
	// Check is deployed
	SFCStatus := SFCObjject.Status
	// if SFCStatus.Translated {

	// }
	for id, deployment := range SFCStatus.ServiceFunctions {
		if false {
			if len(deployment.Placement) > 0 && !deployment.Deployed {
				// Deploy to cluster
				// Commit to git repository
				content := string("aaaa")
				// Call translate
				path := SFCObjject.Name + "/" + SFCObjject.Spec.Links[id].Metadata.Name + "/" + "deployment.yaml"
				CreateAFile(client, ctx, repo, path, &content)
				SFCObjject.Status.ServiceFunctions[id].Deployed = true
			}
		}

	}
	return ctrl.Result{}, nil
}

// Create or
// func (r *ServiceFunctionChainReconciler)
func (r *ServiceFunctionChainReconciler) ReconcilerDelete(ctx context.Context, sfcobjject *sfcv1.ServiceFunctionChain) (ctrl.Result, error) {
	controllerutil.RemoveFinalizer(sfcobjject, SFCFinalizer)
	return ctrl.Result{}, nil
}

func (r *ServiceFunctionChainReconciler) ReconcilerNormal(ctx context.Context, sfcobjject *sfcv1.ServiceFunctionChain) (ctrl.Result, error) {
	// Get or Create SFC

	return ctrl.Result{}, nil
}
func (r *ServiceFunctionChainReconciler) GetOrCreateSFCDeployment(ctx context.Context, sfcobjject *sfcv1.ServiceFunctionChain) (sfcv1.SFCDeploymentList, error) {
	var SFCDeploymentList sfcv1.SFCDeploymentList

	return SFCDeploymentList, nil
}

func (r *ServiceFunctionChainReconciler) GetSFCDeployment() {
	return
}
func (r *ServiceFunctionChainReconciler) DeployServiceFunctionDeployment(ctx context.Context, sfc *sfcv1.ServiceFunctionChain) error {

	for _, serviceDeployment := range sfc.Spec.Links {
		loggerSFC.Info("Start pick a location for Service Function Deployment")
		loggerSFC.Info(serviceDeployment.Metadata.Name)
	}
	return nil

}
func ResourceScoring(sdeployment sfcv1.SFCDeployment, serviceType int, logicalClusterGraph []int) error {

	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ServiceFunctionChainReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&sfcv1.ServiceFunctionChain{}).
		Complete(r)
}

func GetLogicalCLusterAvailableResource(serviceDeploymentName string, logicalCluster *intentv1.LogicalCluster) ([]ClusterResource, error) {
	var logicalClusterResource = []ClusterResource{}

	return logicalClusterResource, nil
}
