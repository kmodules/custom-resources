package servicecatalog

import (
	"unsafe"

	svcat "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"k8s.io/apimachinery/pkg/conversion"
	appcat "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
)

func Convert_AppCatalog_v1alpha1_SecretTransform_To_ServiceCatalog_v1beta1_SecretTransform(in *appcat.SecretTransform, out *svcat.SecretTransform, s conversion.Scope) error {
	out.RenameKey = (*svcat.RenameKeyTransform)(unsafe.Pointer(in.RenameKey))
	out.AddKey = (*svcat.AddKeyTransform)(unsafe.Pointer(in.AddKey))
	out.AddKeysFrom = (*svcat.AddKeysFromTransform)(unsafe.Pointer(in.AddKeysFrom))
	out.RemoveKey = (*svcat.RemoveKeyTransform)(unsafe.Pointer(in.RemoveKey))
	return nil
}
