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
	"reflect"
	"testing"

	"kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"

	"github.com/appscode/go/types"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	fakek8s "k8s.io/client-go/kubernetes/fake"
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

func TestAppBinding_TransformCredentials(t *testing.T) {
	cases := []struct {
		name                   string
		transforms             []v1alpha1.SecretTransform
		credentials            map[string][]byte
		transformedCredentials map[string][]byte
		otherSecret            *core.Secret
	}{
		{
			name: "RenameKeyTransform",
			transforms: []v1alpha1.SecretTransform{
				{
					RenameKey: &v1alpha1.RenameKeyTransform{
						From: "foo",
						To:   "bar",
					},
				},
			},
			credentials: map[string][]byte{
				"foo": []byte("123"),
			},
			transformedCredentials: map[string][]byte{
				"bar": []byte("123"),
			},
		},
		{
			name: "AddKeyTransform with value",
			transforms: []v1alpha1.SecretTransform{
				{
					AddKey: &v1alpha1.AddKeyTransform{
						Key:   "bar",
						Value: []byte("456"),
					},
				},
			},
			credentials: map[string][]byte{
				"foo": []byte("123"),
			},
			transformedCredentials: map[string][]byte{
				"foo": []byte("123"),
				"bar": []byte("456"),
			},
		},
		{
			name: "AddKeyTransform with stringValue",
			transforms: []v1alpha1.SecretTransform{
				{
					AddKey: &v1alpha1.AddKeyTransform{
						Key:         "bar",
						StringValue: types.StringP("456"),
					},
				},
			},
			credentials: map[string][]byte{
				"foo": []byte("123"),
			},
			transformedCredentials: map[string][]byte{
				"foo": []byte("123"),
				"bar": []byte("456"),
			},
		},
		{
			name: "AddKeyTransform stringValue precedence over value",
			transforms: []v1alpha1.SecretTransform{
				{
					AddKey: &v1alpha1.AddKeyTransform{
						Key:         "bar",
						Value:       []byte("456"),
						StringValue: types.StringP("789"),
					},
				},
			},
			credentials: map[string][]byte{
				"foo": []byte("123"),
			},
			transformedCredentials: map[string][]byte{
				"foo": []byte("123"),
				"bar": []byte("789"),
			},
		},
		{
			name: "MergeSecretTransform",
			transforms: []v1alpha1.SecretTransform{
				{
					AddKeysFrom: &v1alpha1.AddKeysFromTransform{
						SecretRef: &v1alpha1.ObjectReference{
							Namespace: "ns",
							Name:      "other-secret",
						},
					},
				},
			},
			credentials: map[string][]byte{
				"foo": []byte("123"),
			},
			otherSecret: &core.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "ns",
					Name:      "other-secret",
				},
				Data: map[string][]byte{
					"bar": []byte("456"),
				},
			},
			transformedCredentials: map[string][]byte{
				"foo": []byte("123"),
				"bar": []byte("456"),
			},
		},
		{
			name: "RemoveKeyTransform",
			transforms: []v1alpha1.SecretTransform{
				{
					RemoveKey: &v1alpha1.RemoveKeyTransform{
						Key: "bar",
					},
				},
			},
			credentials: map[string][]byte{
				"foo": []byte("123"),
				"bar": []byte("456"),
			},
			transformedCredentials: map[string][]byte{
				"foo": []byte("123"),
			},
		},
	}
	var k8sClient kubernetes.Interface
	appBinding := newAppBinding("kubedb.com/postgres")
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			appBinding.Spec.SecretTransforms = tc.transforms
			if tc.otherSecret != nil {
				k8sClient = fakek8s.NewSimpleClientset(tc.otherSecret)
			}
			err := appBinding.TransformSecret(k8sClient, tc.credentials)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(tc.credentials, tc.transformedCredentials) {
				t.Errorf("%v: unexpected transformed secret data; expected: %v; actual: %v", tc.name, tc.transformedCredentials, tc.credentials)
			}
		})
	}
}
