package openshift

import (
	"github.com/hidevopsio/hiboot/pkg/app"
	"github.com/hidevopsio/hioak/starter/kube"
	appsv1 "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	authorizationv1 "github.com/openshift/client-go/authorization/clientset/versioned/typed/authorization/v1"
	buildv1 "github.com/openshift/client-go/build/clientset/versioned/typed/build/v1"
	imagev1 "github.com/openshift/client-go/image/clientset/versioned/typed/image/v1"
	oauthv1 "github.com/openshift/client-go/oauth/clientset/versioned/typed/oauth/v1"
	projectv1 "github.com/openshift/client-go/project/clientset/versioned/typed/project/v1"
	routev1 "github.com/openshift/client-go/route/clientset/versioned/typed/route/v1"
	"github.com/prometheus/common/log"
)

type configuration struct {
	app.Configuration `depends:"kube"`
}

type Oauth interface {
	oauthv1.OauthV1Interface
}

func (c *configuration) Auth(restConfig *kube.RestConfig) (retVal *OAuthAccessToken) {
	clientSet, err := oauthv1.NewForConfig(restConfig.Config)
	if err != nil {
		log.Errorf("oauthv1.NewForConfig %v", err)
		return
	}
	retVal = NewOAuthAccessToken(clientSet)
	return
}

func (c *configuration) Token(restConfig *kube.RestConfig) Token {
	return Token(restConfig.Config.BearerToken)
}

func (c *configuration) DeploymentConfig(restConfig *kube.RestConfig) (retVal *DeploymentConfig) {
	clientSet, err := appsv1.NewForConfig(restConfig.Config)
	if err != nil {
		log.Errorf("appsv1.NewForConfig %v", err)
		return
	}
	retVal = newDeploymentConfig(clientSet)
	return
}

func (c *configuration) ImageStream(restConfig *kube.RestConfig) (retVal *ImageStream) {
	clientSet, err := imagev1.NewForConfig(restConfig.Config)
	if err != nil {
		log.Errorf("imagev1.NewForConfig %v", err)
		return
	}
	retVal = newImageStream(clientSet)
	return
}

func (c *configuration) ImageStreamTag(restConfig *kube.RestConfig) (retVal *ImageStreamTag) {
	clientSet, err := imagev1.NewForConfig(restConfig.Config)
	if err != nil {
		log.Errorf("imagev1.NewForConfig %v", err)
		return
	}
	retVal = newImageStreamTags(clientSet)
	return
}

func (c *configuration) Project(restConfig *kube.RestConfig) (retVal *Project) {
	clientSet, err := projectv1.NewForConfig(restConfig.Config)
	if err != nil {
		log.Errorf("projectv1.NewForConfig %v", err)
		return
	}
	retVal = newProject(clientSet)
	return
}

func (c *configuration) RoleBinding(restConfig *kube.RestConfig) (retVal *RoleBinding) {
	clientSet, err := authorizationv1.NewForConfig(restConfig.Config)
	if err != nil {
		log.Errorf("authorizationv1.NewForConfig %v", err)
		return
	}
	retVal = newRoleBinding(clientSet)
	return
}

func (c *configuration) Route(restConfig *kube.RestConfig) (retVal *Route) {
	clientSet, err := routev1.NewForConfig(restConfig.Config)
	if err != nil {
		log.Errorf("routev1.NewForConfig %v", err)
		return
	}
	retVal = newRoute(clientSet)
	return
}

func (c *configuration) BuildConfig(restConfig *kube.RestConfig) (retVal *BuildConfig) {
	clientSet, err := buildv1.NewForConfig(restConfig.Config)
	if err != nil {
		log.Errorf("routev1.NewForConfig %v", err)
		return
	}
	retVal = newBuildConfig(clientSet)
	return
}
