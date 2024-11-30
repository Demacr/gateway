// Copyright Envoy Gateway Authors
// SPDX-License-Identifier: Apache-2.0
// The full text of the Apache license is available in the LICENSE file at
// the root of the repo.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// KindVirtualBackend is the name of the VirtualBackend kind.
	KindVirtualBackend = "VirtualBackend"
)

// +kubebuilder:validation:Minimum=100
// +kubebuilder:validation:Maximum=599
type StatusCode uint32

// +kubebuilder:validation:Pattern=`^([\w-]+)$`
type ResponseHeader string

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=envoy-gateway,shortName=vb
// +kubebuilder:subresource:status
//
// VirtualBackend defines the configuration for direct response.
type VirtualBackend struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines desired state of VirtualBackend.
	Spec VirtualBackendSpec `json:"spec"`

	// Status defines the current status of SecurityPolicy.
	Status VirtualBackendStatus `json:"status,omitempty"`
}

// VirtualBackendStatus defines the state of virtualbackend
type VirtualBackendStatus struct {
	// Conditions describe the current conditions of the Backend.
	//
	// +optional
	// +listType=map
	// +listMapKey=type
	// +kubebuilder:validation:MaxItems=8
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:validation:XValidation:rule="has(self.statusCode)"
//
// VirtualBackendSpec defines direct response configuration.
type VirtualBackendSpec struct {
	// +optional
	//
	// Body contains data which gateway returns in direct response.
	Body []byte `json:"body,omitempty" yaml:"body,omitempty"`

	// +kubebuilder:default=200
	//
	// StatusCode defines HTTP response status code of direct response. Default value is 200.
	StatusCode StatusCode `json:"statusCode" yaml:"statusCode"`

	// +optional
	//
	// ResponseHeaders defines Header:Value map of additional headers to response.
	ResponseHeaders map[ResponseHeader]string `json:"responseHeaders,omitempty" yaml:"responseHeaders,omitempty"`
}

// +kubebuilder:object:root=true
//
// VirtualBackendList contains a list of VirtualBackend resources.
type VirtualBackendList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualBackend `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualBackend{}, &VirtualBackendList{})
}
