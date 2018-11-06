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
	"github.com/openshift/api/image/v1"
	imagev1 "github.com/openshift/client-go/image/clientset/versioned/typed/image/v1"
	"hidevops.io/hiboot/pkg/log"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ImageStreamInterface interface {
	Create() (*v1.ImageStream, error)
	Get() (*v1.ImageStream, error)
	Delete() error
}

type ImageStream struct {
	clientSet imagev1.ImageV1Interface
}

func newImageStream(clientSet imagev1.ImageV1Interface) *ImageStream {
	return &ImageStream{
		clientSet: clientSet,
	}
}

// @Title NewBuildConfig
// @Description Create new BuildConfig Instance
// @Param namespace, appName, gitUrl, imageTag, s2iImageStream string
// @Return *BuildConfig, error
func (is *ImageStream) CreateImageStream(namespace, name, scmRef, s2iImageStream string, rebuild bool) (*corev1.ObjectReference, error) {

	log.Debug("NewBuildConfig()")

	// TODO: for the sake of decoupling, the image stream creation should be here or not?
	var err error
	var from corev1.ObjectReference
	if !rebuild {
		image, err := is.Get(name, namespace)
		if err != nil {
			return nil, err
		}
		// the images stream is exist with 0 tags, then delete it
		if len(image.Status.Tags) == 0 {
			is.Delete(name, namespace)
			_, err = is.Get(name, namespace)
		}
	}
	from = corev1.ObjectReference{
		Kind:      "ImageStreamTag",
		Name:      s2iImageStream,
		Namespace: "openshift",
	}
	return &from, err
}

// @Title Create
// @Description create imagestream
// @Param
// @Return v1.ImageStream, error
func (is *ImageStream) Create(name, namespace, source, version string) (*v1.ImageStream, error) {
	log.Debug("ImageStream.Create()")

	imageStream := &v1.ImageStream{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels: map[string]string{
				"app":     name,
				"version": version,
			},
		},
	}

	if source != "" {
		imageStream.Spec = v1.ImageStreamSpec{
			Tags: []v1.TagReference{
				{
					Name: version,
					From: &corev1.ObjectReference{
						Kind: "DockerImage",
						Name: source,
					},
				},
			},
		}
	}

	result, err := is.Get(name, namespace)
	message := "create ImageStream"
	switch {
	case err == nil:
		imageStream.ObjectMeta.ResourceVersion = result.ResourceVersion
		result, err = is.clientSet.ImageStreams(namespace).Update(imageStream)
		message = "update ImageStream"

	case errors.IsNotFound(err):
		result, err = is.clientSet.ImageStreams(namespace).Create(imageStream)
	}

	if err != nil {
		log.Errorf("Failed to %v %v.", message, result.Name)
		return nil, err
	}
	log.Infof("Succeed to %v %v.", message, result.Name)
	return result, err
}

// @Title Get
// @Description get imagestream
// @Param
// @Return v1.ImageStream, error
func (is *ImageStream) Get(name, namespace string) (*v1.ImageStream, error) {
	log.Debug("ImageStream.Get()")
	return is.clientSet.ImageStreams(namespace).Get(name, metav1.GetOptions{})
}

// @Title Delete
// @Description delete imagestream
// @Param
// @Return error
func (is *ImageStream) Delete(name, namespace string) error {
	log.Debug("ImageStream.Delete()")
	return is.clientSet.ImageStreams(namespace).Delete(name, &metav1.DeleteOptions{})
}
