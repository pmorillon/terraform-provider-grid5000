resource "local_file" "kube_cluster_yaml" {
  filename = "kube_config_cluster.yml"
  content = rke_cluster.cluster.kube_config_yaml
}
