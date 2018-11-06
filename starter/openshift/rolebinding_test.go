package openshift

import (
	"github.com/magiconair/properties/assert"
	"github.com/openshift/client-go/authorization/clientset/versioned/fake"
	"hidevops.io/hiboot/pkg/log"
	"testing"
)

func TestRoleBindingCrd(t *testing.T) {
	name := "admin"
	namespace := "demo-test"
	clientSet := fake.NewSimpleClientset().AuthorizationV1()
	roleBinding := newRoleBinding(clientSet)
	roleRefName := "User"
	roleRefKind := "chen"
	subjectKind := "User"
	subjectName := "shi"
	role, err := roleBinding.Create(name, namespace, roleRefName, roleRefKind, subjectKind, subjectName)
	log.Debug(role)
	log.Debug(err)

	// Get
	binding, err := roleBinding.Get(name, namespace)
	assert.Equal(t, nil, err)
	assert.Equal(t, name, binding.Name)

	err = roleBinding.Delete(name, namespace)
	assert.Equal(t, nil, err)
}

func TestCreateImagePullers(t *testing.T) {
	name := "system:image-pullers"
	namespace := "demo-test"
	clientSet := fake.NewSimpleClientset().AuthorizationV1()
	roleBinding := newRoleBinding(clientSet)
	roleRefName := "User"
	roleRefKind := "chen"
	subjectKind := "User"
	subjectName := "shi"

	_, err := roleBinding.Create(name, namespace, roleRefName, roleRefKind, subjectKind, subjectName)
	assert.Equal(t, nil, err)

}

func TestCreateImageBuilders(t *testing.T) {

	name := "system:image-builders"
	namespace := "demo-test"
	clientSet := fake.NewSimpleClientset().AuthorizationV1()
	roleBinding := newRoleBinding(clientSet)
	roleRefName := "User"
	roleRefKind := "chen"
	subjectKind := "User"
	subjectName := "shi"
	_, err := roleBinding.Create(name, namespace, roleRefName, roleRefKind, subjectKind, subjectName)
	assert.Equal(t, nil, err)

}

func TestCreateSystemDeployers(t *testing.T) {
	name := "system:deployers"
	namespace := "demo-test"
	clientSet := fake.NewSimpleClientset().AuthorizationV1()
	roleBinding := newRoleBinding(clientSet)
	roleRefName := "User"
	roleRefKind := "chen"
	subjectKind := "User"
	subjectName := "shi"
	_, err := roleBinding.Create(name, namespace, roleRefName, roleRefKind, subjectKind, subjectName)
	assert.Equal(t, nil, err)

}

func TestRoleBindingUpdate(t *testing.T) {
	name := "admin"
	namespace := "default"
	clientSet := fake.NewSimpleClientset().AuthorizationV1()
	roleBinding := newRoleBinding(clientSet)
	roleRefName := "User"
	roleRefKind := "chen"
	subjectKind := "User"
	subjectName := "shi"
	_, err := roleBinding.Create(name, namespace, roleRefName, roleRefKind, subjectKind, subjectName)
	assert.Equal(t, nil, err)

	err = roleBinding.InitImagePullers(namespace, roleRefName, roleRefKind, subjectKind, subjectName)
	assert.Equal(t, nil, err)

	err = roleBinding.InitImageBuilders(namespace, roleRefName, roleRefKind, subjectKind, subjectName)
	assert.Equal(t, nil, err)

	err = roleBinding.InitSystemDeployers(namespace, roleRefName, roleRefKind, subjectKind, subjectName)
	assert.Equal(t, nil, err)
}
