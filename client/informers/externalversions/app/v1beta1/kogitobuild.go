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
// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	time "time"

	appv1beta1 "github.com/kiegroup/kogito-operator/apis/app/v1beta1"
	versioned "github.com/kiegroup/kogito-operator/client/clientset/versioned"
	internalinterfaces "github.com/kiegroup/kogito-operator/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/kiegroup/kogito-operator/client/listers/app/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// KogitoBuildInformer provides access to a shared informer and lister for
// KogitoBuilds.
type KogitoBuildInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.KogitoBuildLister
}

type kogitoBuildInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewKogitoBuildInformer constructs a new informer for KogitoBuild type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKogitoBuildInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKogitoBuildInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredKogitoBuildInformer constructs a new informer for KogitoBuild type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKogitoBuildInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppV1beta1().KogitoBuilds(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppV1beta1().KogitoBuilds(namespace).Watch(context.TODO(), options)
			},
		},
		&appv1beta1.KogitoBuild{},
		resyncPeriod,
		indexers,
	)
}

func (f *kogitoBuildInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKogitoBuildInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *kogitoBuildInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&appv1beta1.KogitoBuild{}, f.defaultInformer)
}

func (f *kogitoBuildInformer) Lister() v1beta1.KogitoBuildLister {
	return v1beta1.NewKogitoBuildLister(f.Informer().GetIndexer())
}
