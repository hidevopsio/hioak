package kube

import (
	"fmt"
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

// CreateCRD takes the representation of a customResourceDefinition and creates it.  Returns the server's representation of the customResourceDefinition, and an error, if there is any.
func (p *CustomResourceDefinition) CreateCRD(crd *apiextensionsv1beta1.CustomResourceDefinition) (*apiextensionsv1beta1.CustomResourceDefinition, error) {
	log.Debugf("create CustomResourceDefinition")
	crd, err := p.clientSet.ApiextensionsV1beta1().CustomResourceDefinitions().Create(crd)

	if err != nil {
		return nil, err
	}
	return crd, nil
}

// WatchCRD returns a watch.Interface that watches the requested customResourceDefinitions.
func (p *CustomResourceDefinition) WatchCRD(listOptions meta_v1.ListOptions) (watch.Interface, error) {
	log.Infof(fmt.Sprintf("watch label for %s CustomResourceDefinition", listOptions.LabelSelector))
	listOptions.Watch = true
	w, err := p.clientSet.ApiextensionsV1beta1().CustomResourceDefinitions().Watch(listOptions)
	if err != nil {
		return nil, err
	}
	return w, nil
}

// GetCRD takes name of the customResourceDefinition, and returns the corresponding customResourceDefinition object, and an error if there is any.
func (p *CustomResourceDefinition) GetCRD(name string, opts meta_v1.GetOptions) (*apiextensionsv1beta1.CustomResourceDefinition, error) {
	log.Infof(fmt.Sprintf("get CustomResourceDefinition %s ", name))

	crd, err := p.clientSet.ApiextensionsV1beta1().CustomResourceDefinitions().Get(name, opts)
	if err != nil {
		return nil, err
	}
	return crd, nil
}

// CRDList takes label and field selectors, and returns the list of CustomResourceDefinitions that match those selectors.
func (p *CustomResourceDefinition) CRDList(opts meta_v1.ListOptions) (*apiextensionsv1beta1.CustomResourceDefinitionList, error) {
	log.Infof(fmt.Sprintf("get CustomResourceDefinition list"))

	crdList, err := p.clientSet.ApiextensionsV1beta1().CustomResourceDefinitions().List(opts)
	if err != nil {
		return nil, err
	}
	return crdList, nil
}

// DeleteCRD takes name of the customResourceDefinition and deletes it. Returns an error if one occurs.
func (p *CustomResourceDefinition) DeleteCRD(name string, opts meta_v1.DeleteOptions) error {
	log.Infof(fmt.Sprintf("delete CustomResourceDefinition %s ", name))

	if err := p.clientSet.ApiextensionsV1beta1().CustomResourceDefinitions().Delete(name, &opts); err != nil {
		return err
	}
	return nil
}

// UpdateCRD takes the representation of a customResourceDefinition and updates it. Returns the server's representation of the customResourceDefinition, and an error, if there is any.
func (p *CustomResourceDefinition) UpdateCRD(crd *apiextensionsv1beta1.CustomResourceDefinition) (*apiextensionsv1beta1.CustomResourceDefinition, error) {
	log.Infof(fmt.Sprintf("update CustomResourceDefinition %s ", crd.Name))

	crd, err := p.clientSet.ApiextensionsV1beta1().CustomResourceDefinitions().Update(crd)
	if err != nil {
		return nil, err
	}
	return crd, nil
}
