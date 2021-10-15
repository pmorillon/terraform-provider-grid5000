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
    # parasilo@rennes : disk5
    # chifflot@lille : disk1
    #
    # ceph_metadata_device = "disk5"

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

When terraform apply is completed, a kubeconfig file must be created in the current directory.

```sh
export KUBECONFIG=${PWD}/kube_config_cluster.yml
```

Get ceph cluster creation status, wait for `Created` status :

```sh
watch kubectl -n rook-ceph get cephcluster rook-ceph -o=jsonpath='{.status.state}'
```

```sh
kubectl -n rook-ceph get pods -o wide
# NAME                                                              READY   STATUS      RESTARTS   AGE     IP            NODE                         
# csi-cephfsplugin-7mkfn                                            3/3     Running     0          3m18s   172.16.97.7   parasilo-7.rennes.grid5000.fr
# csi-cephfsplugin-b56w6                                            3/3     Running     0          3m18s   172.16.97.8   parasilo-8.rennes.grid5000.fr
# csi-cephfsplugin-provisioner-8658f67749-lmfzw                     6/6     Running     0          3m18s   10.42.3.7     parasilo-5.rennes.grid5000.fr
# csi-cephfsplugin-provisioner-8658f67749-p7gt6                     6/6     Running     0          3m18s   10.42.2.5     parasilo-8.rennes.grid5000.fr
# csi-cephfsplugin-s9g89                                            3/3     Running     0          3m18s   172.16.97.5   parasilo-5.rennes.grid5000.fr
# csi-cephfsplugin-tmv9t                                            3/3     Running     0          3m18s   172.16.97.3   parasilo-3.rennes.grid5000.fr
# csi-rbdplugin-6k86k                                               3/3     Running     0          3m19s   172.16.97.7   parasilo-7.rennes.grid5000.fr
# csi-rbdplugin-dxvjw                                               3/3     Running     0          3m19s   172.16.97.3   parasilo-3.rennes.grid5000.fr
# csi-rbdplugin-lhjsr                                               3/3     Running     0          3m19s   172.16.97.5   parasilo-5.rennes.grid5000.fr
# csi-rbdplugin-pcvjt                                               3/3     Running     0          3m19s   172.16.97.8   parasilo-8.rennes.grid5000.fr
# csi-rbdplugin-provisioner-94f699d86-5clfx                         6/6     Running     0          3m19s   10.42.4.5     parasilo-7.rennes.grid5000.fr
# csi-rbdplugin-provisioner-94f699d86-tz6m9                         6/6     Running     0          3m19s   10.42.1.7     parasilo-3.rennes.grid5000.fr
# rook-ceph-crashcollector-parasilo-3.rennes.grid5000.fr-6c9t8glz   1/1     Running     0          83s     172.16.97.3   parasilo-3.rennes.grid5000.fr
# rook-ceph-crashcollector-parasilo-5.rennes.grid5000.fr-6d7m58hx   1/1     Running     0          27s     172.16.97.5   parasilo-5.rennes.grid5000.fr
# rook-ceph-crashcollector-parasilo-7.rennes.grid5000.fr-8fb2q7vf   1/1     Running     0          2m18s   172.16.97.7   parasilo-7.rennes.grid5000.fr
# rook-ceph-crashcollector-parasilo-8.rennes.grid5000.fr-59dw6fs2   1/1     Running     0          29s     172.16.97.8   parasilo-8.rennes.grid5000.fr
# rook-ceph-mds-cephfs-a-6dd7b9df7-qxgj8                            1/1     Running     0          29s     172.16.97.8   parasilo-8.rennes.grid5000.fr
# rook-ceph-mds-cephfs-b-644b4855f7-vhmxr                           1/1     Running     0          27s     172.16.97.5   parasilo-5.rennes.grid5000.fr
# rook-ceph-mgr-a-79569f9f4c-76z82                                  1/1     Running     0          2m4s    172.16.97.3   parasilo-3.rennes.grid5000.fr
# rook-ceph-mon-a-6dc85f7fdd-tzpqb                                  1/1     Running     0          3m12s   172.16.97.3   parasilo-3.rennes.grid5000.fr
# rook-ceph-mon-b-56fccd9468-qv6qh                                  1/1     Running     0          2m48s   172.16.97.5   parasilo-5.rennes.grid5000.fr
# rook-ceph-mon-c-759c86c58d-cd4fw                                  1/1     Running     0          2m19s   172.16.97.7   parasilo-7.rennes.grid5000.fr
# rook-ceph-operator-75d899d57c-zsdjb                               1/1     Running     0          4m21s   10.42.1.4     parasilo-3.rennes.grid5000.fr
# rook-ceph-osd-0-559c5c9748-vw4rr                                  1/1     Running     0          90s     172.16.97.5   parasilo-5.rennes.grid5000.fr
# rook-ceph-osd-1-8455c8878f-m7m6d                                  1/1     Running     0          86s     172.16.97.7   parasilo-7.rennes.grid5000.fr
# rook-ceph-osd-10-57cd96c67d-lmls2                                 1/1     Running     0          85s     172.16.97.7   parasilo-7.rennes.grid5000.fr
# rook-ceph-osd-11-848bc7bb56-lsn54                                 1/1     Running     0          83s     172.16.97.3   parasilo-3.rennes.grid5000.fr
# rook-ceph-osd-12-5787db776f-m8kg4                                 1/1     Running     0          89s     172.16.97.5   parasilo-5.rennes.grid5000.fr
# rook-ceph-osd-13-5df4cfcb5b-4xcnz                                 1/1     Running     0          84s     172.16.97.7   parasilo-7.rennes.grid5000.fr
# rook-ceph-osd-14-bb4c8dd46-qlbf5                                  1/1     Running     0          83s     172.16.97.3   parasilo-3.rennes.grid5000.fr
# rook-ceph-osd-15-bcb6f66fc-mgwgf                                  1/1     Running     0          65s     172.16.97.8   parasilo-8.rennes.grid5000.fr
# rook-ceph-osd-16-84df4c4c99-v9s9t                                 1/1     Running     0          64s     172.16.97.8   parasilo-8.rennes.grid5000.fr
# rook-ceph-osd-17-7fd5d8d7b7-vdq64                                 1/1     Running     0          63s     172.16.97.8   parasilo-8.rennes.grid5000.fr
# rook-ceph-osd-18-cf46c85df-2kkp8                                  1/1     Running     0          62s     172.16.97.8   parasilo-8.rennes.grid5000.fr
# rook-ceph-osd-19-5dd6c7f78b-g69wd                                 1/1     Running     0          62s     172.16.97.8   parasilo-8.rennes.grid5000.fr
# rook-ceph-osd-2-5c98f84f75-54296                                  1/1     Running     0          82s     172.16.97.3   parasilo-3.rennes.grid5000.fr
# rook-ceph-osd-3-5c7d6d7997-nn57t                                  1/1     Running     0          88s     172.16.97.5   parasilo-5.rennes.grid5000.fr
# rook-ceph-osd-4-7f4f8b5dff-zxp86                                  1/1     Running     0          87s     172.16.97.7   parasilo-7.rennes.grid5000.fr
# rook-ceph-osd-5-7cdb467bbd-64tp4                                  1/1     Running     0          81s     172.16.97.3   parasilo-3.rennes.grid5000.fr
# rook-ceph-osd-6-878f796b5-4ltfs                                   1/1     Running     0          91s     172.16.97.5   parasilo-5.rennes.grid5000.fr
# rook-ceph-osd-7-6b8fbf9f97-j7d7d                                  1/1     Running     0          87s     172.16.97.7   parasilo-7.rennes.grid5000.fr
# rook-ceph-osd-8-7d7bd794d6-wbjw5                                  1/1     Running     0          80s     172.16.97.3   parasilo-3.rennes.grid5000.fr
# rook-ceph-osd-9-6b876f697-v5sdl                                   1/1     Running     0          90s     172.16.97.5   parasilo-5.rennes.grid5000.fr
# rook-ceph-osd-prepare-parasilo-3.rennes.grid5000.fr-t5p5s         0/1     Completed   0          2m4s    172.16.97.3   parasilo-3.rennes.grid5000.fr
# rook-ceph-osd-prepare-parasilo-5.rennes.grid5000.fr-zn8f4         0/1     Completed   0          2m3s    172.16.97.5   parasilo-5.rennes.grid5000.fr
# rook-ceph-osd-prepare-parasilo-7.rennes.grid5000.fr-m8kfr         0/1     Completed   0          2m3s    172.16.97.7   parasilo-7.rennes.grid5000.fr
# rook-ceph-osd-prepare-parasilo-8.rennes.grid5000.fr-2nx7r         0/1     Completed   0          2m2s    172.16.97.8   parasilo-8.rennes.grid5000.fr
# rook-ceph-tools-77bf5b9b7d-lclv4                                  1/1     Running     0          3m50s   10.42.4.6     parasilo-7.rennes.grid5000.fr
```

