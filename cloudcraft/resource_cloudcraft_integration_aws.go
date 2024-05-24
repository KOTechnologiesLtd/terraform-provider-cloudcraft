package cloudcraft

import (
	"context"
	"fmt"

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

func resourceIntegrationAwsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	c := m.(*cloudcraft.Client)

	//Get the data from the schema
	name := d.Get("name").(string)
	roleArn := d.Get("rolearn").(string)
	readAccess, err := getStringListFromResourceData(d, "read_access")
	if err != nil {
		return err
	}
	writeAccess, err := getStringListFromResourceData(d, "write_access")
	if err != nil {
		return err
	}

	//Set the request from the data
	accCreate := cloudcraft.AccountIntegrationAws{}
	accCreate.Name = &name
	accCreate.RoleArn = &roleArn
	accCreate.ReadAccess = &readAccess
	accCreate.WriteAccess = &writeAccess

	//Request the action
	if err := c.AccountIntegrationAwsCreate(&accCreate); err != nil {
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
	d.Set("read_access", accountInfo.ReadAccess)
	d.Set("write_access", accountInfo.WriteAccess)
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

	if d.HasChange("name") || d.HasChange("rolearn") || d.HasChange("read_access") ||
		d.HasChange("write_access") {
		newName := d.Get("name").(string)
		//log.Printf("name update%s", newName)
		updates.Name = &newName
		newRoleArn := d.Get("rolearn").(string)
		//log.Printf("rolearn update%s", newRoleArn)
		updates.RoleArn = &newRoleArn

		newReadAccess, err := getStringListFromResourceData(d, "read_access")
		if err != nil {
			return err
		}
		updates.ReadAccess = &newReadAccess

		newWriteAccess, err := getStringListFromResourceData(d, "write_access")
		if err != nil {
			return err
		}
		updates.WriteAccess = &newWriteAccess
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

func getStringListFromResourceData(d *schema.ResourceData, key string) ([]string, diag.Diagnostics) {
	rawData, ok := d.Get(key).([]interface{})
	if !ok {
		return nil, diag.Diagnostics{{
			Severity: diag.Error,
			Summary:  "Type mismatch",
			Detail:   fmt.Sprintf("Expected type of %s to be []interface{}, got %T", key, d.Get(key)),
		}}
	}

	stringList := make([]string, len(rawData))
	for i, rawValue := range rawData {
		strValue, ok := rawValue.(string)
		if !ok {
			return nil, diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Type mismatch in list",
				Detail:   fmt.Sprintf("Expected all items in %s to be strings, got %T", key, rawValue),
			}}
		}
		stringList[i] = strValue
	}

	return stringList, nil
}
