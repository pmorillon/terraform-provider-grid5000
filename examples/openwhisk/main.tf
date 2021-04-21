module "g5k-openwhisk" {
    source = "./modules/g5k-openwhisk"

    username = "username" # Replace by your Grid'5000 username
    nodes_location = "rennes"
    nodes_count = 5
    walltime = "1"

    data_location = "rennes" # rennes or nantes
    ceph_pool_quota = "200G"
}

output "wsk_set_apihost" {
    value = "wsk property set --apihost ${module.g5k-openwhisk.wsk_apihost}"
}

output "wsk_set_auth" {
    value = "wsk property set --auth ${module.g5k-openwhisk.wsk_auth}"
}