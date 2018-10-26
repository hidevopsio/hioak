package kube

import (
	"github.com/hidevopsio/hiboot/pkg/log"
	"github.com/stretchr/testify/assert"
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

	t.Run("Pods should watch succeed", func(t *testing.T) {
		listOptions := metav1.ListOptions{}
		i, err := client.Watch(listOptions, app, namespace)
		log.Infof("i: %v", i)
		assert.Equal(t, nil, err)

	})

	t.Run("pods should list succeed", func(t *testing.T) {
		_,err:= client.GetPodList(namespace,metav1.ListOptions{})
		assert.Equal(t, nil, err)
	})


	t.Run("pod should get succeed", func(t *testing.T) {
		_,err:= client.GetPods(namespace,app,metav1.GetOptions{})
		assert.Equal(t, nil, err)
	})

	t.Run("pod should get failed", func(t *testing.T) {
		_,err:= client.GetPods(namespace,projectName,metav1.GetOptions{})
		assert.NotEqual(t, nil, err)
	})

	t.Run("Pods should get logs succeed", func(t *testing.T) {
		_,err:= client.GetPodLogs(namespace,app,&corev1.PodLogOptions{})
		assert.Equal(t, nil, err)
	})

	t.Run("Pods should get logs failed", func(t *testing.T) {
		_,err:= client.GetPodLogs(namespace,projectName,&corev1.PodLogOptions{})
		assert.NotEqual(t, nil, err)
	})
}