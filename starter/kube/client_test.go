package kube

import (
	"fmt"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type KubernetesAPI struct {
	Suffix string
	Client kubernetes.Interface
}

// NewNamespaceWithPostfix creates a new namespace with a stable postfix
func (k KubernetesAPI) NewNamespaceWithSuffix(namespace string) error {
	ns := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("%s-%s", namespace, k.Suffix),
		},
	}

	_, err := k.Client.CoreV1().Namespaces().Create(ns)

	if err != nil {
		return err
	}

	return nil
}
