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

type KeyVaultConnectionStringObservation struct {
}

type KeyVaultConnectionStringParameters struct {

	// Specifies the name of an existing Key Vault Data Factory Linked Service.
	// +kubebuilder:validation:Required
	LinkedServiceName *string `json:"linkedServiceName" tf:"linked_service_name,omitempty"`

	// Specifies the secret name in Azure Key Vault that stores SQL Server connection string.
	// +kubebuilder:validation:Required
	SecretName *string `json:"secretName" tf:"secret_name,omitempty"`
}

type LinkedServiceAzureSQLDatabaseKeyVaultPasswordObservation struct {
}

type LinkedServiceAzureSQLDatabaseKeyVaultPasswordParameters struct {

	// Specifies the name of an existing Key Vault Data Factory Linked Service.
	// +kubebuilder:validation:Required
	LinkedServiceName *string `json:"linkedServiceName" tf:"linked_service_name,omitempty"`

	// Specifies the secret name in Azure Key Vault that stores SQL Server password.
	// +kubebuilder:validation:Required
	SecretName *string `json:"secretName" tf:"secret_name,omitempty"`
}

type LinkedServiceAzureSQLDatabaseObservation struct {

	// The ID of the Data Factory Azure SQL Database Linked Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LinkedServiceAzureSQLDatabaseParameters struct {

	// A map of additional properties to associate with the Data Factory Linked Service Azure SQL Database.
	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service Azure SQL Database.
	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The connection string in which to authenticate with Azure SQL Database. Exactly one of either connection_string or key_vault_connection_string is required.
	// +kubebuilder:validation:Optional
	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.Factory
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// Reference to a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDRef *v1.Reference `json:"dataFactoryIdRef,omitempty" tf:"-"`

	// Selector for a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDSelector *v1.Selector `json:"dataFactoryIdSelector,omitempty" tf:"-"`

	// The description for the Data Factory Linked Service Azure SQL Database.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service Azure SQL Database.
	// +kubebuilder:validation:Optional
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A key_vault_connection_string block as defined below. Use this argument to store Azure SQL Database connection string in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service. Exactly one of either connection_string or key_vault_connection_string is required.
	// +kubebuilder:validation:Optional
	KeyVaultConnectionString []KeyVaultConnectionStringParameters `json:"keyVaultConnectionString,omitempty" tf:"key_vault_connection_string,omitempty"`

	// A key_vault_password block as defined below. Use this argument to store SQL Server password in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	KeyVaultPassword []LinkedServiceAzureSQLDatabaseKeyVaultPasswordParameters `json:"keyVaultPassword,omitempty" tf:"key_vault_password,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service Azure SQL Database.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The service principal id in which to authenticate against the Azure SQL Database. Required if service_principal_key is set.
	// +kubebuilder:validation:Optional
	ServicePrincipalID *string `json:"servicePrincipalId,omitempty" tf:"service_principal_id,omitempty"`

	// The service principal key in which to authenticate against the Azure SQL Database. Required if service_principal_id is set.
	// +kubebuilder:validation:Optional
	ServicePrincipalKey *string `json:"servicePrincipalKey,omitempty" tf:"service_principal_key,omitempty"`

	// The tenant id or name in which to authenticate against the Azure SQL Database.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Whether to use the Data Factory's managed identity to authenticate against the Azure SQL Database. Incompatible with service_principal_id and service_principal_key
	// +kubebuilder:validation:Optional
	UseManagedIdentity *bool `json:"useManagedIdentity,omitempty" tf:"use_managed_identity,omitempty"`
}

// LinkedServiceAzureSQLDatabaseSpec defines the desired state of LinkedServiceAzureSQLDatabase
type LinkedServiceAzureSQLDatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LinkedServiceAzureSQLDatabaseParameters `json:"forProvider"`
}

// LinkedServiceAzureSQLDatabaseStatus defines the observed state of LinkedServiceAzureSQLDatabase.
type LinkedServiceAzureSQLDatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LinkedServiceAzureSQLDatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceAzureSQLDatabase is the Schema for the LinkedServiceAzureSQLDatabases API. Manages a Linked Service (connection) between Azure SQL Database and Azure Data Factory.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LinkedServiceAzureSQLDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LinkedServiceAzureSQLDatabaseSpec   `json:"spec"`
	Status            LinkedServiceAzureSQLDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceAzureSQLDatabaseList contains a list of LinkedServiceAzureSQLDatabases
type LinkedServiceAzureSQLDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LinkedServiceAzureSQLDatabase `json:"items"`
}

// Repository type metadata.
var (
	LinkedServiceAzureSQLDatabase_Kind             = "LinkedServiceAzureSQLDatabase"
	LinkedServiceAzureSQLDatabase_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LinkedServiceAzureSQLDatabase_Kind}.String()
	LinkedServiceAzureSQLDatabase_KindAPIVersion   = LinkedServiceAzureSQLDatabase_Kind + "." + CRDGroupVersion.String()
	LinkedServiceAzureSQLDatabase_GroupVersionKind = CRDGroupVersion.WithKind(LinkedServiceAzureSQLDatabase_Kind)
)

func init() {
	SchemeBuilder.Register(&LinkedServiceAzureSQLDatabase{}, &LinkedServiceAzureSQLDatabaseList{})
}
