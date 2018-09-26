package openshift

import (
	"github.com/hidevopsio/hiboot/pkg/log"
	"github.com/openshift/api/image/v1"
	imagev1 "github.com/openshift/client-go/image/clientset/versioned/typed/image/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ApiVersion = "apps.openshift.io/v1"
	Kind       = "ImageStreamTag"
)

type ImageTagInterface interface {
	Create(fromNamespace string) (*v1.ImageStreamTag, error)
	Get() (*v1.ImageStreamTag, error)
	Delete() error
	Update(fromNamespace string) (*v1.ImageStreamTag, error)
}

type ImageStreamTag struct {
	clientSet imagev1.ImageV1Interface
}

func newImageStreamTags(clientSet imagev1.ImageV1Interface) *ImageStreamTag {
	return &ImageStreamTag{
		clientSet: clientSet,
	}
}

func (ist *ImageStreamTag) Create(fullName, namespace, fromNamespace string) (*v1.ImageStreamTag, error) {
	log.Debug("ImageStreamTag Create")
	imageTag := &v1.ImageStreamTag{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fullName,
			Namespace: namespace,
		},
		Tag: &v1.TagReference{
			From: &corev1.ObjectReference{
				Name:      fullName,
				Namespace: fromNamespace,
				Kind:      Kind,
			},
		},
	}
	_, err := ist.Get(fullName, namespace)
	if err != nil {
		img, err := ist.clientSet.ImageStreamTags(namespace).Create(imageTag)
		log.Debug("image.tag.create", err)
		return img, err
	}
	err = ist.Delete(fullName, namespace)
	if err != nil {
		log.Error("")
		return nil, err
	}
	img, err := ist.clientSet.ImageStreamTags(namespace).Create(imageTag)
	log.Debug("images.tag.update", err)
	return img, err
}

func (ist *ImageStreamTag) Get(fullName, namespace string) (*v1.ImageStreamTag, error) {
	log.Debug("ImageStreamTag Get")
	option := metav1.GetOptions{
		TypeMeta: metav1.TypeMeta{
			APIVersion: ApiVersion,
			Kind:       Kind,
		},
		IncludeUninitialized: true,
	}
	img, err := ist.clientSet.ImageStreamTags(namespace).Get(fullName, option)
	if err != nil {
		log.Println("imageStreamTag get", err)
		return nil, err
	}
	return img, nil
}

func (ist *ImageStreamTag) Delete(fullName, namespace string) error {
	log.Debug("ImageStreamTag Delete")
	meta := &metav1.DeleteOptions{}
	err := ist.clientSet.ImageStreamTags(namespace).Delete(fullName, meta)
	return err
}

func (ist *ImageStreamTag) Update(fullName, namespace, fromNamespace string) (*v1.ImageStreamTag, error) {
	img := &v1.ImageStreamTag{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fullName,
			Namespace: namespace,
		},
		Tag: &v1.TagReference{
			From: &corev1.ObjectReference{
				Name:      fullName,
				Namespace: fromNamespace,
				Kind:      Kind,
			},
		},
	}
	return ist.clientSet.ImageStreamTags(namespace).Update(img)
}
