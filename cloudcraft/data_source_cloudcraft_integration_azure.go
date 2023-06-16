package cloudcraft

import (
	"context"

	"github.com/KOTechnologiesLtd/go-cloudcraft-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIntegrationAzure() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIntegrationAzureRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"applicationid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"directoryid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subscriptionid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"clientsecret": {
				Type:     schema.TypeString,
				Required: true,
			},
			"externalid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"createdat": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"updatedat": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"creatorid": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIntegrationAzureRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*cloudcraft.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	azureAccountId := d.Get("id").(string)

	accountInfo, err := client.AccountIntegrationAzure(azureAccountId)
	{
		if err != nil {
			return diag.FromErr(err)
		}
	}
	//log.Printf("accounts%v ", accountInfo)
	if (cloudcraft.AccountIntegrationAzure{}) == accountInfo {
		diags := append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "ID Not Found",
			Detail:   "The CloudCraft Azure Integration ID does not exist: " + azureAccountId,
		})
		return diags
	}

	//populate account scheme
	d.Set("name", accountInfo.Name)
	d.Set("applicationid", accountInfo.ApplicationId)
	d.Set("directoryid", accountInfo.DirectoryId)
	d.Set("subscriptionid", accountInfo.SubscriptionId)
	d.Set("clientsecret", accountInfo.ClientSecret)
	d.Set("externalid", accountInfo.ExternalID)
	d.Set("createdat", accountInfo.CreatedAt)
	d.Set("updatedat", accountInfo.UpdatedAt)
	d.Set("creatorid", accountInfo.CreatorID)

	// always run
	d.SetId(azureAccountId)

	return diags
}
