package kube

import (
	"k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

type PersistentVolumeClaim struct {
	clientSet kubernetes.Interface
}

// NewPersistentVolumeClaim initialize construction
func NewPersistentVolumeClaim(clientSet kubernetes.Interface) *PersistentVolumeClaim {
	return &PersistentVolumeClaim{
		clientSet: clientSet,
	}
}

func (p *PersistentVolumeClaim) Create(pvc *v1.PersistentVolumeClaim) (persistentVolumeClaim *v1.PersistentVolumeClaim, err error) {
	persistentVolumeClaim, err = p.clientSet.CoreV1().PersistentVolumeClaims(pvc.Namespace).Create(pvc)
	return
}

func (p *PersistentVolumeClaim) Get(name, namespace string, options meta_v1.GetOptions) (persistentVolumeClaim *v1.PersistentVolumeClaim, err error) {
	persistentVolumeClaim, err = p.clientSet.CoreV1().PersistentVolumeClaims(namespace).Get(name, options)
	return
}

func (p *PersistentVolumeClaim) Delete(name, namespace string, options *meta_v1.DeleteOptions) (err error) {
	err = p.clientSet.CoreV1().PersistentVolumeClaims(namespace).Delete(name, options)
	return
}

func (p *PersistentVolumeClaim) Update(pvc *v1.PersistentVolumeClaim) (persistentVolumeClaim *v1.PersistentVolumeClaim, err error) {
	persistentVolumeClaim, err = p.clientSet.CoreV1().PersistentVolumeClaims(pvc.Namespace).Update(pvc)
	return
}

func (p *PersistentVolumeClaim) Watch(namespace string, options meta_v1.ListOptions) (w watch.Interface, err error) {
	w, err = p.clientSet.CoreV1().PersistentVolumeClaims(namespace).Watch(options)
	return
}

