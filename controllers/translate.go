package controllers

import (
	"context"
	"fmt"

	sfcv1 "github.com/ntnguyencse/Service-Chain-Orchestrator/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
)

func (r *SFCDeploymentReconciler) TranslateSFCtoServiceMeshDeployment(ctx context.Context, req ctrl.Request, sfc sfcv1.LinkService) (string, error) {
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
			Name:      SFCDeploymentName,
			Namespace: sfc.Deployment.Namespace,
		},
		Spec: appsv1.DeploymentSpec(sfcSpec),
	}
	deploymentString = deployment.String()
	return deploymentString, nil
}

func (r *SFCDeploymentReconciler) GetSFCDeploymentByName(ctx context.Context, deploymentName string) (sfcv1.SFCDeployment, error) {
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
	for _, depl := range deploymentList.Items {
		if depl.Name == deploymentName {
			deployment = depl
		}
	}
	if len(deployment.Name) < 2 {
		loggerSFC.Info("Could not list SF deployment")
		err1 := fmt.Errorf("Not found SF Deployment")
		return deployment, err1
	}

	return deployment, nil
}

// func CreateFolderOnGit()
