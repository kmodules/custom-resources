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

package servicecatalog

import (
	"unsafe"

	appcat "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"

	svcat "github.com/kubernetes-sigs/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"k8s.io/apimachinery/pkg/conversion"
)

func Convert_AppCatalog_v1alpha1_SecretTransform_To_ServiceCatalog_v1beta1_SecretTransform(in *appcat.SecretTransform, out *svcat.SecretTransform, s conversion.Scope) error {
	out.RenameKey = (*svcat.RenameKeyTransform)(unsafe.Pointer(in.RenameKey))
	out.AddKey = (*svcat.AddKeyTransform)(unsafe.Pointer(in.AddKey))
	out.AddKeysFrom = (*svcat.AddKeysFromTransform)(unsafe.Pointer(in.AddKeysFrom))
	out.RemoveKey = (*svcat.RemoveKeyTransform)(unsafe.Pointer(in.RemoveKey))
	return nil
}
