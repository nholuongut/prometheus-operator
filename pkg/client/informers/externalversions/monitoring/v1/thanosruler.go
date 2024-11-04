// Copyright The Nho Luong DevOps
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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	monitoringv1 "github.com/nholuongut/prometheus-operator/pkg/apis/monitoring/v1"
	internalinterfaces "github.com/nholuongut/prometheus-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/nholuongut/prometheus-operator/pkg/client/listers/monitoring/v1"
	versioned "github.com/nholuongut/prometheus-operator/pkg/client/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ThanosRulerInformer provides access to a shared informer and lister for
// ThanosRulers.
type ThanosRulerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ThanosRulerLister
}

type thanosRulerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewThanosRulerInformer constructs a new informer for ThanosRuler type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewThanosRulerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredThanosRulerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredThanosRulerInformer constructs a new informer for ThanosRuler type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredThanosRulerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MonitoringV1().ThanosRulers(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MonitoringV1().ThanosRulers(namespace).Watch(context.TODO(), options)
			},
		},
		&monitoringv1.ThanosRuler{},
		resyncPeriod,
		indexers,
	)
}

func (f *thanosRulerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredThanosRulerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *thanosRulerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&monitoringv1.ThanosRuler{}, f.defaultInformer)
}

func (f *thanosRulerInformer) Lister() v1.ThanosRulerLister {
	return v1.NewThanosRulerLister(f.Informer().GetIndexer())
}
