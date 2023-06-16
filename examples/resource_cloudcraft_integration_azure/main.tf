terraform {
 required_providers {
   cloudcraft = {
     source  = "github.com/kotechnologiesltd/cloudcraft"
     version = "2.0.0"
   }
 }
}

provider "cloudcraft" {
    apikey = ""
}

resource "cloudcraft_integration_azure" "cloudcraftintegrationazure" {
  id = ""
  applicationid = ""
  directoryid = ""
  subscriptionid = ""
  clientsecret = ""
}

output "ccazureaccountid" {
  value = cloudcraft_integration_azure.cloudcraftintegrationazure.id
}

output "ccazureaccountextid" {
  value = cloudcraft_integration_azure.cloudcraftintegrationazure.externalid
}
