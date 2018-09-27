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
	"github.com/openshift/client-go/image/clientset/versioned/fake"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImageStreamCrud(t *testing.T) {
	imageStreamName := "is-test"
	namespace := "openshift"
	clientSet := fake.NewSimpleClientset().ImageV1()
	imageStream := newImageStream(clientSet)
	version := "v1"
	source := "docker.io/hidevops/s2i-java:latest"
	// create imagestream
	is, err := imageStream.Create(imageStreamName, namespace, source, version)
	assert.Equal(t, nil, err)
	assert.Equal(t, imageStreamName, is.ObjectMeta.Name)

	// get imagestream

	is, err = imageStream.Get(imageStreamName, namespace)
	assert.Equal(t, nil, err)
	assert.Equal(t, imageStreamName, is.ObjectMeta.Name)

	// delete imagestream
	err = imageStream.Delete(imageStreamName, namespace)
	assert.Equal(t, nil, err)
}
