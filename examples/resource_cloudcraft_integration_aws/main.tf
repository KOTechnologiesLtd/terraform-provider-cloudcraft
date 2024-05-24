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

resource "cloudcraft_integration_aws" "cloudcraftintegrationaws" {
  name    = ""
  rolearn = ""
  //read_access = ["team/GUID"]
  //write_access = ["team/GUID"]
}

output "ccawsaccountid" {
  value = cloudcraft_integration_aws.cloudcraftintegrationaws.id
}

output "ccawsaccountextid" {
  value = cloudcraft_integration_aws.cloudcraftintegrationaws.externalid
}
