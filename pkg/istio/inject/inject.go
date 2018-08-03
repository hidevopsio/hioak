// Copyright 2018 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package inject

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"reflect"
	"strconv"
	"strings"
	"text/template"

	"github.com/ghodss/yaml"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/duration"
	"k8s.io/api/batch/v2alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	yamlDecoder "k8s.io/apimachinery/pkg/util/yaml"

	meshconfig "istio.io/api/mesh/v1alpha1"
	"istio.io/istio/pkg/log"
)

// per-sidecar policy and status
const (
	sidecarAnnotationPolicyKey                        = "sidecar.istio.io/inject"
	sidecarAnnotationStatusKey                        = "sidecar.istio.io/status"
	sidecarAnnotationProxyImageOverride               = "sidecar.istio.io/proxyImage"
	sidecarAnnotationInterceptionModeKey              = "sidecar.istio.io/interceptionMode"
	sidecarAnnotationIncludeOutboundIPRangesPolicyKey = "traffic.sidecar.istio.io/includeOutboundIPRanges"
	sidecarAnnotationExcludeOutboundIPRangesPolicyKey = "traffic.sidecar.istio.io/excludeOutboundIPRanges"
	sidecarAnnotationIncludeInboundPortsPolicyKey     = "traffic.sidecar.istio.io/includeInboundPorts"
	sidecarAnnotationExcludeInboundPortsPolicyKey     = "traffic.sidecar.istio.io/excludeInboundPorts"
)

// InjectionPolicy determines the policy for injecting the
// sidecar proxy into the watched namespace(s).
type InjectionPolicy string

const (
	// InjectionPolicyDisabled specifies that the sidecar injector
	// will not inject the sidecar into resources by default for the
	// namespace(s) being watched. Resources can enable injection
	// using the "sidecar.istio.io/inject" annotation with value of
	// true.
	InjectionPolicyDisabled InjectionPolicy = "disabled"

	// InjectionPolicyEnabled specifies that the sidecar injector will
	// inject the sidecar into resources by default for the
	// namespace(s) being watched. Resources can disable injection
	// using the "sidecar.istio.io/inject" annotation with value of
	// false.
	InjectionPolicyEnabled InjectionPolicy = "enabled"
)

// Defaults values for injecting istio proxy into kubernetes
// resources.
const (
	DefaultSidecarProxyUID     = uint64(1337)
	DefaultVerbosity           = 2
	DefaultImagePullPolicy     = "IfNotPresent"
	DefaultIncludeIPRanges     = "*"
	DefaultIncludeInboundPorts = "*"
)

const (
	// ProxyContainerName is used by e2e integration tests for fetching logs
	ProxyContainerName = "istio-proxy"
)

// SidecarInjectionSpec collects all container types and volumes for
// sidecar mesh injection
type SidecarInjectionSpec struct {
	InitContainers   []corev1.Container            `yaml:"initContainers"`
	Containers       []corev1.Container            `yaml:"containers"`
	Volumes          []corev1.Volume               `yaml:"volumes"`
	ImagePullSecrets []corev1.LocalObjectReference `yaml:"imagePullSecrets"`
}

// SidecarTemplateData is the data object to which the templated
// version of `SidecarInjectionSpec` is applied.
type SidecarTemplateData struct {
	ObjectMeta  *metav1.ObjectMeta
	Spec        *corev1.PodSpec
	ProxyConfig *meshconfig.ProxyConfig
	MeshConfig  *meshconfig.MeshConfig
}

// InitImageName returns the fully qualified image name for the istio
// init image given a docker hub and tag and debug flag
func InitImageName(hub string, tag string, _ bool) string {
	return hub + "/proxy_init:" + tag
}

// ProxyImageName returns the fully qualified image name for the istio
// proxy image given a docker hub and tag and whether to use debug or not.
func ProxyImageName(hub string, tag string, debug bool) string {
	// Allow overriding the proxy image.
	if debug {
		return hub + "/proxy_debug:" + tag
	}
	return hub + "/proxyv2:" + tag
}

