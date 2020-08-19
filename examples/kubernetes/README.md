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


# Show deployed application
kubectl -n monitoring get all         
NAME                                                         READY   STATUS    RESTARTS   AGE
pod/alertmanager-prom-op-prometheus-operato-alertmanager-0   2/2     Running   0          119s
pod/prom-op-grafana-775bb4dcbf-kcv84                         2/2     Running   0          2m12s
pod/prom-op-kube-state-metrics-5b69759966-9jt6w              1/1     Running   0          2m12s
pod/prom-op-prometheus-node-exporter-9rb7d                   1/1     Running   0          2m12s
pod/prom-op-prometheus-node-exporter-mhs8v                   1/1     Running   0          2m12s
pod/prom-op-prometheus-node-exporter-zrggt                   1/1     Running   0          2m12s
pod/prom-op-prometheus-operato-operator-5f7ccbc9b7-wtfxr     2/2     Running   0          2m12s
pod/prometheus-prom-op-prometheus-operato-prometheus-0       3/3     Running   1          109s

NAME                                              TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)                      AGE
service/alertmanager-operated                     ClusterIP   None            <none>        9093/TCP,9094/TCP,9094/UDP   119s
service/prom-op-grafana                           ClusterIP   10.43.225.48    <none>        80/TCP                       2m12s
service/prom-op-kube-state-metrics                ClusterIP   10.43.229.88    <none>        8080/TCP                     2m12s
service/prom-op-prometheus-node-exporter          ClusterIP   10.43.231.69    <none>        9100/TCP                     2m12s
service/prom-op-prometheus-operato-alertmanager   ClusterIP   10.43.111.112   <none>        9093/TCP                     2m12s
service/prom-op-prometheus-operato-operator       ClusterIP   10.43.215.68    <none>        8080/TCP,443/TCP             2m12s
service/prom-op-prometheus-operato-prometheus     ClusterIP   10.43.175.156   <none>        9090/TCP                     2m12s
service/prometheus-operated                       ClusterIP   None            <none>        9090/TCP                     109s

NAME                                              DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR   AGE
daemonset.apps/prom-op-prometheus-node-exporter   3         3         3       3            3           <none>          2m12s

NAME                                                  READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/prom-op-grafana                       1/1     1            1           2m12s
deployment.apps/prom-op-kube-state-metrics            1/1     1            1           2m12s
deployment.apps/prom-op-prometheus-operato-operator   1/1     1            1           2m12s

NAME                                                             DESIRED   CURRENT   READY   AGE
replicaset.apps/prom-op-grafana-775bb4dcbf                       1         1         1       2m12s
replicaset.apps/prom-op-kube-state-metrics-5b69759966            1         1         1       2m12s
replicaset.apps/prom-op-prometheus-operato-operator-5f7ccbc9b7   1         1         1       2m12s

NAME                                                                    READY   AGE
statefulset.apps/alertmanager-prom-op-prometheus-operato-alertmanager   1/1     119s
statefulset.apps/prometheus-prom-op-prometheus-operato-prometheus       1/1     109s


# To destroy the cluster and delete the OAR job
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