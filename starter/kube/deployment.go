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
	"hidevops.io/hiboot/pkg/log"
	"hidevops.io/hiboot/pkg/utils/copier"
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

// NewDeployment ConfigMaps initialize construction
func NewDeployment(clientSet kubernetes.Interface) *Deployment {
	return &Deployment{
		clientSet: clientSet,
	}
}

type DeployRequest struct {
	App            string
	Namespace      string
	Version        string
	Tag            string
	DockerRegistry string
	Env            []corev1.EnvVar
	Labels         map[string]string
	Ports          []corev1.ContainerPort
	Replicas       *int32
	NodeSelector   map[string]string
	ReadinessProbe *corev1.Probe
	LivenessProbe  *corev1.Probe
	Volumes        []corev1.Volume
	VolumeMounts   []corev1.VolumeMount
}

// @Title Deploy
// @Description deploy application
// @Param pipeline
// @Return error
func (d *Deployment) Deploy(request *DeployRequest) (*extensionsV1beta1.Deployment, error) {
	runAsRoot := false
	log.Debug("Deployment.Deploy()")
	deploySpec := &extensionsV1beta1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-%s", request.App, request.Version),
			Namespace: request.Namespace,
			Labels: map[string]string{
				"app":     request.App,
				"version": request.Version,
			},
		},
		Spec: extensionsV1beta1.DeploymentSpec{
			Replicas: request.Replicas,
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
					Name: request.App,
					Labels: map[string]string{
						"app":     request.App,
						"version": request.Version,
					},
				},
				Spec: corev1.PodSpec{
					SecurityContext: &corev1.PodSecurityContext{
						RunAsNonRoot: &runAsRoot,
					},
					NodeSelector: request.NodeSelector,
					Volumes:      request.Volumes,
					Containers: []corev1.Container{
						{
							Name:            request.App,
							Image:           request.DockerRegistry + "/" + request.Namespace + "/" + request.App + ":" + request.Tag,
							Ports:           request.Ports,
							Env:             request.Env,
							ImagePullPolicy: corev1.PullAlways,
							ReadinessProbe:  request.ReadinessProbe,
							LivenessProbe:   request.LivenessProbe,
							VolumeMounts:    request.VolumeMounts,
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
	deployments := d.clientSet.ExtensionsV1beta1().Deployments(request.Namespace)
	log.Info("Update or Create Deployment...")
	result, err := deployments.Update(deploySpec)
	switch {
	case err == nil:
		log.Info("Deployment updated")
	case err != nil:
		result, err = deployments.Create(deploySpec)
		log.Info("deploy: ", err)
	default:
		return result, fmt.Errorf("could not update deployment controller: %s", err)
	}

	return result, err
}

func (d *Deployment) ExtensionsV1beta1Deploy(app, project, imageTag, dockerRegistry string, env interface{}, labels map[string]string, ports interface{}, replicas int32, force bool, healthEndPoint, nodeSelector string) (string, error) {

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
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":     app,
					"version": imageTag,
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

type DeployData struct {
	Name           string
	NameSpace      string
	Replicas       int32
	Labels         map[string]string
	Image          string
	Ports          []int
	Envs           map[string]string
	HostPathVolume map[string]string
	NodeSelector   map[string]string
	NodeName       string
}

func (d *Deployment) DeployNode(deployData *DeployData) (string, error) {
	log.Debug("Deployment.Deploy()")

	//port
	var containerPorts []corev1.ContainerPort
	for _, port := range deployData.Ports {
		containerPorts = append(containerPorts, corev1.ContainerPort{
			Name:          fmt.Sprintf("http-%d", port),
			Protocol:      corev1.ProtocolTCP,
			ContainerPort: int32(port),
		})
	}

	//env
	var envs []corev1.EnvVar
	for k, v := range deployData.Envs {
		envs = append(envs, corev1.EnvVar{
			Name:  k,
			Value: v,
		})
	}

	//volume
	var Volumes []corev1.Volume
	var VolumeMounts []corev1.VolumeMount
	i := 0
	for k, v := range deployData.HostPathVolume {
		i++
		volumeName := fmt.Sprintf("volume%d", i)
		Volumes = append(Volumes, corev1.Volume{
			Name: volumeName,
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: k,
				},
			},
		})

		VolumeMounts = append(VolumeMounts, corev1.VolumeMount{
			Name:      volumeName,
			MountPath: v,
		})
	}

	deploySpec := &extensionsV1beta1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "extensions/v1beta1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployData.Name,
			Namespace: deployData.NameSpace,
		},
		Spec: extensionsV1beta1.DeploymentSpec{
			Replicas: int32Ptr(deployData.Replicas),
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
					Name:   deployData.Name,
					Labels: deployData.Labels,
				},
				Spec: corev1.PodSpec{
					NodeSelector: deployData.NodeSelector,
					Containers: []corev1.Container{
						{
							Name:            deployData.Name,
							Image:           deployData.Image,
							Ports:           containerPorts,
							Env:             envs,
							ImagePullPolicy: corev1.PullIfNotPresent,
							VolumeMounts:    VolumeMounts,
						},
					},
					Volumes: Volumes,
				},
			},
		},
	}
	if deployData.NodeName != "" {
		deploySpec.Spec.Template.Spec.NodeName = deployData.NodeName
	}
	// Create Deployment
	deployment, err := d.clientSet.ExtensionsV1beta1().Deployments(deployData.NameSpace).Create(deploySpec)
	if err != nil {
		return "", err
	}
	deploymentJson, _ := json.Marshal(deployment)
	return string(deploymentJson), nil
}

func (d *Deployment) Delete(name, namespace string, option *metav1.DeleteOptions) error {
	log.Debugf("delete deployment name :%v, namespace :%v", name, namespace)
	err := d.clientSet.ExtensionsV1beta1().Deployments(namespace).Delete(name, option)
	return err
}

func (d *Deployment) Update(deployment *extensionsV1beta1.Deployment) error {
	log.Debugf("update deployment name :%v, namespace :%v", deployment.Name, deployment.Namespace)
	_, err := d.clientSet.ExtensionsV1beta1().Deployments(deployment.Namespace).Update(deployment)
	return err
}

func (d *Deployment) Get(name, namespace string, option metav1.GetOptions) (*extensionsV1beta1.Deployment, error) {
	log.Debugf("get deployment name :%v, namespace :%v", name, namespace)
	deploy, err := d.clientSet.ExtensionsV1beta1().Deployments(namespace).Get(name, option)
	return deploy, err
}

func (d *Deployment) List(namespace string, option metav1.ListOptions) (*extensionsV1beta1.DeploymentList, error) {
	log.Debugf("get deployment namespace :%v", namespace)
	deploys, err := d.clientSet.ExtensionsV1beta1().Deployments(namespace).List(option)
	return deploys, err
}