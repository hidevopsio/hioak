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

package openshift

import (
	"github.com/openshift/api/project/v1"
	projectv1 "github.com/openshift/client-go/project/clientset/versioned/typed/project/v1"
	"hidevops.io/hiboot/pkg/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const NodeSelector = "openshift.io/node-selector"

type Project struct {
	clientSet projectv1.ProjectV1Interface
}

func newProject(clientSet projectv1.ProjectV1Interface) *Project {

	return &Project{
		clientSet: clientSet,
	}
}

func (p *Project) Create(name, nodeSelector string) (*v1.Project, error) {
	log.Debug("Project.Create()")
	annotations := map[string]string{
		NodeSelector: nodeSelector,
	}
	ps := &v1.Project{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Labels: map[string]string{
				"project": name,
			},
			Annotations: annotations,
		},
	}
	project, err := p.Get(name)
	if err == nil {
		return project, err
	}
	// create project
	return p.clientSet.Projects().Create(ps)
}

func (p *Project) Get(name string) (*v1.Project, error) {
	log.Debug("Project.Get()")
	return p.clientSet.Projects().Get(name, metav1.GetOptions{})
}

func (p *Project) List() (*v1.ProjectList, error) {
	log.Debug("Project.List()")
	return p.clientSet.Projects().List(metav1.ListOptions{})
}

func (p *Project) Delete(name string) error {
	log.Debug("Project.Delete()")
	return p.clientSet.Projects().Delete(name, &metav1.DeleteOptions{})
}
