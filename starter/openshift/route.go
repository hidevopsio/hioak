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
	"fmt"
	"github.com/openshift/api/route/v1"
	routev1 "github.com/openshift/client-go/route/clientset/versioned/typed/route/v1"
	"hidevops.io/hiboot/pkg/log"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type Route struct {
	clientSet routev1.RouteV1Interface
}

func newRoute(clientSet routev1.RouteV1Interface) *Route {
	log.Debug("NewRoute()")
	return &Route{
		clientSet: clientSet,
	}
}

func (r *Route) Create(name, namespace string, port int32) (string, error) {
	log.Debug("Route.Create()")
	upstreamUrl := ""
	cfg := &v1.Route{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Labels: map[string]string{
				"app": name,
			},
		},
		Spec: v1.RouteSpec{
			To: v1.RouteTargetReference{
				Kind: "Service",
				Name: name,
			},
			Port: &v1.RoutePort{
				TargetPort: intstr.IntOrString{
					IntVal: port,
				},
			},
		},
	}

	result, err := r.clientSet.Routes(namespace).Get(name, metav1.GetOptions{})
	switch {
	case err == nil:
		cfg.ObjectMeta.ResourceVersion = result.ResourceVersion
		result, err = r.clientSet.Routes(namespace).Update(cfg)
		if err == nil {
			upstreamUrl = result.Spec.Host
			log.Infof("Updated Route %v", result.Name)
		} else {
			return upstreamUrl, err
		}
		break
	case errors.IsNotFound(err):
		route, err := r.clientSet.Routes(namespace).Create(cfg)
		if err != nil {
			return upstreamUrl, err
		}
		upstreamUrl = route.Spec.Host
		log.Infof("Created Route %q.\n", route.Name)
		break
	default:
		return upstreamUrl, fmt.Errorf("failed to create Route: %s", err)
	}
	return upstreamUrl, nil
}

func (r *Route) Get(name, namespace string) (*v1.Route, error) {
	log.Debug("Route.get()")

	return r.clientSet.Routes(namespace).Get(name, metav1.GetOptions{})
}

func (r *Route) Delete(name, namespace string) error {
	log.Debug("Route.Delete()")

	return r.clientSet.Routes(namespace).Delete(name, &metav1.DeleteOptions{})
}
