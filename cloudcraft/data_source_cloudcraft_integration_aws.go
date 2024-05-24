package cloudcraft

import (
	"context"
	"reflect"

	"github.com/KOTechnologiesLtd/go-cloudcraft-api/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIntegrationAws() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIntegrationAwsRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"rolearn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"externalid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"read_access": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"write_access": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

func dataSourceIntegrationAwsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*cloudcraft.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	awsAccountId := d.Get("id").(string)

	accountInfo, err := client.AccountIntegrationAws(awsAccountId)
	{
		if err != nil {
			return diag.FromErr(err)
		}
	}
	//log.Printf("accounts%v ", accountInfo)
	if reflect.DeepEqual(cloudcraft.AccountIntegrationAws{}, accountInfo) {
		diags := append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "ID Not Found",
			Detail:   "The CloudCraft AWS Integration ID does not exist: " + awsAccountId,
		})
		return diags
	}

	//populate account scheme
	d.Set("name", accountInfo.Name)
	d.Set("rolearn", accountInfo.RoleArn)
	d.Set("externalid", accountInfo.ExternalID)
	d.Set("read_access", accountInfo.ReadAccess)
	d.Set("write_access", accountInfo.WriteAccess)
	d.Set("createdat", accountInfo.CreatedAt)
	d.Set("updatedat", accountInfo.UpdatedAt)
	d.Set("creatorid", accountInfo.CreatorID)

	// always run
	d.SetId(awsAccountId)

	return diags
}
