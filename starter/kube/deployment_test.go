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
	"github.com/magiconair/properties/assert"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

func TestExtensionsV1beta1Deploy(t *testing.T) {
	clientSet := fake.NewSimpleClientset()
	app := "hello-world"
	project := "demo-dev"
	dockerRegistry := "docker.vpclub.cn"
	imageTag := "v1"
	deploy := Deployment{
		clientSet: clientSet,
	}
	_, err := deploy.ExtensionsV1beta1Deploy(app, project, imageTag, dockerRegistry, nil, nil, nil, 0, false, "", "")
	assert.Equal(t, nil, err)
}

func TestDeploy(t *testing.T) {
	clientSet := fake.NewSimpleClientset()
	app := "hello-world"
	project := "demo-dev"
	dockerRegistry := "docker.vpclub.cn"
	imageTag := "v1"
	deploy := Deployment{
		clientSet: clientSet,
	}
	request := &DeployRequest{
		App:       app,
		Namespace: project,
		Version:imageTag,
		DockerRegistry:dockerRegistry,
		Replicas:int32Ptr(1),
	}
	_, err := deploy.Deploy(request)
	assert.Equal(t, nil, err)
}

//should return err is nil
func TestDeployment(t *testing.T) {
	clientSet := fake.NewSimpleClientset()
	deployDate := &DeployData{
		Name:           "hello-world",
		NameSpace:      "demo-dev",
		Replicas:       int32(1),
		Labels:         map[string]string{"app": "hello-world"},
		Image:          "demo:0.1",
		Ports:          []int{8080},
		Envs:           map[string]string{"ENVTEST": "ENVTEST"},
		HostPathVolume: map[string]string{"/var": "var"},
	}

	deploy := Deployment{
		clientSet: clientSet,
	}

	_, err := deploy.DeployNode(deployDate)
	assert.Equal(t, nil, err)

}

func TestDeleteDeployment(t *testing.T) {
	clientSet := fake.NewSimpleClientset()
	deploy := Deployment{
		clientSet: clientSet,
	}
	name := "hello-world"
	namespace := "demo-dev"
	option := &v1.DeleteOptions{}
	deployDate := &DeployData{
		Name:           name,
		NameSpace:      namespace,
		Replicas:       int32(1),
		Labels:         map[string]string{"app": "hello-world"},
		Image:          "demo:0.1",
		Ports:          []int{8080},
		Envs:           map[string]string{"ENVTEST": "ENVTEST"},
		HostPathVolume: map[string]string{"/var": "var"},
	}
	_, err := deploy.DeployNode(deployDate)
	assert.Equal(t, nil, err)
	err = deploy.Delete(name, namespace, option)
	assert.Equal(t, nil, err)

}
