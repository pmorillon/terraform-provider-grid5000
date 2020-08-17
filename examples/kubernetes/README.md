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
cd terraform-grid5000-provider/examples/kubernetes

# init : will automatically download required Terraform providers
terraform init

# apply
terraform apply

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
NAME                              STATUS   ROLES               AGE   VERSION
paravance-66.rennes.grid5000.fr   Ready    controlplane,etcd   28m   v1.18.3
paravance-67.rennes.grid5000.fr   Ready    worker              28m   v1.18.3
paravance-7.rennes.grid5000.fr    Ready    worker              28m   v1.18.3
paravance-9.rennes.grid5000.fr    Ready    worker              28m   v1.18.3

# To destroy the cluster and delete the OAR job
unset KUBECONFIG
terraform destroy
# ...
helm_release.prometheus: Destroying... [id=prom-op]
helm_release.prometheus: Destruction complete after 4s
kubernetes_namespace.monitoring: Destroying... [id=monitoring]
kubernetes_namespace.monitoring: Still destroying... [id=monitoring, 10s elapsed]
kubernetes_namespace.monitoring: Destruction complete after 13s
rke_cluster.cluster: Destroying... [id=9f7be6f0-01a4-4be6-b178-76a3459cb9eb]
rke_cluster.cluster: Still destroying... [id=9f7be6f0-01a4-4be6-b178-76a3459cb9eb, 10s elapsed]
rke_cluster.cluster: Destruction complete after 10s
null_resource.docker_install[3]: Destroying... [id=6174148238043559366]
null_resource.docker_install[0]: Destroying... [id=3829418478221242807]
null_resource.docker_install[3]: Destruction complete after 0s
null_resource.docker_install[0]: Destruction complete after 0s
null_resource.docker_install[1]: Destroying... [id=7000171430741773912]
null_resource.docker_install[2]: Destroying... [id=5248345062053011481]
null_resource.docker_install[1]: Destruction complete after 0s
null_resource.docker_install[2]: Destruction complete after 0s
grid5000_deployment.my_deployment: Destroying... [id=D-a1976333-465e-4b7e-97b0-012e79f9e30a]
grid5000_deployment.my_deployment: Destruction complete after 0s
grid5000_job.my_job: Destroying... [id=1282973]
grid5000_job.my_job: Destruction complete after 1s
```