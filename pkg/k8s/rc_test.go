package k8s

import (
	"testing"
	"github.com/hidevopsio/hiboot/pkg/log"
	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/kubernetes/fake"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func TestReplicationControllerWatching(t *testing.T) {
	log.Debug("TestServiceDeletion()")

	projectName := "demo"
	profile     := "dev"
	namespace   := projectName + "-" + profile
	app         := "hello-world"
	clientSet := fake.NewSimpleClientset()
	rc := NewReplicationController(clientSet, app, namespace)
	go func() {
		err := rc.Watch(func() error {
			log.Debug("Completed!")
			return nil
		})
		assert.Equal(t, nil, err)

	}()
	assert.Equal(t, app, rc.Name)
}
