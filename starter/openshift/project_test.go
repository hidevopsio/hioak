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
	"github.com/hidevopsio/hiboot/pkg/log"
	"github.com/openshift/client-go/project/clientset/versioned/fake"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProjectLit(t *testing.T) {
	clientSet := fake.NewSimpleClientset().ProjectV1()
	project := newProject(clientSet)
	pl, err := project.List()
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, len(pl.Items))
	log.Debugf("There are %d projects in the cluster", len(pl.Items))

	for i, p := range pl.Items {
		log.Debugf("index %d: project: %s", i, p.Name)
	}
}

func TestProjectCrud(t *testing.T) {
	projectName := "project-crud"
	fake.NewSimpleClientset().ProjectV1()
	clientSet := fake.NewSimpleClientset().ProjectV1()
	project := newProject(clientSet)
	// create project
	p, err := project.Create(projectName, projectName)
	assert.Equal(t, nil, err)
	assert.Equal(t, projectName, p.Name)

	// read project
	p, err = project.Get(projectName)
	assert.Equal(t, nil, err)
	assert.Equal(t, projectName, p.Name)

	// delete project
	err = project.Delete(projectName)
	assert.Equal(t, nil, err)

}
