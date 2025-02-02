/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type IOTHubEnrichmentObservation struct {

	// The ID of the IoTHub Enrichment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type IOTHubEnrichmentParameters struct {

	// The list of endpoints which will be enriched.
	// +kubebuilder:validation:Required
	EndpointNames []*string `json:"endpointNames" tf:"endpoint_names,omitempty"`

	// The IoTHub name of the enrichment. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/devices/v1beta1.IOTHub
	// +kubebuilder:validation:Optional
	IOTHubName *string `json:"iothubName,omitempty" tf:"iothub_name,omitempty"`

	// Reference to a IOTHub in devices to populate iothubName.
	// +kubebuilder:validation:Optional
	IOTHubNameRef *v1.Reference `json:"iothubNameRef,omitempty" tf:"-"`

	// Selector for a IOTHub in devices to populate iothubName.
	// +kubebuilder:validation:Optional
	IOTHubNameSelector *v1.Selector `json:"iothubNameSelector,omitempty" tf:"-"`

	// The key of the enrichment. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// The name of the resource group under which the IoTHub resource is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The value of the enrichment. Value can be any static string, the name of the IoT hub sending the message (use $iothubname) or information from the device twin (ex: $twin.tags.latitude)
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// IOTHubEnrichmentSpec defines the desired state of IOTHubEnrichment
type IOTHubEnrichmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IOTHubEnrichmentParameters `json:"forProvider"`
}

// IOTHubEnrichmentStatus defines the observed state of IOTHubEnrichment.
type IOTHubEnrichmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IOTHubEnrichmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubEnrichment is the Schema for the IOTHubEnrichments API. Manages an IotHub Enrichment
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IOTHubEnrichment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IOTHubEnrichmentSpec   `json:"spec"`
	Status            IOTHubEnrichmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubEnrichmentList contains a list of IOTHubEnrichments
type IOTHubEnrichmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IOTHubEnrichment `json:"items"`
}

// Repository type metadata.
var (
	IOTHubEnrichment_Kind             = "IOTHubEnrichment"
	IOTHubEnrichment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IOTHubEnrichment_Kind}.String()
	IOTHubEnrichment_KindAPIVersion   = IOTHubEnrichment_Kind + "." + CRDGroupVersion.String()
	IOTHubEnrichment_GroupVersionKind = CRDGroupVersion.WithKind(IOTHubEnrichment_Kind)
)

func init() {
	SchemeBuilder.Register(&IOTHubEnrichment{}, &IOTHubEnrichmentList{})
}
