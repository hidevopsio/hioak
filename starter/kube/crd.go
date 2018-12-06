package kube

import (
	"hidevops.io/hiboot/pkg/log"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

// CustomResourceDefinition implements CustomResourceDefinitionInterface
type CustomResourceDefinition struct {
	clientSet apiextensionsclient.Interface
}

// NewCustomResourceDefinition returns a clientSet
func NewCustomResourceDefinition(clientSet apiextensionsclient.Interface) *CustomResourceDefinition {
	return &CustomResourceDefinition{
		clientSet: clientSet,
	}
}

// Create takes the representation of a customResourceDefinition and creates it.  Returns the server's representation of the customResourceDefinition, and an error, if there is any.
func (p *CustomResourceDefinition) Create(crd *apiextensionsv1beta1.CustomResourceDefinition) (*apiextensionsv1beta1.CustomResourceDefinition, error) {
	log.Debugf("create CustomResourceDefinition")
	return p.clientSet.ApiextensionsV1beta1().CustomResourceDefinitions().Create(crd)
}

// Watch returns a watch.Interface that watches the requested customResourceDefinitions.
func (p *CustomResourceDefinition) Watch(listOptions meta_v1.ListOptions) (watch.Interface, error) {
	log.Debugf("watch label for %s CustomResourceDefinition", listOptions.LabelSelector)
	return p.clientSet.ApiextensionsV1beta1().CustomResourceDefinitions().Watch(listOptions)
}

// Get takes name of the customResourceDefinition, and returns the corresponding customResourceDefinition object, and an error if there is any.
func (p *CustomResourceDefinition) Get(name string, opts meta_v1.GetOptions) (*apiextensionsv1beta1.CustomResourceDefinition, error) {
	log.Debugf("get CustomResourceDefinition %s ", name)
	return p.clientSet.ApiextensionsV1beta1().CustomResourceDefinitions().Get(name, opts)
}

// List takes label and field selectors, and returns the list of CustomResourceDefinitions that match those selectors.
func (p *CustomResourceDefinition) List(opts meta_v1.ListOptions) (*apiextensionsv1beta1.CustomResourceDefinitionList, error) {
	log.Debugf("get CustomResourceDefinition list")
	return p.clientSet.ApiextensionsV1beta1().CustomResourceDefinitions().List(opts)
}

// Delete takes name of the customResourceDefinition and deletes it. Returns an error if one occurs.
func (p *CustomResourceDefinition) Delete(name string, opts meta_v1.DeleteOptions) error {
	log.Debugf("delete CustomResourceDefinition %s ", name)
	return p.clientSet.ApiextensionsV1beta1().CustomResourceDefinitions().Delete(name, &opts)
}

// Update takes the representation of a customResourceDefinition and updates it. Returns the server's representation of the customResourceDefinition, and an error, if there is any.
func (p *CustomResourceDefinition) Update(crd *apiextensionsv1beta1.CustomResourceDefinition) (*apiextensionsv1beta1.CustomResourceDefinition, error) {
	log.Debugf("update CustomResourceDefinition %s ", crd.Name)
	return p.clientSet.ApiextensionsV1beta1().CustomResourceDefinitions().Update(crd)
}
