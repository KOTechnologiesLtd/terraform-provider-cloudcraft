package cloudcraft

import (
	"github.com/KOTechnologiesLtd/go-cloudcraft-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider configuration for the cloudcraft terraform provider.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{

			"apikey": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLOUDCRAFT_APITOKEN", nil),
				Description: "apikey for cloudcraft",
				Sensitive:   true,
			},
			"baseurl": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLOUDCRAFT_HOST", "api.cloudcraft.co"),
				Description: "Host URL for cloudcraft",
			},
			"max_retries": {
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLOUDCRAFT_MAX_RETRIES", 1),
				Description: "Max Retries",
			},
		},
		ConfigureFunc: providerConfigure,

		ResourcesMap: map[string]*schema.Resource{
			"cloudcraft_integration_aws": resourceIntegrationAws(),
			"cloudcraft_blueprint":       resourceBlueprint(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"cloudcraft_integration_aws": dataSourceIntegrationAws(),
			"cloudcraft_blueprint":       datasourceBlueprint(),
		},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	apiKey := d.Get("apikey").(string)
	baseurl := d.Get("baseurl").(string)
	max_retries := d.Get("max_retries").(int)

	client := cloudcraft.NewClient(apiKey, baseurl, max_retries)

	return client, nil
}
