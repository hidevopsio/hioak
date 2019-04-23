package kube

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"hidevops.io/hiboot/pkg/log"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"testing"
)

func TestNewAutoConfigure(t *testing.T) {
	c := newConfiguration()
	clientSet, _ := kubernetes.NewForConfig(&rest.Config{})
	restConfig := &RestConfig{
		Config: &rest.Config{
			BearerToken: "",
		},
	}

	apiExtensionsClient, _ := apiextensionsclient.NewForConfig(&rest.Config{})
	crd := c.CustomResourceDefinition(apiExtensionsClient)
	fmt.Println(crd)

	testCases := []struct {
		expected interface{}
		actual   interface{}
	}{
		{NewConfigMaps(clientSet), c.ConfigMaps(clientSet)},
		{NewPod(clientSet), c.Pod(clientSet)},
		{NewService(clientSet), c.Service(clientSet)},
		{NewDeployment(clientSet), c.Deployment(clientSet)},
		{NewReplicaSet(clientSet), c.ReplicaSet(clientSet)},
		{NewSecret(clientSet), c.Secret(clientSet)},
		{NewNamespace(clientSet), c.Namespace(clientSet)},
		{NewReplicationController(clientSet), c.ReplicationController(clientSet)},
		{NewEvents(clientSet), c.Events(clientSet)},
		{NewCustomResourceDefinition(apiExtensionsClient), c.CustomResourceDefinition(apiExtensionsClient)},
		{Token(""), c.Token(restConfig)},
		{(*ConfigMaps)(nil), c.ConfigMaps(nil)},
		{(Token)(""), c.Token(nil)},
		{(*Pod)(nil), c.Pod(nil)},
		{(*Secret)(nil), c.Secret(nil)},
		{(*Service)(nil), c.Service(nil)},
		{(*Deployment)(nil), c.Deployment(nil)},
		{(*ReplicationController)(nil), c.ReplicationController(nil)},
		{(*ReplicaSet)(nil), c.ReplicaSet(nil)},
		{(*Events)(nil), c.Events(nil)},
		{(*CustomResourceDefinition)(nil), c.CustomResourceDefinition(nil)},
		{(*Ingress)(nil), c.Ingress(nil)},
		{(*Namespace)(nil), c.Namespace(nil)},
	}

	for _, item := range testCases {
		assert.Equal(t, item.expected, item.actual)
	}
}

func TestConfigurationClientSet(t *testing.T) {
	c := newConfiguration()
	//config := c.RestConfig(nil)
	config := &RestConfig{
		Config: &rest.Config{},
	}
	clientSet := c.ClientSet(nil)
	assert.Equal(t, nil, clientSet)

	extensionsClient := c.ApiExtensionsClient(nil)
	assert.Equal(t, nil, extensionsClient)

	c.ClientSet(config)
	c.ApiExtensionsClient(config)
	_, err := apiextensionsclient.NewForConfig(&rest.Config{})
	assert.Equal(t, nil, err)
}

func TestConfiguration_ClientConfig(t *testing.T) {
	c := newConfiguration()
	clientConfig := c.ClientConfig()
	log.Info(clientConfig)
}