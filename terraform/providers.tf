terraform {
  required_providers {
    twc = {
      source = "tf/timeweb-cloud"
      version = "~> 1.0.0"
    }
  }
}

provider "twc" {
  token = var.timeweb_token
}
