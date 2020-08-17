locals {
  site        = "rennes"
  nodes_count = 4
}

resource "grid5000_job" "my_job" {
  name      = "Terraform RKE"
  site      = local.site
  command   = "sleep 1h"
  resources = "/nodes=${local.nodes_count}"
  types     = ["deploy"]
}

resource "grid5000_deployment" "my_deployment" {
  site        = local.site
  environment = "debian10-x64-base"
  nodes       = grid5000_job.my_job.assigned_nodes
  key         = file("~/.ssh/id_rsa.pub")
}