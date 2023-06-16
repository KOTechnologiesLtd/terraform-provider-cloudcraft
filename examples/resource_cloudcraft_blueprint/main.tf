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

resource "cloudcraft_blueprint" "blueprint" {
    name = "blueprint"
    grid = "standard"
}
