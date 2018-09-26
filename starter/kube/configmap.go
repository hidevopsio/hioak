package kube

import (
	"k8s.io/client-go/kubernetes/typed/core/v1"
	core_v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/hidevopsio/hiboot/pkg/log"
	"k8s.io/client-go/kubernetes"
)

type ConfigMaps struct {
	clientSet kubernetes.Interface
}

func newConfigMaps(clientSet kubernetes.Interface) *ConfigMaps {
	return &ConfigMaps{
		clientSet: clientSet,
	}
}


func (c *ConfigMaps) Create() (*core_v1.ConfigMap, error) {
	log.Debug("config map create :", c)
	configMap := &core_v1.ConfigMap{
		ObjectMeta: meta_v1.ObjectMeta{
			Name: c.Name,
		},
		Data: c.Data,
	}
	cm, err := c.Get()
	log.Debug("config map get :", cm)
	if err == nil {
		nc, err := c.Update(configMap)
		return nc, err
	}
	config, er := c.Interface.Create(configMap)
	if er != nil {
		return nil, er
	}

	return config, nil
}

func (c *ConfigMaps) Get() (config *core_v1.ConfigMap, err error) {
	log.Info("get config map :", c.Name)
	result, err := c.Interface.Get(c.Name, meta_v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ConfigMaps) Delete() error {
	log.Info("get config map :", c.Name)
	err := c.Interface.Delete(c.Name, &meta_v1.DeleteOptions{})
	return err
}

func (c *ConfigMaps) Update(configMap *core_v1.ConfigMap) (*core_v1.ConfigMap, error) {
	log.Info("get config map :", c.Name)
	result, err := c.Interface.Update(configMap)
	return result, err
}
