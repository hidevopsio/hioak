package kube

import (
	"flag"
	"hidevops.io/hiboot/pkg/app"
	"hidevops.io/hiboot/pkg/at"
	"hidevops.io/hiboot/pkg/log"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
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

type ClientConfig clientcmd.ClientConfig

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

type ApiExtensionsClient interface {
	apiextensionsclient.Interface
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

func (c *configuration) ClientSet(restConfig *RestConfig) ClientSet {
	if restConfig != nil {
		clientSet, err := kubernetes.NewForConfig(restConfig.Config)
		if err != nil {
			return nil
		}
		return clientSet
	}
	return nil
}

func (c *configuration) ApiExtensionsClient(restConfig *RestConfig) ApiExtensionsClient {
	if restConfig != nil {
		clientSet, err := apiextensionsclient.NewForConfig(restConfig.Config)
		if err != nil {
			return nil
		}
		return clientSet
	}
	return nil
}

//ConfigMaps autoConfigure deployment need initialize construction
func (c *configuration) ConfigMaps(clientSet ClientSet) *ConfigMaps {
	if clientSet != nil {
		return NewConfigMaps(clientSet)
	}
	return nil
}

//Deployment autoConfigure deployment need initialize construction
func (c *configuration) Deployment(clientSet ClientSet) *Deployment {
	if clientSet != nil {
		return NewDeployment(clientSet)
	}
	return nil
}

//ReplicationController autoConfigure deployment need initialize construction
func (c *configuration) ReplicationController(clientSet ClientSet) *ReplicationController {
	if clientSet != nil {
		return NewReplicationController(clientSet)
	}
	return nil

}

//Secret autoConfigure deployment need initialize construction
func (c *configuration) Secret(clientSet ClientSet) *Secret {
	if clientSet != nil {
		return NewSecret(clientSet)
	}
	return nil

}

//Service autoConfigure deployment need initialize construction
func (c *configuration) Service(clientSet ClientSet) *Service {
	if clientSet != nil {
		return NewService(clientSet)
	}
	return nil

}

//Pod autoConfigure deployment need initialize construction
func (c *configuration) Pod(clientSet ClientSet) *Pod {
	if clientSet != nil {
		return NewPod(clientSet)
	}
	return nil

}

//Token autoConfigure deployment need initialize construction
func (c *configuration) Token(restConfig *RestConfig) Token {
	if restConfig != nil {
		return Token(restConfig.Config.BearerToken)
	}
	return ""
}

//ReplicaSet autoConfigure deployment need initialize construction
func (c *configuration) ReplicaSet(clientSet ClientSet) *ReplicaSet {
	if clientSet != nil {
		return NewReplicaSet(clientSet)
	}
	return nil
}

//Events autoConfigure deployment need initialize construction
func (c *configuration) Events(clientSet ClientSet) *Events {
	if clientSet != nil {
		return NewEvents(clientSet)
	}
	return nil
}

//CustomResourceDefinition autoConfigure deployment need initialize construction
func (c *configuration) CustomResourceDefinition(apiExtensionsClient ApiExtensionsClient) *CustomResourceDefinition {
	if apiExtensionsClient != nil {
		return NewCustomResourceDefinition(apiExtensionsClient)
	}
	return nil

}

//Ingress autoConfigure deployment need initialize construction
func (c *configuration) Ingress(clientSet ClientSet) *Ingress {
	if clientSet != nil {
		return NewIngress(clientSet)
	}
	return nil

}

//Namespace autoConfigure deployment need initialize construction
func (c *configuration) Namespace(clientSet ClientSet) *Namespace {
	if clientSet != nil {
		return NewNamespace(clientSet)
	}
	return nil

}

//ClusterRoleBinding autoConfigure deployment need initialize construction
func (c *configuration) ClusterRoleBinding(clientSet ClientSet) *ClusterRoleBining {
	if clientSet != nil {
		return NewClusterRoleBining(clientSet)
	}
	return nil

}

//ClusterRole autoConfigure deployment need initialize construction
func (c *configuration) ClusterRole(clientSet ClientSet) *ClusterRole {
	if clientSet != nil {
		return NewClusterRole(clientSet)
	}
	return nil

}

//ClientConfig creates a ConfigClientClientConfig using the passed context name
func (c *configuration) ClientConfig() ClientConfig {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
}