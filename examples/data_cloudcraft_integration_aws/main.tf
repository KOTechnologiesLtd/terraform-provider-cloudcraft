#terraform 13 syntax
#terraform {
#  required_providers {
#    cloudcraft = {
#      source  = "github.com/kotechnologiesltd/cloudcraft"
#      versions = ["0.1"]
#    }
#  }
#}

provider "cloudcraft" {
    apikey = "APIKEY"
  }

data "cloudcraft_integration_aws" "integratedawsaccount" {
  id = "INTEGRATIONID"
}

output "account_name" {
  value = data.cloudcraft_integration_aws.integratedawsaccount.name
}

output "accountInfo" {
  value = data.cloudcraft_integration_aws.integratedawsaccount
}