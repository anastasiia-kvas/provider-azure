package botservice

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures the botservice group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_bot_service_azure_bot", func(r *config.Resource) {
		r.References["developer_app_insights_api_key"] = config.Reference{
			Type: "github.com/upbound/provider-azure/apis/insights/v1beta1.ApplicationInsightsAPIKey",
		}
	})
}
