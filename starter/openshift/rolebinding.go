package openshift

import (
	authorization_v1 "github.com/openshift/api/authorization/v1"
	"github.com/openshift/client-go/authorization/clientset/versioned/typed/authorization/v1"
	"hidevops.io/hiboot/pkg/log"
	corev1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type RoleBinding struct {
	clientSet v1.AuthorizationV1Interface
}

func newRoleBinding(client v1.AuthorizationV1Interface) *RoleBinding {
	log.Debug("NewPolicy()")
	r := &RoleBinding{
		clientSet: client,
	}
	return r
}

func (rb *RoleBinding) Get(name, namespace string) (*authorization_v1.RoleBinding, error) {
	log.Debug("get RoleBinding:")
	role, err := rb.clientSet.RoleBindings(namespace).Get(name, meta_v1.GetOptions{})
	if err != nil {
		log.Error("get policy err :", err)
		return nil, err
	}
	return role, nil
}

func (rb *RoleBinding) Create(name, namespace, roleRefName, roleRefKind, subjectKind, subjectName string) (*authorization_v1.RoleBinding, error) {
	log.Debug("create role binding")
	roleBinding := &authorization_v1.RoleBinding{
		ObjectMeta: meta_v1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		RoleRef: corev1.ObjectReference{
			Name: roleRefName,
			Kind: roleRefKind,
		},
		Subjects: []corev1.ObjectReference{
			{
				Kind:      subjectKind,
				Name:      subjectName,
				Namespace: namespace,
			},
		},
	}
	_, err := rb.clientSet.RoleBindings(namespace).Get(name, meta_v1.GetOptions{})
	if err == nil {
		result, err := rb.Update(name, namespace, roleRefName, roleRefKind, subjectKind, subjectName)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	result, err := rb.clientSet.RoleBindings(namespace).Create(roleBinding)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (rb *RoleBinding) Delete(name, namespace string) error {
	log.Debug("get RoleBinding:")
	err := rb.clientSet.RoleBindings(namespace).Delete(name, &meta_v1.DeleteOptions{})
	return err
}

func (rb *RoleBinding) Update(name, namespace, roleRefName, roleRefKind, subjectKind, subjectName string) (*authorization_v1.RoleBinding, error) {
	log.Debug("get RoleBinding:")
	roleBinding := &authorization_v1.RoleBinding{
		ObjectMeta: meta_v1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		RoleRef: corev1.ObjectReference{
			Name: roleRefName,
			Kind: roleRefKind,
		},
		Subjects: []corev1.ObjectReference{
			{
				Kind:      subjectKind,
				Name:      subjectName,
				Namespace: namespace,
			},
		},
	}
	result, err := rb.clientSet.RoleBindings(namespace).Update(roleBinding)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (rb *RoleBinding) InitImagePullers(namespace, roleRefName, roleRefKind, subjectKind, subjectName string) error {
	name := "system:image-pullers"
	roleBinding := &authorization_v1.RoleBinding{
		ObjectMeta: meta_v1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		RoleRef: corev1.ObjectReference{
			Name: "system:image-puller",
			Kind: "ClusterRole",
		},
		Subjects: []corev1.ObjectReference{
			{
				Kind:      "Group",
				Name:      "system:serviceaccounts:" + namespace,
				Namespace: namespace,
			},
		},
	}
	_, err := rb.clientSet.RoleBindings(namespace).Create(roleBinding)
	return err
}

func (rb *RoleBinding) InitImageBuilders(namespace, roleRefName, roleRefKind, subjectKind, subjectName string) error {
	name := "system:image-builders"
	roleBinding := &authorization_v1.RoleBinding{
		ObjectMeta: meta_v1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		RoleRef: corev1.ObjectReference{
			Name: "system:image-builder",
			Kind: "ClusterRole",
		},
		Subjects: []corev1.ObjectReference{
			{
				Kind:      "ServiceAccount",
				Name:      "builder",
				Namespace: namespace,
			},
		},
	}
	_, err := rb.clientSet.RoleBindings(namespace).Create(roleBinding)
	return err
}

func (rb *RoleBinding) InitSystemDeployers(namespace, roleRefName, roleRefKind, subjectKind, subjectName string) error {
	name := "system:deployers"
	roleBinding := &authorization_v1.RoleBinding{
		ObjectMeta: meta_v1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		RoleRef: corev1.ObjectReference{
			Name: "system:deployer",
			Kind: "ClusterRole",
		},
		Subjects: []corev1.ObjectReference{
			{
				Kind:      "ServiceAccount",
				Name:      "deployer",
				Namespace: namespace,
			},
		},
	}
	_, err := rb.clientSet.RoleBindings(namespace).Create(roleBinding)
	return err
}
