package kube

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestNewConfigMaps(t *testing.T) {
	c := newConfiguration()
	cm := c.ConfigMaps(nil)
	assert.Equal(t, &ConfigMaps{}, cm)
	pod := c.Pod(nil)
	assert.Equal(t, &Pod{}, pod)
	secret := c.Secret(nil)
	assert.Equal(t, &Secret{}, secret)
	service := c.Service(nil)
	assert.Equal(t, &Service{}, service)
	replicaSet := c.ReplicaSet(nil)
	assert.Equal(t, &ReplicaSet{}, replicaSet)
	deployment := c.Deployment(nil)
	assert.Equal(t, &Deployment{}, deployment)
	replicationController := c.ReplicationController(nil)
	assert.Equal(t, &ReplicationController{}, replicationController)
}
