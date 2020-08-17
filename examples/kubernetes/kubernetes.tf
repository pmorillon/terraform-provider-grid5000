resource "null_resource" "docker_install" {
  depends_on = [grid5000_deployment.my_deployment]

  count = local.nodes_count 

  connection {
    host          = element(sort(grid5000_deployment.my_deployment.nodes), count.index) 
    type          = "ssh"
    user          = "root"
    private_key   = file("~/.ssh/id_rsa")
  }

  provisioner "file" {
    source = "install-docker.sh"
    destination = "/root/install-docker.sh"
  }

  provisioner "remote-exec" {
    inline = [
      "sh /root/install-docker.sh >/dev/null 2>&1",
    ]
  }
}

resource "rke_cluster" "cluster" {
  depends_on = [null_resource.docker_install]

  dynamic "nodes" {
    for_each = [for s in range(local.nodes_count): {
      address = sort(grid5000_deployment.my_deployment.nodes)[s] 
      internal_address = sort(grid5000_deployment.my_deployment.nodes)[s]
      role = s > 0 ? ["worker"] : ["controlplane", "etcd"]
    }]

    content {
      address = nodes.value.address
      internal_address = nodes.value.internal_address
      user = "root"
      role = nodes.value.role
      ssh_key = file("~/.ssh/id_rsa") 
    }
  }
}

provider "kubernetes" {
  host     = rke_cluster.cluster.api_server_url
  username = rke_cluster.cluster.kube_admin_user

  client_certificate     = rke_cluster.cluster.client_cert
  client_key             = rke_cluster.cluster.client_key
  cluster_ca_certificate = rke_cluster.cluster.ca_crt
}

resource "kubernetes_namespace" "monitoring" {
  metadata {
    name = "monitoring"
  }
}
