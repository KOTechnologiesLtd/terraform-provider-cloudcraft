package cloudcraft

import (
	"context"
	"log"

	"github.com/KOTechnologiesLtd/go-cloudcraft-api/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIntegrationAwsIamParams() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIntegrationAwsIamParamsRead,
		Schema: map[string]*schema.Schema{
			"accountid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"externalid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"awsconsoleurl": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIntegrationAwsIamParamsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*cloudcraft.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	accountInfo, err := client.AccountIntegrationAwsIamParams()
	{
		if err != nil {
			return diag.FromErr(err)
		}
	}
	//log.Printf("accounts%v ", accountInfo)
	if (cloudcraft.AccountIntegrationAwsIamParams{}) == accountInfo {
		diags := append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "ID Not Found",
			Detail:   "The CloudCraft AWS Integration ID does not exist: ",
		})
		return diags
	}

	//populate account scheme
	log.Printf("HERE accounts%s ", *accountInfo.AccountID)
	log.Printf("HERE ExternalID%s ", *accountInfo.ExternalID)
	log.Printf("HERE awsconsoleurl%s ", *accountInfo.AwsConsoleUrl)
	d.Set("accountid", accountInfo.AccountID)
	d.Set("externalid", accountInfo.ExternalID)
	d.Set("awsconsoleurl", accountInfo.AwsConsoleUrl)

	// always run
	d.SetId(*accountInfo.AccountID)

	return diags
}
