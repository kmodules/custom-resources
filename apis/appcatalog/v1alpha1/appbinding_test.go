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

func TestGroupResource(t *testing.T) {
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
			g, r := c.x.GroupResource()
			if g != c.group || r != c.resource {
				t.Errorf("Failed GroupResource test for '%v': expected (%v,%v), got (%v,%v)", c.x.Spec.Type, c.group, c.resource, g, r)
			}
		})
	}
}
