terraform {
  required_providers {
    cloudcraft = {
      source  = "github.com/kotechnologiesltd/cloudcraft"
      version = "2.1.0"
    }
  }
}

provider "cloudcraft" {
  apikey = ""
}

data "cloudcraft_integration_azure" "integratedazureaccount" {
  id = ""
}

output "account_name" {
  value = data.cloudcraft_integration_azure.integratedazureaccount.name
}

output "accountInfo" {
  value = data.cloudcraft_integration_azure.integratedazureaccount
}