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

resource "cloudcraft_blueprint" "blueprint" {
  name = "blueprint"
  grid = "standard"
}
