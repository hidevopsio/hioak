package kube

import (
	"github.com/stretchr/testify/assert"
	"hidevops.io/hiboot/pkg/log"
	"hidevops.io/hioak/starter"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func TestServiceCreation(t *testing.T) {
	log.Debug("TestServiceCreation()")

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
	err := service.Create(app, namespace, p)
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

	err = service.Delete(app, namespace)
	assert.Equal(t, nil, err)
}
