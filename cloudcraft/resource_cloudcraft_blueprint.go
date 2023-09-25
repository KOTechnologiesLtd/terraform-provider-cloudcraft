package cloudcraft

import (
	"context"

	"github.com/KOTechnologiesLtd/go-cloudcraft-api/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBlueprint() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBlueprintCreate,
		ReadContext:   resourceBlueprintRead,
		UpdateContext: resourceBlueprintUpdate,
		DeleteContext: resourceBlueprintDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
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
				Required: true,
			},
			"grid": {
				Type:     schema.TypeString,
				Required: true,
				//ValidateFunc: validateGridType,
			},
			"version": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceBlueprintCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	c := m.(*cloudcraft.Client)

	//Set the request from the data
	dataname := d.Get("name").(string)
	datagrid := d.Get("grid").(string)

	bpInfo := cloudcraft.BluePrint{}

	bpData := cloudcraft.DataDetails{}
	bpData.Name = &dataname
	bpData.Grid = &datagrid

	bpInfo.Data = &bpData

	err := c.BluePrintCreate(&bpInfo)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*bpInfo.ID)
	resourceBlueprintRead(ctx, d, m)

	return diags
}

func resourceBlueprintRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	client := m.(*cloudcraft.Client)
	bpInfo := cloudcraft.BluePrint{}
	bpID := d.Id()
	bpInfo.ID = &bpID

	err := client.BluePrintGet(&bpInfo)
	{
		if err != nil {
			return diag.FromErr(err)
		}
	}

	//populate blueprint scheme
	d.Set("lastuserid", bpInfo.LastUserID)
	d.Set("createdat", bpInfo.CreatedAt)
	d.Set("updatedat", bpInfo.UpdatedAt)
	d.Set("creatorid", bpInfo.CreatorID)
	d.Set("version", bpInfo.Data.Version)
	d.Set("name", bpInfo.Data.Name)
	d.Set("grid", bpInfo.Data.Grid)

	return diags
}

func resourceBlueprintUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	client := m.(*cloudcraft.Client)
	updates := cloudcraft.BluePrint{}
	bpID := d.Id()
	updates.ID = &bpID

	updatedbpData := cloudcraft.DataDetails{}

	if d.HasChange("name") || d.HasChange("grid") {
		newName := d.Get("name").(string)
		updatedbpData.Name = &newName
		newGrid := d.Get("grid").(string)
		updatedbpData.Grid = &newGrid
		updates.Data = &updatedbpData
	}

	err := client.BluePrintUpdate(&updates)
	if err != nil {
		return diag.FromErr(err)
	}
	resourceBlueprintRead(ctx, d, m)

	return diags
}

func resourceBlueprintDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	client := m.(*cloudcraft.Client)
	deletebp := cloudcraft.BluePrint{}
	bpID := d.Id()
	deletebp.ID = &bpID

	err := client.BluePrintDelete(&deletebp)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
