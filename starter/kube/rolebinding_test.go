package kube

import (
	"github.com/magiconair/properties/assert"
	"k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

func TestRoleBining_Crud(t *testing.T) {
	name := "test"
	namespace := "hidevopsio"
	clientSet := fake.NewSimpleClientset()
	crb := NewRoleBinding(clientSet)
	roleBinding := &v1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Subjects: []v1.Subject{
			v1.Subject{
				Name: name,
			},
		},
		RoleRef: v1.RoleRef{},
	}
	result, err := crb.Create(namespace, roleBinding)
	assert.Equal(t, nil, err)
	assert.Equal(t, name, result.Name)
	options := metav1.GetOptions{}
	result, err = crb.Get(name, namespace, options)
	assert.Equal(t, nil, err)
	assert.Equal(t, name, result.Name)

	result, err = crb.Update(namespace, roleBinding)
	assert.Equal(t, nil, err)
	assert.Equal(t, name, result.Name)
	dOptions := &metav1.DeleteOptions{}
	err = crb.Delete(name, namespace, dOptions)
	assert.Equal(t, nil, err)
}
