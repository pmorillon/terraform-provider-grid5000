# How to deploy Kubernetes cluster on Grid'5000 using Terraform and Terraform Grid'5000 and RKE providers

This repository is an example for building a Kubernetes cluster using Terraform and Terraform RKE provider on [Grid'5000](https://www.grid5000.fr).

## How to use

### Requirements

* A Grid'5000 account
* [terraform](https://terraform.io/) v.0.13.x on a Grid'5000 frontend
* [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) on a Grid'5000 frontend

### Deploy Kubernetes cluster on Grid'5000

On a Grid'5000 frontend :

```sh
# Clone this repo
git clone https://github.com/pmorillon/terraform-provider-grid5000.git
cd terraform-provider-grid5000/examples/kubernetes


# init : will automatically download required Terraform providers
terraform init
```

Edit `main.tf` file to describe your deployment, see available [input values](https://registry.terraform.io/modules/pmorillon/k8s-cluster/grid5000/latest?tab=inputs)

```tf
module "k8s_cluster" {
    source = "pmorillon/k8s-cluster/grid5000"
    version = "~> 0.0.1"

    walltime = "2"
    site = "rennes"
}
```

```sh
# apply
terraform apply
# ...
module.k8s_cluster.grid5000_job.k8s: Creating...
module.k8s_cluster.grid5000_job.k8s: Still creating... [10s elapsed]
module.k8s_cluster.grid5000_job.k8s: Creation complete after 18s [id=1457758]
module.k8s_cluster.grid5000_deployment.k8s: Creating...
module.k8s_cluster.grid5000_deployment.k8s: Still creating... [10s elapsed]
# ...
module.k8s_cluster.grid5000_deployment.k8s: Creation complete after 4m51s [id=D-3c3d18c3-3f21-4f63-8c96-282fff9ea135]
module.k8s_cluster.null_resource.docker_install[0]: Creating...
module.k8s_cluster.null_resource.docker_install[0]: Provisioning with 'file'...
module.k8s_cluster.null_resource.docker_install[1]: Creating...
module.k8s_cluster.null_resource.docker_install[3]: Creating...
module.k8s_cluster.null_resource.docker_install[2]: Creating...
module.k8s_cluster.null_resource.docker_install[3]: Provisioning with 'file'...
module.k8s_cluster.null_resource.docker_install[1]: Provisioning with 'file'...
module.k8s_cluster.null_resource.docker_install[2]: Provisioning with 'file'...
module.k8s_cluster.null_resource.docker_install[2]: Provisioning with 'remote-exec'...
# ...
module.k8s_cluster.null_resource.docker_install[0]: Still creating... [10s elapsed]
module.k8s_cluster.null_resource.docker_install[1]: Still creating... [10s elapsed]
module.k8s_cluster.null_resource.docker_install[3]: Still creating... [10s elapsed]
module.k8s_cluster.null_resource.docker_install[2]: Still creating... [10s elapsed]
# ...
module.k8s_cluster.null_resource.docker_install[1]: Creation complete after 1m40s [id=2735933040659486651]
module.k8s_cluster.null_resource.docker_install[0]: Creation complete after 1m40s [id=5229444562703927069]
module.k8s_cluster.null_resource.docker_install[2]: Creation complete after 1m41s [id=4065656999805753798]
module.k8s_cluster.null_resource.docker_install[3]: Creation complete after 1m45s [id=362653944675258782]
module.k8s_cluster.rke_cluster.cluster: Creating...
module.k8s_cluster.rke_cluster.cluster: Still creating... [10s elapsed]
# ...
module.k8s_cluster.rke_cluster.cluster: Creation complete after 3m54s [id=edd11299-1b9c-46aa-9027-9b49a8d8724d]
module.k8s_cluster.local_file.kube_cluster_yaml: Creating...
module.k8s_cluster.local_file.kube_cluster_yaml: Creation complete after 0s [id=067ddeeab04eb55eb19756d705f74f4481033df6]

Apply complete! Resources: 8 added, 0 changed, 0 destroyed.

# When terraform apply is completed, a kubeconfig file must be
# created in the current directory
export KUBECONFIG=${PWD}/kube_config_cluster.yml


# Check component statuses
kubectl get cs
NAME                 STATUS    MESSAGE             ERROR
scheduler            Healthy   ok                  
controller-manager   Healthy   ok                  
etcd-0               Healthy   {"health":"true"}


# List cluster nodes
kubectl get nodes        
NAME                              STATUS   ROLES               AGE     VERSION
paranoia-1.rennes.grid5000.fr     Ready    controlplane,etcd   5m19s   v1.19.6
paranoia-5.rennes.grid5000.fr     Ready    worker              5m16s   v1.19.6
paranoia-6.rennes.grid5000.fr     Ready    worker              5m16s   v1.19.6
paravance-63.rennes.grid5000.fr   Ready    worker              5m15s   v1.19.6


# Show deployed resources
kubectl get all -A 
NAMESPACE       NAME                                           READY   STATUS      RESTARTS   AGE
ingress-nginx   pod/default-http-backend-65dd5949d9-4p2mc      1/1     Running     0          5m5s
ingress-nginx   pod/nginx-ingress-controller-hxzhw             1/1     Running     0          5m5s
ingress-nginx   pod/nginx-ingress-controller-sqq7s             1/1     Running     0          5m5s
ingress-nginx   pod/nginx-ingress-controller-z98mz             1/1     Running     0          5m5s
kube-system     pod/calico-kube-controllers-7fbff695b4-52fgf   1/1     Running     0          5m32s
kube-system     pod/canal-6x6kk                                2/2     Running     0          5m32s
kube-system     pod/canal-fwkcv                                2/2     Running     0          5m32s
kube-system     pod/canal-n6pkv                                2/2     Running     0          5m32s
kube-system     pod/canal-wtnxg                                2/2     Running     0          5m32s
kube-system     pod/coredns-6f85d5fb88-7fqzf                   1/1     Running     0          4m39s
kube-system     pod/coredns-6f85d5fb88-9gpx7                   1/1     Running     0          5m26s
kube-system     pod/coredns-autoscaler-79599b9dc6-2jvrq        1/1     Running     0          5m25s
kube-system     pod/metrics-server-8449844bf-d8drl             1/1     Running     0          5m15s
kube-system     pod/rke-coredns-addon-deploy-job-gl6n9         0/1     Completed   0          5m30s
kube-system     pod/rke-ingress-controller-deploy-job-qpmzf    0/1     Completed   0          5m9s
kube-system     pod/rke-metrics-addon-deploy-job-vwxdt         0/1     Completed   0          5m20s
kube-system     pod/rke-network-plugin-deploy-job-4f89s        0/1     Completed   0          5m45s

NAMESPACE       NAME                           TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)                  AGE
default         service/kubernetes             ClusterIP   10.43.0.1      <none>        443/TCP                  6m33s
ingress-nginx   service/default-http-backend   ClusterIP   10.43.240.19   <none>        80/TCP                   5m5s
kube-system     service/kube-dns               ClusterIP   10.43.0.10     <none>        53/UDP,53/TCP,9153/TCP   5m26s
kube-system     service/metrics-server         ClusterIP   10.43.49.31    <none>        443/TCP                  5m15s

NAMESPACE       NAME                                      DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR            AGE
ingress-nginx   daemonset.apps/nginx-ingress-controller   3         3         3       3            3           <none>                   5m5s
kube-system     daemonset.apps/canal                      4         4         4       4            4           kubernetes.io/os=linux   5m33s

NAMESPACE       NAME                                      READY   UP-TO-DATE   AVAILABLE   AGE
ingress-nginx   deployment.apps/default-http-backend      1/1     1            1           5m5s
kube-system     deployment.apps/calico-kube-controllers   1/1     1            1           5m33s
kube-system     deployment.apps/coredns                   2/2     2            2           5m27s
kube-system     deployment.apps/coredns-autoscaler        1/1     1            1           5m26s
kube-system     deployment.apps/metrics-server            1/1     1            1           5m15s

NAMESPACE       NAME                                                 DESIRED   CURRENT   READY   AGE
ingress-nginx   replicaset.apps/default-http-backend-65dd5949d9      1         1         1       5m5s
kube-system     replicaset.apps/calico-kube-controllers-7fbff695b4   1         1         1       5m33s
kube-system     replicaset.apps/coredns-6f85d5fb88                   2         2         2       5m27s
kube-system     replicaset.apps/coredns-autoscaler-79599b9dc6        1         1         1       5m26s
kube-system     replicaset.apps/metrics-server-8449844bf             1         1         1       5m15s

NAMESPACE     NAME                                          COMPLETIONS   DURATION   AGE
kube-system   job.batch/rke-coredns-addon-deploy-job        1/1           6s         5m30s
kube-system   job.batch/rke-ingress-controller-deploy-job   1/1           6s         5m9s
kube-system   job.batch/rke-metrics-addon-deploy-job        1/1           6s         5m20s
kube-system   job.batch/rke-network-plugin-deploy-job       1/1           13s        5m45s


# To destroy the cluster and delete the OAR job
terraform destroy
# ...
module.k8s_cluster.local_file.kube_cluster_yaml: Destroying... [id=067ddeeab04eb55eb19756d705f74f4481033df6]
module.k8s_cluster.local_file.kube_cluster_yaml: Destruction complete after 0s
module.k8s_cluster.rke_cluster.cluster: Destroying... [id=edd11299-1b9c-46aa-9027-9b49a8d8724d]
module.k8s_cluster.rke_cluster.cluster: Destruction complete after 10s
module.k8s_cluster.null_resource.docker_install[1]: Destroying... [id=2735933040659486651]
module.k8s_cluster.null_resource.docker_install[1]: Destruction complete after 0s
module.k8s_cluster.null_resource.docker_install[3]: Destroying... [id=362653944675258782]
module.k8s_cluster.null_resource.docker_install[0]: Destroying... [id=5229444562703927069]
module.k8s_cluster.null_resource.docker_install[2]: Destroying... [id=4065656999805753798]
module.k8s_cluster.null_resource.docker_install[3]: Destruction complete after 0s
module.k8s_cluster.null_resource.docker_install[2]: Destruction complete after 0s
module.k8s_cluster.null_resource.docker_install[0]: Destruction complete after 0s
module.k8s_cluster.grid5000_deployment.k8s: Destroying... [id=D-3c3d18c3-3f21-4f63-8c96-282fff9ea135]
module.k8s_cluster.grid5000_deployment.k8s: Destruction complete after 0s
module.k8s_cluster.grid5000_job.k8s: Destroying... [id=1457758]
module.k8s_cluster.grid5000_job.k8s: Destruction complete after 1s

Destroy complete! Resources: 8 destroyed.
```