// Params describes configurable parameters for injecting istio proxy
// into a kubernetes resource.
type Params struct {
	InitImage       string                 `json:"initImage"`
	ProxyImage      string                 `json:"proxyImage"`
	Verbosity       int                    `json:"verbosity"`
	SidecarProxyUID uint64                 `json:"sidecarProxyUID"`
	Version         string                 `json:"version"`
	EnableCoreDump  bool                   `json:"enableCoreDump"`
	DebugMode       bool                   `json:"debugMode"`
	Mesh            *meshconfig.MeshConfig `json:"-"`
	ImagePullPolicy string                 `json:"imagePullPolicy"`
	// Comma separated list of IP ranges in CIDR form. If set, only redirect outbound traffic to Envoy for these IP
	// ranges. All outbound traffic can be redirected with the wildcard character "*". Defaults to "*".
	IncludeIPRanges string `json:"includeIPRanges"`
	// Comma separated list of IP ranges in CIDR form. If set, outbound traffic will not be redirected for
	// these IP ranges. Exclusions are only applied if configured to redirect all outbound traffic. By default,
	// no IP ranges are excluded.
	ExcludeIPRanges string `json:"excludeIPRanges"`
	// Comma separated list of inbound ports for which traffic is to be redirected to Envoy. All ports can be
	// redirected with the wildcard character "*". Defaults to "*".
	IncludeInboundPorts string `json:"includeInboundPorts"`
	// Comma separated list of inbound ports. If set, inbound traffic will not be redirected for those ports.
	// Exclusions are only applied if configured to redirect all inbound traffic. By default, no ports are excluded.
	ExcludeInboundPorts string `json:"excludeInboundPorts"`
}

// Validate validates the parameters and returns an error if there is configuration issue.
func (p *Params) Validate() error {
	if err := ValidateIncludeIPRanges(p.IncludeIPRanges); err != nil {
		return err
	}
	if err := ValidateExcludeIPRanges(p.ExcludeIPRanges); err != nil {
		return err
	}
	if err := ValidateIncludeInboundPorts(p.IncludeInboundPorts); err != nil {
		return err
	}
	return ValidateExcludeInboundPorts(p.ExcludeInboundPorts)
}

// Config specifies the sidecar injection configuration This includes
// the sidear template and cluster-side injection policy. It is used
// by kube-inject, sidecar injector, and http endpoint.
type Config struct {
	Policy InjectionPolicy `json:"policy"`

	// Template is the templated version of `SidecarInjectionSpec` prior to
	// expansion over the `SidecarTemplateData`.
	Template string `json:"template"`
}

func validateCIDRList(cidrs string) error {
	if len(cidrs) > 0 {
		for _, cidr := range strings.Split(cidrs, ",") {
			if _, _, err := net.ParseCIDR(cidr); err != nil {
				return fmt.Errorf("failed parsing cidr '%s': %v", cidr, err)
			}
		}
	}
	return nil
}

func validatePortList(ports string) error {
	if len(ports) > 0 {
		for _, port := range strings.Split(ports, ",") {
			if _, err := strconv.ParseInt(port, 10, 16); err != nil {
				return fmt.Errorf("failed parsing port '%s': %v", port, err)
			}
		}
	}
	return nil
}

// ValidateInterceptionMode validates the interceptionMode annotation
func ValidateInterceptionMode(mode string) error {
	switch mode {
	case meshconfig.ProxyConfig_REDIRECT.String():
	case meshconfig.ProxyConfig_TPROXY.String():
	default:
		return fmt.Errorf("interceptionMode invalid: %v", mode)
	}
	return nil
}

// ValidateIncludeIPRanges validates the includeIPRanges parameter
func ValidateIncludeIPRanges(ipRanges string) error {
	if ipRanges != "*" {
		if e := validateCIDRList(ipRanges); e != nil {
			return fmt.Errorf("includeIPRanges invalid: %v", e)
		}
	}
	return nil
}

// ValidateExcludeIPRanges validates the excludeIPRanges parameter
func ValidateExcludeIPRanges(ipRanges string) error {
	if e := validateCIDRList(ipRanges); e != nil {
		return fmt.Errorf("excludeIPRanges invalid: %v", e)
	}
	return nil
}