Checks cluster status :

```sh
kubectl -n rook-ceph exec -ti $(k -n rook-ceph get pods -l app=rook-ceph-tools -o name | cut -d"/" -f2) -- ceph -s
#  cluster:
#    id:     95cfc872-caef-4a0c-b0ba-219ea40f9a70
#    health: HEALTH_OK
# 
#  services:
#    mon: 3 daemons, quorum a,b,c (age 5m)
#    mgr: a(active, since 5m)
#    mds: cephfs:1 {0=cephfs-a=up:active} 1 up:standby-replay
#    osd: 20 osds: 20 up (since 4m), 20 in (since 4m)
# 
#  data:
#    pools:   5 pools, 129 pgs
#    objects: 22 objects, 2.2 KiB
#    usage:   20 GiB used, 9.4 TiB / 9.5 TiB avail
#    pgs:     129 active+clean
# 
#  io:
#    client:   853 B/s rd, 1 op/s rd, 0 op/s wr
```

Get ceph cluster configuration file :

```sh
kubectl -n rook-ceph exec -ti $(k -n rook-ceph get pods -l app=rook-ceph-tools -o name | cut -d"/" -f2) -- cat /etc/ceph/ceph.conf
# [global]
# mon_host = 172.16.97.3:6789,172.16.97.5:6789,172.16.97.7:6789
# 
# [client.admin]
# keyring = /etc/ceph/keyring
```

