package openshift

import (
	"github.com/magiconair/properties/assert"
	"github.com/openshift/client-go/image/clientset/versioned/fake"
	"testing"
)

const (
	name          = "hiweb"
	namespace     = "hidevopsio"
	fromNamespace = "hidevopsio-dev"
	version       = "v1"
	fullName      = name + ":" + version
)

func TestCrudTags(t *testing.T) {
	clientSet := fake.NewSimpleClientset().ImageV1()
	ist := newImageStreamTags(clientSet)
	is, err := ist.Create(fullName, namespace, fromNamespace)
	assert.Equal(t, nil, err)
	assert.Equal(t, fullName, is.Name)

	is, err = ist.Get(fullName, namespace)
	assert.Equal(t, nil, err)
	assert.Equal(t, fullName, is.Name)

	_, err = ist.Update(fullName, namespace, fromNamespace)
	assert.Equal(t, nil, err)

	err = ist.Delete(fullName, namespace)
	assert.Equal(t, nil, err)

}
