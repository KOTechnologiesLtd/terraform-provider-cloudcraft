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

resource "cloudcraft_blueprint" "blueprint" {
    name = "blueprint"
    grid = "standard"
}
