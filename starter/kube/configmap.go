package kube

import (
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


func (c *ConfigMaps) Create(name, namespace string, data map[string]string) (*core_v1.ConfigMap, error) {
	log.Debug("config map create :", c)
	configMap := &core_v1.ConfigMap{
		ObjectMeta: meta_v1.ObjectMeta{
			Name: name,
		},
		Data: data,
	}
	cm, err := c.Get(name, namespace)
	log.Debug("config map get :", cm)
	if err == nil {
		nc, err := c.Update(name, namespace, configMap)
		return nc, err
	}
	config, er := c.clientSet.CoreV1().ConfigMaps(namespace).Create(configMap)
	if er != nil {
		return nil, er
	}

	return config, nil
}

func (c *ConfigMaps) Get(name, namespace string) (config *core_v1.ConfigMap, err error) {
	log.Info("get config map :", name)
	result, err := c.clientSet.CoreV1().ConfigMaps(namespace).Get(name, meta_v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ConfigMaps) Delete(name, namespace string) error {
	log.Info("get config map :", name)
	err := c.clientSet.CoreV1().ConfigMaps(namespace).Delete(name, &meta_v1.DeleteOptions{})
	return err
}

func (c *ConfigMaps) Update(name, namespace string, configMap *core_v1.ConfigMap) (*core_v1.ConfigMap, error) {
	log.Info("get config map :", name)
	result, err := c.clientSet.CoreV1().ConfigMaps(namespace).Update(configMap)
	return result, err
}
