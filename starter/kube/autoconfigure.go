package kube

import (
	"github.com/hidevopsio/hiboot/pkg/app"
	"k8s.io/client-go/util/homedir"
	"flag"
	"path/filepath"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/rest"
	"github.com/hidevopsio/hiboot/pkg/log"
	"k8s.io/client-go/kubernetes"
)

type properties struct {
	ExternalClient bool `json:"external_client" default:"false"`
}

// define type configuration
type configuration struct{
	app.Configuration

	Properties properties `json:"properties" mapstruct:"k8s"`
}

func newConfiguration() *configuration {
	return &configuration{}
}

type Config struct {
	*string
}

type RestConfig struct {
	*rest.Config
}

type ClientSet interface {
	kubernetes.Interface
}


func init() {
	app.AutoConfiguration(newConfiguration)
}

func (c *configuration) KubeRestConfig(kubeConfig *Config) *RestConfig {
	retVal := new(RestConfig)
	var err error
	if c.Properties.ExternalClient {
		retVal.Config, err = clientcmd.BuildConfigFromFlags("", string(*kubeConfig))
	} else {
		retVal.Config, err = rest.InClusterConfig()
	}
	if err != nil {
		return nil
	}
	return retVal
}

func (c *configuration) KubeConfig() *Config {
	kc := new(Config)
	if c.Properties.ExternalClient {
		log.Info("Kubernetes External Client Mode")
		if home := homedir.HomeDir(); home != "" {
			kc.string = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kc.string = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
	} else {
		log.Info("Kubernetes Internal Client Mode")
	}
	return kc
}

func (c *configuration) KubeClientSet(kubeRestConfig *RestConfig) ClientSet {
	clientSet, err := kubernetes.NewForConfig(kubeRestConfig.Config)
	if err != nil {
		return nil
	}
	return clientSet
}

func (c *configuration) KubeConfigMaps(clientSet ClientSet) *ConfigMaps {
	return newConfigMaps(clientSet)
}