// ValidateIncludeInboundPorts validates the includeInboundPorts parameter
func ValidateIncludeInboundPorts(ports string) error {
	if ports != "*" {
		if e := validatePortList(ports); e != nil {
			return fmt.Errorf("includeInboundPorts invalid: %v", e)
		}
	}
	return nil
}

// ValidateExcludeInboundPorts validates the excludeInboundPorts parameter
func ValidateExcludeInboundPorts(ports string) error {
	if e := validatePortList(ports); e != nil {
		return fmt.Errorf("excludeInboundPorts invalid: %v", e)
	}
	return nil
}

func injectRequired(ignored []string, namespacePolicy InjectionPolicy, podSpec *corev1.PodSpec, metadata *metav1.ObjectMeta) bool { // nolint: lll
	// Skip injection when host networking is enabled. The problem is
	// that the iptable changes are assumed to be within the pod when,
	// in fact, they are changing the routing at the host level. This
	// often results in routing failures within a node which can
	// affect the network provider within the cluster causing
	// additional pod failures.
	if podSpec.HostNetwork {
		return false
	}

	// skip special kubernetes system namespaces
	for _, namespace := range ignored {
		if metadata.Namespace == namespace {
			return false
		}
	}

	annotations := metadata.GetAnnotations()
	if annotations == nil {
		annotations = map[string]string{}
	}

	var useDefault bool
	var inject bool
	switch strings.ToLower(annotations[sidecarAnnotationPolicyKey]) {
	// http://yaml.org/type/bool.html
	case "y", "yes", "true", "on":
		inject = true
	case "":
		useDefault = true
	}

	var required bool
	switch namespacePolicy {
	default: // InjectionPolicyOff
		required = false
	case InjectionPolicyDisabled:
		if useDefault {
			required = false
		} else {
			required = inject
		}
	case InjectionPolicyEnabled:
		if useDefault {
			required = true
		} else {
			required = inject
		}
	}

	log.Debugf("Sidecar injection policy for %v/%v: namespacePolicy:%v useDefault:%v inject:%v status:%q proxyImage:%q"+
		" interceptionMode:%v required:%v"+
		" includeOutboundIPRanges:%v excludeOutboundIPRanges:%v includeInboundPorts:%v excludeInboundPorts:%v",
		metadata.Namespace,
		metadata.Name,
		namespacePolicy,
		useDefault,
		inject,
		annotations[sidecarAnnotationStatusKey],
		annotations[sidecarAnnotationProxyImageOverride],
		annotationString(annotations, sidecarAnnotationInterceptionModeKey),
		required,
		annotationString(annotations, sidecarAnnotationIncludeOutboundIPRangesPolicyKey),
		annotationString(annotations, sidecarAnnotationExcludeOutboundIPRangesPolicyKey),
		annotationString(annotations, sidecarAnnotationIncludeInboundPortsPolicyKey),
		annotationString(annotations, sidecarAnnotationExcludeInboundPortsPolicyKey))

	return required
}

func formatDuration(in *duration.Duration) string {
	dur, err := ptypes.Duration(in)
	if err != nil {
		return "1s"
	}
	return dur.String()
}

func isset(m map[string]string, key string) bool {
	_, ok := m[key]
	return ok
}

func annotationString(annotations map[string]string, key string) string {
	if val, ok := annotations[key]; ok {
		return val
	}
	return "(unset)"
}

func validateAnnotation(annotations map[string]string, key string, validateFunc func(string) error) error {
	if val, ok := annotations[key]; ok {
		if err := validateFunc(val); err != nil {
			return fmt.Errorf("injection failed. Invalid value for annotation %s: %s. Error: %v", key, val, err)
		}
	}
	return nil
}

