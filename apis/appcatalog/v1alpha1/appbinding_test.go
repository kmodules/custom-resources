/*
Copyright The Kmodules Authors.

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

package v1alpha1_test

import (
	"testing"

	"kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
)

func newAppBinding(t string) *v1alpha1.AppBinding {
	return &v1alpha1.AppBinding{
		Spec: v1alpha1.AppBindingSpec{
			Type: v1alpha1.AppType(t),
		},
	}
}

func TestAppGroupResource(t *testing.T) {
	cases := []struct {
		x        *v1alpha1.AppBinding
		group    string
		resource string
	}{
		{newAppBinding("kubedb.com/postgres"), "kubedb.com", "postgres"},
		{newAppBinding("postgres"), "", "postgres"},
		{newAppBinding("kubedb.com"), "", "kubedb.com"},
		{newAppBinding("kubedb.com/xyz/postgres"), "kubedb.com/xyz", "postgres"},
		{newAppBinding("xyz.kubedb.com/postgres"), "xyz.kubedb.com", "postgres"},
		{newAppBinding("kubedb.com/"), "kubedb.com", ""},
		{newAppBinding("/postgres"), "", "postgres"},
		{newAppBinding(""), "", ""},
	}

	for _, c := range cases {
		t.Run(string(c.x.Spec.Type), func(t *testing.T) {
			g, r := c.x.AppGroupResource()
			if g != c.group || r != c.resource {
				t.Errorf("Failed to parse AppGroupResource for '%v': expected (%v,%v), got (%v,%v)", c.x.Spec.Type, c.group, c.resource, g, r)
			}
		})
	}
}
