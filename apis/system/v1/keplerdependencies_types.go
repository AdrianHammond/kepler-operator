/*
Copyright 2022.

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

package v1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

//Enums

// +kubebuilder:validation:Enum=FooOne;FooTwo
type ModelPipeline string

// +kubebuilder:validation:Enum=enable;disable
type Cgroupv2 string

const (
	// random list of pipelines
	FooOnePipeline ModelPipeline = "FooOne"
	FooTwoPipeline ModelPipeline = "FooTwo"
	// enum for cgroup options
	EnableCgroupv2  Cgroupv2 = "enable"
	DisableCgroupv2 Cgroupv2 = "disable"
)

// Collector CRD Requirements

type RatioMetrics struct {
	Global string `json:"global,omitempty"`
	Core   string `json:"core,omitempty"`
	Uncore string `json:"uncore,omitempty"`
	Dram   string `json:"dram,omitempty"`
}

type Sources struct {
	Cgroupv2 Cgroupv2 `json:"cgroupv2,omitempty"`
	Bpf      string   `json:"bpf,omitempty"`
	Counters string   `json:"counters,omitempty"`
	Kubelet  string   `json:"kubelet,omitempty"`
}

type CollectorSpec struct {
	Image        string       `json:"image,omitempty"`
	Sources      Sources      `json:"sources,omitempty"`
	RatioMetrics RatioMetrics `json:"ratioMetrics,omitempty"`
}

// Estimater CRD Requirements

type EstimatorSpec struct {
	// +kubebuilder:default=true
	Enable   bool                  `json:"enable,omitempty"`
	Image    string                `json:"image,omitempty"`
	Strategy []v1.NodeSelectorTerm `json:"strategy,omitempty"`
}

// Model Server CRD Requirements
type ModelStorage struct {
	// +kubebuilder:default="local"
	Type     string `json:"type,omitempty"`
	HostPath string `json:"hostPath,omitempty"`
}

type ModelServerSpec struct {
	// +kubebuilder:default=false
	Install          bool            `json:"install,omitempty"`
	Image            string          `json:"image,omitempty"`
	QueryStep        int64           `json:"queryStep,omitempty"`
	SamplingInterval int64           `json:"samplingInterval,omitempty"`
	EnablePipelines  []ModelPipeline `json:"enablePipelines"`
	ModelStorage     ModelStorage    `json:"modelStorage,omitempty"`
}

// KeplerDependenciesSpec defines the desired state of KeplerDependencies
type KeplerDependenciesSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	ModelServer ModelServerSpec `json:"modelServer,omitempty"`
	Estimator   EstimatorSpec   `json:"estimatorSpec,omitempty"`
	Collector   CollectorSpec   `json:"collectorSpec,omitempty"`
}

// KeplerDependenciesStatus defines the observed state of KeplerDependencies
type KeplerDependenciesStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// KeplerDependencies is the Schema for the keplerdependencies API
type KeplerDependencies struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KeplerDependenciesSpec   `json:"spec,omitempty"`
	Status KeplerDependenciesStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// KeplerDependenciesList contains a list of KeplerDependencies
type KeplerDependenciesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeplerDependencies `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KeplerDependencies{}, &KeplerDependenciesList{})
}
