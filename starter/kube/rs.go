package kube

import (
	"hidevops.io/hiboot/pkg/log"
	"k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ReplicaSet struct {
	clientSet kubernetes.Interface
}

// NewReplicaSet ConfigMaps initialize construction
func NewReplicaSet(clientSet kubernetes.Interface) *ReplicaSet {
	return &ReplicaSet{
		clientSet: clientSet,
	}
}

func (rs *ReplicaSet) Create(replica *v1.ReplicaSet) (*v1.ReplicaSet, error) {
	log.Infof("create replica set name %v, namespace %v", replica.Name, replica.Namespace)
	r, err := rs.clientSet.AppsV1().ReplicaSets(replica.Namespace).Create(replica)
	return r, err
}

func (rs *ReplicaSet) Delete(name, namespace string, option *metav1.DeleteOptions) error {
	log.Infof("delete replica set name %v, namespace %v", name, namespace)
	err := rs.clientSet.AppsV1().ReplicaSets(namespace).Delete(name, option)
	return err
}

func (rs *ReplicaSet) List(name, namespace string, option metav1.ListOptions) (*v1.ReplicaSetList, error) {
	log.Infof("list replica set name %v, namespace %v", name, namespace)
	replicaSets, err := rs.clientSet.AppsV1().ReplicaSets(namespace).List(option)
	return replicaSets, err
}
