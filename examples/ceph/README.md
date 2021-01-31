# How to deploy Ceph cluster on Grid5000 using Terraform

This repository is an example to deploy a Ceph cluster with Prometheus monitoring stack, using Terraform. Tools involved :
* [Grid'5000 Terraform provider](https://registry.terraform.io/providers/pmorillon/grid5000/latest/docs), based on [gog5k Go package](https://pkg.go.dev/gitlab.inria.fr/pmorillo/gog5k#section-documentation)
* [Rancher Kubernetes Engine Terraform provider](https://registry.terraform.io/providers/rancher/rke/latest/docs)
* [Rook Ceph kubernetes operator](https://rook.io)
* [Kube Prometheus Stack](https://github.com/prometheus-community/helm-charts/tree/main/charts/kube-prometheus-stack)

## How to use

### Requirements

* A Grid'5000 account
* [terraform](https://terraform.io/) v.0.13.x on a Grid'5000 frontend
* [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) on a Grid'5000 frontend

### Deploy Ceph cluster on Grid'5000

On a Grid'5000 frontend :

```sh
# Clone this repo
git clone https://github.com/pmorillon/terraform-provider-grid5000.git
cd terraform-grid5000-provider/examples/ceph

# init : will automatically download required Terraform providers
terraform init
```

Edit `main.tf` file to describe your deployment :

```tf
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
    # rook_ceph_version = "v1.5.5"
    # ceph_version = "v15.2.8-20201217"
}
```

```sh
terraform apply
```