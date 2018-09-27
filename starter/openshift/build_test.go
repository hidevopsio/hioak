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
	"github.com/openshift/api/build/v1"
	"github.com/openshift/client-go/build/clientset/versioned/fake"
	imageFake "github.com/openshift/client-go/image/clientset/versioned/fake"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func TestBuildCreation(t *testing.T) {
	log.Debug("TestBuildCreate()")

	// put below configs in yaml file
	project := "demo"
	profile := "stage"
	namespace := project + "-" + profile
	appName := "hello-world"
	scmUrl := os.Getenv("SCM_URL") + "/" + project + "/" + appName + ".git"
	scmRef := "master"
	secret := "test-secret"
	version := "v1"
	s2iImageStream := "s2i-java:latest"
	repoUrl := os.Getenv("MAVEN_MIRROR_URL")
	clientSet := fake.NewSimpleClientset().BuildV1()
	log.Debug(repoUrl)
	log.Debug(scmUrl)

	log.Debugf("workDir: %v", os.Getenv("PWD"))
	var err error
	var buildConfig *BuildConfig
	imageClient := imageFake.NewSimpleClientset().ImageV1()
	imageStream := newImageStream(imageClient)
	from, err := imageStream.CreateImageStream(namespace, appName, scmRef, s2iImageStream, true)
	t.Run("should create buildConfig instance", func(t *testing.T) {
		buildConfig = newBuildConfig(clientSet)
	})

	var bc *v1.BuildConfig
	t.Run("should create buildConfig", func(t *testing.T) {
		bc, err = buildConfig.Create(appName, namespace, scmUrl, scmRef, version, secret, from)
		assert.Equal(t, nil, err)
		assert.Equal(t, appName, bc.Name)
	})

	// Get build config
	bc, err = buildConfig.Get(appName, namespace)
	assert.Equal(t, nil, err)
	assert.Equal(t, appName, bc.Name)

	// Build image stream
	//_, err = buildConfig.Build(env)
	assert.Equal(t, nil, err)

}
