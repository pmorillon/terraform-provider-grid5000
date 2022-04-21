module "rook_ceph" {
    source = "./modules/rook_ceph"

    # Grid'5000 site
    site = "rennes"

    # Choose a Grid'5000 cluster with several reservable disks, eg. :
    # parasilo@rennes : 4 rotational disks and 1 ssd
    # chifflot@lille : 4 rotational disks and 1 ssd
    # ... Refer to the Grid'5000 Reference repository and hardware wiki page.
    cluster_selector = "parasilo"

    # Uncomment ceph_metadata_device if you want to use an ssd disk for Bluestore metadata.
    # parasilo@rennes : sdf
    # chifflot@lille : sdb
    #
    # ceph_metadata_device = "sdf"

    # Number of reserved nodes, 4 minimum (1 Kubernetes controlplane/etcd node, 3 workers 
    # in order to satisfy the Ceph monitors quorum).
    nodes_count = 5

    # OAR job duration.
    walltime = "3"

    # Other defaults variables :
    # rook_ceph_version = "v1.9.0"
    # ceph_version = "v16.2.7"
}
