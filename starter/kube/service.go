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
	"fmt"
	"github.com/jinzhu/copier"
	"hidevops.io/hiboot/pkg/log"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Service struct {
	clientSet kubernetes.Interface
}

func NewService(clientSet kubernetes.Interface) *Service {
	return &Service{
		clientSet: clientSet,
	}
}

func (s *Service) Create(name, namespace string, ports interface{}) error {

	p := make([]corev1.ServicePort, 0)
	copier.Copy(&p, ports)

	// create service
	serviceSpec := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Labels: map[string]string{
				"app": name,
			},
		},
		Spec: corev1.ServiceSpec{
			Type:  corev1.ServiceTypeClusterIP,
			Ports: p,
			Selector: map[string]string{
				"app": name,
			},
		},
	}

	svc, err := s.clientSet.CoreV1().Services(namespace).Get(name, metav1.GetOptions{})
	switch {
	case err == nil:
		serviceSpec.ObjectMeta.ResourceVersion = svc.ObjectMeta.ResourceVersion
		serviceSpec.Spec.ClusterIP = svc.Spec.ClusterIP
		_, err = s.clientSet.CoreV1().Services(namespace).Update(serviceSpec)
		if err != nil {
			return fmt.Errorf("failed to update service: %s", err)
		}
		log.Info("service updated")
	case errors.IsNotFound(err):
		_, err = s.clientSet.CoreV1().Services(namespace).Create(serviceSpec)
		if err != nil {
			return fmt.Errorf("failed to create service")
		}
		log.Info("service created")
	default:
		return fmt.Errorf("upexected error: %s", err)
	}
	return nil
}

func (s *Service) CreateService(name, namespace string, ports []corev1.ServicePort) error {

	// create service
	serviceSpec := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,

			Labels: map[string]string{
				"app": name,
			},
		},
		Spec: corev1.ServiceSpec{
			Type:  corev1.ServiceTypeClusterIP,
			Ports: ports,
			Selector: map[string]string{
				"app": name,
			},
		},
	}

	svc, err := s.clientSet.CoreV1().Services(namespace).Get(name, metav1.GetOptions{})
	switch {
	case err == nil:
		serviceSpec.ObjectMeta.ResourceVersion = svc.ObjectMeta.ResourceVersion
		serviceSpec.Spec.ClusterIP = svc.Spec.ClusterIP
		_, err = s.clientSet.CoreV1().Services(namespace).Update(serviceSpec)
		if err != nil {
			return fmt.Errorf("failed to update service: %s", err)
		}
		log.Info("service updated")
	case errors.IsNotFound(err):
		_, err = s.clientSet.CoreV1().Services(namespace).Create(serviceSpec)
		if err != nil {
			return fmt.Errorf("failed to create service")
		}
		log.Info("service created")
	default:
		return fmt.Errorf("upexected error: %s", err)
	}
	return nil
}

func (s *Service) Delete(name, namespace string) error {
	return s.clientSet.CoreV1().Services(namespace).Delete(name, &metav1.DeleteOptions{})
}

func (s *Service) Get(name, namespace string) (*corev1.Service, error) {
	return s.clientSet.CoreV1().Services(namespace).Get(name, metav1.GetOptions{})
}
