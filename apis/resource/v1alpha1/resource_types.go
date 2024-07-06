/*
Copyright 2022 The Crossplane Authors.

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
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// ResourceParameters are the configurable fields of a Resource.
type ResourceParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// ResourceObservation are the observable fields of a Resource.
type ResourceObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A ResourceSpec defines the desired state of a Resource.
type ResourceSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ResourceParameters `json:"forProvider"`
}

// A ResourceStatus represents the observed state of a Resource.
type ResourceStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ResourceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Resource is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,rest}
type Resource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ResourceSpec   `json:"spec"`
	Status ResourceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceList contains a list of Resource
type ResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Resource `json:"items"`
}

// Resource type metadata.
var (
	ResourceKind             = reflect.TypeOf(Resource{}).Name()
	ResourceGroupKind        = schema.GroupKind{Group: Group, Kind: ResourceKind}.String()
	ResourceKindAPIVersion   = ResourceKind + "." + SchemeGroupVersion.String()
	ResourceGroupVersionKind = SchemeGroupVersion.WithKind(ResourceKind)
)

func init() {
	SchemeBuilder.Register(&Resource{}, &ResourceList{})
}
