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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	operatorsv1 "github.com/operator-framework/operator-marketplace/pkg/apis/operators/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCatalogSourceConfigs implements CatalogSourceConfigInterface
type FakeCatalogSourceConfigs struct {
	Fake *FakeOperatorsV1
	ns   string
}

var catalogsourceconfigsResource = schema.GroupVersionResource{Group: "operators.coreos.com", Version: "v1", Resource: "catalogsourceconfigs"}

var catalogsourceconfigsKind = schema.GroupVersionKind{Group: "operators.coreos.com", Version: "v1", Kind: "CatalogSourceConfig"}

// Get takes name of the catalogSourceConfig, and returns the corresponding catalogSourceConfig object, and an error if there is any.
func (c *FakeCatalogSourceConfigs) Get(name string, options v1.GetOptions) (result *operatorsv1.CatalogSourceConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(catalogsourceconfigsResource, c.ns, name), &operatorsv1.CatalogSourceConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*operatorsv1.CatalogSourceConfig), err
}

// List takes label and field selectors, and returns the list of CatalogSourceConfigs that match those selectors.
func (c *FakeCatalogSourceConfigs) List(opts v1.ListOptions) (result *operatorsv1.CatalogSourceConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(catalogsourceconfigsResource, catalogsourceconfigsKind, c.ns, opts), &operatorsv1.CatalogSourceConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &operatorsv1.CatalogSourceConfigList{ListMeta: obj.(*operatorsv1.CatalogSourceConfigList).ListMeta}
	for _, item := range obj.(*operatorsv1.CatalogSourceConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested catalogSourceConfigs.
func (c *FakeCatalogSourceConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(catalogsourceconfigsResource, c.ns, opts))

}

// Create takes the representation of a catalogSourceConfig and creates it.  Returns the server's representation of the catalogSourceConfig, and an error, if there is any.
func (c *FakeCatalogSourceConfigs) Create(catalogSourceConfig *operatorsv1.CatalogSourceConfig) (result *operatorsv1.CatalogSourceConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(catalogsourceconfigsResource, c.ns, catalogSourceConfig), &operatorsv1.CatalogSourceConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*operatorsv1.CatalogSourceConfig), err
}

// Update takes the representation of a catalogSourceConfig and updates it. Returns the server's representation of the catalogSourceConfig, and an error, if there is any.
func (c *FakeCatalogSourceConfigs) Update(catalogSourceConfig *operatorsv1.CatalogSourceConfig) (result *operatorsv1.CatalogSourceConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(catalogsourceconfigsResource, c.ns, catalogSourceConfig), &operatorsv1.CatalogSourceConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*operatorsv1.CatalogSourceConfig), err
}

// Delete takes name of the catalogSourceConfig and deletes it. Returns an error if one occurs.
func (c *FakeCatalogSourceConfigs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(catalogsourceconfigsResource, c.ns, name), &operatorsv1.CatalogSourceConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCatalogSourceConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(catalogsourceconfigsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &operatorsv1.CatalogSourceConfigList{})
	return err
}

// Patch applies the patch and returns the patched catalogSourceConfig.
func (c *FakeCatalogSourceConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *operatorsv1.CatalogSourceConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(catalogsourceconfigsResource, c.ns, name, pt, data, subresources...), &operatorsv1.CatalogSourceConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*operatorsv1.CatalogSourceConfig), err
}
