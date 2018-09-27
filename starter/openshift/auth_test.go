package openshift

import (
	"github.com/magiconair/properties/assert"
	"github.com/openshift/client-go/oauth/clientset/versioned/fake"
	"github.com/prometheus/common/log"
	"testing"
)

func TestOAuthAccessToken_Get(t *testing.T) {
	clientSet := fake.NewSimpleClientset().OauthV1()
	token := NewOAuthAccessToken(clientSet)
	to, err := token.Create()
	to, err = token.Get(to.Name)
	assert.Equal(t, nil, err)
	log.Info(to)
}

func TestOAuthAccessToken_List(t *testing.T) {
	clientSet := fake.NewSimpleClientset().OauthV1()
	token := NewOAuthAccessToken(clientSet)
	to, err := token.List()
	assert.Equal(t, nil, err)
	log.Info(to)
}
