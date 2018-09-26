package openshift

import (
	"github.com/hidevopsio/hioak/starter"
	"github.com/openshift/client-go/oauth/clientset/versioned/fake"
	oauthv1 "github.com/openshift/client-go/oauth/clientset/versioned/typed/oauth/v1"
	"github.com/prometheus/common/log"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/openshift/api/oauth/v1"
)

type OAuthAccessToken struct {
	Interface oauthv1.OAuthAccessTokenInterface

}

func NewOAuthAccessTokenClientSet() (oauthv1.OauthV1Interface, error) {

	cli := orch.GetClientInstance()
	// get the fake ClientSet for testing
	if cli.IsTestRunning() {
		return fake.NewSimpleClientset().OauthV1(), nil
	}

	// get the real ClientSet
	clientSet, err := oauthv1.NewForConfig(cli.Config())

	return clientSet, err
}

func NewOAuthAccessToken(clientSet oauthv1.OauthV1Interface) (*OAuthAccessToken, error) {
	return &OAuthAccessToken{
		Interface: clientSet.OAuthAccessTokens(),
	}, nil
}

func (o *OAuthAccessToken) Create() (*v1.OAuthAccessToken, error) {
	log.Debug("openshift get OAuthAccessToken")
	token := &v1.OAuthAccessToken{

	}
	token, err := o.Interface.Create(token)
	return token, err
}

func (o *OAuthAccessToken) Get(name string) (*v1.OAuthAccessToken, error) {
	log.Debug("openshift get OAuthAccessToken")
	opt := meta_v1.GetOptions{}
	token, err := o.Interface.Get(name, opt)
	return token, err
}

func (o *OAuthAccessToken) List() (*v1.OAuthAccessTokenList, error) {
	log.Debug("openshift get OAuthAccessToken")
	opt := meta_v1.ListOptions{}
	tokens, err := o.Interface.List(opt)
	return tokens, err
}

func GetToken()  {
	clientSet, err := NewOAuthAccessTokenClientSet()
	log.Info(err)
	clientSet.RESTClient()
}