package cloudcraft

import (
	"context"

	"github.com/KOTechnologiesLtd/go-cloudcraft-api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIntegrationAzure() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIntegrationAzureCreate,
		ReadContext:   resourceIntegrationAzureRead,
		UpdateContext: resourceIntegrationAzureUpdate,
		DeleteContext: resourceIntegrationAzureDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"applicationid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"directoryid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"subscriptionid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"clientsecret": &schema.Schema{
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

func resourceIntegrationAzureCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	c := m.(*cloudcraft.Client)

	//Get the data from the schema
	name := d.Get("name").(string)
	applicationId := d.Get("applicationid").(string)
	directoryId := d.Get("directoryid").(string)
	subscriptionId := d.Get("subscriptionid").(string)
	clientSecret := d.Get("clientsecret").(string)

	//Set the request from the data
	accCreate := cloudcraft.AccountIntegrationAzure{}
	accCreate.Name = &name
	accCreate.ApplicationId = &applicationId
	accCreate.DirectoryId = &directoryId
	accCreate.SubscriptionId = &subscriptionId
	accCreate.ClientSecret = &clientSecret

	//Request the action
	err := c.AccountIntegrationAzureCreate(&accCreate)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*accCreate.ID)
	resourceIntegrationAzureRead(ctx, d, m)

	return diags
}

func resourceIntegrationAzureRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	client := m.(*cloudcraft.Client)
	accountID := d.Id()

	accountInfo, err := client.AccountIntegrationAzure(accountID)
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
	d.Set("applicationid", accountInfo.ApplicationId)
	d.Set("directoryid", accountInfo.DirectoryId)
	d.Set("subscriptionid", accountInfo.SubscriptionId)
	d.Set("clientsecret", accountInfo.ClientSecret)
	d.Set("name", accountInfo.Name)

	return diags
}

func resourceIntegrationAzureUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	client := m.(*cloudcraft.Client)
	updates := cloudcraft.AccountIntegrationAzure{}
	accountID := d.Id()
	updates.ID = &accountID

	if d.HasChange("name") || d.HasChange("applicationid") || d.HasChange("directoryid") || d.HasChange("subscriptionid") || d.HasChange("clientsecret") {
		newName := d.Get("name").(string)
		//log.Printf("name update%s", newName)
		updates.Name = &newName

		newApplicationId := d.Get("applicationsd").(string)
		newDirectoryId := d.Get("directorysd").(string)
		newSubscriptionId := d.Get("subscriptionsd").(string)
		newClientSecret := d.Get("clientsecret").(string)

		//log.Printf("azure update%s", newApplicationId, newDirectoryId, newSubscriptionId, newClientSecret)
		updates.ApplicationId = &newApplicationId
		updates.DirectoryId = &newDirectoryId
		updates.SubscriptionId = &newSubscriptionId
		updates.ClientSecret = &newClientSecret
	}

	err := client.AccountIntegrationAzureUpdate(&updates)
	if err != nil {
		return diag.FromErr(err)
	}
	resourceIntegrationAzureRead(ctx, d, m)

	return diags
}

func resourceIntegrationAzureDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	client := m.(*cloudcraft.Client)
	accountID := d.Id()

	accountInfo, err := client.AccountIntegrationAzure(accountID)
	err = client.AccountIntegrationAzureDelete(&accountInfo)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
