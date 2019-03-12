package kube

import (
	"github.com/stretchr/testify/assert"
	"k8s.io/api/extensions/v1beta1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

//TestNewIngress test ingress ORCD
func TestIngress(t *testing.T) {
	name := "Name"
	namespace := "Namespace"
	ingress := &v1beta1.Ingress{
		ObjectMeta: meta_v1.ObjectMeta{Name: name, Namespace: namespace},
	}

	clientSet := fake.NewSimpleClientset()
	client := NewIngress(clientSet)

	t.Run("should ingress create success", func(t *testing.T) {
		_, err := client.Create(ingress)
		assert.Equal(t, nil, err)
	})

	t.Run("should ingress get success", func(t *testing.T) {
		_, err := client.Get(namespace, name, meta_v1.GetOptions{})
		assert.Equal(t, nil, err)
	})

	t.Run("should ingress list success", func(t *testing.T) {
		_, err := client.List(namespace, meta_v1.ListOptions{})
		assert.Equal(t, nil, err)
	})

	t.Run("should ingress update success", func(t *testing.T) {
		_, err := client.Update(ingress)
		assert.Equal(t, nil, err)
	})

	t.Run("should ingress watch success", func(t *testing.T) {
		_, err := client.Watch(namespace, meta_v1.ListOptions{})
		assert.Equal(t, nil, err)
	})

	t.Run("should ingress delete success", func(t *testing.T) {
		err := client.Delete(namespace, name, &meta_v1.DeleteOptions{})
		assert.Equal(t, nil, err)
	})

}

