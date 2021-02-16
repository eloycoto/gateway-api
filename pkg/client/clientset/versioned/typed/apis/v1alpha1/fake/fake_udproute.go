/*
Copyright 2021 The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "sigs.k8s.io/service-apis/apis/v1alpha1"
)

// FakeUDPRoutes implements UDPRouteInterface
type FakeUDPRoutes struct {
	Fake *FakeNetworkingV1alpha1
	ns   string
}

var udproutesResource = schema.GroupVersionResource{Group: "networking.x-k8s.io", Version: "v1alpha1", Resource: "udproutes"}

var udproutesKind = schema.GroupVersionKind{Group: "networking.x-k8s.io", Version: "v1alpha1", Kind: "UDPRoute"}

// Get takes name of the uDPRoute, and returns the corresponding uDPRoute object, and an error if there is any.
func (c *FakeUDPRoutes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.UDPRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(udproutesResource, c.ns, name), &v1alpha1.UDPRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UDPRoute), err
}

// List takes label and field selectors, and returns the list of UDPRoutes that match those selectors.
func (c *FakeUDPRoutes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.UDPRouteList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(udproutesResource, udproutesKind, c.ns, opts), &v1alpha1.UDPRouteList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.UDPRouteList{ListMeta: obj.(*v1alpha1.UDPRouteList).ListMeta}
	for _, item := range obj.(*v1alpha1.UDPRouteList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested uDPRoutes.
func (c *FakeUDPRoutes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(udproutesResource, c.ns, opts))

}

// Create takes the representation of a uDPRoute and creates it.  Returns the server's representation of the uDPRoute, and an error, if there is any.
func (c *FakeUDPRoutes) Create(ctx context.Context, uDPRoute *v1alpha1.UDPRoute, opts v1.CreateOptions) (result *v1alpha1.UDPRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(udproutesResource, c.ns, uDPRoute), &v1alpha1.UDPRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UDPRoute), err
}

// Update takes the representation of a uDPRoute and updates it. Returns the server's representation of the uDPRoute, and an error, if there is any.
func (c *FakeUDPRoutes) Update(ctx context.Context, uDPRoute *v1alpha1.UDPRoute, opts v1.UpdateOptions) (result *v1alpha1.UDPRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(udproutesResource, c.ns, uDPRoute), &v1alpha1.UDPRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UDPRoute), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeUDPRoutes) UpdateStatus(ctx context.Context, uDPRoute *v1alpha1.UDPRoute, opts v1.UpdateOptions) (*v1alpha1.UDPRoute, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(udproutesResource, "status", c.ns, uDPRoute), &v1alpha1.UDPRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UDPRoute), err
}

// Delete takes name of the uDPRoute and deletes it. Returns an error if one occurs.
func (c *FakeUDPRoutes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(udproutesResource, c.ns, name), &v1alpha1.UDPRoute{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeUDPRoutes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(udproutesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.UDPRouteList{})
	return err
}

// Patch applies the patch and returns the patched uDPRoute.
func (c *FakeUDPRoutes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.UDPRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(udproutesResource, c.ns, name, pt, data, subresources...), &v1alpha1.UDPRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UDPRoute), err
}
