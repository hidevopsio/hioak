package kube

import (
	"flag"
	"hidevops.io/hiboot/pkg/app"
	"hidevops.io/hiboot/pkg/at"
	"hidevops.io/hiboot/pkg/log"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

type properties struct {
	ServiceHost string `json:"service_host" default:"${KUBERNETES_SERVICE_HOST}"`
}

// define type configuration
type configuration struct {
	at.AutoConfiguration

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
	app.Register(newConfiguration)
}

func (c *configuration) RestConfig(Config *Config) *RestConfig {
	retVal := new(RestConfig)
	var err error
	if c.Properties.ServiceHost == "" {
		retVal.Config, err = clientcmd.BuildConfigFromFlags("", *Config.string)
	} else {
		retVal.Config, err = rest.InClusterConfig()
	}
	if err != nil {
		return nil
	}
	return retVal
}

func (c *configuration) Config() *Config {
	kc := new(Config)
	if c.Properties.ServiceHost == "" {
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

func (c *configuration) ClientSet(RestConfig *RestConfig) ClientSet {
	clientSet, err := kubernetes.NewForConfig(RestConfig.Config)
	if err != nil {
		return nil
	}
	return clientSet
}

//ConfigMaps autoConfigure deployment need initialize construction
func (c *configuration) ConfigMaps(clientSet ClientSet) *ConfigMaps {
	return NewConfigMaps(clientSet)
}

//Deployment autoConfigure deployment need initialize construction
func (c *configuration) Deployment(clientSet ClientSet) *Deployment {
	return NewDeployment(clientSet)
}

//ReplicationController autoConfigure deployment need initialize construction
func (c *configuration) ReplicationController(clientSet ClientSet) *ReplicationController {
	return NewReplicationController(clientSet)
}

//Secret autoConfigure deployment need initialize construction
func (c *configuration) Secret(clientSet ClientSet) *Secret {
	return NewSecret(clientSet)
}

//Service autoConfigure deployment need initialize construction
func (c *configuration) Service(clientSet ClientSet) *Service {
	return NewService(clientSet)
}

//Pod autoConfigure deployment need initialize construction
func (c *configuration) Pod(clientSet ClientSet) *Pod {
	return NewPod(clientSet)
}

//Token autoConfigure deployment need initialize construction
func (c *configuration) Token(restConfig *RestConfig) Token {
	return Token(restConfig.Config.BearerToken)
}

//ReplicaSet autoConfigure deployment need initialize construction
func (c *configuration) ReplicaSet(clientSet ClientSet) *ReplicaSet {
	return NewReplicaSet(clientSet)
}
