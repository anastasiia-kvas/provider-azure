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

type NotificationRecipientUserObservation struct {

	// The ID of the API Management Notification Recipient User.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type NotificationRecipientUserParameters struct {

	// The ID of the API Management Service from which to create this Notification Recipient User. Changing this forces a new API Management Notification Recipient User to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.Management
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	APIManagementID *string `json:"apiManagementId,omitempty" tf:"api_management_id,omitempty"`

	// Reference to a Management in apimanagement to populate apiManagementId.
	// +kubebuilder:validation:Optional
	APIManagementIDRef *v1.Reference `json:"apiManagementIdRef,omitempty" tf:"-"`

	// Selector for a Management in apimanagement to populate apiManagementId.
	// +kubebuilder:validation:Optional
	APIManagementIDSelector *v1.Selector `json:"apiManagementIdSelector,omitempty" tf:"-"`

	// The Notification Name to be received. Changing this forces a new API Management Notification Recipient User to be created. Possible values are AccountClosedPublisher, BCC, NewApplicationNotificationMessage, NewIssuePublisherNotificationMessage, PurchasePublisherNotificationMessage, QuotaLimitApproachingPublisherNotificationMessage, and RequestPublisherNotificationMessage.
	// +kubebuilder:validation:Required
	NotificationType *string `json:"notificationType" tf:"notification_type,omitempty"`

	// The recipient user ID. Changing this forces a new API Management Notification Recipient User to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.User
	// +kubebuilder:validation:Optional
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// Reference to a User in apimanagement to populate userId.
	// +kubebuilder:validation:Optional
	UserIDRef *v1.Reference `json:"userIdRef,omitempty" tf:"-"`

	// Selector for a User in apimanagement to populate userId.
	// +kubebuilder:validation:Optional
	UserIDSelector *v1.Selector `json:"userIdSelector,omitempty" tf:"-"`
}

// NotificationRecipientUserSpec defines the desired state of NotificationRecipientUser
type NotificationRecipientUserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NotificationRecipientUserParameters `json:"forProvider"`
}

// NotificationRecipientUserStatus defines the observed state of NotificationRecipientUser.
type NotificationRecipientUserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NotificationRecipientUserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NotificationRecipientUser is the Schema for the NotificationRecipientUsers API. Manages a API Management Notification Recipient User.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type NotificationRecipientUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NotificationRecipientUserSpec   `json:"spec"`
	Status            NotificationRecipientUserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NotificationRecipientUserList contains a list of NotificationRecipientUsers
type NotificationRecipientUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NotificationRecipientUser `json:"items"`
}

// Repository type metadata.
var (
	NotificationRecipientUser_Kind             = "NotificationRecipientUser"
	NotificationRecipientUser_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NotificationRecipientUser_Kind}.String()
	NotificationRecipientUser_KindAPIVersion   = NotificationRecipientUser_Kind + "." + CRDGroupVersion.String()
	NotificationRecipientUser_GroupVersionKind = CRDGroupVersion.WithKind(NotificationRecipientUser_Kind)
)

func init() {
	SchemeBuilder.Register(&NotificationRecipientUser{}, &NotificationRecipientUserList{})
}
