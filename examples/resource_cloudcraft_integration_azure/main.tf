terraform {
  required_providers {
    cloudcraft = {
      source  = "github.com/kotechnologiesltd/cloudcraft"
      version = "2.1.2"
    }
  }
}

provider "cloudcraft" {
  apikey = ""
}

resource "cloudcraft_integration_azure" "cloudcraftintegrationazure" {
  name           = ""
  applicationid  = ""
  directoryid    = ""
  subscriptionid = ""
  clientsecret   = ""
}

output "ccazureaccountid" {
  value = cloudcraft_integration_azure.cloudcraftintegrationazure.id
}