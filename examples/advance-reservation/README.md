# How to use OAR advance reservation

Two solutions.

## Solution 1 : Use terraform import

Prepare your experimentation by reserving resources manually :

```sh
oarsub -r "2020-10-13 14:00:00" -l {"type='disk' or type='default' and disk_reservation_count > 0 and cluster='parasilo'"}/host=1 -n test
```

Define an empty job :

```hcl
resource "grid5000_job" "test" {

}
```

When job is running :

```sh
terraform import grid5000_job.test 1370063@rennes
# grid5000_job.test: Importing from ID "1370063@rennes"...
# grid5000_job.test: Import prepared!
#   Prepared grid5000_job for import
# grid5000_job.test: Refreshing state... [id=1370063]
# 
# Import successful!
# 
# The resources that were imported are shown above. These resources are now in
# your Terraform state and will henceforth be managed by Terraform.

terraform state show grid5000_job.test
# # grid5000_job.test:
# resource "grid5000_job" "test" {
#     assigned_nodes     = [
#         "parasilo-1.rennes.grid5000.fr",
#     ]
#     disks_resources    = [
#         {
#             device   = "sde"
#             hostname = "parasilo-1.rennes.grid5000.fr"
#         },
#         {
#             device   = "sdd"
#             hostname = "parasilo-1.rennes.grid5000.fr"
#         },
#         {
#             device   = "sdc"
#             hostname = "parasilo-1.rennes.grid5000.fr"
#         },
#         {
#             device   = "sdb"
#             hostname = "parasilo-1.rennes.grid5000.fr"
#         },
#         {
#             device   = "sdf"
#             hostname = "parasilo-1.rennes.grid5000.fr"
#         },
#     ]
#     id                 = "1370063"
#     need_state         = "running"
#     scheduled_at_limit = "5m"
#     site               = "rennes"
#     state              = "running"
#     subnets_resources  = []
#     vlans_resources    = []
# }
```

## Solution 2 : Use grid5000_job reservation argument

Define you Grid'5000 OAR job :

```hcl
resource "grid5000_job" "job" {
  name = "Terraform test"
  resources = "/nodes=1,walltime=1"
  command = "sleep 1h"
  site = "rennes"
  reservation = "2020-10-07 7:55:00"
  types = [
    "deploy"
  ]
}

resource "grid5000_deployment" "deploy" {
  environment = "debian10-x64-base"
  site = grid5000_job.job.site
  key = file("~/.ssh/id_rsa.pub")
  nodes = grid5000_job.job.assigned_nodes
}
```

```sh
terraform apply
# ...
# grid5000_job.job: Creating...
# 
# Error: oar job 1357359 is in waiting state, it will be scheduled at 2020-10-07 07:55:00 +0200 CEST, restart terraform apply later
# ...
```

The job is created and terraform stop. The resource is automatically tainted. When the job is running :

```sh
terraform plan
# ...
# # grid5000_job.job is tainted, so must be replaced
# -/+ resource "grid5000_job" "job" {
# ...
```

We need to untaint the grid5000_job resource and apply :

```sh
terraform untaint grid5000_job.job
# Resource instance grid5000_job.job has been successfully untainted.
terraform apply
# ...
# Terraform will perform the following actions:
#
#   # grid5000_deployment.deploy will be created
#   + resource "grid5000_deployment" "deploy" {
# ...
# grid5000_deployment.deploy: Creating...
# grid5000_deployment.deploy: Still creating... [10s elapsed]
# ...
# grid5000_deployment.deploy: Creation complete after 3m54s [id=D-2c3e6305-7bc4-42e6-a71d-9be5a84ef2ce]
#
# Apply complete! Resources: 1 added, 0 changed, 0 destroyed.
``` 