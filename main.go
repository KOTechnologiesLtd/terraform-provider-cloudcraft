package main

import (
	"terraform-provider-cloudcraft/cloudcraft"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: cloudcraft.Provider})
}
