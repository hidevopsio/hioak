package kube_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"hidevops.io/hiboot/pkg/log"
	"hidevops.io/hioak/starter/kube"
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
	appName := "hello-world"
	labels := map[string]string{
		"app": appName,
	}
	clientSet := fake.NewSimpleClientset()
	client := kube.NewPod(clientSet)

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      appName,
			Namespace: namespace,
			Labels:    labels,
		},
	}
	p, err := client.Create(pod)
	log.Infof("pod :%v", p)
	assert.Equal(t, nil, err)

	t.Run("Pods should watch succeed", func(t *testing.T) {
		listOptions := metav1.ListOptions{LabelSelector: fmt.Sprintf("app=%s", appName)}
		i, err := client.Watch(listOptions, namespace)
		log.Infof("i: %v", i)
		assert.Equal(t, nil, err)

	})

	t.Run("pods should list succeed", func(t *testing.T) {
		_, err := client.GetPodList(namespace, metav1.ListOptions{})
		assert.Equal(t, nil, err)
	})

	t.Run("pod should get succeed", func(t *testing.T) {
		_, err := client.GetPods(namespace, appName, metav1.GetOptions{})
		assert.Equal(t, nil, err)
	})

	t.Run("pod should get failed", func(t *testing.T) {
		_, err := client.GetPods(namespace, projectName, metav1.GetOptions{})
		assert.NotEqual(t, nil, err)
	})

	t.Run("Pods should get logs succeed", func(t *testing.T) {
		_, err := client.GetPodLogs(namespace, appName, &corev1.PodLogOptions{})
		assert.Equal(t, nil, err)
	})

	t.Run("Pods should get logs failed", func(t *testing.T) {
		_, err := client.GetPodLogs(namespace, projectName, &corev1.PodLogOptions{})
		assert.NotEqual(t, nil, err)
	})
}
