package kube

import (
	"k8s.io/client-go/kubernetes"
	corev1 "k8s.io/api/core/v1"
	"github.com/prometheus/common/log"
	"fmt"
	"k8s.io/apimachinery/pkg/watch"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Pod struct {
	clientSet kubernetes.Interface
}

func NewPod(clientSet kubernetes.Interface) *Pod {
	return &Pod{
		clientSet: clientSet,
	}
}

func (p *Pod) Create(pod *corev1.Pod) (*corev1.Pod, error) {
	log.Debugf("create pod :")
	return p.clientSet.CoreV1().Pods(pod.Namespace).Create(pod)
}

func (p *Pod) Watch(listOptions metav1.ListOptions, namespace, name string) (watch.Interface, error) {
	log.Info(fmt.Sprintf("watch pod app %s in namespace %s:", name, namespace))

	listOptions.LabelSelector = fmt.Sprintf("app=%s", name)
	listOptions.Watch = true

	w, err := p.clientSet.CoreV1().Pods(namespace).Watch(listOptions)
	if err != nil {
		return nil, err
	}
	return w, nil
}
