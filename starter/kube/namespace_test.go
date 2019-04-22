package kube

import (
	"github.com/stretchr/testify/assert"
	"k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

func TestNamespaceCurd(t *testing.T) {
	name := "test"
	//namespace := "demo-dev"
	clientSet := fake.NewSimpleClientset()
	ns := NewNamespace(clientSet)
	n := &v1.Namespace{
		ObjectMeta: meta_v1.ObjectMeta{Name: name},
		Spec: v1.NamespaceSpec{
			Finalizers: []v1.FinalizerName{
				"default",
			},
		},
	}
	result, err := ns.Create(n)
	assert.Equal(t, nil, err)
	assert.Equal(t, name, result.Name)
	options := meta_v1.GetOptions{}
	result, err = ns.Get(name, options)
	assert.Equal(t, nil, err)
	assert.Equal(t, name, result.Name)

	result, err = ns.Update(n)
	assert.Equal(t, nil, err)
	assert.Equal(t, name, result.Name)
	op := &meta_v1.DeleteOptions{}
	err = ns.Delete(name, op)
	assert.Equal(t, nil, err)

}
