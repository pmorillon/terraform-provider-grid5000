output "wsk_apihost" {
    value = "https://${module.k8s-cluster.worker_hosts[0]}:31001"
}

output "wsk_auth" {
    value = "23bc46b1-71f6-4ed5-8c54-816aa4f8c502:123zO3xZCLrMN6v2BKK1dXYFpXlPkccOFqm12CdAsMgRU4VrNZ9lyGVCGuMDGIwP"
}