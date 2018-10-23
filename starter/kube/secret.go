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
	"github.com/hidevopsio/hiboot/pkg/log"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/typed/core/v1"
)

type Secret struct {
	clientSet kubernetes.Interface
	Interface v1.SecretInterface
}

// Create new instance of type Secret
func NewSecret(clientSet kubernetes.Interface) *Secret {
	s := &Secret{
		clientSet: clientSet,
	}
	return s
}

// Create takes the representation of a secret and creates it.  Returns the server's representation of the secret, and an error, if there is any.
func (s *Secret) Create(username, password, token, name, namespace string) error {
	log.Debug("Secret.Create()")
	var data map[string][]byte
	if username != "" {
		data = map[string][]byte{
			corev1.BasicAuthUsernameKey:   []byte(username),
			corev1.BasicAuthPasswordKey:   []byte(password),
			corev1.ServiceAccountTokenKey: []byte(token),
		}
	} else {
		data = map[string][]byte{
			corev1.BasicAuthPasswordKey:   []byte(password),
			corev1.ServiceAccountTokenKey: []byte(token),
		}
	}
	// k8s.io/api/core/v1/types.go
	coreSecret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Data: data,
		Type: corev1.SecretTypeBasicAuth,
	}
	var err error

	_, err = s.Get(name, namespace)
	if errors.IsNotFound(err) {
		_, err = s.clientSet.CoreV1().Secrets(namespace).Create(coreSecret)
	} else {
		_, err = s.clientSet.CoreV1().Secrets(namespace).Update(coreSecret)
	}

	return err
}

// Get takes name of the secret, and returns the corresponding secret object, and an error if there is any.
func (s *Secret) Get(name, namespace string) (*corev1.Secret, error) {
	log.Debug("Secret.Get()")
	var err error
	secret, err := s.clientSet.CoreV1().Secrets(namespace).Get(name, metav1.GetOptions{})

	return secret, err
}

// Delete takes name of the secret and deletes it. Returns an error if one occurs.
func (s *Secret) Delete(name, namespace string) error {
	log.Debug("Secret.Delete()")
	var err error
	err = s.clientSet.CoreV1().Secrets(namespace).Delete(name, &metav1.DeleteOptions{})

	return err
}
