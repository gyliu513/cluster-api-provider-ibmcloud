/*
Copyright 2019 The Kubernetes Authors.

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

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IbmcloudMachineProviderSpec defines the desired state of IbmcloudMachineProvider
// +k8s:openapi-gen=true
type IbmcloudMachineProviderSpec struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Hostname string `json:"hostname,omitempty"`

	Domain string `json:"domain,omitempty"`

	MaxMemory int `json:"maxMemory,omitempty"`
	StartCpus int `json:"startCpus,omitempty"`

	Datacenter string `json:"dataCenter,omitempty"`

	OSReferenceCode string `json:"osReferenceCode,omitempty"`

	LocalDiskFlag bool `json:"localDiskFlag,omitempty"`

	HourlyBillingFlag bool `json:"hourlyBillingFlag,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IbmcloudMachineProviderSpecList contains a list of IbmcloudMachineProviderSpec
type IbmcloudMachineProviderSpecList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IbmcloudMachineProviderSpec `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IbmcloudMachineProviderSpec{}, &IbmcloudMachineProviderSpecList{})
}
