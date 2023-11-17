package controllers

import (
	"context"
	"fmt"

	sfcv1 "github.com/ntnguyencse/Service-Chain-Orchestrator/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	yaml "sigs.k8s.io/yaml"
)

func (r *ServiceFunctionChainReconciler) TranslateSFCtoServiceMeshDeployment(ctx context.Context, req ctrl.Request, sfc sfcv1.LinkService) (string, error) {
	var deploymentString string
	SFCDeploymentName := sfc.Deployment.Name
	SFCDeployment, err := r.GetSFCDeploymentByName(ctx, SFCDeploymentName)
	if err != nil {
		loggerSFC.Error(err, "Error when get SFC Deployment")
	}
	// Translate LinkService to SFC Deployment
	sfcSpec := SFCDeployment.DeepCopy().Spec
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sfc.Deployment.Name,
			Namespace: sfc.Deployment.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: sfcSpec.Replicas,
			Selector: sfcSpec.Selector,
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: sfcSpec.Template.Labels,
				},
				Spec: sfcSpec.Template.Spec,
			},
		},
	}
	deploymentByte, _ := yaml.Marshal(deployment)
	deploymentString = string(deploymentByte)
	return deploymentString, nil
}

func (r *ServiceFunctionChainReconciler) GetSFCDeploymentByName(ctx context.Context, deploymentName string) (sfcv1.SFCDeployment, error) {
	var deploymentList sfcv1.SFCDeploymentList
	var deployment sfcv1.SFCDeployment

	if err := r.Client.List(ctx, &deploymentList); err != nil {
		if apierrors.IsNotFound(err) {
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			loggerSFC.Info("Could not list SF deployment")
			return deployment, err
		}

		// Error reading the object - requeue the request.
		return deployment, err
	}
	if len(deploymentList.Items) < 1 {
		loggerSFC.Info("List SFC Deployment is Empty")
	}
	for _, depl := range deploymentList.Items {
		if depl.Name == deploymentName {
			deployment = depl
			return deployment, nil
		}
	}

	err1 := fmt.Errorf("Not found SF Deployment")
	return deployment, err1

}

// func CreateFolderOnGit()
