package kube

import (
	"github.com/prometheus/common/log"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"testing"
)

func TestAutoConfigure(t *testing.T) {
	c := newConfiguration()
	cm := c.ConfigMaps(nil)
	log.Info(cm)
	pod := c.Pod(nil)
	log.Info(pod)
	secret := c.Secret(nil)
	log.Info(secret)
	service := c.Service(nil)
	log.Info(service)
	replicaSet := c.ReplicaSet(nil)
	log.Info(replicaSet)
	deployment := c.Deployment(nil)
	log.Info(deployment)
	replicationController := c.ReplicationController(nil)
	log.Info(replicationController)
}


func TestNewAutoConfigure(t *testing.T) {
	c := newConfiguration()
	clientSet, _ := kubernetes.NewForConfig(&rest.Config{})
	cm := c.ConfigMaps(clientSet)
	log.Info(cm)
	pod := c.Pod(clientSet)
	log.Info(pod)
	secret := c.Secret(clientSet)
	log.Info(secret)
	service := c.Service(clientSet)
	log.Info(service)
	replicaSet := c.ReplicaSet(clientSet)
	log.Info(replicaSet)
	deployment := c.Deployment(clientSet)
	log.Info(deployment)
	token := c.Token(nil)
	log.Info(token)
	replicationController := c.ReplicationController(clientSet)
	log.Info(replicationController)
}
