package kube

import (
	"hidevops.io/hiboot/pkg/log"
	"k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)


type Ingress struct {
	clientSet kubernetes.Interface
}

func NewIngress(clientSet kubernetes.Interface) *Ingress {
	return &Ingress{
		clientSet: clientSet,
	}
}

func (i *Ingress) Get(namespace, name string, opts v1.GetOptions) (*v1beta1.Ingress, error) {
	log.Debugf("get Ingress %s in namespace %s", name, namespace)
	return i.clientSet.ExtensionsV1beta1().Ingresses(namespace).Get(name, opts)
}

func (i *Ingress) List(namespace string, opts v1.ListOptions) (*v1beta1.IngressList, error) {
	log.Debugf("get Ingress list by label %s in namespace %s", opts.LabelSelector, namespace)
	return i.clientSet.ExtensionsV1beta1().Ingresses(namespace).List(opts)
}

func (i *Ingress) Watch(namespace string, opts v1.ListOptions) (watch.Interface, error) {
	log.Debugf("watch Ingress by label %s in namespace %s", opts.LabelSelector, namespace)
	return i.clientSet.ExtensionsV1beta1().Ingresses(namespace).Watch(opts)
}

func (i *Ingress) Create(ingress *v1beta1.Ingress) (*v1beta1.Ingress, error) {
	log.Debugf("create Ingress %s in namespace %s", ingress.Name, ingress.Namespace)
	return i.clientSet.ExtensionsV1beta1().Ingresses(ingress.Namespace).Create(ingress)
}

func (i *Ingress) Update(ingress *v1beta1.Ingress) (*v1beta1.Ingress, error) {
	log.Debugf("update Ingress %s in namespace %s", ingress.Name, ingress.Namespace)
	return i.clientSet.ExtensionsV1beta1().Ingresses(ingress.Namespace).Update(ingress)
}

func (i *Ingress) Delete(namespace, name string, opts *v1.DeleteOptions) error {
	log.Debugf("delete Ingress %s in namespace %s", name, namespace)
	return i.clientSet.ExtensionsV1beta1().Ingresses(namespace).Delete(name, opts)
}
