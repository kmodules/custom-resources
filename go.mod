module kmodules.xyz/custom-resources

go 1.12

require (
	github.com/appscode/go v0.0.0-20200323182826-54e98e09185a
	github.com/go-openapi/spec v0.19.3
	github.com/gogo/protobuf v1.3.1
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/google/gofuzz v1.1.0
	github.com/json-iterator/go v1.1.8
	github.com/kubernetes-incubator/service-catalog v0.2.3
	github.com/kubernetes-sigs/service-catalog v0.2.3 // indirect
	k8s.io/api v0.18.3
	k8s.io/apiextensions-apiserver v0.18.3
	k8s.io/apimachinery v0.18.3
	k8s.io/client-go v0.18.3
	k8s.io/kube-openapi v0.0.0-20200410145947-61e04a5be9a6
	kmodules.xyz/client-go v0.0.0-20200521005126-35ce6bd4ed46
	kmodules.xyz/crd-schema-fuzz v0.0.0-20200521005638-2433a187de95
	sigs.k8s.io/yaml v1.2.0
)

replace (
	k8s.io/apimachinery => github.com/kmodules/apimachinery v0.19.0-alpha.0.0.20200520235721-10b58e57a423
	k8s.io/apiserver => github.com/kmodules/apiserver v0.18.4-0.20200521000930-14c5f6df9625
)