func validateAnnotations(metadata *metav1.ObjectMeta) error {
	// Validate injection annotations, if present.
	annotations := metadata.GetAnnotations()
	if err := validateAnnotation(annotations, sidecarAnnotationInterceptionModeKey, ValidateInterceptionMode); err != nil {
		return err
	}
	if err := validateAnnotation(annotations, sidecarAnnotationIncludeOutboundIPRangesPolicyKey, ValidateIncludeIPRanges); err != nil {
		return err
	}
	if err := validateAnnotation(annotations, sidecarAnnotationExcludeOutboundIPRangesPolicyKey, ValidateExcludeIPRanges); err != nil {
		return err
	}
	if err := validateAnnotation(annotations, sidecarAnnotationIncludeInboundPortsPolicyKey, ValidateIncludeInboundPorts); err != nil {
		return err
	}
	return validateAnnotation(annotations, sidecarAnnotationExcludeInboundPortsPolicyKey, ValidateExcludeInboundPorts)
}

func injectionData(sidecarTemplate, version string, spec *corev1.PodSpec, metadata *metav1.ObjectMeta, proxyConfig *meshconfig.ProxyConfig, meshConfig *meshconfig.MeshConfig) (*SidecarInjectionSpec, string, error) { // nolint: lll
	if err := validateAnnotations(metadata); err != nil {
		return nil, "", err
	}

	data := SidecarTemplateData{
		ObjectMeta:  metadata,
		Spec:        spec,
		ProxyConfig: proxyConfig,
		MeshConfig:  meshConfig,
	}

	funcMap := template.FuncMap{
		"formatDuration": formatDuration,
		"isset":          isset,
	}

	var tmpl bytes.Buffer
	temp := template.New("inject").Delims(sidecarTemplateDelimBegin, sidecarTemplateDelimEnd)
	t := template.Must(temp.Funcs(funcMap).Parse(sidecarTemplate))
	if err := t.Execute(&tmpl, &data); err != nil {
		return nil, "", err
	}

	var sic SidecarInjectionSpec
	if err := yaml.Unmarshal(tmpl.Bytes(), &sic); err != nil {
		log.Warnf("Failed to unmarshall template %v %s", err, string(tmpl.Bytes()))
		return nil, "", err
	}

	status := &SidecarInjectionStatus{Version: version}
	for _, c := range sic.InitContainers {
		status.InitContainers = append(status.InitContainers, c.Name)
	}
	for _, c := range sic.Containers {
		status.Containers = append(status.Containers, c.Name)
	}
	for _, c := range sic.Volumes {
		status.Volumes = append(status.Volumes, c.Name)
	}
	for _, c := range sic.ImagePullSecrets {
		status.ImagePullSecrets = append(status.ImagePullSecrets, c.Name)
	}
	statusAnnotationValue, err := json.Marshal(status)
	if err != nil {
		return nil, "", fmt.Errorf("error encoded injection status: %v", err)
	}
	return &sic, string(statusAnnotationValue), nil
}

// IntoResourceFile injects the istio proxy into the specified
// kubernetes YAML file.
func IntoResourceFile(sidecarTemplate string, meshconfig *meshconfig.MeshConfig, in io.Reader, out io.Writer) error {
	reader := yamlDecoder.NewYAMLReader(bufio.NewReaderSize(in, 4096))
	for {
		raw, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		obj, err := fromRawToObject(raw)
		if err != nil && !runtime.IsNotRegisteredError(err) {
			return err
		}

		var updated []byte
		if err == nil {
			outObject, err := IntoObject(sidecarTemplate, meshconfig, obj) // nolint: vetshadow
			if err != nil {
				return err
			}
			if updated, err = yaml.Marshal(outObject); err != nil {
				return err
			}
		} else {
			updated = raw // unchanged
		}

		if _, err = out.Write(updated); err != nil {
			return err
		}
		if _, err = fmt.Fprint(out, "---\n"); err != nil {
			return err
		}
	}
	return nil
}

