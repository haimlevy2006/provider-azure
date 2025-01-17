/*
Copyright 2019 The Crossplane Authors.

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

package v1alpha3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2022-06-01/containerservice"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

const (
	// DefaultNodeCount is the default node count for a cluster.
	DefaultNodeCount = 1
)

// AKSClusterParameters define the desired state of an Azure Kubernetes Engine
// cluster.
type AKSClusterParameters struct {
	// ResourceGroupName is the name of the resource group that the cluster will
	// be created in
	ResourceGroupName string `json:"resourceGroupName,omitempty"`

	// ResourceGroupNameRef - A reference to a ResourceGroup to retrieve its
	// name
	ResourceGroupNameRef *xpv1.Reference `json:"resourceGroupNameRef,omitempty"`

	// ResourceGroupNameSelector - Select a reference to a ResourceGroup to
	// retrieve its name
	ResourceGroupNameSelector *xpv1.Selector `json:"resourceGroupNameSelector,omitempty"`

	// Location is the Azure location that the cluster will be created in
	Location string `json:"location"`

	// Version is the Kubernetes version that will be deployed to the cluster
	Version string `json:"version"`

	// VnetSubnetID is the subnet to which the cluster will be deployed.
	// +optional
	VnetSubnetID string `json:"vnetSubnetID,omitempty"`

	// VnetSubnetIDRef - A reference to a Subnet to retrieve its ID
	VnetSubnetIDRef *xpv1.Reference `json:"vnetSubnetIDRef,omitempty"`

	// VnetSubnetIDSelector - Select a reference to a Subnet to retrieve
	// its ID
	VnetSubnetIDSelector *xpv1.Selector `json:"vnetSubnetIDSelector,omitempty"`

	// AgentPools Specs
	AgentPoolProfiles []AgentPoolProfiles `json:"agentPoolProfiles"`

	// NetworkProfile Specs
	NetworkProfile NetworkProfile `json:"networkProfile"`

	// DNSNamePrefix is the DNS name prefix to use with the hosted Kubernetes
	// API server FQDN. You will use this to connect to the Kubernetes API when
	// managing containers after creating the cluster.
	// +optional
	DNSNamePrefix string `json:"dnsNamePrefix"`

	// DisableRBAC determines whether RBAC will be disabled or enabled in the
	// cluster.
	// +optional
	DisableRBAC bool `json:"disableRBAC,omitempty"`
}

// An AgentPool Specs defines the desired state of an AgentPool.
type AgentPoolProfiles struct {
	Name       *string                        `json:"name"`
	NodeCount  *int32                         `json:"nodeCount"`
	NodeVMSize *string                        `json:"nodeVMSize"`
	Mode       containerservice.AgentPoolMode `json:"mode"`
	Type       containerservice.AgentPoolType `json:"type"`
	MaxPods    *int32                         `json:"maxPods,omitempty"`
}

// AKS Cluster Network Profile Specs
type NetworkProfile struct {
	NetworkPlugin containerservice.NetworkPlugin `json:"networkPlugin,omitempty"`
	NetworkPolicy containerservice.NetworkPolicy `json:"networkPolicy,omitempty"`
}

// An AKSClusterSpec defines the desired state of a AKSCluster.
type AKSClusterSpec struct {
	xpv1.ResourceSpec    `json:",inline"`
	AKSClusterParameters `json:",inline"`
}

// An AKSClusterStatus represents the observed state of an AKSCluster.
type AKSClusterStatus struct {
	xpv1.ResourceStatus `json:",inline"`

	// State is the current state of the cluster.
	State string `json:"state,omitempty"`

	// ProviderID is the external ID to identify this resource in the cloud
	// provider.
	ProviderID string `json:"providerID,omitempty"`

	// Endpoint is the endpoint where the cluster can be reached
	Endpoint string `json:"endpoint,omitempty"`
}

// +kubebuilder:object:root=true

// An AKSCluster is a managed resource that represents an Azure Kubernetes
// Engine cluster.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="ENDPOINT",type="string",JSONPath=".status.endpoint"
// +kubebuilder:printcolumn:name="LOCATION",type="string",JSONPath=".spec.location"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
// +kubebuilder:subresource:status
type AKSCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AKSClusterSpec   `json:"spec"`
	Status AKSClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AKSClusterList contains a list of AKSCluster.
type AKSClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AKSCluster `json:"items"`
}
