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
  apikey = "/kiR68dHWzu/ZUB32cHRSMXWDnfTD/SFNsMS1ltApjw="
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
