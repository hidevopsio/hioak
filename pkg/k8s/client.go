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


package k8s


import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"
	"github.com/hidevopsio/hioak/pkg"
)


func NewClientSet() kubernetes.Interface {

	cli := orch.GetClientInstance()

	// get the fake ClientSet for testing
	if cli.IsTestRunning() {
		return fake.NewSimpleClientset()
	}

	// get the real ClientSet
	clientSet, err := kubernetes.NewForConfig(cli.Config())
	if err != nil {
		panic(err.Error())
	}
	return clientSet
}