func fromRawToObject(raw []byte) (runtime.Object, error) {
	var typeMeta metav1.TypeMeta
	if err := yaml.Unmarshal(raw, &typeMeta); err != nil {
		return nil, err
	}

	gvk := schema.FromAPIVersionAndKind(typeMeta.APIVersion, typeMeta.Kind)
	obj, err := injectScheme.New(gvk)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(raw, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func IntoObject(sidecarTemplate string, meshconfig *meshconfig.MeshConfig, in runtime.Object) (interface{}, error) {
	out := in.DeepCopyObject()

	var metadata *metav1.ObjectMeta
	var podSpec *corev1.PodSpec

	// Handle Lists
	if list, ok := out.(*corev1.List); ok {
		result := list

		for i, item := range list.Items {
			obj, err := fromRawToObject(item.Raw)
			if runtime.IsNotRegisteredError(err) {
				continue
			}
			if err != nil {
				return nil, err
			}

			r, err := IntoObject(sidecarTemplate, meshconfig, obj) // nolint: vetshadow
			if err != nil {
				return nil, err
			}

			re := runtime.RawExtension{}
			re.Object = r.(runtime.Object)
			result.Items[i] = re
		}
		return result, nil
	}

	// CronJobs have JobTemplates in them, instead of Templates, so we
	// special case them.
	if job, ok := out.(*v2alpha1.CronJob); ok {
		metadata = &job.Spec.JobTemplate.ObjectMeta
		podSpec = &job.Spec.JobTemplate.Spec.Template.Spec
	} else if pod, ok := out.(*corev1.Pod); ok {
		metadata = &pod.ObjectMeta
		podSpec = &pod.Spec
	} else {
		// `in` is a pointer to an Object. Dereference it.
		outValue := reflect.ValueOf(out).Elem()

		templateValue := outValue.FieldByName("Spec").FieldByName("Template")
		// `Template` is defined as a pointer in some older API
		// definitions, e.g. ReplicationController
		if templateValue.Kind() == reflect.Ptr {
			templateValue = templateValue.Elem()
		}
		metadata = templateValue.FieldByName("ObjectMeta").Addr().Interface().(*metav1.ObjectMeta)
		podSpec = templateValue.FieldByName("Spec").Addr().Interface().(*corev1.PodSpec)
	}

	// Skip injection when host networking is enabled. The problem is
	// that the iptable changes are assumed to be within the pod when,
	// in fact, they are changing the routing at the host level. This
	// often results in routing failures within a node which can
	// affect the network provider within the cluster causing
	// additional pod failures.
	if podSpec.HostNetwork {
		fmt.Fprintf(os.Stderr, "Skipping injection because %q has host networking enabled\n", metadata.Name)
		return out, nil
	}

	spec, status, err := injectionData(
		sidecarTemplate,
		sidecarTemplateVersionHash(sidecarTemplate),
		podSpec,
		metadata,
		meshconfig.DefaultConfig,
		meshconfig)
	if err != nil {
		return nil, err
	}

	podSpec.InitContainers = append(podSpec.InitContainers, spec.InitContainers...)
	podSpec.Containers = append(podSpec.Containers, spec.Containers...)
	podSpec.Volumes = append(podSpec.Volumes, spec.Volumes...)

	if metadata.Annotations == nil {
		metadata.Annotations = make(map[string]string)
	}
	metadata.Annotations[sidecarAnnotationStatusKey] = status

	return out, nil
}

// GenerateTemplateFromParams generates a sidecar template from the legacy injection parameters
func GenerateTemplateFromParams(params *Params) (string, error) {
	// Validate the parameters before we go any farther.
	if err := params.Validate(); err != nil {
		return "", err
	}
	var tmp bytes.Buffer
	err := template.Must(template.New("inject").Parse(parameterizedTemplate)).Execute(&tmp, params)
	return tmp.String(), err
}

// SidecarInjectionStatus contains basic information about the
// injected sidecar. This includes the names of added containers and
// volumes.
type SidecarInjectionStatus struct {
	Version          string   `json:"version"`
	InitContainers   []string `json:"initContainers"`
	Containers       []string `json:"containers"`
	Volumes          []string `json:"volumes"`
	ImagePullSecrets []string `json:"imagePullSecrets"`
}

// helper function to generate a template version identifier from a
// hash of the un-executed template contents.
func sidecarTemplateVersionHash(in string) string {
	hash := sha256.Sum256([]byte(in))
	return hex.EncodeToString(hash[:])
}
