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
	"testing"
	"github.com/magiconair/properties/assert"
	"k8s.io/client-go/kubernetes/fake"
)


func TestDeploy(t *testing.T)  {
	clientSet := fake.NewSimpleClientset()
	deploy := Deployment{
		App: "hello-world",
		Project: "demo-dev",
		Profile: "master",
		DockerRegistry: "docker.vpclub.cn",
		ImageTag: "v1",
		clientSet:clientSet,

	}
	_, err := deploy.ExtensionsV1beta1Deploy(nil, nil, nil, 0, false, "", "")
	assert.Equal(t, nil, err)
}
