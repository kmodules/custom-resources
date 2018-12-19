package v1alpha1

import (
	"bytes"

	crdutils "github.com/appscode/kutil/apiextensions/v1beta1"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/jsonpath"
)

func (p AppBinding) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	return crdutils.NewCustomResourceDefinition(crdutils.Config{
		Group:         SchemeGroupVersion.Group,
		Plural:        ResourceApps,
		Singular:      ResourceApp,
		Kind:          ResourceKindApp,
		Categories:    []string{"catalog", "appscode", "all"},
		ResourceScope: string(apiextensions.NamespaceScoped),
		Versions: []apiextensions.CustomResourceDefinitionVersion{
			{
				Name:    SchemeGroupVersion.Version,
				Served:  true,
				Storage: true,
			},
		},
		Labels: crdutils.Labels{
			LabelsMap: map[string]string{"app": "catalog"},
		},
		SpecDefinitionName:      "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1.AppBinding",
		EnableValidation:        true,
		GetOpenAPIDefinitions:   GetOpenAPIDefinitions,
		EnableStatusSubresource: false,
		AdditionalPrinterColumns: []apiextensions.CustomResourceColumnDefinition{
			{
				Name:     "Age",
				Type:     "date",
				JSONPath: ".metadata.creationTimestamp",
			},
		},
	})
}

// ref: https://github.com/kubernetes-incubator/service-catalog/blob/37b874716ad709a175e426f5f5638322a600849f/pkg/controller/controller_binding.go#L588
func TransformCredentials(kc kubernetes.Interface, transforms []SecretTransform, credentials map[string]interface{}) error {
	for _, t := range transforms {
		switch {
		case t.AddKey != nil:
			var value interface{}
			if t.AddKey.JSONPathExpression != nil {
				result, err := evaluateJSONPath(*t.AddKey.JSONPathExpression, credentials)
				if err != nil {
					return err
				}
				value = result
			} else if t.AddKey.StringValue != nil {
				value = *t.AddKey.StringValue
			} else {
				value = t.AddKey.Value
			}
			credentials[t.AddKey.Key] = value
		case t.RenameKey != nil:
			value, ok := credentials[t.RenameKey.From]
			if ok {
				credentials[t.RenameKey.To] = value
				delete(credentials, t.RenameKey.From)
			}
		case t.AddKeysFrom != nil:
			secret, err := kc.CoreV1().
				Secrets(t.AddKeysFrom.SecretRef.Namespace).
				Get(t.AddKeysFrom.SecretRef.Name, metav1.GetOptions{})
			if err != nil {
				return err // TODO: if the Secret doesn't exist yet, can we perform the transform when it does?
			}
			for k, v := range secret.Data {
				credentials[k] = v
			}
		case t.RemoveKey != nil:
			delete(credentials, t.RemoveKey.Key)
		}
	}
	return nil
}

func evaluateJSONPath(jsonPath string, credentials map[string]interface{}) (string, error) {
	j := jsonpath.New("expression")
	buf := new(bytes.Buffer)
	if err := j.Parse(jsonPath); err != nil {
		return "", err
	}
	if err := j.Execute(buf, credentials); err != nil {
		return "", err
	}
	return buf.String(), nil
}
