package kube

import (
	"github.com/magiconair/properties/assert"
	"hidevops.io/hiboot/pkg/log"
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
			Namespace: namespace,
			Name:      app,
			Labels: map[string]string{
				"app": app,
			},
		},
	}
	_, err := rs.Create(replica)
	assert.Equal(t, nil, err)
	option := metav1.ListOptions{
		LabelSelector: "app=" + app,
	}
	_, err = rs.List(app, namespace, option)
	assert.Equal(t, nil, err)

	err = rs.Delete(app, namespace, &metav1.DeleteOptions{})
	assert.Equal(t, nil, err)
}
