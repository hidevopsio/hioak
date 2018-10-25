package kube

import (
	"github.com/hidevopsio/hiboot/pkg/log"
	"github.com/magiconair/properties/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

func TestPodWatching(t *testing.T) {
	log.Debug("TestServiceDeletion()")

	projectName := "demo"
	profile := "dev"
	namespace := projectName + "-" + profile
	app := "hello-world"
	labels := map[string]string{
		"app": app,
	}
	clientSet := fake.NewSimpleClientset()
	client := NewPod(clientSet)

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      app,
			Namespace: namespace,
			Labels:    labels,
		},
	}
	p, err := client.Create(pod)
	log.Infof("pod :%v", p)
	assert.Equal(t, nil, err)
	listOptions := metav1.ListOptions{}
	i, err := client.Watch(listOptions, app, namespace)
	log.Infof("i: %v", i)
	assert.Equal(t, nil, err)
}
