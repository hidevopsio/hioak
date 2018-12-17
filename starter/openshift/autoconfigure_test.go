package openshift

import (
	"github.com/magiconair/properties/assert"
	appsv1 "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	imagev1 "github.com/openshift/client-go/image/clientset/versioned/typed/image/v1"
	oauthv1 "github.com/openshift/client-go/oauth/clientset/versioned/typed/oauth/v1"
	"hidevops.io/hioak/starter/kube"
	"k8s.io/client-go/rest"
	"testing"
)

func TestConfiguration(t *testing.T) {
	c := newConfiguration()
	restConfig := &kube.RestConfig{
		Config: &rest.Config{
			BearerToken: "",
		},
	}
	c.Auth(restConfig)
	_, err := oauthv1.NewForConfig(restConfig.Config)
	assert.Equal(t, nil, err)

	c.DeploymentConfig(restConfig)
	_, err = appsv1.NewForConfig(restConfig.Config)
	assert.Equal(t, nil, err)

	c.ImageStream(restConfig)
	_, err = imagev1.NewForConfig(restConfig.Config)
	assert.Equal(t, nil, err)

	c.ImageStreamTag(restConfig)
	_, err = imagev1.NewForConfig(restConfig.Config)
	assert.Equal(t, nil, err)

	c.Project(restConfig)
	_, err = imagev1.NewForConfig(restConfig.Config)
	assert.Equal(t, nil, err)

	c.RoleBinding(restConfig)
	_, err = imagev1.NewForConfig(restConfig.Config)
	assert.Equal(t, nil, err)

	c.Route(restConfig)
	_, err = imagev1.NewForConfig(restConfig.Config)
	assert.Equal(t, nil, err)

	c.BuildConfig(restConfig)
	_, err = imagev1.NewForConfig(restConfig.Config)
	assert.Equal(t, nil, err)
}
