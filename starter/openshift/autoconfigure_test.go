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

	rest := &kube.RestConfig{
		Config: &rest.Config{
			BearerToken: "",
		},
	}
	rest.GroupVersion = nil

	c.Auth(restConfig)
	c.Auth(rest)
	_, err := oauthv1.NewForConfig(restConfig.Config)
	assert.Equal(t, nil, err)

	c.DeploymentConfig(restConfig)
	c.DeploymentConfig(rest)
	_, err = appsv1.NewForConfig(restConfig.Config)
	assert.Equal(t, nil, err)

	c.ImageStream(restConfig)
	c.ImageStream(rest)
	_, err = imagev1.NewForConfig(restConfig.Config)
	assert.Equal(t, nil, err)

	c.ImageStreamTag(restConfig)
	c.ImageStreamTag(rest)
	_, err = imagev1.NewForConfig(restConfig.Config)
	assert.Equal(t, nil, err)

	c.Project(restConfig)
	c.Project(rest)
	_, err = imagev1.NewForConfig(restConfig.Config)
	assert.Equal(t, nil, err)

	c.RoleBinding(restConfig)
	c.RoleBinding(rest)
	_, err = imagev1.NewForConfig(restConfig.Config)
	assert.Equal(t, nil, err)

	c.Route(restConfig)
	c.Route(rest)
	_, err = imagev1.NewForConfig(restConfig.Config)
	assert.Equal(t, nil, err)

	c.BuildConfig(restConfig)
	c.BuildConfig(rest)
	_, err = imagev1.NewForConfig(restConfig.Config)
	assert.Equal(t, nil, err)

	auth := c.Auth(nil)
	assert.Equal(t, (*OAuthAccessToken)(nil) , auth)

	deploymentConfig := c.DeploymentConfig(nil)
	assert.Equal(t, (*DeploymentConfig)(nil) , deploymentConfig)

	imageStream := c.ImageStream(nil)
	assert.Equal(t, (*ImageStream)(nil) , imageStream)

	imageStreamTag := c.ImageStreamTag(nil)
	assert.Equal(t, (*ImageStreamTag)(nil) , imageStreamTag)

	project := c.Project(nil)
	assert.Equal(t, (*Project)(nil) , project)

	roleBinding := c.RoleBinding(nil)
	assert.Equal(t, (*RoleBinding)(nil) , roleBinding)

	route := c.Route(nil)
	assert.Equal(t, (*Route)(nil) , route)

	buildConfig := c.BuildConfig(nil)
	assert.Equal(t, (*BuildConfig)(nil) , buildConfig)

}
