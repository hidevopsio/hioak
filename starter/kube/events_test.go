package kube

import (
	"github.com/stretchr/testify/assert"
	"k8s.io/api/events/v1beta1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

//TestNewEvents test events ORCD
func TestNewEvents(t *testing.T) {
	eventName := "eventName"
	eventNamespace := "eventNamespace"
	event := &v1beta1.Event{
		ObjectMeta: meta_v1.ObjectMeta{Name: eventName, Namespace: eventNamespace},
		Reason:     "Unhealthy",
		Type:       "Warning",
	}

	clientSet := fake.NewSimpleClientset()
	client := NewEvents(clientSet)

	t.Run("should event create success", func(t *testing.T) {
		_, err := client.Create(event)
		assert.Equal(t, nil, err)
	})

	t.Run("should event get success", func(t *testing.T) {
		_, err := client.Get(eventNamespace, eventName, meta_v1.GetOptions{})
		assert.Equal(t, nil, err)
	})

	t.Run("should event list success", func(t *testing.T) {
		_, err := client.List(eventNamespace, meta_v1.ListOptions{})
		assert.Equal(t, nil, err)
	})

	t.Run("should event update success", func(t *testing.T) {
		_, err := client.Update(event)
		assert.Equal(t, nil, err)
	})

	t.Run("should event watch success", func(t *testing.T) {
		_, err := client.Watch(eventNamespace, meta_v1.ListOptions{})
		assert.Equal(t, nil, err)
	})

	t.Run("should event delete success", func(t *testing.T) {
		err := client.Delete(eventNamespace, eventName, &meta_v1.DeleteOptions{})
		assert.Equal(t, nil, err)
	})

}
