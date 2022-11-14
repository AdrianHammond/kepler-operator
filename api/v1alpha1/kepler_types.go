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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type Cgroupv2 string

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

type ModelServerSpec struct {
	Foo string `json:"foo,omitempty"`
}

type EstimatorSpec struct {
	Foo string `json:"foo,omitempty"`
}

// KeplerSpec defines the desired state of Kepler
type KeplerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Kepler. Edit kepler_types.go to remove/update
	ModelServer ModelServerSpec `json:"modelServer,omitempty"`
	Estimator   EstimatorSpec   `json:"estimatorSpec,omitempty"`
	Collector   CollectorSpec   `json:"collectorSpec,omitempty"`
}

// KeplerStatus defines the observed state of Kepler
type KeplerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Kepler is the Schema for the keplers API
type Kepler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KeplerSpec   `json:"spec,omitempty"`
	Status KeplerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// KeplerList contains a list of Kepler
type KeplerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Kepler `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Kepler{}, &KeplerList{})
}
