package kube

import (
	"flag"
	"github.com/hidevopsio/hiboot/pkg/app"
	"github.com/hidevopsio/hiboot/pkg/log"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

type properties struct {
	KubeServiceHost string `json:"kube_service_host" default:"${KUBERNETES_SERVICE_HOST}"`
}

// define type configuration
type configuration struct {
	app.Configuration

	Properties properties `json:"properties" mapstructure:"kube"`
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
	if c.Properties.KubeServiceHost == "" {
		retVal.Config, err = clientcmd.BuildConfigFromFlags("", *kubeConfig.string)
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
	if c.Properties.KubeServiceHost == "" {
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

func (c *configuration) KubeDeployment(clientSet ClientSet) *Deployment {
	return newDeployment(clientSet)
}

func (c *configuration) KubeReplicationController(clientSet ClientSet) *ReplicationController {
	return NewReplicationController(clientSet)
}

func (c *configuration) KubeSecret(clientSet ClientSet) *Secret {
	return NewSecret(clientSet)
}

func (c *configuration) KubeService(clientSet ClientSet) *Service {
	return NewService(clientSet)
}
