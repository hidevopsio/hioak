// Copyright 2018 John Deng (hi.devops.io@gmail.com).
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kube

import (
	"encoding/json"
	"fmt"
	"github.com/hidevopsio/hiboot/pkg/log"
	"github.com/hidevopsio/hiboot/pkg/utils/copier"
	"k8s.io/api/apps/v1beta1"
	corev1 "k8s.io/api/core/v1"
	extensionsV1beta1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"strings"
)

type Deployment struct {
	clientSet kubernetes.Interface
}

func int32Ptr(i int32) *int32 { return &i }

func newDeployment(clientSet kubernetes.Interface) *Deployment {
	return &Deployment{
		clientSet: clientSet,
	}
}

// @Title Deploy
// @Description deploy application
// @Param pipeline
// @Return error
func (d *Deployment) Deploy(app, project, profile, imageTag, dockerRegistry string, env interface{}, labels map[string]string, ports interface{}, replicas int32, force bool, healthEndPoint, nodeSelector string) (string, error) {

	log.Debug("Deployment.Deploy()")
	e := make([]corev1.EnvVar, 0)
	copier.Copy(&e, env)
	selector := map[string]string{}
	if nodeSelector != "" {
		selector[strings.Split(nodeSelector, "=")[0]] = strings.Split(nodeSelector, "=")[1]
	}
	p := make([]corev1.ContainerPort, 0)
	copier.Copy(&p, ports)
	deploySpec := &v1beta1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      app,
			Namespace: project,
		},
		Spec: v1beta1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Strategy: v1beta1.DeploymentStrategy{
				Type: v1beta1.RollingUpdateDeploymentStrategyType,
				RollingUpdate: &v1beta1.RollingUpdateDeployment{
					MaxUnavailable: &intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: int32(0),
					},
					MaxSurge: &intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: int32(1),
					},
				},
			},
			RevisionHistoryLimit: int32Ptr(10),
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:   app,
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            app,
							Image:           dockerRegistry + "/" + project + "/" + app + ":" + imageTag,
							Ports:           p,
							Env:             e,
							ImagePullPolicy: corev1.PullAlways,
						},
					},
				},
			},
		},
	}
	log.Debug(deploySpec)
	j, err := json.Marshal(deploySpec)
	log.Debug("json", string(j))
	// Create Deployment
	//Client.ClientSet.ExtensionsV1beta1().Deployments()
	deployments := d.clientSet.AppsV1beta1().Deployments(project)
	log.Info("Update or Create Deployment...")
	result, err := deployments.Update(deploySpec)
	var retVal string
	switch {
	case err == nil:
		log.Info("Deployment updated")
	case err != nil:
		_, err = deployments.Create(deploySpec)
		retVal = fmt.Sprintf("Created deployment %q.\n", result)
		log.Info("retval:", err)
	default:
		return retVal, fmt.Errorf("could not update deployment controller: %s", err)
	}

	return retVal, err
}

func (d *Deployment) ExtensionsV1beta1Deploy(app, project, profile, imageTag, dockerRegistry string, env interface{}, labels map[string]string, ports interface{}, replicas int32, force bool, healthEndPoint, nodeSelector string) (string, error) {

	log.Debug("Deployment.Deploy()")
	e := make([]corev1.EnvVar, 0)
	copier.Copy(&e, env)
	selector := map[string]string{}
	if nodeSelector != "" {
		selector[strings.Split(nodeSelector, "=")[0]] = strings.Split(nodeSelector, "=")[1]
	}
	p := make([]corev1.ContainerPort, 0)
	copier.Copy(&p, ports)
	deploySpec := &extensionsV1beta1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      app,
			Namespace: project,
		},
		Spec: extensionsV1beta1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Strategy: extensionsV1beta1.DeploymentStrategy{
				Type: extensionsV1beta1.RollingUpdateDeploymentStrategyType,
				RollingUpdate: &extensionsV1beta1.RollingUpdateDeployment{
					MaxUnavailable: &intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: int32(0),
					},
					MaxSurge: &intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: int32(1),
					},
				},
			},
			RevisionHistoryLimit: int32Ptr(10),
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:   app,
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            app,
							Image:           dockerRegistry + "/" + project + "/" + app + ":" + imageTag,
							Ports:           p,
							Env:             e,
							ImagePullPolicy: corev1.PullAlways,
						},
					},
				},
			},
		},
	}
	log.Debug(deploySpec)
	j, err := json.Marshal(deploySpec)
	log.Debug("json", string(j))
	// Create Deployment
	//Client.ClientSet.ExtensionsV1beta1().Deployments()
	deployments := d.clientSet.ExtensionsV1beta1().Deployments(project)
	log.Info("Update or Create Deployment...")
	result, err := deployments.Update(deploySpec)
	var retVal string
	switch {
	case err == nil:
		log.Info("Deployment updated")
	case err != nil:
		_, err = deployments.Create(deploySpec)
		retVal = fmt.Sprintf("Created deployment %q.\n", result)
		log.Info("retval:", err)
	default:
		return retVal, fmt.Errorf("could not update deployment controller: %s", err)
	}

	return retVal, err
}

func (d *Deployment) DeployNode(app, project, dockerRegistry, imageTag, profile string) (string, error) {
	log.Debug("Deployment.Deploy()")
	deploySpec := &v1beta1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      app,
			Namespace: project,
		},
		Spec: v1beta1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Strategy: v1beta1.DeploymentStrategy{
				Type: v1beta1.RollingUpdateDeploymentStrategyType,
				RollingUpdate: &v1beta1.RollingUpdateDeployment{
					MaxUnavailable: &intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: int32(0),
					},
					MaxSurge: &intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: int32(1),
					},
				},
			},
			RevisionHistoryLimit: int32Ptr(10),
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name: app,
					Labels: map[string]string{
						"app": app,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  app,
							Image: dockerRegistry + "/" + project + "/" + app + ":" + imageTag,
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									Protocol:      corev1.ProtocolTCP,
									ContainerPort: 7575,
								},
							},
							Env: []corev1.EnvVar{
								{
									Name:  "APP_PROFILES_ACTIVE",
									Value: profile,
								},
							},
							ImagePullPolicy: corev1.PullIfNotPresent,
						},
					},
				},
			},
		},
	}
	log.Debug(deploySpec)
	j, err := json.Marshal(deploySpec)
	log.Debug("json", string(j))
	// Create Deployment
	deployments := d.clientSet.AppsV1beta1().Deployments(project)
	log.Info("Update or Create Deployment...")
	result, err := deployments.Update(deploySpec)
	var retVal string
	switch {
	case err == nil:
		log.Info("Deployment updated")
	case err != nil:
		_, err = deployments.Create(deploySpec)
		retVal = fmt.Sprintf("Created deployment %q.\n", result)
		log.Info("retval:", err)
	default:
		return retVal, fmt.Errorf("could not update deployment controller: %s", err)
	}

	return retVal, err
}
