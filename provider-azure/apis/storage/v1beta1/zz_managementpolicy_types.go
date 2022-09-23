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

type ActionsObservation struct {
}

type ActionsParameters struct {

	// A base_blob block as documented below.
	// +kubebuilder:validation:Optional
	BaseBlob []BaseBlobParameters `json:"baseBlob,omitempty" tf:"base_blob,omitempty"`

	// A snapshot block as documented below.
	// +kubebuilder:validation:Optional
	Snapshot []SnapshotParameters `json:"snapshot,omitempty" tf:"snapshot,omitempty"`

	// A version block as documented below.
	// +kubebuilder:validation:Optional
	Version []VersionParameters `json:"version,omitempty" tf:"version,omitempty"`
}

type BaseBlobObservation struct {
}

type BaseBlobParameters struct {

	// The age in days after last access time to delete the blob. Must be between 0 and 99999.
	// +kubebuilder:validation:Optional
	DeleteAfterDaysSinceLastAccessTimeGreaterThan *float64 `json:"deleteAfterDaysSinceLastAccessTimeGreaterThan,omitempty" tf:"delete_after_days_since_last_access_time_greater_than,omitempty"`

	// The age in days after last modification to delete the blob. Must be between 0 and 99999.
	// +kubebuilder:validation:Optional
	DeleteAfterDaysSinceModificationGreaterThan *float64 `json:"deleteAfterDaysSinceModificationGreaterThan,omitempty" tf:"delete_after_days_since_modification_greater_than,omitempty"`

	// The age in days after last access time to tier blobs to archive storage. Supports blob currently at Hot or Cool tier. Must be between 0 and 99999`.
	// +kubebuilder:validation:Optional
	TierToArchiveAfterDaysSinceLastAccessTimeGreaterThan *float64 `json:"tierToArchiveAfterDaysSinceLastAccessTimeGreaterThan,omitempty" tf:"tier_to_archive_after_days_since_last_access_time_greater_than,omitempty"`

	// The age in days after last modification to tier blobs to archive storage. Supports blob currently at Hot or Cool tier. Must be between 0 and 99999.
	// +kubebuilder:validation:Optional
	TierToArchiveAfterDaysSinceModificationGreaterThan *float64 `json:"tierToArchiveAfterDaysSinceModificationGreaterThan,omitempty" tf:"tier_to_archive_after_days_since_modification_greater_than,omitempty"`

	// The age in days after last access time to tier blobs to cool storage. Supports blob currently at Hot tier. Must be between 0 and 99999.
	// +kubebuilder:validation:Optional
	TierToCoolAfterDaysSinceLastAccessTimeGreaterThan *float64 `json:"tierToCoolAfterDaysSinceLastAccessTimeGreaterThan,omitempty" tf:"tier_to_cool_after_days_since_last_access_time_greater_than,omitempty"`

	// The age in days after last modification to tier blobs to cool storage. Supports blob currently at Hot tier. Must be between 0 and 99999.
	// +kubebuilder:validation:Optional
	TierToCoolAfterDaysSinceModificationGreaterThan *float64 `json:"tierToCoolAfterDaysSinceModificationGreaterThan,omitempty" tf:"tier_to_cool_after_days_since_modification_greater_than,omitempty"`
}

type FiltersObservation struct {
}

type FiltersParameters struct {

	// An array of predefined values. Valid options are blockBlob and appendBlob.
	// +kubebuilder:validation:Required
	BlobTypes []*string `json:"blobTypes" tf:"blob_types,omitempty"`

	// A match_blob_index_tag block as defined below. The block defines the blob index tag based filtering for blob objects.
	// +kubebuilder:validation:Optional
	MatchBlobIndexTag []MatchBlobIndexTagParameters `json:"matchBlobIndexTag,omitempty" tf:"match_blob_index_tag,omitempty"`

	// An array of strings for prefixes to be matched.
	// +kubebuilder:validation:Optional
	PrefixMatch []*string `json:"prefixMatch,omitempty" tf:"prefix_match,omitempty"`
}

type ManagementPolicyObservation struct {

	// The ID of the Storage Account Management Policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ManagementPolicyParameters struct {

	// A rule block as documented below.
	// +kubebuilder:validation:Optional
	Rule []RuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`

	// Specifies the id of the storage account to apply the management policy to.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/storage/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// Reference to a Account in storage to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDRef *v1.Reference `json:"storageAccountIdRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDSelector *v1.Selector `json:"storageAccountIdSelector,omitempty" tf:"-"`
}

