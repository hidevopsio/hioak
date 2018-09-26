package openshift

import (
	"testing"
	"github.com/magiconair/properties/assert"
	"github.com/prometheus/common/log"
	"github.com/openshift/client-go/oauth/clientset/versioned/fake"
)

func TestOAuthAccessToken_Get(t *testing.T) {
	clientSet := fake.NewSimpleClientset().OauthV1()
	token, err := NewOAuthAccessToken(clientSet)
	assert.Equal(t, nil, err)
	to, err := token.Create()
	to, err = token.Get(to.Name)
	assert.Equal(t, nil, err)
	log.Info(to)
}

func TestOAuthAccessToken_List(t *testing.T) {
	clientSet := fake.NewSimpleClientset().OauthV1()
	token, err := NewOAuthAccessToken(clientSet)
	assert.Equal(t, nil, err)
	to, err := token.List()
	assert.Equal(t, nil, err)
	log.Info(to)
}