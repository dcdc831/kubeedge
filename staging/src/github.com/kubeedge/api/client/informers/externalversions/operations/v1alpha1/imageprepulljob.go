/*
Copyright The KubeEdge Authors.

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
	time "time"

	operationsv1alpha1 "github.com/kubeedge/api/apis/operations/v1alpha1"
	versioned "github.com/kubeedge/api/client/clientset/versioned"
	internalinterfaces "github.com/kubeedge/api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kubeedge/api/client/listers/operations/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ImagePrePullJobInformer provides access to a shared informer and lister for
// ImagePrePullJobs.
type ImagePrePullJobInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ImagePrePullJobLister
}

type imagePrePullJobInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewImagePrePullJobInformer constructs a new informer for ImagePrePullJob type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewImagePrePullJobInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredImagePrePullJobInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredImagePrePullJobInformer constructs a new informer for ImagePrePullJob type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredImagePrePullJobInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperationsV1alpha1().ImagePrePullJobs().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperationsV1alpha1().ImagePrePullJobs().Watch(context.TODO(), options)
			},
		},
		&operationsv1alpha1.ImagePrePullJob{},
		resyncPeriod,
		indexers,
	)
}

func (f *imagePrePullJobInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredImagePrePullJobInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *imagePrePullJobInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&operationsv1alpha1.ImagePrePullJob{}, f.defaultInformer)
}

func (f *imagePrePullJobInformer) Lister() v1alpha1.ImagePrePullJobLister {
	return v1alpha1.NewImagePrePullJobLister(f.Informer().GetIndexer())
}
