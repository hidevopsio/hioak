// Copyright 2018 John Deng (hi.devops.io@gmail.com).
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

package openshift

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/openshift/api/build/v1"
	buildv1 "github.com/openshift/client-go/build/clientset/versioned/typed/build/v1"
	"hidevops.io/hiboot/pkg/log"
	"hidevops.io/hiboot/pkg/system"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

type Scm struct {
	Url    string
	Ref    string
	Secret string
}

type BuildConfig struct {
	clientSet buildv1.BuildV1Interface
}

func newBuildConfig(clientSet buildv1.BuildV1Interface) *BuildConfig {
	return &BuildConfig{
		clientSet: clientSet,
	}
}

// @Title Create
// @Description Create new BuildConfig
// @Param
// @Return *v1.BuildConfig, error
func (b *BuildConfig) Create(name, namespace, url, ref, version, secret string, from *corev1.ObjectReference) (*v1.BuildConfig, error) {
	log.Debug("BuildConfig.Create()")
	// buildConfig
	buildConfig := &v1.BuildConfig{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Labels: map[string]string{
				"app": name,
			},
		},
		Spec: v1.BuildConfigSpec{
			// The runPolicy field controls whether builds created from this build configuration can be run simultaneously.
			// The default value is Serial, which means new builds will run sequentially, not simultaneously.
			RunPolicy: v1.BuildRunPolicy("Serial"),
			CommonSpec: v1.CommonSpec{

				Source: v1.BuildSource{
					Type: v1.BuildSourceType(v1.BuildSourceGit),
					Git: &v1.GitBuildSource{
						URI: url,
						Ref: ref,
					},
					SourceSecret: &corev1.LocalObjectReference{
						Name: secret,
					},
				},
				Strategy: v1.BuildStrategy{
					Type: v1.BuildStrategyType(v1.SourceBuildStrategyType),
					SourceStrategy: &v1.SourceBuildStrategy{
						From: *from,
					},
				},
				Output: v1.BuildOutput{
					To: &corev1.ObjectReference{
						Kind: "ImageStreamTag",
						Name: name + ":" + version,
					},
				},
			},
		},
	}

	bc, err := b.clientSet.BuildConfigs(namespace).Get(name, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		bc, err = b.clientSet.BuildConfigs(namespace).Create(buildConfig)
		if nil == err {
			log.Infof("Created BuildConfig %v", bc.Name)
		}
	} else {
		buildConfig.ResourceVersion = bc.ResourceVersion
		bc, err = b.clientSet.BuildConfigs(namespace).Update(buildConfig)
		if nil == err {
			log.Infof("Updated BuildConfig %v", bc.Name)
		}
	}

	return bc, err
}

// @Title Get
// @Description Get BuildConfig
// @Param
// @Return *v1.BuildConfig, error
func (b *BuildConfig) Get(name, namespace string) (*v1.BuildConfig, error) {
	log.Debug("BuildConfig.Get()")
	return b.clientSet.BuildConfigs(namespace).Get(name, metav1.GetOptions{})
}

// @Title Delete
// @Description Delete BuildConfig
// @Param
// @Return error
func (b *BuildConfig) Delete(name, namespace string) error {
	log.Debug("BuildConfig.Delet()")
	return b.clientSet.BuildConfigs(namespace).Delete(name, &metav1.DeleteOptions{})
}

// @Title Build
// @Description Start build according to previous build config settings, it will produce new image build
// @Param repo string, buildCmd string
// @Return *v1.Build, error
func (b *BuildConfig) Build(name, namespace, version string, env []system.Env, from *corev1.ObjectReference) (*v1.Build, error) {
	log.Debug("BuildConfig.Build()")

	e := make([]corev1.EnvVar, 0)
	copier.Copy(&e, env)

	incremental := false
	buildTriggerCauseManualMsg := "Manually triggered"
	buildRequest := v1.BuildRequest{
		TypeMeta: metav1.TypeMeta{
			Kind:       "BuildRequest",
			APIVersion: "build.openshift.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Labels: map[string]string{
				"app":     name,
				"version": version,
			},
		},
		TriggeredBy: append([]v1.BuildTriggerCause{},
			v1.BuildTriggerCause{
				Message: buildTriggerCauseManualMsg,
			},
		),
		SourceStrategyOptions: &v1.SourceStrategyOptions{
			Incremental: &incremental,
		},
		Env:  e,
		From: from,
	}

	build, err := b.clientSet.BuildConfigs(namespace).Instantiate(name, &buildRequest)
	if nil != err {
		log.Error("b.BuildConfigs.Instantiate err", err)
		return nil, err
	}
	return build, err
}

func (b *BuildConfig) Watch(name, namespace string, build *v1.Build, completedHandler func() error) error {
	w, err := b.clientSet.Builds(namespace).Watch(metav1.ListOptions{
		LabelSelector: "app=" + name,
		Watch:         true,
	})

	if nil != err {
		log.Error("BuildConfig.Watch err ", err)
		return err
	}

	for {
		select {
		case event, ok := <-w.ResultChan():
			if !ok {
				log.Info("resultChan: ", ok)
				return fmt.Errorf("resultChan: %v", ok)
			}
			switch event.Type {
			case watch.Added:
				bld := event.Object.(*v1.Build)
				log.Info("Added new build ", bld.Name)
			case watch.Modified:

				bld := event.Object.(*v1.Build)
				if bld.Name == build.Name {
					//log.Info("Modified: ", event.Object)
					log.Debugf("bld.Status.Phase: %v", bld.Status.Phase)
					switch bld.Status.Phase {
					case v1.BuildPhaseComplete:
						log.Info("bld.Status.Phase", bld.Status.Phase)
						var err error
						if nil != completedHandler {
							err = completedHandler()
						}
						w.Stop()
						log.Error("bld.Status.Phase completedHandler", err)
						return err
					case v1.BuildPhaseError, v1.BuildPhaseCancelled, v1.BuildPhaseFailed:
						w.Stop()
						log.Error("bld.Status.Phase BuildPhaseError", fmt.Errorf(bld.Status.Message))
						return fmt.Errorf(bld.Status.Message)

					}
				}

			case watch.Deleted:
				log.Info("Deleted: ", event.Object)
			default:
				log.Error("Failed")
			}
		}
	}
	log.Infof("build.watch :%v", err)
	return err
}

// @Title GetBuild
// @Description Get current build
// @Param
// @Return *v1.Build, error
func (b *BuildConfig) GetBuild(name, namespace string) (*v1.Build, error) {
	log.Debug("BuildConfig.GetBuild()")
	return b.clientSet.Builds(namespace).Get(name, metav1.GetOptions{})
}

// @Title GetBuildStatus
// @Description Get current build status
// @Param
// @Return v1.BuildPhase, error
func (b *BuildConfig) GetBuildStatus(name, namespace string) (v1.BuildPhase, error) {
	log.Debug("BuildConfig.GetBuildStatus()")
	build, err := b.GetBuild(name, namespace)
	return build.Status.Phase, err
}
