/*
Copyright AppsCode Inc. and Contributors

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

	v1alpha1 "kmodules.xyz/custom-resources/apis/auditor/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSiteInfos implements SiteInfoInterface
type FakeSiteInfos struct {
	Fake *FakeAuditorV1alpha1
}

var siteinfosResource = schema.GroupVersionResource{Group: "auditor.appscode.com", Version: "v1alpha1", Resource: "siteinfos"}

var siteinfosKind = schema.GroupVersionKind{Group: "auditor.appscode.com", Version: "v1alpha1", Kind: "SiteInfo"}

// Get takes name of the siteInfo, and returns the corresponding siteInfo object, and an error if there is any.
func (c *FakeSiteInfos) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SiteInfo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(siteinfosResource, name), &v1alpha1.SiteInfo{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SiteInfo), err
}

// List takes label and field selectors, and returns the list of SiteInfos that match those selectors.
func (c *FakeSiteInfos) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SiteInfoList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(siteinfosResource, siteinfosKind, opts), &v1alpha1.SiteInfoList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SiteInfoList{ListMeta: obj.(*v1alpha1.SiteInfoList).ListMeta}
	for _, item := range obj.(*v1alpha1.SiteInfoList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested siteInfos.
func (c *FakeSiteInfos) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(siteinfosResource, opts))
}

// Create takes the representation of a siteInfo and creates it.  Returns the server's representation of the siteInfo, and an error, if there is any.
func (c *FakeSiteInfos) Create(ctx context.Context, siteInfo *v1alpha1.SiteInfo, opts v1.CreateOptions) (result *v1alpha1.SiteInfo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(siteinfosResource, siteInfo), &v1alpha1.SiteInfo{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SiteInfo), err
}

// Update takes the representation of a siteInfo and updates it. Returns the server's representation of the siteInfo, and an error, if there is any.
func (c *FakeSiteInfos) Update(ctx context.Context, siteInfo *v1alpha1.SiteInfo, opts v1.UpdateOptions) (result *v1alpha1.SiteInfo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(siteinfosResource, siteInfo), &v1alpha1.SiteInfo{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SiteInfo), err
}

// Delete takes name of the siteInfo and deletes it. Returns an error if one occurs.
func (c *FakeSiteInfos) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(siteinfosResource, name), &v1alpha1.SiteInfo{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSiteInfos) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(siteinfosResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SiteInfoList{})
	return err
}

// Patch applies the patch and returns the patched siteInfo.
func (c *FakeSiteInfos) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SiteInfo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(siteinfosResource, name, pt, data, subresources...), &v1alpha1.SiteInfo{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SiteInfo), err
}
