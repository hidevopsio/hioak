package kube

import "testing"
import (
	"github.com/magiconair/properties/assert"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestConfigMapsCurd(t *testing.T) {
	name := "test"
	namespace := "demo-dev"
	data := map[string]string{}
	clientSet := fake.NewSimpleClientset()
	configMaps := NewConfigMaps(clientSet)
	result, err := configMaps.Create(name, namespace, data)
	assert.Equal(t, nil, err)
	assert.Equal(t, name, result.Name)

	result, err = configMaps.Create(name, namespace, data)
	assert.Equal(t, nil, err)
	assert.Equal(t, name, result.Name)

	result, err = configMaps.Get(name, namespace)
	assert.Equal(t, nil, err)
	assert.Equal(t, name, result.Name)

	configMap := &core.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Data: map[string]string{
			"default": "{a}",
		},
	}
	result, err = configMaps.Update(name, namespace, configMap)
	assert.Equal(t, nil, err)
	assert.Equal(t, name, result.Name)

	err = configMaps.Delete(name, namespace)
	assert.Equal(t, nil, err)

}
