/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// This file was automatically generated by informer-gen

package v1beta1

import (
	kubernetes "github.com/caicloud/clientset/kubernetes"
	v1beta1 "github.com/caicloud/clientset/listers/resource/v1beta1"
	resource_v1beta1 "github.com/caicloud/clientset/pkg/apis/resource/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	client_go_kubernetes "k8s.io/client-go/kubernetes"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// ClusterInformer provides access to a shared informer and lister for
// Clusters.
type ClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.ClusterLister
}

type clusterInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newClusterInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.ResourceV1beta1().Clusters().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.ResourceV1beta1().Clusters().Watch(options)
			},
		},
		&resource_v1beta1.Cluster{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *clusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&resource_v1beta1.Cluster{}, func(client client_go_kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
		// panic if client is not *kubernetes.Clientset
		return newClusterInformer(client.(kubernetes.Interface), resyncPeriod)
	})
}

func (f *clusterInformer) Lister() v1beta1.ClusterLister {
	return v1beta1.NewClusterLister(f.Informer().GetIndexer())
}
