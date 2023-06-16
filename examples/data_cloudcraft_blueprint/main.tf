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

data "cloudcraft_blueprint" "blueprint" {
  id = ""
}

output "blueprint_name" {
  value = data.cloudcraft_blueprint.blueprint.name
}

output "blueprint_grid" {
  value = data.cloudcraft_blueprint.blueprint.grid
}