terraform {
  required_version = ">= 0.13.0"
  required_providers {
    grid5000 = {
      source = "pmorillon/grid5000"
      version = "0.0.7"
    } # Used for Ceph pool management
  }
}

# Kubernetes cluster deployment
#
module "k8s-cluster" {
  source  = "pmorillon/k8s-cluster/grid5000"
  version = "0.0.4"

  walltime        = var.walltime
  site            = var.nodes_location
  nodes_count     = var.nodes_count
  nodes_selector  = var.nodes_selector
  oar_job_name    = "openwhisk-tf"
}

# Configure Kubernetes provider
#
provider "kubernetes" {
  host                    = module.k8s-cluster.api_server_url
  username                = module.k8s-cluster.kube_admin_user
  client_certificate      = module.k8s-cluster.client_cert
  client_key              = module.k8s-cluster.client_key
  cluster_ca_certificate  = module.k8s-cluster.ca_crt
}

# Configure Helm provider
#
provider "helm" {
  kubernetes {
    host                    = module.k8s-cluster.api_server_url
    username                = module.k8s-cluster.kube_admin_user
    client_certificate      = module.k8s-cluster.client_cert
    client_key              = module.k8s-cluster.client_key
    cluster_ca_certificate  = module.k8s-cluster.ca_crt
  }
}

# Create ceph-csi-rbd kubernetes namespace
#
resource "kubernetes_namespace" "ceph-csi-rbd" {
  depends_on = [ module.k8s-cluster ]
  metadata {
    name = "ceph-csi-rbd"
  }
}

# Deploy Ceph CSI RBD driver
#
resource "helm_release" "ceph-csi-rbd" {
  depends_on = [ kubernetes_namespace.ceph-csi-rbd ]
  name = "ceph-sci-rbd"
  namespace = "ceph-csi-rbd"
  repository = "https://ceph.github.io/csi-charts"
  chart = "ceph-csi-rbd"
  version = "v3.2.1"

  values = [
    <<EOT
csiConfig:
- clusterID: g5k-ceph
  monitors:
  - ceph0.${var.data_location}.grid5000.fr
  - ceph1.${var.data_location}.grid5000.fr
  - ceph2.${var.data_location}.grid5000.fr
EOT
  ]
}

# Get Ceph auth key
#
data "grid5000_ceph_auth" "k8s" {
  site = var.data_location
}

# Create Ceph pool
#
resource "grid5000_ceph_pool" "k8s" {
  site = var.data_location
  pool_name = "k8s"
  quota = var.ceph_pool_quota
}

resource "kubernetes_secret" "csi-rbd-secret" {
  metadata {
    name = "csi-rbd-secret"
    namespace = "ceph-csi-rbd"
  }
  data = {
    "userID" = var.username
    "userKey" = data.grid5000_ceph_auth.k8s.key
    "encryptionPassphrase" = "test_passphrase"
  }
}

resource "kubernetes_storage_class" "csi-rbd-sc" {
  metadata {
    name = "csid-rbd-sc"
    annotations = {
      "storageclass.kubernetes.io/is-default-class" = "true"
    }
  }
  storage_provisioner = "rbd.csi.ceph.com"
  parameters = {
    "clusterID" = "g5k-ceph"
    "pool" = grid5000_ceph_pool.k8s.real_pool_name
    "imageFeatures" = "layering"
    "csi.storage.k8s.io/provisioner-secret-name" = "csi-rbd-secret"
    "csi.storage.k8s.io/provisioner-secret-namespace" = "ceph-csi-rbd"
    "csi.storage.k8s.io/controller-expand-secret-name" = "csi-rbd-secret"
    "csi.storage.k8s.io/controller-expand-secret-namespace" = "ceph-csi-rbd"
    "csi.storage.k8s.io/node-stage-secret-name" = "csi-rbd-secret"
    "csi.storage.k8s.io/node-stage-secret-namespace" = "ceph-csi-rbd"
    "csi.storage.k8s.io/fstype" = "ext4"
  }
  reclaim_policy = "Delete"
  allow_volume_expansion = true
  mount_options = ["discard"]
}

# Create openwhisk kubernetes namespace
#
resource "kubernetes_namespace" "openwhisk" {
  depends_on = [ module.k8s-cluster ]
  metadata {
    name = "openwhisk"
  }
}


# Deploy Openwhisk
#
resource "helm_release" "ow" {
  depends_on = [ kubernetes_namespace.openwhisk ]
  name = "ow"
  namespace = "openwhisk"
  repository = "https://openwhisk.apache.org/charts"
  chart = "openwhisk"
  version = "1.0.0"
  wait = false

  values = [
    <<EOT
whisk:
  ingress:
    type: NodePort
    apiHostName: ${module.k8s-cluster.worker_hosts[0]}
    apiHostPort: 31001
k8s:
  persistence:
    explicitStorageClass: csid-rbd-sc

kafka:
  replicaCount: ${var.kafka_replicas}
  persistence:
    size: ${var.kafka_persistence_size}

db:
  persistence:
    size: 10Gi

redis:
  persistence:
    size: 2Gi
EOT
  ]
}