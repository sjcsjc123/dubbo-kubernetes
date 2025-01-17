/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	dubboapacheorgv1alpha1 "github.com/apache/dubbo-kubernetes/pkg/core/gen/apis/dubbo.apache.org/v1alpha1"
	versioned "github.com/apache/dubbo-kubernetes/pkg/core/gen/generated/clientset/versioned"
	internalinterfaces "github.com/apache/dubbo-kubernetes/pkg/core/gen/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/apache/dubbo-kubernetes/pkg/core/gen/generated/listers/dubbo.apache.org/v1alpha1"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TagRouteInformer provides access to a shared informer and lister for
// TagRoutes.
type TagRouteInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.TagRouteLister
}

type tagRouteInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTagRouteInformer constructs a new informer for TagRoute type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTagRouteInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTagRouteInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTagRouteInformer constructs a new informer for TagRoute type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTagRouteInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DubboV1alpha1().TagRoutes(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DubboV1alpha1().TagRoutes(namespace).Watch(context.TODO(), options)
			},
		},
		&dubboapacheorgv1alpha1.TagRoute{},
		resyncPeriod,
		indexers,
	)
}

func (f *tagRouteInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTagRouteInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *tagRouteInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&dubboapacheorgv1alpha1.TagRoute{}, f.defaultInformer)
}

func (f *tagRouteInformer) Lister() v1alpha1.TagRouteLister {
	return v1alpha1.NewTagRouteLister(f.Informer().GetIndexer())
}
