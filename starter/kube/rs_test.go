package kube

import (
	"github.com/hidevopsio/hiboot/pkg/log"
	"github.com/magiconair/properties/assert"
	"k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

func TestReplicaSetCreate(t *testing.T) {
	log.Debug("TestReplicaSetCreate()")
	projectName := "demo"
	profile := "dev"
	namespace := projectName + "-" + profile
	app := "hello-world"
	clientSet := fake.NewSimpleClientset()
	rs := NewReplicaSet(clientSet)
	replica := &v1.ReplicaSet{
		ObjectMeta: metav1.ObjectMeta{
			Namespace:namespace,
			Name:app,
		},
	}
	_, err := rs.Create(replica)
	assert.Equal(t, nil, err)
	err = rs.Delete(app, namespace, &metav1.DeleteOptions{})
	assert.Equal(t, nil, err)
}
