# Terraform provider for Grid5000

## Requirements

* [Terraform](https://www.terraform.io/downloads.html) >= 0.12.x

## Installation

### Terraform >= 0.13.x

Add this into your Terraform configuration, then run `terraform init` :

```tf
terraform {
  required_providers {
    grid5000 = {
      source = "pmorillon/grid5000"
      version = "0.0.8"
    }
  }
}
```

### Terraform = 0.12.x

Downloadable packages are available at [Grid5000 provider releases](https://gitlab.inria.fr/pmorillo/terraform-provider-grid5000/-/releases).

How to install:
* Get binary for your platform.
* Place provider binary (into the archive) on your terraform plugin directory :

| Directory | Purpose |
|-|-|
| . | In case the provider is only used in a single Terraform project. |
| ~/.terraform.d/plugins | The user plugins directory. |

## Usage

* Provider configuration for External usage :

```tf
provider "grid5000" {
  # Basic Auth
  username = "john"
  password = "xxx"
}
```
Or :

```tf
provider "grid5000" {
  # Uses Restfully credentials
  restfully_file = "/home/user/.restfully/api.grid5000.fr.yml"
}
```

* OAR job

```tf
resource "grid5000_job" "my_job" {
  name      = "terraform"
  site      = "rennes"
  command   = "sleep 1h"
  resources = "/nodes=2"
  types     = ["deploy"]
}
```

* Kadeploy deployment

```tf
resource "grid5000_deployment" "my_deployment" {
  site        = "rennes"
  environment = "debian10-x64-base"
  nodes       = grid5000_job.my_job.assigned_nodes
  key         = file("~/.ssh/id_rsa.pub")
}
```
