package kube

import (
	"k8s.io/api/rbac/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

type ClusterRoleBining struct {
	clientSet kubernetes.Interface
}

// NewClusterRoleBining ConfigMaps initialize construction
func NewClusterRoleBining(clientSet kubernetes.Interface) *ClusterRoleBining {
	return &ClusterRoleBining{
		clientSet: clientSet,
	}
}

func (c *ClusterRoleBining) Create(crb *v1.ClusterRoleBinding) (clusterRoleBinding *v1.ClusterRoleBinding, err error) {
	clusterRoleBinding, err = c.clientSet.RbacV1().ClusterRoleBindings().Create(crb)
	return
}

func (c *ClusterRoleBining) Get(name string, options meta_v1.GetOptions) (clusterRoleBinding *v1.ClusterRoleBinding, err error) {
	clusterRoleBinding, err = c.clientSet.RbacV1().ClusterRoleBindings().Get(name, options)
	return
}

func (c *ClusterRoleBining) Delete(name string, options *meta_v1.DeleteOptions) (err error) {
	err = c.clientSet.RbacV1().ClusterRoleBindings().Delete(name, options)
	return
}

func (c *ClusterRoleBining) Update(crb *v1.ClusterRoleBinding) (clusterRoleBinding *v1.ClusterRoleBinding, err error) {
	clusterRoleBinding, err = c.clientSet.RbacV1().ClusterRoleBindings().Update(crb)
	return
}

func (c *ClusterRoleBining) Watch(options meta_v1.ListOptions) (w watch.Interface, err error) {
	w, err = c.clientSet.RbacV1().ClusterRoleBindings().Watch(options)
	return
}
