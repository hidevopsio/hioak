package kube

import (
	"fmt"
	"github.com/prometheus/common/log"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
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
	log.Infof(fmt.Sprintf("watch pod app %s in namespace %s:", name, namespace))

	//listOptions.LabelSelector = fmt.Sprintf("app=%s", name)
	listOptions.Watch = true

	w, err := p.clientSet.CoreV1().Pods(namespace).Watch(listOptions)
	if err != nil {
		return nil, err
	}
	return w, nil
}

func (p *Pod) GetPodLogs(namespace, name string, opts *corev1.PodLogOptions) (*restclient.Request, error) {
	log.Infof(fmt.Sprintf("get pod %s logs in namespace %s", name, namespace))
	var err error
	if _, err = p.clientSet.CoreV1().Pods(namespace).Get(name, metav1.GetOptions{}); err != nil {
		return nil, err
	}
	request := p.clientSet.CoreV1().Pods(namespace).GetLogs(name, opts)
	return request, nil
}

func (p *Pod) GetPods(namespace, name string, opts metav1.GetOptions) (*corev1.Pod, error) {
	log.Infof(fmt.Sprintf("get pod %s in namespace %s", name, namespace))
	var pod *corev1.Pod
	var err error
	if pod, err = p.clientSet.CoreV1().Pods(namespace).Get(name, opts); err != nil {
		return nil, err
	}
	return pod, nil
}

func (p *Pod) GetPodList(namespace string, opts metav1.ListOptions) (*corev1.PodList, error) {
	log.Infof(fmt.Sprintf("get pod list in namespace %s", namespace))
	var podList *corev1.PodList
	var err error
	if podList, err = p.clientSet.CoreV1().Pods(namespace).List(opts); err != nil {
		return nil, err
	}
	return podList, nil
}
