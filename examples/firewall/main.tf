# Terraform Firewall exemple
#

# Terraform dependencies configuration
#
terraform {
  required_version = ">= 0.13.0"
  required_providers {
    grid5000 = {
      source  = "pmorillon/grid5000"
      version = "~> 0.0.8"
    }
  }
}

# Locals
#
locals {
  site = "rennes"
}

# OAR resources reservation
#
resource "grid5000_job" "job1" {
  site      = local.site
  command   = "sleep 1d"
  resources = "/nodes=1,walltime=1"
  types     = ["deploy"]
}

# Kadeploy bare-metal deployment
#
resource "grid5000_deployment" "debian" {
  site        = local.site
  environment = "debian10-x64-base"
  nodes       = grid5000_job.job1.assigned_nodes
  key         = file("~/.ssh/id_rsa.pub")
}

data "grid5000_node" "node" {
  name = element(sort(grid5000_job.job1.assigned_nodes), 0)
  site = local.site
}

# nginx install
#
resource "null_resource" "nginx_install" {
  count = 1

  connection {
    host        = element(sort(grid5000_deployment.debian.nodes), count.index)
    type        = "ssh"
    user        = "root"
  }

  provisioner "remote-exec" {
    inline = [
      "apt install -y nginx >/dev/null 2>&1",
      "dhclient -6 ${data.grid5000_node.node.primary_network_interface} >/dev/null 2>&1",
    ]
  }
}

data "grid5000_ipv6_nodelist" "ipv6list" {
  nodelist = grid5000_job.job1.assigned_nodes 
}

resource "grid5000_firewall" "f1" {
    depends_on = [
      null_resource.nginx_install
    ]

    site = local.site
    job_id = grid5000_job.job1.id

    rule {
      dest = data.grid5000_ipv6_nodelist.ipv6list.result
      ports = [80]
    }

    rule {
      dest = data.grid5000_ipv6_nodelist.ipv6list.result
      ports = [443]
    }
    
}