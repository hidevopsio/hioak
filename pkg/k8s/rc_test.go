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
	_, err := rc.Create(1)
	assert.Equal(t, nil, err)
}
