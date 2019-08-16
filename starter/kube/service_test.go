package kube

import (
	"github.com/stretchr/testify/assert"
	"hidevops.io/hiboot/pkg/log"
	"hidevops.io/hioak/starter"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func TestServiceCreation(t *testing.T) {
	log.Debug("TestServiceCreation()")
	name := "demo"
	projectName := "demo"
	profile := "dev"
	namespace := projectName + "-" + profile
	app := "hello-world"

	p := []orch.Ports{
		{
			Name: "8080-tcp",
			Port: 8080,
		},
		{
			Name: "7575-tcp",
			Port: 7575,
		},
	}
	clientSet := fake.NewSimpleClientset()
	service := NewService(clientSet)
	err := service.Create(app, name, namespace, p)
	assert.Equal(t, nil, err)

	svc, err := service.Get(app, namespace)
	assert.Equal(t, nil, err)
	assert.Equal(t, app, svc.Name)

	err = service.Delete(app, namespace)
	assert.Equal(t, nil, err)
}

func TestServiceCreate(t *testing.T) {
	log.Debug("TestServiceCreation()")

	projectName := "demo"
	profile := "dev"
	namespace := projectName + "-" + profile
	app := "hello-world"

	p := []corev1.ServicePort{
		{
			Name: "8080-tcp",
			Port: 8080,
		},
		{
			Name: "7575-tcp",
			Port: 7575,
		},
	}
	clientSet := fake.NewSimpleClientset()
	service := NewService(clientSet)
	err := service.CreateService(app, namespace, p)
	assert.Equal(t, nil, err)

	svc, err := service.Get(app, namespace)
	assert.Equal(t, nil, err)
	assert.Equal(t, app, svc.Name)

	_, err = service.List(namespace, metav1.ListOptions{})
	assert.Equal(t, nil, err)

	err = service.Delete(app, namespace)
	assert.Equal(t, nil, err)
}

func TestSvcCreate(t *testing.T) {
	log.Debug("TestServiceCreation()")
	p := []corev1.ServicePort{
		{
			Name: "8080-tcp",
			Port: 8080,
		},
		{
			Name: "7575-tcp",
			Port: 7575,
		},
	}
	projectName := "demo"
	profile := "dev"
	namespace := projectName + "-" + profile
	serviceSpec := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: projectName,
			Labels: map[string]string{
				"app":  projectName,
				"name": projectName,
			},
		},
		Spec: corev1.ServiceSpec{
			Type:  corev1.ServiceTypeClusterIP,
			Ports: p,
			Selector: map[string]string{
				"app": profile,
			},
		},
	}
	clientSet := fake.NewSimpleClientset()
	service := NewService(clientSet)
	_, err := service.CreateSvc(namespace, serviceSpec)
	assert.Equal(t, nil, err)

}
