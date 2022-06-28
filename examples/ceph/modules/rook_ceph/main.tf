# Terraform Ceph module
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
    rke = {
      source  = "rancher/rke"
      version = "~> 1.3.0"
    }
  }
}

# OAR resources reservation
#
resource "grid5000_job" "k8s" {
  name      = "Terraform Rook Ceph"
  site      = var.site
  command   = "sleep 1d"
  resources = "{(type='disk' OR type='default') AND disk_reservation_count > 0 AND cluster='${var.cluster_selector}'}/nodes=${var.nodes_count},walltime=${var.walltime}"
  types     = ["deploy"]
}

# Kadeploy bare-metal deployment
#
resource "grid5000_deployment" "k8s" {
  site        = var.site
  environment = "debian10-x64-base"
  nodes       = grid5000_job.k8s.assigned_nodes
  key         = file("~/.ssh/id_rsa.pub")
}

# Docker installation and local registry configuration
#
resource "null_resource" "docker_install" {
  depends_on = [grid5000_deployment.k8s]

  count = var.nodes_count

  connection {
    host        = element(sort(grid5000_deployment.k8s.nodes), count.index)
    type        = "ssh"
    user        = "root"
    private_key = file("~/.ssh/id_rsa")
  }

  provisioner "file" {
    source      = "${path.module}/files/install-docker.sh"
    destination = "/root/install-docker.sh"
  }

  provisioner "remote-exec" {
    inline = [
      "sh /root/install-docker.sh >/dev/null 2>&1",
    ]
  }
}

# Kubernetes deployment with Rancher Kubernetes Engine
#
resource "rke_cluster" "cluster" {
  depends_on = [null_resource.docker_install]

  dynamic "nodes" {
    for_each = [for s in range(var.nodes_count) : {
      address          = sort(grid5000_deployment.k8s.nodes)[s]
      internal_address = sort(grid5000_deployment.k8s.nodes)[s]
      role             = s > 0 ? ["worker"] : ["controlplane", "etcd"]
    }]

    content {
      address          = nodes.value.address
      internal_address = nodes.value.internal_address
      user             = "root"
      role             = nodes.value.role
      ssh_key          = file("~/.ssh/id_rsa")
    }
  }
}

# Generate Kubernetes access config file
#
resource "local_file" "kube_cluster_yaml" {
  filename = "kube_config_cluster.yml"
  content  = rke_cluster.cluster.kube_config_yaml
}

# Kubernetes Terraform provider configuration
#
provider "kubernetes" {
  host                    = rke_cluster.cluster.api_server_url
  username                = rke_cluster.cluster.kube_admin_user

  client_certificate      = rke_cluster.cluster.client_cert
  client_key              = rke_cluster.cluster.client_key
  cluster_ca_certificate  = rke_cluster.cluster.ca_crt
}

# Helm Terraform provider configuration
#
provider "helm" {
  kubernetes {
    host                   = rke_cluster.cluster.api_server_url
    username               = rke_cluster.cluster.kube_admin_user
    client_certificate     = rke_cluster.cluster.client_cert
    client_key             = rke_cluster.cluster.client_key
    cluster_ca_certificate = rke_cluster.cluster.ca_crt
  }
}

# Create rook-ceph Kubernetes namespace
#
resource "kubernetes_namespace" "rook-ceph" {
  depends_on = [ local_file.kube_cluster_yaml ]
  metadata {
    name = "rook-ceph"
  }
  provisioner "local-exec" {
    when = destroy
    command = "sleep 10"
  }
}

# Rook Ceph Kubernetes operator deployment with Helm
#
resource "helm_release" "rook-ceph" {
  depends_on  = [ kubernetes_namespace.rook-ceph ]

  name        = "rook-ceph"
  repository  = "https://charts.rook.io/release"
  chart       = "rook-ceph"
  namespace   = "rook-ceph"
  version     = var.rook_ceph_version
}

