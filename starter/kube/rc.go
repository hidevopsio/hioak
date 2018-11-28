package kube

import (
	"fmt"
	"hidevops.io/hiboot/pkg/log"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

type ReplicationController struct {
	clientSet kubernetes.Interface
}

// NewReplicationController ConfigMaps initialize construction
func NewReplicationController(clientSet kubernetes.Interface) *ReplicationController {
	return &ReplicationController{
		clientSet: clientSet,
	}
}

func (rc *ReplicationController) Create(name, namespace string, replicas int32) (*corev1.ReplicationController, error) {
	crc := &corev1.ReplicationController{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Labels: map[string]string{
				"app": name,
			},
		},
		Spec: corev1.ReplicationControllerSpec{
			Replicas: &replicas,
			Template: &corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{},
			},
		},
	}

	return rc.clientSet.CoreV1().ReplicationControllers(namespace).Create(crc)
}

func (rc *ReplicationController) Watch(name, namespace string, completedHandler func() error) error {
	w, err := rc.clientSet.CoreV1().ReplicationControllers(namespace).Watch(metav1.ListOptions{
		LabelSelector: "app=" + name,
		Watch:         true,
	})

	if err != nil {
		return err
	}

	for {
		select {
		case event, ok := <-w.ResultChan():
			if !ok {
				log.Errorf("failed on RC watching %v", ok)
				return fmt.Errorf("failed on RC watching %v", ok)
			}
			switch event.Type {
			case watch.Added:
				//log.Info("Added: ", event.Object)
				//o := event.Object
				rc := event.Object.(*corev1.ReplicationController)
				log.Debug(rc.Name)
			case watch.Modified:
				rc := event.Object.(*corev1.ReplicationController)
				log.Debugf("RC: %s, Replicas: %d, AvailableReplicas: %d", rc.Name, rc.Status.Replicas, rc.Status.AvailableReplicas)
				if rc.Status.Replicas != 0 && rc.Status.AvailableReplicas == rc.Status.Replicas {
					var err error
					if nil != completedHandler {
						err = completedHandler()
					}
					w.Stop()
					return err
				}

			case watch.Deleted:
				log.Info("Deleted: ", event.Object)
			default:
				log.Error("Failed")
			}
		}
	}
}

func (rc *ReplicationController) Delete(name, namespace string, option *metav1.DeleteOptions) (err error) {
	err = rc.clientSet.CoreV1().ReplicationControllers(namespace).Delete(name, option)
	return err
}
