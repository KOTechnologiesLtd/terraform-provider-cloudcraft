terraform {
  required_providers {
    cloudcraft = {
      source = "github.com/KOTechnologiesLtd/cloudcraft"
      #source = "KOTechnologiesLtd/cloudcraft"
      version = "2.1.3"
    }
  }
}
provider "cloudcraft" {
  apikey = ""
}

data "cloudcraft_integration_aws_iamparams" "params" {
}

output "account_id" {
  value = data.cloudcraft_integration_aws_iamparams.params.accountid
}

output "extid" {
    value = data.cloudcraft_integration_aws_iamparams.params.externalid
}

output "uri" {
    value = data.cloudcraft_integration_aws_iamparams.params.awsconsoleurl
}
