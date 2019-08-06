package openshift

import (
	"github.com/openshift/api/oauth/v1"
	oauthv1 "github.com/openshift/client-go/oauth/clientset/versioned/typed/oauth/v1"
	"hidevops.io/hiboot/pkg/log"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type OAuthAccessToken struct {
	Interface oauthv1.OAuthAccessTokenInterface
}

func NewOAuthAccessToken(clientSet oauthv1.OauthV1Interface) *OAuthAccessToken {
	return &OAuthAccessToken{
		Interface: clientSet.OAuthAccessTokens(),
	}
}

func (o *OAuthAccessToken) Create() (*v1.OAuthAccessToken, error) {
	log.Debug("openshift get OAuthAccessToken")
	token := &v1.OAuthAccessToken{}
	token, err := o.Interface.Create(token)
	return token, err
}

func (o *OAuthAccessToken) Get(name string) (*v1.OAuthAccessToken, error) {
	log.Debug("openshift get OAuthAccessToken")
	opt := metaV1.GetOptions{}
	token, err := o.Interface.Get(name, opt)
	return token, err
}

func (o *OAuthAccessToken) List() (*v1.OAuthAccessTokenList, error) {
	log.Debug("openshift get OAuthAccessToken")
	opt := metaV1.ListOptions{}
	tokens, err := o.Interface.List(opt)
	return tokens, err
}
