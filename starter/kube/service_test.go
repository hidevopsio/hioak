package kube

import (
	"github.com/hidevopsio/hiboot/pkg/log"
	"github.com/hidevopsio/hioak/starter"
	"github.com/stretchr/testify/assert"
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
