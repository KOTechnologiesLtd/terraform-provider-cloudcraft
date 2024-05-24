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