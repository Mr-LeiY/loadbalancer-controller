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

// TagsGetter has a method to return a TagInterface.
// A group's client should implement this interface.
type TagsGetter interface {
	Tags() TagInterface
}

// TagInterface has methods to work with Tag resources.
type TagInterface interface {
	Create(*v1beta1.Tag) (*v1beta1.Tag, error)
	Update(*v1beta1.Tag) (*v1beta1.Tag, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.Tag, error)
	List(opts v1.ListOptions) (*v1beta1.TagList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Tag, err error)
	TagExpansion
}

// tags implements TagInterface
type tags struct {
	client rest.Interface
}

// newTags returns a Tags
func newTags(c *ResourceV1beta1Client) *tags {
	return &tags{
		client: c.RESTClient(),
	}
}

// Create takes the representation of a tag and creates it.  Returns the server's representation of the tag, and an error, if there is any.
func (c *tags) Create(tag *v1beta1.Tag) (result *v1beta1.Tag, err error) {
	result = &v1beta1.Tag{}
	err = c.client.Post().
		Resource("tags").
		Body(tag).
		Do().
		Into(result)
	return
}

// Update takes the representation of a tag and updates it. Returns the server's representation of the tag, and an error, if there is any.
func (c *tags) Update(tag *v1beta1.Tag) (result *v1beta1.Tag, err error) {
	result = &v1beta1.Tag{}
	err = c.client.Put().
		Resource("tags").
		Name(tag.Name).
		Body(tag).
		Do().
		Into(result)
	return
}

// Delete takes name of the tag and deletes it. Returns an error if one occurs.
func (c *tags) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("tags").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *tags) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("tags").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the tag, and returns the corresponding tag object, and an error if there is any.
func (c *tags) Get(name string, options v1.GetOptions) (result *v1beta1.Tag, err error) {
	result = &v1beta1.Tag{}
	err = c.client.Get().
		Resource("tags").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Tags that match those selectors.
func (c *tags) List(opts v1.ListOptions) (result *v1beta1.TagList, err error) {
	result = &v1beta1.TagList{}
	err = c.client.Get().
		Resource("tags").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tags.
func (c *tags) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("tags").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Patch applies the patch and returns the patched tag.
func (c *tags) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Tag, err error) {
	result = &v1beta1.Tag{}
	err = c.client.Patch(pt).
		Resource("tags").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
