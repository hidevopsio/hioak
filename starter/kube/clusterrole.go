package kube

import (
	"k8s.io/api/rbac/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

type ClusterRole struct {
	clientSet kubernetes.Interface
}

// NewClusterRole ConfigMaps initialize construction
func NewClusterRole(clientSet kubernetes.Interface) *ClusterRole {
	return &ClusterRole{
		clientSet: clientSet,
	}
}

func (c *ClusterRole) Create(cr *v1.ClusterRole) (clusterRole *v1.ClusterRole, err error) {
	clusterRole, err = c.clientSet.RbacV1().ClusterRoles().Create(cr)
	return
}

func (c *ClusterRole) Delete(name string, options *meta_v1.DeleteOptions) (err error) {
	err = c.clientSet.RbacV1().ClusterRoles().Delete(name, options)
	return
}

func (c *ClusterRole) Update(cr *v1.ClusterRole) (clusterRole *v1.ClusterRole, err error) {
	clusterRole, err = c.clientSet.RbacV1().ClusterRoles().Update(cr)
	return
}

func (c *ClusterRole) Watch(options meta_v1.ListOptions) (w watch.Interface, err error) {
	w, err = c.clientSet.RbacV1().ClusterRoles().Watch(options)
	return
}

