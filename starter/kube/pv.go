package kube

import (
	"k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

type PersistentVolume struct {
	clientSet kubernetes.Interface
}

// NewPersistentVolume initialize construction
func NewPersistentVolume(clientSet kubernetes.Interface) *PersistentVolume {
	return &PersistentVolume{
		clientSet: clientSet,
	}
}

func (p *PersistentVolume) Create(pv *v1.PersistentVolume) (persistentVolume *v1.PersistentVolume, err error) {
	persistentVolume, err = p.clientSet.CoreV1().PersistentVolumes().Create(pv)
	return
}

func (p *PersistentVolume) Get(name string, options meta_v1.GetOptions) (persistentVolume *v1.PersistentVolume, err error) {
	persistentVolume, err = p.clientSet.CoreV1().PersistentVolumes().Get(name, options)
	return
}

func (p *PersistentVolume) Delete(name string, options *meta_v1.DeleteOptions) (err error) {
	err = p.clientSet.CoreV1().PersistentVolumes().Delete(name, options)
	return
}

func (p *PersistentVolume) Update(pvc *v1.PersistentVolume) (persistentVolume *v1.PersistentVolume, err error) {
	persistentVolume, err = p.clientSet.CoreV1().PersistentVolumes().Update(pvc)
	return
}

func (p *PersistentVolume) Watch(options meta_v1.ListOptions) (w watch.Interface, err error) {
	w, err = p.clientSet.CoreV1().PersistentVolumes().Watch(options)
	return
}

