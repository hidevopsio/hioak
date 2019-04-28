package kube

import (
	"k8s.io/api/rbac/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

type Role struct {
	clientSet kubernetes.Interface
}

// NewRole ConfigMaps initialize construction
func NewRole(clientSet kubernetes.Interface) *Role {
	return &Role{
		clientSet: clientSet,
	}
}

func (c *Role) Create(namespace string, cr *v1.Role) (Role *v1.Role, err error) {
	Role, err = c.clientSet.RbacV1().Roles(namespace).Create(cr)
	return
}

func (c *Role) Get(name, namespace string, options meta_v1.GetOptions) (Role *v1.Role, err error) {
	Role, err = c.clientSet.RbacV1().Roles(namespace).Get(name, options)
	return
}

func (c *Role) Delete(name string, namespace string, options *meta_v1.DeleteOptions) (err error) {
	err = c.clientSet.RbacV1().Roles(namespace).Delete(name, options)
	return
}

func (c *Role) Update(namespace string, cr *v1.Role) (Role *v1.Role, err error) {
	Role, err = c.clientSet.RbacV1().Roles(namespace).Update(cr)
	return
}

func (c *Role) Watch(namespace string, options meta_v1.ListOptions) (w watch.Interface, err error) {
	w, err = c.clientSet.RbacV1().Roles(namespace).Watch(options)
	return
}

