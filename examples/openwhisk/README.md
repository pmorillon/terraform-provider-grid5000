# How to deploy OpenWhisk on Grid'5000 using Terraform

[OpenWhisk](https://openwhisk.apache.org) is an open source serverless cloud platform. 

## How to use

### Requirements

* A Grid'5000 account
* [terraform](https://terraform.io/) v.0.13.x on a Grid'5000 frontend
* [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/),[wsk](https://openwhisk.apache.org/documentation.html#wsk-cli) on a Grid'5000 frontend

### Deploy OpenWhisk on Grid'5000

On a Grid'5000 frontend :

```sh
# Clone this repo
> git clone https://github.com/pmorillon/terraform-provider-grid5000.git
> cd terraform-grid5000-provider/examples/openwhisk


# init : will automatically download required Terraform providers
> terraform init
```

Edit `main.tf` file to describe your deployment, see available input values into file `g5k-openwhisk/variables.tf` :

```tf
module "g5k-openwhisk" {
    source = "./modules/g5k-openwhisk"

    username = "username" # Replace by your Grid'5000 username
    nodes_location = "rennes"
    nodes_count = 5
    walltime = "1"

    data_location = "rennes" # rennes or nantes
    ceph_pool_quota = "200G"
}
```

```sh
# apply
> terraform apply
module.g5k-openwhisk.grid5000_ceph_pool.k8s: Creating...
module.g5k-openwhisk.module.k8s-cluster.grid5000_job.k8s: Creating...
module.g5k-openwhisk.grid5000_ceph_pool.k8s: Creation complete after 7s [id=k8s]
module.g5k-openwhisk.module.k8s-cluster.grid5000_job.k8s: Still creating... [10s elapsed]
module.g5k-openwhisk.module.k8s-cluster.grid5000_job.k8s: Still creating... [20s elapsed]
module.g5k-openwhisk.module.k8s-cluster.grid5000_job.k8s: Still creating... [30s elapsed]
module.g5k-openwhisk.module.k8s-cluster.grid5000_job.k8s: Creation complete after 34s [id=1543904]
module.g5k-openwhisk.module.k8s-cluster.grid5000_deployment.k8s: Creating...
module.g5k-openwhisk.module.k8s-cluster.grid5000_deployment.k8s: Still creating... [10s elapsed]
# ...
module.g5k-openwhisk.module.k8s-cluster.grid5000_deployment.k8s: Creation complete after 3m31s [id=D-47a616f6-b4bf-477c-822b-54833ede25d7]
module.g5k-openwhisk.module.k8s-cluster.null_resource.docker_install[3]: Creating...
module.g5k-openwhisk.module.k8s-cluster.null_resource.docker_install[2]: Creating...
module.g5k-openwhisk.module.k8s-cluster.null_resource.docker_install[0]: Creating...
module.g5k-openwhisk.module.k8s-cluster.null_resource.docker_install[1]: Creating...
module.g5k-openwhisk.module.k8s-cluster.null_resource.docker_install[4]: Creating...
# ...
module.g5k-openwhisk.module.k8s-cluster.null_resource.docker_install[3]: Creation complete after 1m38s [id=4582518092104065079]
module.g5k-openwhisk.module.k8s-cluster.null_resource.docker_install[0]: Creation complete after 1m39s [id=8211829712503359637]
module.g5k-openwhisk.module.k8s-cluster.null_resource.docker_install[2]: Creation complete after 1m40s [id=6649728824178764707]
module.g5k-openwhisk.module.k8s-cluster.null_resource.docker_install[1]: Still creating... [1m40s elapsed]
module.g5k-openwhisk.module.k8s-cluster.null_resource.docker_install[4]: Still creating... [1m40s elapsed]
module.g5k-openwhisk.module.k8s-cluster.null_resource.docker_install[4]: Creation complete after 1m46s [id=504146819526374970]
module.g5k-openwhisk.module.k8s-cluster.null_resource.docker_install[1]: Creation complete after 1m48s [id=7222830882574490417]
module.g5k-openwhisk.module.k8s-cluster.rke_cluster.cluster: Creating...
module.g5k-openwhisk.module.k8s-cluster.rke_cluster.cluster: Still creating... [10s elapsed]
# ...
module.g5k-openwhisk.module.k8s-cluster.rke_cluster.cluster: Creation complete after 4m9s [id=2f26602c-cacd-41e2-96fb-2eb09b44885b]
module.g5k-openwhisk.module.k8s-cluster.local_file.kube_cluster_yaml: Creating...
module.g5k-openwhisk.module.k8s-cluster.local_file.kube_cluster_yaml: Creation complete after 1s [id=9b3444477cdbc382646bd4b5944bc24590454309]
module.g5k-openwhisk.kubernetes_namespace.openwhisk: Creating...
module.g5k-openwhisk.kubernetes_namespace.ceph-csi-rbd: Creating...
module.g5k-openwhisk.kubernetes_secret.csi-rbd-secret: Creating...
module.g5k-openwhisk.kubernetes_storage_class.csi-rbd-sc: Creating...
module.g5k-openwhisk.kubernetes_storage_class.csi-rbd-sc: Creation complete after 0s [id=csid-rbd-sc]
module.g5k-openwhisk.kubernetes_namespace.ceph-csi-rbd: Creation complete after 0s [id=ceph-csi-rbd]
module.g5k-openwhisk.kubernetes_namespace.openwhisk: Creation complete after 0s [id=openwhisk]
module.g5k-openwhisk.kubernetes_secret.csi-rbd-secret: Creation complete after 1s [id=ceph-csi-rbd/csi-rbd-secret]
module.g5k-openwhisk.helm_release.ceph-csi-rbd: Creating...
module.g5k-openwhisk.helm_release.ceph-csi-rbd: Still creating... [10s elapsed]
module.g5k-openwhisk.helm_release.ow: Creating...
module.g5k-openwhisk.helm_release.ceph-csi-rbd: Still creating... [20s elapsed]
module.g5k-openwhisk.helm_release.ow: Still creating... [10s elapsed]
module.g5k-openwhisk.helm_release.ceph-csi-rbd: Still creating... [30s elapsed]
module.g5k-openwhisk.helm_release.ow: Still creating... [20s elapsed]
module.g5k-openwhisk.helm_release.ow: Creation complete after 22s [id=ow]
module.g5k-openwhisk.helm_release.ceph-csi-rbd: Still creating... [40s elapsed]
module.g5k-openwhisk.helm_release.ceph-csi-rbd: Still creating... [50s elapsed]
module.g5k-openwhisk.helm_release.ceph-csi-rbd: Creation complete after 59s [id=ceph-sci-rbd]

Apply complete! Resources: 16 added, 0 changed, 0 destroyed.

Outputs:

wsk_set_apihost = "wsk property set --apihost https://parasilo-11.rennes.grid5000.fr:31001"
wsk_set_auth = "wsk property set --auth 23bc46b1-71f6-4ed5-8c54-816aa4f8c502:123zO3xZCLrMN6v2BKK1dXYFpXlPkccOFqm12CdAsMgRU4VrNZ9lyGVCGuMDGIwP"
```

```sh
# When terraform apply is completed, a kubeconfig file must be
# created in the current directory
> export KUBECONFIG=${PWD}/kube_config_cluster.yml

# Add node label openwhisk-role
> kubectl label nodes --all openwhisk-role=invoker
node/parasilo-10.rennes.grid5000.fr labeled
node/parasilo-11.rennes.grid5000.fr labeled
node/parasilo-12.rennes.grid5000.fr labeled
node/parasilo-13.rennes.grid5000.fr labeled
node/parasilo-14.rennes.grid5000.fr labeled
> kubectl -n openwhisk wait --for=condition=complete job/ow-install-packages
job.batch/ow-install-packages condition met
> kubectl -n openwhisk get pods,pvc,job,svc
NAME                                                          READY   STATUS      RESTARTS   AGE
pod/ow-alarmprovider-5466dfc55b-k5wxq                         1/1     Running     0          9m16s
pod/ow-apigateway-6bc54688fd-2k5hj                            1/1     Running     0          9m16s
pod/ow-controller-0                                           1/1     Running     0          9m16s
pod/ow-couchdb-547b88587f-mdbzh                               1/1     Running     0          9m16s
pod/ow-gen-certs-jpx22                                        0/1     Completed   0          9m16s
pod/ow-init-couchdb-cghrz                                     0/1     Completed   0          9m16s
pod/ow-install-packages-v5sf2                                 0/1     Completed   0          9m16s
pod/ow-invoker-0                                              1/1     Running     0          9m16s
pod/ow-kafka-0                                                1/1     Running     0          9m16s
pod/ow-kafka-1                                                1/1     Running     0          9m16s
pod/ow-kafka-2                                                1/1     Running     0          9m16s
pod/ow-kafkaprovider-8467b78b8b-v74ks                         1/1     Running     0          9m16s
pod/ow-nginx-5d97d56cd5-hnbnz                                 1/1     Running     0          9m16s
pod/ow-redis-5f95f4679-l79xz                                  1/1     Running     0          9m16s
pod/ow-wskadmin                                               1/1     Running     0          9m16s
pod/ow-zookeeper-0                                            1/1     Running     0          9m16s
pod/wskow-invoker-00-1-prewarm-nodejs10                       1/1     Running     0          6m36s
pod/wskow-invoker-00-2-prewarm-nodejs10                       1/1     Running     0          6m36s
pod/wskow-invoker-00-3-whisksystem-invokerhealthtestaction0   1/1     Running     0          6m36s

NAME                                             STATUS   VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS   AGE
persistentvolumeclaim/ow-alarmprovider-pvc       Bound    pvc-9d445e06-0c50-4bd3-acbd-9af6da69eed5   1Gi        RWO            csid-rbd-sc    9m16s
persistentvolumeclaim/ow-couchdb-pvc             Bound    pvc-53545452-b30d-475d-a0b8-b13f5dbeb44d   10Gi       RWO            csid-rbd-sc    9m16s
persistentvolumeclaim/ow-kafka-pvc-ow-kafka-0    Bound    pvc-69e8c791-2345-4c28-aedd-59f399b6637b   20Gi       RWO            csid-rbd-sc    9m16s
persistentvolumeclaim/ow-kafka-pvc-ow-kafka-1    Bound    pvc-56cc58b3-1817-4576-b86a-a2ead27d2fb8   20Gi       RWO            csid-rbd-sc    9m16s
persistentvolumeclaim/ow-kafka-pvc-ow-kafka-2    Bound    pvc-c110a3a8-43c1-4535-a629-0219c13b2180   20Gi       RWO            csid-rbd-sc    9m16s
persistentvolumeclaim/ow-redis-pvc               Bound    pvc-e6213cc8-9a56-41ed-bbba-43c769a6ef28   2Gi        RWO            csid-rbd-sc    9m16s
persistentvolumeclaim/ow-zookeeper-pvc-data      Bound    pvc-26436a3e-2b0c-4ce0-b414-226a177c15ee   256Mi      RWO            csid-rbd-sc    9m16s
persistentvolumeclaim/ow-zookeeper-pvc-datalog   Bound    pvc-26c7dacd-cd11-43c1-ac47-6f3e692eef58   256Mi      RWO            csid-rbd-sc    9m16s

NAME                            COMPLETIONS   DURATION   AGE
job.batch/ow-gen-certs          1/1           56s        9m16s
job.batch/ow-init-couchdb       1/1           103s       9m16s
job.batch/ow-install-packages   1/1           4m18s      9m16s

NAME                    TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)                      AGE
service/ow-apigateway   ClusterIP   10.43.182.42    <none>        8080/TCP,9000/TCP            9m16s
service/ow-controller   ClusterIP   10.43.129.46    <none>        8080/TCP                     9m16s
service/ow-couchdb      ClusterIP   10.43.8.61      <none>        5984/TCP                     9m16s
service/ow-kafka        ClusterIP   None            <none>        9092/TCP                     9m16s
service/ow-nginx        NodePort    10.43.244.210   <none>        80:32593/TCP,443:31001/TCP   9m16s
service/ow-redis        ClusterIP   10.43.2.4       <none>        6379/TCP                     9m16s
service/ow-zookeeper    ClusterIP   None            <none>        2181/TCP,2888/TCP,3888/TCP   9m16s
```

```sh
> wsk property set --apihost https://parasilo-11.rennes.grid5000.fr:31001
> wsk property set --auth 23bc46b1-71f6-4ed5-8c54-816aa4f8c502:123zO3xZCLrMN6v2BKK1dXYFpXlPkccOFqm12CdAsMgRU4VrNZ9lyGVCGuMDGIwP
> wsk --insecure package list /whisk.system                                                                   
packages
/whisk.system/messaging                                                shared
/whisk.system/alarms                                                   shared
/whisk.system/github                                                   shared
/whisk.system/slack                                                    shared
/whisk.system/weather                                                  shared
/whisk.system/websocket                                                shared
/whisk.system/samples                                                  shared
/whisk.system/utils                                                    shared
```