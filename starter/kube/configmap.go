package kube

import (
	"hidevops.io/hiboot/pkg/log"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ConfigMaps struct {
	clientSet kubernetes.Interface
}

// NewConfigMaps need add test
func NewConfigMaps(clientSet kubernetes.Interface) *ConfigMaps {
	return &ConfigMaps{
		clientSet: clientSet,
	}
}

func (c *ConfigMaps) Create(name, namespace string, data map[string]string) (*coreV1.ConfigMap, error) {
	log.Debug("config map create :", c)
	configMap := &coreV1.ConfigMap{
		ObjectMeta: metaV1.ObjectMeta{
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

func (c *ConfigMaps) Get(name, namespace string) (config *coreV1.ConfigMap, err error) {
	log.Info("get config map :", name)
	result, err := c.clientSet.CoreV1().ConfigMaps(namespace).Get(name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ConfigMaps) Delete(name, namespace string) error {
	log.Info("get config map :", name)
	err := c.clientSet.CoreV1().ConfigMaps(namespace).Delete(name, &metaV1.DeleteOptions{})
	return err
}

func (c *ConfigMaps) Update(name, namespace string, configMap *coreV1.ConfigMap) (*coreV1.ConfigMap, error) {
	log.Info("get config map :", name)
	result, err := c.clientSet.CoreV1().ConfigMaps(namespace).Update(configMap)
	return result, err
}
