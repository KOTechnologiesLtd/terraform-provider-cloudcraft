terraform {
  required_providers {
    cloudcraft = {
      source  = "github.com/kotechnologiesltd/cloudcraft"
      version = "2.1.1"
    }
  }
}

provider "cloudcraft" {
  apikey = ""
}

data "cloudcraft_integration_aws" "integratedawsaccount" {
  id = ""
}

output "account_name" {
  value = data.cloudcraft_integration_aws.integratedawsaccount.name
}

output "accountInfo" {
  value = data.cloudcraft_integration_aws.integratedawsaccount
}