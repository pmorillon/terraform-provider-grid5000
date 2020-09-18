# Need terraform 0.13
terraform {
  required_providers {
    rke = {
      source = "rancher/rke"
      version = "~> 1.0.1"
    }
    grid5000 = {
      source = "pmorillon/grid5000"
      version = "~> 0.0.4"
    }
  }
}