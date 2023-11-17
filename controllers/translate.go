package controllers

import (
	"context"
	"fmt"

	sfcv1 "github.com/ntnguyencse/Service-Chain-Orchestrator/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	yaml "sigs.k8s.io/yaml"
)

func (r *ServiceFunctionChainReconciler) TranslateSFCtoServiceMeshService(ctx context.Context, req ctrl.Request, sfc sfcv1.LinkService) (string, error) {
	var serviceString string
	SFCServiceName := sfc.Service.ServiceRef.Name
	SFCService, err := r.GetSFCServiceByName(ctx, SFCServiceName)
	if err != nil {
		loggerSFC.Error(err, "Error when get SFC Deployment")
	} else {
		service := &corev1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name:      SFCService.Name,
				Namespace: SFCService.Namespace,
				Labels:    SFCService.Labels,
			},
			Spec: corev1.ServiceSpec{
				Ports:    SFCService.Spec.Ports,
				Selector: SFCService.Spec.Selector,
			},
		}
		serviceByte, _ := yaml.Marshal(service)
		serviceString = "apiVersion: v1\nkind: Service\n" + string(serviceByte)
		return serviceString, nil

	}

	return serviceString, err
}
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
			Labels: map[string]string{
				"app": sfc.Deployment.Name,
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: sfcSpec.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: sfcSpec.Selector.MatchLabels,
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: sfcSpec.Selector.MatchLabels,
				},
				Spec: sfcSpec.Template.Spec,
			},
		},
	}
	deploymentByte, _ := yaml.Marshal(deployment)
	deploymentString = "apiVersion: apps/v1\nkind: Deployment\n" + string(deploymentByte)
	return deploymentString, nil
}

func (r *ServiceFunctionChainReconciler) GetSFCServiceByName(ctx context.Context, deploymentName string) (sfcv1.SFCService, error) {
	var serviceList sfcv1.SFCServiceList
	var service sfcv1.SFCService

	if err := r.Client.List(ctx, &serviceList); err != nil {
		if apierrors.IsNotFound(err) {
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			loggerSFC.Info("Could not list Service")
			return service, err
		}

		// Error reading the object - requeue the request.
		return service, err
	}
	if len(serviceList.Items) < 1 {
		loggerSFC.Info("List SFC Service is Empty")
	}
	for _, depl := range serviceList.Items {
		if depl.Name == deploymentName {
			service = depl
			return service, nil
		}
	}

	err1 := fmt.Errorf("Not found SF Service")
	return service, err1

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
