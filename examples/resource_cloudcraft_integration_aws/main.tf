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

resource "cloudcraft_integration_aws" "cloudcraftintegrationaws" {
    name = "INTEGRATIONNAME"
    rolearn = "YOURAWSROLEARN"
}

output "ccawsaccountid" {
  value = cloudcraft_integration_aws.cloudcraftintegrationaws.id
}

output "ccawsaccountextid" {
  value = cloudcraft_integration_aws.cloudcraftintegrationaws.externalid
}
