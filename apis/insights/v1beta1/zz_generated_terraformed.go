/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this ApplicationInsights
func (mg *ApplicationInsights) GetTerraformResourceType() string {
	return "azurerm_application_insights"
}

// GetConnectionDetailsMapping for this ApplicationInsights
func (tr *ApplicationInsights) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"connection_string": "status.atProvider.connectionString", "instrumentation_key": "status.atProvider.instrumentationKey"}
}

// GetObservation of this ApplicationInsights
func (tr *ApplicationInsights) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ApplicationInsights
func (tr *ApplicationInsights) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ApplicationInsights
func (tr *ApplicationInsights) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ApplicationInsights
func (tr *ApplicationInsights) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ApplicationInsights
func (tr *ApplicationInsights) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ApplicationInsights using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ApplicationInsights) LateInitialize(attrs []byte) (bool, error) {
	params := &ApplicationInsightsParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ApplicationInsights) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this ApplicationInsightsAPIKey
func (mg *ApplicationInsightsAPIKey) GetTerraformResourceType() string {
	return "azurerm_application_insights_api_key"
}

// GetConnectionDetailsMapping for this ApplicationInsightsAPIKey
func (tr *ApplicationInsightsAPIKey) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"api_key": "status.atProvider.apiKey"}
}

// GetObservation of this ApplicationInsightsAPIKey
func (tr *ApplicationInsightsAPIKey) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ApplicationInsightsAPIKey
func (tr *ApplicationInsightsAPIKey) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ApplicationInsightsAPIKey
func (tr *ApplicationInsightsAPIKey) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ApplicationInsightsAPIKey
func (tr *ApplicationInsightsAPIKey) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ApplicationInsightsAPIKey
func (tr *ApplicationInsightsAPIKey) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ApplicationInsightsAPIKey using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ApplicationInsightsAPIKey) LateInitialize(attrs []byte) (bool, error) {
	params := &ApplicationInsightsAPIKeyParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ApplicationInsightsAPIKey) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this MonitorActionGroup
func (mg *MonitorActionGroup) GetTerraformResourceType() string {
	return "azurerm_monitor_action_group"
}

// GetConnectionDetailsMapping for this MonitorActionGroup
func (tr *MonitorActionGroup) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this MonitorActionGroup
func (tr *MonitorActionGroup) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this MonitorActionGroup
func (tr *MonitorActionGroup) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this MonitorActionGroup
func (tr *MonitorActionGroup) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this MonitorActionGroup
func (tr *MonitorActionGroup) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this MonitorActionGroup
func (tr *MonitorActionGroup) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this MonitorActionGroup using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *MonitorActionGroup) LateInitialize(attrs []byte) (bool, error) {
	params := &MonitorActionGroupParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *MonitorActionGroup) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this MonitorMetricAlert
func (mg *MonitorMetricAlert) GetTerraformResourceType() string {
	return "azurerm_monitor_metric_alert"
}

// GetConnectionDetailsMapping for this MonitorMetricAlert
func (tr *MonitorMetricAlert) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this MonitorMetricAlert
func (tr *MonitorMetricAlert) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this MonitorMetricAlert
func (tr *MonitorMetricAlert) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this MonitorMetricAlert
func (tr *MonitorMetricAlert) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this MonitorMetricAlert
func (tr *MonitorMetricAlert) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this MonitorMetricAlert
func (tr *MonitorMetricAlert) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this MonitorMetricAlert using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *MonitorMetricAlert) LateInitialize(attrs []byte) (bool, error) {
	params := &MonitorMetricAlertParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *MonitorMetricAlert) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this MonitorPrivateLinkScope
func (mg *MonitorPrivateLinkScope) GetTerraformResourceType() string {
	return "azurerm_monitor_private_link_scope"
}

// GetConnectionDetailsMapping for this MonitorPrivateLinkScope
func (tr *MonitorPrivateLinkScope) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this MonitorPrivateLinkScope
func (tr *MonitorPrivateLinkScope) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this MonitorPrivateLinkScope
func (tr *MonitorPrivateLinkScope) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this MonitorPrivateLinkScope
func (tr *MonitorPrivateLinkScope) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this MonitorPrivateLinkScope
func (tr *MonitorPrivateLinkScope) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this MonitorPrivateLinkScope
func (tr *MonitorPrivateLinkScope) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this MonitorPrivateLinkScope using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *MonitorPrivateLinkScope) LateInitialize(attrs []byte) (bool, error) {
	params := &MonitorPrivateLinkScopeParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *MonitorPrivateLinkScope) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this MonitorPrivateLinkScopedService
func (mg *MonitorPrivateLinkScopedService) GetTerraformResourceType() string {
	return "azurerm_monitor_private_link_scoped_service"
}

// GetConnectionDetailsMapping for this MonitorPrivateLinkScopedService
func (tr *MonitorPrivateLinkScopedService) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this MonitorPrivateLinkScopedService
func (tr *MonitorPrivateLinkScopedService) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this MonitorPrivateLinkScopedService
func (tr *MonitorPrivateLinkScopedService) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this MonitorPrivateLinkScopedService
func (tr *MonitorPrivateLinkScopedService) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this MonitorPrivateLinkScopedService
func (tr *MonitorPrivateLinkScopedService) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this MonitorPrivateLinkScopedService
func (tr *MonitorPrivateLinkScopedService) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this MonitorPrivateLinkScopedService using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *MonitorPrivateLinkScopedService) LateInitialize(attrs []byte) (bool, error) {
	params := &MonitorPrivateLinkScopedServiceParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *MonitorPrivateLinkScopedService) GetTerraformSchemaVersion() int {
	return 0
}
