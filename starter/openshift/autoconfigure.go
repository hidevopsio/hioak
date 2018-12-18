package openshift

import (
	appsv1 "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	authorizationv1 "github.com/openshift/client-go/authorization/clientset/versioned/typed/authorization/v1"
	buildv1 "github.com/openshift/client-go/build/clientset/versioned/typed/build/v1"
	imagev1 "github.com/openshift/client-go/image/clientset/versioned/typed/image/v1"
	oauthv1 "github.com/openshift/client-go/oauth/clientset/versioned/typed/oauth/v1"
	projectv1 "github.com/openshift/client-go/project/clientset/versioned/typed/project/v1"
	routev1 "github.com/openshift/client-go/route/clientset/versioned/typed/route/v1"
	"hidevops.io/hiboot/pkg/app"
	"hidevops.io/hiboot/pkg/at"
	"hidevops.io/hioak/starter/kube"
)

type configuration struct {
	at.AutoConfiguration
}

type Oauth interface {
	oauthv1.OauthV1Interface
}

func init() {
	app.Register(newConfiguration)
}


func newConfiguration() *configuration {
	return &configuration{}
}

func (c *configuration) Auth(restConfig *kube.RestConfig) (retVal *OAuthAccessToken) {
	if restConfig != nil {
		clientSet, _ := oauthv1.NewForConfig(restConfig.Config)
		retVal = NewOAuthAccessToken(clientSet)
		return
	}
	return
}

func (c *configuration) DeploymentConfig(restConfig *kube.RestConfig) (retVal *DeploymentConfig) {
	if restConfig != nil {
		clientSet, _ := appsv1.NewForConfig(restConfig.Config)
		retVal = newDeploymentConfig(clientSet)
		return
	}
	return
}

func (c *configuration) ImageStream(restConfig *kube.RestConfig) (retVal *ImageStream) {
	if restConfig != nil {
		clientSet, _ := imagev1.NewForConfig(restConfig.Config)
		retVal = newImageStream(clientSet)
		return
	}
	return
}

func (c *configuration) ImageStreamTag(restConfig *kube.RestConfig) (retVal *ImageStreamTag) {
	if restConfig != nil {
		clientSet, _ := imagev1.NewForConfig(restConfig.Config)
		retVal = newImageStreamTags(clientSet)
		return
	}
	return
}

func (c *configuration) Project(restConfig *kube.RestConfig) (retVal *Project) {
	if restConfig != nil {
		clientSet, _ := projectv1.NewForConfig(restConfig.Config)
		retVal = newProject(clientSet)
		return
	}
	return
}

func (c *configuration) RoleBinding(restConfig *kube.RestConfig) (retVal *RoleBinding) {
	if restConfig != nil {
		clientSet, _ := authorizationv1.NewForConfig(restConfig.Config)
		retVal = newRoleBinding(clientSet)
		return
	}
	return
}

func (c *configuration) Route(restConfig *kube.RestConfig) (retVal *Route) {
	if restConfig != nil {
		clientSet, _ := routev1.NewForConfig(restConfig.Config)
		retVal = newRoute(clientSet)
		return
	}
	return
}

func (c *configuration) BuildConfig(restConfig *kube.RestConfig) (retVal *BuildConfig) {
	if restConfig != nil {
		clientSet, _ := buildv1.NewForConfig(restConfig.Config)
		retVal = newBuildConfig(clientSet)
		return
	}
	return
}