Get OSD tree :

```sh
kubectl -n rook-ceph exec -ti $(k -n rook-ceph get pods -l app=rook-ceph-tools -o name | cut -d"/" -f2) -- ceph osd tree          
# ID   CLASS  WEIGHT   TYPE NAME                               STATUS  REWEIGHT  PRI-AFF
#  -1         9.46021  root default                                                     
# -10         2.36505      host parasilo-3-rennes-grid5000-fr                           
#   2    hdd  0.54579          osd.2                               up   1.00000  1.00000
#   5    hdd  0.54579          osd.5                               up   1.00000  1.00000
#   8    hdd  0.54579          osd.8                               up   1.00000  1.00000
#  11    hdd  0.54579          osd.11                              up   1.00000  1.00000
#  14    ssd  0.18188          osd.14                              up   1.00000  1.00000
#  -3         2.36505      host parasilo-5-rennes-grid5000-fr                           
#   3    hdd  0.54579          osd.3                               up   1.00000  1.00000
#   6    hdd  0.54579          osd.6                               up   1.00000  1.00000
#   9    hdd  0.54579          osd.9                               up   1.00000  1.00000
#  12    hdd  0.54579          osd.12                              up   1.00000  1.00000
#   0    ssd  0.18188          osd.0                               up   1.00000  1.00000
#  -7         2.36505      host parasilo-7-rennes-grid5000-fr                           
#   1    hdd  0.54579          osd.1                               up   1.00000  1.00000
#   4    hdd  0.54579          osd.4                               up   1.00000  1.00000
#   7    hdd  0.54579          osd.7                               up   1.00000  1.00000
#  10    hdd  0.54579          osd.10                              up   1.00000  1.00000
#  13    ssd  0.18188          osd.13                              up   1.00000  1.00000
# -13         2.36505      host parasilo-8-rennes-grid5000-fr                           
#  15    hdd  0.54579          osd.15                              up   1.00000  1.00000
#  17    hdd  0.54579          osd.17                              up   1.00000  1.00000
#  18    hdd  0.54579          osd.18                              up   1.00000  1.00000
#  19    hdd  0.54579          osd.19                              up   1.00000  1.00000
#  16    ssd  0.18188          osd.16                              up   1.00000  1.00000
 ```

 ### Access to ingresses

 ```sh
kubectl get ingress -A -o=custom-columns='DATA:spec.rules[*].host'
# DATA
# grafana.172.16.97.3.xip.io
# alertmanager.172.16.97.3.xip.io
# prometheus.172.16.97.3.xip.io
# ceph.172.16.97.3.xip.io
 ```

Use [Grid'5000 VPN](https://www.grid5000.fr/mediawiki/index.php/VPN) or [socks proxy](https://www.grid5000.fr/mediawiki/index.php/SSH#Using_OpenSSH_SOCKS_proxy) for external access.

#### Ceph dashboard

Get admin password :

```sh
kubectl -n rook-ceph get secret rook-ceph-dashboard-password -o jsonpath="{['data']['password']}" | base64 --decode && echo
# f>g46W\YC@r7<z6]J*]y
```

#### Graphana

user/password : admin/admin

#### Prometheus data

Eg. : get sum of osd read bytes during last minute :

```sh
curl -kg "https://prometheus.172.16.97.3.xip.io/api/v1/query_range?query=sum(ceph_osd_op_r_out_bytes{})&start=$(date --date='1 min ago' +%s)&end=$(date +%s)&step=10" | jq
# {
#   "status": "success",
#   "data": {
#     "resultType": "matrix",
#     "result": [
#       {
#         "metric": {},
#         "values": [
#           [
#             1612082841,
#             "203225"
#           ],
#           [
#             1612082851,
#             "204815"
#           ],
#           [
#             1612082861,
#             "204815"
#           ],
#           [
#             1612082871,
#             "206405"
#           ],
#           [
#             1612082881,
#             "207995"
#           ],
#           [
#             1612082891,
#             "207995"
#           ],
#           [
#             1612082901,
#             "209585"
#           ]
#         ]
#       }
#     ]
#   }
# }
```