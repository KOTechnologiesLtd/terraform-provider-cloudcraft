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

resource "cloudcraft_integration_aws" "cloudcraftintegrationaws" {
  name    = ""
  rolearn = ""
}

output "ccawsaccountid" {
  value = cloudcraft_integration_aws.cloudcraftintegrationaws.id
}

output "ccawsaccountextid" {
  value = cloudcraft_integration_aws.cloudcraftintegrationaws.externalid
}