type MatchBlobIndexTagObservation struct {
}

type MatchBlobIndexTagParameters struct {

	// A rule name can contain any combination of alpha numeric characters. Rule name is case-sensitive. It must be unique within a policy.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The comparison operator which is used for object comparison and filtering. Possible value is ==. Defaults to ==.
	// +kubebuilder:validation:Optional
	Operation *string `json:"operation,omitempty" tf:"operation,omitempty"`

	// The filter tag value used for tag based filtering for blob objects.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type RuleObservation struct {
}

type RuleParameters struct {

	// An actions block as documented below.
	// +kubebuilder:validation:Required
	Actions []ActionsParameters `json:"actions" tf:"actions,omitempty"`

	// Boolean to specify whether the rule is enabled.
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// A filter block as documented below.
	// +kubebuilder:validation:Optional
	Filters []FiltersParameters `json:"filters,omitempty" tf:"filters,omitempty"`

	// A rule name can contain any combination of alpha numeric characters. Rule name is case-sensitive. It must be unique within a policy.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type SnapshotObservation struct {
}

type SnapshotParameters struct {

	// The age in days after creation to tier blob snapshot to archive storage. Must be between 0 and 99999.
	// +kubebuilder:validation:Optional
	ChangeTierToArchiveAfterDaysSinceCreation *float64 `json:"changeTierToArchiveAfterDaysSinceCreation,omitempty" tf:"change_tier_to_archive_after_days_since_creation,omitempty"`

	// The age in days after creation to tier blob snapshot to cool storage. Must be between 0 and 99999.
	// +kubebuilder:validation:Optional
	ChangeTierToCoolAfterDaysSinceCreation *float64 `json:"changeTierToCoolAfterDaysSinceCreation,omitempty" tf:"change_tier_to_cool_after_days_since_creation,omitempty"`

	// The age in days after creation to delete the blob snapshot. Must be between 0 and 99999.
	// +kubebuilder:validation:Optional
	DeleteAfterDaysSinceCreationGreaterThan *float64 `json:"deleteAfterDaysSinceCreationGreaterThan,omitempty" tf:"delete_after_days_since_creation_greater_than,omitempty"`
}

type VersionObservation struct {
}

type VersionParameters struct {

	// The age in days after creation to tier blob snapshot to archive storage. Must be between 0 and 99999.
	// +kubebuilder:validation:Optional
	ChangeTierToArchiveAfterDaysSinceCreation *float64 `json:"changeTierToArchiveAfterDaysSinceCreation,omitempty" tf:"change_tier_to_archive_after_days_since_creation,omitempty"`

	// The age in days after creation to tier blob snapshot to cool storage. Must be between 0 and 99999.
	// +kubebuilder:validation:Optional
	ChangeTierToCoolAfterDaysSinceCreation *float64 `json:"changeTierToCoolAfterDaysSinceCreation,omitempty" tf:"change_tier_to_cool_after_days_since_creation,omitempty"`

	// The age in days after creation to delete the blob version. Must be between 0 and 99999.
	// +kubebuilder:validation:Optional
	DeleteAfterDaysSinceCreation *float64 `json:"deleteAfterDaysSinceCreation,omitempty" tf:"delete_after_days_since_creation,omitempty"`
}

// ManagementPolicySpec defines the desired state of ManagementPolicy
type ManagementPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagementPolicyParameters `json:"forProvider"`
}

// ManagementPolicyStatus defines the observed state of ManagementPolicy.
type ManagementPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagementPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagementPolicy is the Schema for the ManagementPolicys API. Manages an Azure Storage Account Management Policy.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ManagementPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagementPolicySpec   `json:"spec"`
	Status            ManagementPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagementPolicyList contains a list of ManagementPolicys
type ManagementPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagementPolicy `json:"items"`
}

// Repository type metadata.
var (
	ManagementPolicy_Kind             = "ManagementPolicy"
	ManagementPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagementPolicy_Kind}.String()
	ManagementPolicy_KindAPIVersion   = ManagementPolicy_Kind + "." + CRDGroupVersion.String()
	ManagementPolicy_GroupVersionKind = CRDGroupVersion.WithKind(ManagementPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagementPolicy{}, &ManagementPolicyList{})
}
