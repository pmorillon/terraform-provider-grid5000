# Terraform Firewall exemple
#

# Terraform dependencies configuration
#
terraform {
  required_version = ">= 0.13.0"
  required_providers {
    grid5000 = {
      source  = "pmorillon/grid5000"
      version = "~> 0.0.9"
    }
  }
}

# Grid5000 provider configuration
#
provider "grid5000" {
  restfully_file = "/Users/pmorillon/.restfully/api.grid5000.fr.yml"
}

# OAR resources reservation
#
resource "grid5000_job" "firewall" {
  name      = "Terraform Firewall"
  site      = "rennes"
  command   = "sleep 1d"
  resources = "/nodes=1"
  types     = ["deploy"]
}

# Kadeploy bare-metal deployment
#
resource "grid5000_deployment" "debian" {
  site        = "rennes"
  environment = "debian10-x64-base"
  nodes       = grid5000_job.firewall.assigned_nodes
  key         = file("~/.ssh/id_rsa.pub")
}

# nginx install
#
resource "null_resource" "nginx_install" {
  count = 1

  connection {
    host        = element(sort(grid5000_deployment.debian.nodes), count.index)
    type        = "ssh"
    user        = "root"
    bastion_host  = "access.grid5000.fr"
    bastion_user  = "pmorillo"
  }

  provisioner "remote-exec" {
    inline = [
      "apt install -y nginx >/dev/null 2>&1",
      "dhclient -6 eno1 >/dev/null 2>&1",
    ]
  }
}

data "grid5000_ipv6_nodelist" "ipv6list" {
  nodelist = grid5000_job.firewall.assigned_nodes 
}

resource "grid5000_firewall" "f1" {
    depends_on = [
      grid5000_deployment.debian
    ]

    site = "rennes"
    job_id = grid5000_job.firewall.id
    address = data.grid5000_ipv6_nodelist.ipv6list.result
    port = 80
}