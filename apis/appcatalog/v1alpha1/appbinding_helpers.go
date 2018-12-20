package v1alpha1

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	crdutils "github.com/appscode/kutil/apiextensions/v1beta1"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
)

func (a AppBinding) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
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

func (a AppBinding) URL() (string, error) {
	c := a.Spec.ClientConfig
	if c.URL != nil {
		return *c.URL, nil
	} else if c.Service != nil {
		u := url.URL{
			Scheme:   c.Service.Scheme,
			Host:     fmt.Sprintf("%s.%s.svc:%d", c.Service.Name, a.Namespace, c.Service.Port),
			Path:     c.Service.Path,
			RawQuery: c.Service.Query,
		}
		return u.String(), nil
	}
	return "", errors.New("connection url is missing")
}

const (
	KeyUsername = "username"
	KeyPassword = "password"
)

func (a AppBinding) URLTemplate() (string, error) {
	rawurl, err := a.URL()
	if err != nil {
		return "", err
	}
	auth := fmt.Sprintf("{{%s}}:{{%s}}@", KeyUsername, KeyPassword)

	i := strings.Index(rawurl, "://")
	if i < 0 {
		return auth + rawurl, nil
	}
	return fmt.Sprintf(rawurl[:i+3] + auth + rawurl[i+3:]), nil
}
