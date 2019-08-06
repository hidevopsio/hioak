package kube

import (
	"github.com/stretchr/testify/assert"
	 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

func TestPersistentVolumeCrud(t *testing.T) {
	name := "test"
	clientSet := fake.NewSimpleClientset()
	crb := NewPersistentVolume(clientSet)
	pvc := &v1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
	result, err := crb.Create(pvc)
	assert.Equal(t, nil, err)
	assert.Equal(t, name, result.Name)
	options := metav1.GetOptions{}
	result, err = crb.Get(name, options)
	assert.Equal(t, nil, err)
	assert.Equal(t, name, result.Name)

	result, err = crb.Update(pvc)
	assert.Equal(t, nil, err)
	assert.Equal(t, name, result.Name)
	dOptions := &metav1.DeleteOptions{}
	err = crb.Delete(name, dOptions)
	assert.Equal(t, nil, err)
}