package cloudcraft

import (
	"context"

	"github.com/KOTechnologiesLtd/go-cloudcraft-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceBlueprint() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBlueprintRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"lastuserid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"creatorid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"updatedat": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"createdat": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"grid": {
				Type:     schema.TypeString,
				Computed: true,
				//ValidateFunc: validateGridType,
			},
			"version": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceBlueprintRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*cloudcraft.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	bpInfo := cloudcraft.BluePrint{}

	bpID := d.Get("id").(string)

	bpInfo.ID = &bpID

	err := client.BluePrintGet(&bpInfo)
	{
		if err != nil {
			return diag.FromErr(err)
		}
	}
	//log.Printf("accounts%v ", accountInfo)
	if (cloudcraft.BluePrint{}) == bpInfo {
		diags := append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "ID Not Found",
			Detail:   "The CloudCraft Blueprint ID does not exist: " + bpID,
		})
		return diags
	}

	//populate account scheme
	d.Set("name", bpInfo.Data.Name)
	d.Set("grid", bpInfo.Data.Grid)
	d.Set("version", bpInfo.Data.Version)
	d.Set("creatorid", bpInfo.CreatorID)
	d.Set("createdat", bpInfo.CreatedAt)
	d.Set("updatedat", bpInfo.UpdatedAt)
	d.Set("lastuserid", bpInfo.LastUserID)

	// always run
	d.SetId(*bpInfo.ID)

	return diags
}
