package kube

import (
	"hidevops.io/hiboot/pkg/log"
	"k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Namespace struct {
	clientSet kubernetes.Interface
}

//NewNamespace create namespaces
func NewNamespace(clientSet kubernetes.Interface) *Namespace {
	return &Namespace{
		clientSet: clientSet,
	}
}

func (n *Namespace) Create(ns *v1.Namespace) (namespace *v1.Namespace, err error) {
	log.Infof("create namespace: %s , name %s", ns.Name, ns.Namespace)
	namespace, err = n.clientSet.CoreV1().Namespaces().Create(ns)
	return
}

func (n *Namespace) Get(name string, options meta_v1.GetOptions) (namespace *v1.Namespace, err error) {
	log.Infof("get namespace get: %s", name)
	namespace, err = n.clientSet.CoreV1().Namespaces().Get(name, options)
	return
}


func (n *Namespace) Delete(name string, options *meta_v1.DeleteOptions) (err error) {
	log.Infof("get namespace get: %s", name)
	err = n.clientSet.CoreV1().Namespaces().Delete(name, options)
	return
}

func (n *Namespace) Update(ns *v1.Namespace) (namespace *v1.Namespace, err error) {
	log.Infof("get namespace get: %s", ns.Name)
	namespace, err = n.clientSet.CoreV1().Namespaces().Update(ns)
	return
}