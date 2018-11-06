package kube

import (
	"github.com/stretchr/testify/assert"
	"hidevops.io/hiboot/pkg/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func TestReplicationControllerWatching(t *testing.T) {
	log.Debug("TestServiceDeletion()")
	projectName := "demo"
	profile := "dev"
	namespace := projectName + "-" + profile
	app := "hello-world"
	clientSet := fake.NewSimpleClientset()
	rc := NewReplicationController(clientSet)
	_, err := rc.Create(app, namespace, 1)
	assert.Equal(t, nil, err)
}

func TestReplicationControllerDelete(t *testing.T) {
	log.Debug("TestServiceDeletion()")
	projectName := "demo"
	profile := "dev"
	namespace := projectName + "-" + profile
	app := "hello-world"
	clientSet := fake.NewSimpleClientset()
	rc := NewReplicationController(clientSet)
	_, err := rc.Create(app, namespace, 1)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, err)
	option := &metav1.DeleteOptions{}
	err = rc.Delete(app, namespace, option)
	assert.Equal(t, nil, err)
}
