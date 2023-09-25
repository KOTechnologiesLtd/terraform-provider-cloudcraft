package cloudcraft

import (
	"context"

	"github.com/KOTechnologiesLtd/go-cloudcraft-api/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIntegrationAws() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIntegrationAwsCreate,
		ReadContext:   resourceIntegrationAwsRead,
		UpdateContext: resourceIntegrationAwsUpdate,
		DeleteContext: resourceIntegrationAwsDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"rolearn": &schema.Schema{
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

func resourceIntegrationAwsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	c := m.(*cloudcraft.Client)

	//Get the data from the schema
	name := d.Get("name").(string)
	roleArn := d.Get("rolearn").(string)

	//Set the request from the data
	accCreate := cloudcraft.AccountIntegrationAws{}
	accCreate.Name = &name
	accCreate.RoleArn = &roleArn

	//Request the action
	err := c.AccountIntegrationAwsCreate(&accCreate)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*accCreate.ID)
	resourceIntegrationAwsRead(ctx, d, m)

	return diags
}

func resourceIntegrationAwsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	client := m.(*cloudcraft.Client)
	accountID := d.Id()

	accountInfo, err := client.AccountIntegrationAws(accountID)
	{
		if err != nil {
			return diag.FromErr(err)
		}
	}

	//populate account scheme
	d.Set("externalid", accountInfo.ExternalID)
	d.Set("createdat", accountInfo.CreatedAt)
	d.Set("updatedat", accountInfo.UpdatedAt)
	d.Set("creatorid", accountInfo.CreatorID)
	d.Set("rolearn", accountInfo.RoleArn)
	d.Set("name", accountInfo.Name)

	return diags
}

func resourceIntegrationAwsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	client := m.(*cloudcraft.Client)
	updates := cloudcraft.AccountIntegrationAws{}
	accountID := d.Id()
	updates.ID = &accountID

	if d.HasChange("name") || d.HasChange("rolearn") {
		newName := d.Get("name").(string)
		//log.Printf("name update%s", newName)
		updates.Name = &newName
		newRoleArn := d.Get("rolearn").(string)
		//log.Printf("rolearn update%s", newRoleArn)
		updates.RoleArn = &newRoleArn
	}

	err := client.AccountIntegrationAwsUpdate(&updates)
	if err != nil {
		return diag.FromErr(err)
	}
	resourceIntegrationAwsRead(ctx, d, m)

	return diags
}

func resourceIntegrationAwsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	client := m.(*cloudcraft.Client)
	accountID := d.Id()

	accountInfo, err := client.AccountIntegrationAws(accountID)
	err = client.AccountIntegrationAwsDelete(&accountInfo)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
