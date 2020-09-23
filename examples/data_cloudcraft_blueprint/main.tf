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

data "cloudcraft_blueprint" "blueprint" {
  id = "BLUEPRINTID"
}

output "blueprint_name" {
  value = data.cloudcraft_blueprint.blueprint.name
}

output "blueprint_grid" {
  value = data.cloudcraft_blueprint.blueprint.grid
}