provider "helm" {
  kubernetes {
    host     = rke_cluster.cluster.api_server_url
    username = rke_cluster.cluster.kube_admin_user

    client_certificate     = rke_cluster.cluster.client_cert
    client_key             = rke_cluster.cluster.client_key
    cluster_ca_certificate = rke_cluster.cluster.ca_crt
  }
}

resource "helm_release" "prometheus" {
  depends_on = [kubernetes_namespace.monitoring]

  name       = "prom-op"
  repository = "https://kubernetes-charts.storage.googleapis.com"
  chart      = "prometheus-operator"
  namespace  = "monitoring"
}