/*
Copyright 2020 caicloud authors. All rights reserved.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	customclient "github.com/caicloud/clientset/customclient"
	internalinterfaces "github.com/caicloud/clientset/custominformers/internalinterfaces"
	v1alpha1 "github.com/caicloud/clientset/listers/config/v1alpha1"
	configv1alpha1 "github.com/caicloud/clientset/pkg/apis/config/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ConfigClaimInformer provides access to a shared informer and lister for
// ConfigClaims.
type ConfigClaimInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ConfigClaimLister
}

type configClaimInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewConfigClaimInformer constructs a new informer for ConfigClaim type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewConfigClaimInformer(client customclient.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredConfigClaimInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredConfigClaimInformer constructs a new informer for ConfigClaim type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredConfigClaimInformer(client customclient.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConfigV1alpha1().ConfigClaims(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConfigV1alpha1().ConfigClaims(namespace).Watch(options)
			},
		},
		&configv1alpha1.ConfigClaim{},
		resyncPeriod,
		indexers,
	)
}

func (f *configClaimInformer) defaultInformer(client customclient.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredConfigClaimInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *configClaimInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&configv1alpha1.ConfigClaim{}, f.defaultInformer)
}

func (f *configClaimInformer) Lister() v1alpha1.ConfigClaimLister {
	return v1alpha1.NewConfigClaimLister(f.Informer().GetIndexer())
}
