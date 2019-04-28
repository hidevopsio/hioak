package kube

import (
	"k8s.io/api/rbac/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

type RoleBinding struct {
	clientSet kubernetes.Interface
}

// NewRoleBinding ConfigMaps initialize construction
func NewRoleBinding(clientSet kubernetes.Interface) *RoleBinding {
	return &RoleBinding{
		clientSet: clientSet,
	}
}

func (c *RoleBinding) Create(namespace string, crb *v1.RoleBinding) (roleBinding *v1.RoleBinding, err error) {
	roleBinding, err = c.clientSet.RbacV1().RoleBindings(namespace).Create(crb)
	return
}

func (c *RoleBinding) Get(name string, namespace string, options meta_v1.GetOptions) (roleBinding *v1.RoleBinding, err error) {
	roleBinding, err = c.clientSet.RbacV1().RoleBindings(namespace).Get(name, options)
	return
}

func (c *RoleBinding) Delete(name string, namespace string, options *meta_v1.DeleteOptions) (err error) {
	err = c.clientSet.RbacV1().RoleBindings(namespace).Delete(name, options)
	return
}

func (c *RoleBinding) Update(namespace string, crb *v1.RoleBinding) (roleBinding *v1.RoleBinding, err error) {
	roleBinding, err = c.clientSet.RbacV1().RoleBindings(namespace).Update(crb)
	return
}

func (c *RoleBinding) Watch(namespace string, options meta_v1.ListOptions) (w watch.Interface, err error) {
	w, err = c.clientSet.RbacV1().RoleBindings(namespace).Watch(options)
	return
}
