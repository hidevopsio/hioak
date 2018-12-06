package kube

import (
	"hidevops.io/hiboot/pkg/log"
	"k8s.io/api/events/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

// Events implements EventInterface
type Events struct {
	clientSet kubernetes.Interface
}

// NewEvents returns a clientSet
func NewEvents(clientSet kubernetes.Interface) *Events {
	return &Events{
		clientSet: clientSet,
	}
}

// Get takes name of the event, and returns the corresponding event object, and an error if there is any.
func (e *Events) Get(namespace, name string, opts v1.GetOptions) (*v1beta1.Event, error) {
	log.Debugf("get events %s in namespace %s", name, namespace)
	return e.clientSet.EventsV1beta1().Events(namespace).Get(name, opts)
}

// List takes label and field selectors, and returns the list of Events that match those selectors.
func (e *Events) List(namespace string, opts v1.ListOptions) (*v1beta1.EventList, error) {
	log.Debugf("get events list by label %s in namespace %s", opts.LabelSelector, namespace)
	return e.clientSet.EventsV1beta1().Events(namespace).List(opts)
}

// Watch returns a watch.Interface that watches the requested events.
func (e *Events) Watch(namespace string, opts v1.ListOptions) (watch.Interface, error) {
	log.Debugf("watch events by label %s in namespace %s", opts.LabelSelector, namespace)
	return e.clientSet.EventsV1beta1().Events(namespace).Watch(opts)
}

// Create takes the representation of a event and creates it.  Returns the server's representation of the event, and an error, if there is any.
func (e *Events) Create(event *v1beta1.Event) (*v1beta1.Event, error) {
	log.Debugf("create events %s in namespace %s", event.Name, event.Namespace)
	return e.clientSet.EventsV1beta1().Events(event.Namespace).Create(event)
}

// Update takes the representation of a event and updates it. Returns the server's representation of the event, and an error, if there is any.
func (e *Events) Update(event *v1beta1.Event) (*v1beta1.Event, error) {
	log.Debugf("update events %s in namespace %s", event.Name, event.Namespace)
	return e.clientSet.EventsV1beta1().Events(event.Namespace).Update(event)
}

// Delete takes name of the event and deletes it. Returns an error if one occurs.
func (e *Events) Delete(namespace, name string, opts *v1.DeleteOptions) error {
	log.Debugf("delete events %s in namespace %s", name, namespace)
	return e.clientSet.EventsV1beta1().Events(namespace).Delete(name, opts)
}
