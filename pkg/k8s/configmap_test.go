package k8s

import "testing"
import (
	core "k8s.io/api/core/v1"
	"github.com/magiconair/properties/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestConfigMapsCurd(t *testing.T) {
	name := "test"
	namespace := "demo-dev"
	data := map[string]string{}
	clientSet := fake.NewSimpleClientset()
	configMaps := NewConfigMaps(clientSet, name, namespace, data)
	configMaps.Data = map[string]string{
	}
	result, err := configMaps.Create()
	assert.Equal(t, nil, err)
	assert.Equal(t, name, result.Name)

	result, err = configMaps.Create()
	assert.Equal(t, nil, err)
	assert.Equal(t, name, result.Name)

	result, err = configMaps.Get()
	assert.Equal(t, nil, err)
	assert.Equal(t, name, result.Name)

	configMap := &core.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Data: map[string]string{
			"default":"{a}",
		},
	}
	result, err = configMaps.Update(configMap)
	assert.Equal(t, nil, err)
	assert.Equal(t, name, result.Name)

	err = configMaps.Delete()
	assert.Equal(t, nil, err)


}
