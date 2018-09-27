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
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

func TestExtensionsV1beta1Deploy(t *testing.T) {
	clientSet := fake.NewSimpleClientset()
	app := "hello-world"
	project := "demo-dev"
	profile := "master"
	dockerRegistry := "docker.vpclub.cn"
	imageTag := "v1"
	deploy := Deployment{
		clientSet: clientSet,
	}
	_, err := deploy.ExtensionsV1beta1Deploy(app, project, profile, imageTag, dockerRegistry, nil, nil, nil, 0, false, "", "")
	assert.Equal(t, nil, err)
}

func TestDeploy(t *testing.T) {
	clientSet := fake.NewSimpleClientset()
	app := "hello-world"
	project := "demo-dev"
	profile := "master"
	dockerRegistry := "docker.vpclub.cn"
	imageTag := "v1"
	deploy := Deployment{
		clientSet: clientSet,
	}
	_, err := deploy.Deploy(app, project, profile, imageTag, dockerRegistry, nil, nil, nil, 0, false, "", "")
	assert.Equal(t, nil, err)
}
