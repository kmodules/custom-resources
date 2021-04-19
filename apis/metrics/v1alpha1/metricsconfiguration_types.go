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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindMetricsConfiguration     = "MetricsConfiguration"
	ResourceSingularMetricsConfiguration = "metricsconfiguration"
	ResourcePluralMetricsConfiguration   = "metricsconfigurations"
)

// MetricsConfiguration defines a generic prometheus metrics configuration

// +genclient
// +genclient:nonNamespaced
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=metricsconfigurations,singular=metricsconfiguration,scope=Cluster,categories={metrics,appscode,all}
// +kubebuilder:printcolumn:name="Group",type="string",JSONPath=".spec.targetRef.group"
// +kubebuilder:printcolumn:name="Resource",type="string",JSONPath=".spec.targetRef.resource"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
type MetricsConfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              MetricsConfigurationSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// MetricsConfigurationSpec is the spec of MetricsConfiguration
type MetricsConfigurationSpec struct {
	// TargetRef defines the object for which metrics will be collected
	TargetRef TargetRef `json:"targetRef" protobuf:"bytes,1,opt,name=targetRef"`

	// List of Metrics configuration for the resource defined in TargetRef
	Metrics []Metrics `json:"metrics" protobuf:"bytes,2,rep,name=metrics"`
}

// TargetRef defines the Object's group, version, resource
type TargetRef struct {
	// Group defines the group of the object
	Group string `json:"group" protobuf:"bytes,1,opt,name=group"`

	// Resource defines the resource of the object
	Resource string `json:"resource" protobuf:"bytes,2,opt,name=resource"`
}

// Metrics contains the configuration of a metric in prometheus style
type Metrics struct {
	// Name defines the metrics name. Example: kube_deployment_spec_replicas
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`

	// Help is used to describe the metrics. Example: "Number of desired pods for a deployment."
	Help string `json:"help" protobuf:"bytes,2,opt,name=help"`

	// Type defines the metrics type. Example: gauge
	Type string `json:"type" protobuf:"bytes,3,opt,name=type"`

	// Field defines the metric value path of the manifest file and the type of that value
	// +optional
	Field Field `json:"field,omitempty" protobuf:"bytes,4,opt,name=field"`

	// Labels defines the metric labels as a key-value pair
	// +optional
	Labels []Label `json:"labels,omitempty" protobuf:"bytes,5,rep,name=labels"`

	// Params is list of parameters configuration used in expression evaluation
	// +optional
	Params []Parameter `json:"params,omitempty" protobuf:"bytes,6,rep,name=params"`

	// States handles metrics with label cardinality
	// +optional
	States State `json:"states,omitempty" protobuf:"bytes,7,opt,name=states"`

	// MetricValue defines the configuration for the metric value
	// +optional
	MetricValue MetricValue `json:"metricValue,omitempty" protobuf:"bytes,8,opt,name=metricValue"`
}

type FieldType string

const (
	FieldTypeInteger    FieldType = "Integer"
	FieldTypePercentage FieldType = "Percentage"
	FieldTypeDateTime   FieldType = "DateTime"
	FieldTypeArray      FieldType = "Array"
)

// Field contains the information of a manifest field for which metric is collected
type Field struct {
	// Path defines the manifest file path, example: spec.replicas
	Path string `json:"path" protobuf:"bytes,1,opt,name=path"`

	// Type defines the type of the value in the given Path
	Type FieldType `json:"type" protobuf:"bytes,2,opt,name=type"`
}

// Label contains the information of a metric label
type Label struct {
	// Key defines the label key
	Key string `json:"key" protobuf:"bytes,1,opt,name=key"`

	// Value defines the hard coded label value
	// +optional
	Value string `json:"value,omitempty" protobuf:"bytes,2,opt,name=value"`

	// ValuePath defines the label value path. example: spec.replicas
	// +optional
	ValuePath string `json:"valuePath" protobuf:"bytes,3,opt,name=valuePath"`
}

// Parameter contains the information of a parameter used in expression evaluation
type Parameter struct {
	// Key defines the parameter's key
	Key string `json:"key" protobuf:"bytes,1,opt,name=key"`

	// Value defines the parameter's value
	// +optional
	Value string `json:"value,omitempty" protobuf:"bytes,2,opt,name=value"`

	// ValuePath defines the manifest field path for the parameter's value. example: spec.replicas
	// +optional
	ValuePath string `json:"valuePath,omitempty" protobuf:"bytes,3,opt,name=valuePath"`
}

// State contains the configuration for generating all the time series
// of a metric with label cardinality is greater than 1
type State struct {
	// LabelKey defines the label key of the label which label cardinality is greater than one
	// example: labelKey = phase
	LabelKey string `json:"labelKey" protobuf:"bytes,1,opt,name=labelKey"`

	// Values contains the list of state values
	Values []StateValues `json:"values" protobuf:"bytes,2,rep,name=values"`
}

// StateValues contains the information of a state value
type StateValues struct {
	// LabelValue defines the value of the label.
	// Example: For labelKey phase label value can be "Running"
	LabelValue string `json:"labelValue" protobuf:"bytes,1,opt,name=labelValue"`

	// MetricValue defines the configuration of the metric value for the corresponding LabelValue
	MetricValue MetricValue `json:"metricValue" protobuf:"bytes,2,opt,name=metricValue"`
}

// MetricValue contains the configuration to obtain the value for a metric
type MetricValue struct {
	// Value contains the metric value. This is the default value of a metric.
	// Most of the time it is equal to one.
	// +optional
	Value string `json:"value,omitempty" protobuf:"bytes,1,opt,name=value"`

	// ValueFromPath contains the field path of the manifest file of a object.
	// example: spec.replicas
	// +optional
	ValueFromPath string `json:"valueFromPath,omitempty" protobuf:"bytes,2,opt,name=valueFromPath"`

	// ValueFromExpression contains an expression for the metric value
	// expression can be a function as well.
	// Used expression evaluation functions are:
	//
	// toInt() returns 1 if the expression is true otherwise 0,
	// example: toInt(phase == 'Running')
	//
	// evaluatePercentage(a, b) returns the value of a * b%
	// example: evaluatePercentage(replicas, maxUnavailable)
	//
	// calculateCPU() returns the cpu in unit core
	// example: calculateCPU(cpu), for cpu value 150m, it will return 0.15
	//
	// calculateMemory() returns the memory size in byte
	// example: calculateMemory(memory), for memory value 1 ki, it will return 1024
	//
	// toUnix() returns the DateTime string into unix format.
	// example: toUnix(dateTime) will return the corresponding unix value for the given dateTime
	//
	// in above examples phase, replicas, maxUnavailable, cpu, memory, dateTime are Parameter's key
	// those values will come from corresponding Parameter's value
	// +optional
	ValueFromExpression string `json:"valueFromExpression,omitempty" protobuf:"bytes,3,opt,name=valueFromExpression"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MetricsConfigurationList is a list of MetricsConfiguration
type MetricsConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []MetricsConfiguration `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}
