package kube

import (
	"fmt"
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

// GetEvents takes name of the event, and returns the corresponding event object, and an error if there is any.
func (e *Events) GetEvents(namespace, name string, opts v1.GetOptions) (*v1beta1.Event, error) {
	log.Infof(fmt.Sprintf("get events %s in namespace %s", name, namespace))
	event, err := e.clientSet.EventsV1beta1().Events(namespace).Get(name, opts)
	if err != nil {
		return nil, err
	}
	return event, nil
}

// EventsList takes label and field selectors, and returns the list of Events that match those selectors.
func (e *Events) EventsList(namespace string, opts v1.ListOptions) (*v1beta1.EventList, error) {
	log.Infof(fmt.Sprintf("get events list by label %s in namespace %s", opts.LabelSelector, namespace))
	eventList, err := e.clientSet.EventsV1beta1().Events(namespace).List(opts)
	if err != nil {
		return nil, err
	}
	return eventList, nil
}

// EventsWatch returns a watch.Interface that watches the requested events.
func (e *Events) EventsWatch(namespace string, opts v1.ListOptions) (watch.Interface, error) {
	log.Infof(fmt.Sprintf("watch events by label %s in namespace %s", opts.LabelSelector, namespace))
	w, err := e.clientSet.EventsV1beta1().Events(namespace).Watch(opts)
	if err != nil {
		return nil, err
	}
	return w, nil
}

// CreateEvents takes the representation of a event and creates it.  Returns the server's representation of the event, and an error, if there is any.
func (e *Events) CreateEvents(event *v1beta1.Event) (*v1beta1.Event, error) {
	log.Infof(fmt.Sprintf("create events %s in namespace %s", event.Name, event.Namespace))
	event, err := e.clientSet.EventsV1beta1().Events(event.Namespace).Create(event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

// UpdateEvents takes the representation of a event and updates it. Returns the server's representation of the event, and an error, if there is any.
func (e *Events) UpdateEvents(event *v1beta1.Event) (*v1beta1.Event, error) {
	log.Infof(fmt.Sprintf("update events %s in namespace %s", event.Name, event.Namespace))
	event, err := e.clientSet.EventsV1beta1().Events(event.Namespace).Update(event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

// DeleteEvents takes name of the event and deletes it. Returns an error if one occurs.
func (e *Events) DeleteEvents(namespace, name string, opts *v1.DeleteOptions) error {
	log.Infof(fmt.Sprintf("delete events %s in namespace %s", name, namespace))
	err := e.clientSet.EventsV1beta1().Events(namespace).Delete(name, opts)
	if err != nil {
		return err
	}
	return nil
}
