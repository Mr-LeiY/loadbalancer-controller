/*
Copyright 2018 caicloud authors. All rights reserved.
*/

package v1beta1

import (
	scheme "github.com/caicloud/clientset/kubernetes/scheme"
	v1beta1 "github.com/caicloud/clientset/pkg/apis/resource/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// StorageTypesGetter has a method to return a StorageTypeInterface.
// A group's client should implement this interface.
type StorageTypesGetter interface {
	StorageTypes() StorageTypeInterface
}

// StorageTypeInterface has methods to work with StorageType resources.
type StorageTypeInterface interface {
	Create(*v1beta1.StorageType) (*v1beta1.StorageType, error)
	Update(*v1beta1.StorageType) (*v1beta1.StorageType, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.StorageType, error)
	List(opts v1.ListOptions) (*v1beta1.StorageTypeList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.StorageType, err error)
	StorageTypeExpansion
}

// storageTypes implements StorageTypeInterface
type storageTypes struct {
	client rest.Interface
}

// newStorageTypes returns a StorageTypes
func newStorageTypes(c *ResourceV1beta1Client) *storageTypes {
	return &storageTypes{
		client: c.RESTClient(),
	}
}

// Create takes the representation of a storageType and creates it.  Returns the server's representation of the storageType, and an error, if there is any.
func (c *storageTypes) Create(storageType *v1beta1.StorageType) (result *v1beta1.StorageType, err error) {
	result = &v1beta1.StorageType{}
	err = c.client.Post().
		Resource("storagetypes").
		Body(storageType).
		Do().
		Into(result)
	return
}

// Update takes the representation of a storageType and updates it. Returns the server's representation of the storageType, and an error, if there is any.
func (c *storageTypes) Update(storageType *v1beta1.StorageType) (result *v1beta1.StorageType, err error) {
	result = &v1beta1.StorageType{}
	err = c.client.Put().
		Resource("storagetypes").
		Name(storageType.Name).
		Body(storageType).
		Do().
		Into(result)
	return
}

// Delete takes name of the storageType and deletes it. Returns an error if one occurs.
func (c *storageTypes) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("storagetypes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *storageTypes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("storagetypes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the storageType, and returns the corresponding storageType object, and an error if there is any.
func (c *storageTypes) Get(name string, options v1.GetOptions) (result *v1beta1.StorageType, err error) {
	result = &v1beta1.StorageType{}
	err = c.client.Get().
		Resource("storagetypes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of StorageTypes that match those selectors.
func (c *storageTypes) List(opts v1.ListOptions) (result *v1beta1.StorageTypeList, err error) {
	result = &v1beta1.StorageTypeList{}
	err = c.client.Get().
		Resource("storagetypes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested storageTypes.
func (c *storageTypes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("storagetypes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Patch applies the patch and returns the patched storageType.
func (c *storageTypes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.StorageType, err error) {
	result = &v1beta1.StorageType{}
	err = c.client.Patch(pt).
		Resource("storagetypes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
