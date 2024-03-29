// Copyright 2020 Red Hat, Inc. and/or its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package framework

import (
	"fmt"

	"github.com/kiegroup/kogito-operator/apis/app/v1beta1"
	v1 "github.com/kiegroup/kogito-operator/apis/rhpam/v1"

	api "github.com/kiegroup/kogito-operator/apis"

	"github.com/kiegroup/kogito-operator/core/client/kubernetes"
	"github.com/kiegroup/kogito-operator/core/framework"
	"github.com/kiegroup/kogito-operator/test/pkg/config"
	"github.com/kiegroup/kogito-operator/test/pkg/framework/mappers"
	bddtypes "github.com/kiegroup/kogito-operator/test/pkg/types"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// DeployKogitoBuild deploy a KogitoBuild
func DeployKogitoBuild(namespace string, installerType InstallerType, buildHolder *bddtypes.KogitoBuildHolder) error {
	GetLogger(namespace).Info(fmt.Sprintf("%s deploy %s example %s with name %s and native %v", installerType, buildHolder.KogitoBuild.GetSpec().GetRuntime(), buildHolder.KogitoBuild.GetSpec().GetGitSource().GetContextDir(), buildHolder.KogitoBuild.GetName(), buildHolder.KogitoBuild.GetSpec().IsNative()))

	switch installerType {
	case CLIInstallerType:
		return cliDeployKogitoBuild(buildHolder)
	case CRInstallerType:
		return crDeployKogitoBuild(buildHolder)
	default:
		panic(fmt.Errorf("Unknown installer type %s", installerType))
	}
}

func crDeployKogitoBuild(buildHolder *bddtypes.KogitoBuildHolder) error {
	if err := kubernetes.ResourceC(kubeClient).CreateIfNotExists(buildHolder.KogitoBuild); err != nil {
		return fmt.Errorf("Error creating example build %s: %v", buildHolder.KogitoBuild.GetName(), err)
	}
	if err := kubernetes.ResourceC(kubeClient).CreateIfNotExists(buildHolder.KogitoService); err != nil {
		return fmt.Errorf("Error creating example service %s: %v", buildHolder.KogitoService.GetName(), err)
	}
	return nil
}

func cliDeployKogitoBuild(buildHolder *bddtypes.KogitoBuildHolder) error {
	cmd := []string{"deploy", buildHolder.KogitoBuild.GetName()}

	// If GIT URI is defined then it needs to be appended as second parameter
	if gitURI := buildHolder.KogitoBuild.GetSpec().GetGitSource().GetURI(); len(gitURI) > 0 {
		cmd = append(cmd, gitURI)
	} else if len(buildHolder.BuiltBinaryFolder) > 0 {
		cmd = append(cmd, buildHolder.BuiltBinaryFolder)
	}

	cmd = append(cmd, mappers.GetBuildCLIFlags(buildHolder.KogitoBuild)...)
	cmd = append(cmd, mappers.GetServiceCLIFlags(buildHolder.KogitoServiceHolder)...)

	_, err := ExecuteCliCommandInNamespace(buildHolder.KogitoBuild.GetNamespace(), cmd...)
	return err
}

// GetKogitoBuildStub Get basic KogitoBuild stub with all needed fields initialized
func GetKogitoBuildStub(namespace, runtimeType, name string) api.KogitoBuildInterface {
	var kogitoBuild api.KogitoBuildInterface = &v1beta1.KogitoBuild{ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: namespace}}
	if config.UseProductOperator() {
		kogitoBuild = &v1.KogitoBuild{ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: namespace}}
	}
	kogitoBuild.GetSpec().SetRuntime(api.RuntimeType(runtimeType))
	kogitoBuild.GetSpec().SetMavenMirrorURL(config.GetMavenMirrorURL())

	if len(config.GetCustomMavenRepoURL()) > 0 {
		kogitoBuild.GetSpec().SetEnv(framework.EnvOverride(kogitoBuild.GetSpec().GetEnv(), corev1.EnvVar{Name: "MAVEN_REPO_URL", Value: config.GetCustomMavenRepoURL()}))
	}

	if config.IsMavenIgnoreSelfSignedCertificate() {
		kogitoBuild.GetSpec().SetEnv(framework.EnvOverride(kogitoBuild.GetSpec().GetEnv(), corev1.EnvVar{Name: "MAVEN_IGNORE_SELF_SIGNED_CERTIFICATE", Value: "true"}))
	}

	return kogitoBuild
}

// GetKogitoBuild returns the KogitoBuild type
func GetKogitoBuild(namespace, buildName string) (api.KogitoBuildInterface, error) {
	var build api.KogitoBuildInterface = &v1beta1.KogitoBuild{}
	if config.UseProductOperator() {
		build = &v1.KogitoBuild{}
	}
	if exists, err := kubernetes.ResourceC(kubeClient).FetchWithKey(types.NamespacedName{Name: buildName, Namespace: namespace}, build); err != nil && !errors.IsNotFound(err) {
		return nil, fmt.Errorf("Error while trying to look for KogitoBuild %s: %v ", buildName, err)
	} else if errors.IsNotFound(err) || !exists {
		return nil, nil
	}
	return build, nil
}

// SetupKogitoBuildImageStreams sets the correct images for the KogitoBuild
func SetupKogitoBuildImageStreams(kogitoBuild api.KogitoBuildInterface) {
	kogitoBuild.GetSpec().SetBuildImage(GetKogitoBuildS2IImage())
	kogitoBuild.GetSpec().SetRuntimeImage(GetKogitoBuildRuntimeImage(kogitoBuild.GetSpec().IsNative()))
}