# Prepare reserved disks for Ceph cluster deployment
#
resource "null_resource" "format_disks" {
  depends_on = [null_resource.docker_install]
  count = var.nodes_count

  connection {
    host          = element(sort(grid5000_deployment.k8s.nodes), count.index) 
    type          = "ssh"
    user          = "root"
    private_key   = file("~/.ssh/id_rsa")
  }

  provisioner "file" {
    content = templatefile("${path.module}/files/disk-format.sh.tmpl",
   { disks = [for d in grid5000_job.k8s.disks_resources : d["device"] if d["hostname"] == element(sort(grid5000_deployment.k8s.nodes), count.index)] })
    destination = "/root/format-disks.sh"
  }

  provisioner "remote-exec" {
    inline = [
      "sh /root/format-disks.sh >/dev/null 2>&1",
    ]
  }
}

# Get Grid5000 Reference repository information for the first worker node
# Used to get the IP address in order to generate nip.io hostnames for Kubernetes Ingresses
#
data "grid5000_node" "worker" {
  site = var.site
  name = element(sort(grid5000_deployment.k8s.nodes), 1)
}

# Deploy Ceph cluster including a ceph filesystem, 2 rados block devices pool (hdd and ssd), ceph dashboard, ...
#
resource "helm_release" "rook-ceph-config" {
  depends_on = [ helm_release.rook-ceph ]

  name  = "rook-ceph-config"
  namespace = "rook-ceph"
  chart = "${path.module}/charts/rook-ceph-config"
  wait  = false

  values = [
    <<EOT
context: staging
rook:
  version: ${helm_release.rook-ceph.version}
ceph:
  version: ${var.ceph_version}
storage:
  config:
    metadataDevice: "${var.ceph_metadata_device}"
  nodes:
%{ for h in distinct([for d in grid5000_job.k8s.disks_resources : d["hostname"]]) ~}
%{ if !contains([ for i in rke_cluster.cluster.control_plane_hosts : i["address" ] ], h) ~}
  - name: ${h}
    devices:
%{ for d in grid5000_job.k8s.disks_resources ~}
%{ if d["hostname"] == h && d.device != var.ceph_metadata_device ~}
    - name: /dev/${d.device}
%{ endif ~}
%{ endfor ~}
%{ endif ~}
%{ endfor ~}
EOT
  ]

  set {
    name = "dashboard.hostname"
    value = "ceph.${data.grid5000_node.worker.ip}.nip.io"
  }

}

# Create monitoring Kubernetes namespace
#
resource "kubernetes_namespace" "monitoring" {
  depends_on = [ local_file.kube_cluster_yaml ]
  metadata {
    name = "monitoring"
  }
  provisioner "local-exec" {
    when = destroy
    command = "sleep 10"
  }
}

# Deploy kube-prometheus-stack with Helm (Kubernetes Prometheus operator, Prometheus server, Grafana, Alertmanager, ...)
#
resource "helm_release" "prometheus" {
  depends_on = [ kubernetes_namespace.monitoring ]

  name       = "prom-op"
  repository = "https://prometheus-community.github.io/helm-charts"
  chart      = "kube-prometheus-stack"
  namespace  = "monitoring"
  version    = "13.2.1"
  wait       = false

  values = [
      <<EOT
commonLabels:
  team: rook
prometheus:
  prometheusSpec:
    serviceMonitorSelector:
      matchLabels:
        team: rook
  ingress:
    enabled: true
    hosts:
      - prometheus.${data.grid5000_node.worker.ip}.nip.io
    tls:
      - secretName: prometheus-tls
        hosts:
          - prometheus.${data.grid5000_node.worker.ip}.nip.io
alertmanager:
  ingress:
    enabled: true
    hosts:
      - alertmanager.${data.grid5000_node.worker.ip}.nip.io
    tls:
      - secretName: alertmanager-tls
        hosts:
          - alertmanager.${data.grid5000_node.worker.ip}.nip.io
grafana:
  adminPassword: admin
  sidecar:
    dashboards:
      enabled: true
  ingress:
    enabled: true
    hosts:
      - grafana.${data.grid5000_node.worker.ip}.nip.io
    tls:
      - secretName: grafana-tls
        hosts:
          - grafana.${data.grid5000_node.worker.ip}.nip.io
EOT
  ]
}

# Add Prometheus servicemonitors for Ceph and CSI drivers
#
resource "helm_release" "monitoring-config" {
  depends_on = [ helm_release.prometheus ]

  name  = "monitoring-config"
  namespace = "rook-ceph"
  chart = "${path.module}/charts/monitoring-config"
}