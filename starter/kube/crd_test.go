package kube

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/fake"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)
//TestCustomResourceDefinition_Create test crd ORCD
func TestCustomResourceDefinition_Create(t *testing.T) {
	clientSet := fake.NewSimpleClientset()
	client := NewCustomResourceDefinition(clientSet)

	crdName := "foos.samplecontroller.k8s.io"
	crd := &apiextensionsv1beta1.CustomResourceDefinition{
		ObjectMeta: meta_v1.ObjectMeta{Name: crdName},
		Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
			Group:   "samplecontroller.k8s.io",
			Version: "v1alpha1",
			Scope:   apiextensionsv1beta1.NamespaceScoped,
			Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
				Plural: "foos",
				Kind:   "Foo",
			},
		},
	}

	t.Run("should crd create success", func(t *testing.T) {
		_, err := client.CreateCRD(crd)
		assert.Equal(t, nil, err)
	})

	t.Run("should crd get success", func(t *testing.T) {
		crd, err := client.GetCRD(crdName, meta_v1.GetOptions{})
		fmt.Println(crd)
		assert.Equal(t, nil, err)
	})

	t.Run("should crd list success", func(t *testing.T) {
		_, err := client.CRDList(meta_v1.ListOptions{})
		assert.Equal(t, nil, err)
	})

	t.Run("should crd update success", func(t *testing.T) {
		_, err := client.UpdateCRD(crd)
		assert.Equal(t, nil, err)
	})

	t.Run("should crd watch success", func(t *testing.T) {
		_, err := client.WatchCRD(meta_v1.ListOptions{})
		assert.Equal(t, nil, err)
	})

	t.Run("should crd delete success", func(t *testing.T) {
		err := client.DeleteCRD(crdName, meta_v1.DeleteOptions{})
		assert.Equal(t, nil, err)
	})
}
