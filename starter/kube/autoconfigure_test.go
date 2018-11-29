package kube

import (
	"github.com/prometheus/common/log"
	"testing"
)

func TestNewConfigMaps(t *testing.T) {
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
