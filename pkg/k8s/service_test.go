package k8s

import (
	"github.com/hidevopsio/hiboot/pkg/log"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/hidevopsio/hioak/pkg"
	"k8s.io/client-go/kubernetes/fake"
)

func init()  {
	log.SetLevel(log.DebugLevel)
}

func TestServiceCreation(t *testing.T) {
	log.Debug("TestServiceCreation()")

	projectName := "demo"
	profile     := "dev"
	namespace   := projectName + "-" + profile
	app         := "hello-world"

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
	service := NewService(clientSet, app, namespace)
	err := service.Create(p)
	assert.Equal(t, nil, err)

	svc, err := service.Get()
	assert.Equal(t, nil, err)
	assert.Equal(t, app, svc.Name)

	err = service.Delete()
	assert.Equal(t, nil, err)
}


