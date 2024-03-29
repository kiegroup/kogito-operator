// Copyright 2021 Red Hat, Inc. and/or its affiliates
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

package test

import (
	"github.com/kiegroup/kogito-operator/apis/app/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateFakeKogitoRuntime ...
func CreateFakeKogitoRuntime(namespace string) *v1beta1.KogitoRuntime {
	replicas := int32(1)
	return &v1beta1.KogitoRuntime{
		ObjectMeta: v1.ObjectMeta{
			Name:      "test-kogito-runtime",
			Namespace: namespace,
		},
		Spec: v1beta1.KogitoRuntimeSpec{
			KogitoServiceSpec: v1beta1.KogitoServiceSpec{Replicas: &replicas},
		},
	}
}
