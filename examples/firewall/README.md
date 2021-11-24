# How to use Grid5000 Firewall API with Terraform

This example will :
* Reserve a server
* Debian 10 bare-metal deployment on this server
* Install NGINX and configure IPv6 on the primary network interface
* Open 80,443 ports

## How to use

### Requirements

* [terraform](https://terraform.io/) on a Grid'5000 frontend

### Apply Terraform config

On a Grid'5000 frontend :

```sh
# Clone this repo
❯ git clone https://github.com/pmorillon/terraform-provider-grid5000.git
❯ cd terraform-provider-grid5000/examples/firewall


# init : will automatically download required Terraform providers
❯ terraform init

Initializing the backend...

Initializing provider plugins...
- Finding pmorillon/grid5000 versions matching "~> 0.0.8"...
- Finding latest version of hashicorp/null...
- Installing pmorillon/grid5000 v0.0.8...
- Installed pmorillon/grid5000 v0.0.8 (self-signed, key ID 01D0BAEDBCC797AE)
- Installing hashicorp/null v3.1.0...
- Installed hashicorp/null v3.1.0 (self-signed, key ID 34365D9472D7468F)

# ...

Terraform has been successfully initialized!
```

```sh
❯ terraform apply
# ...
grid5000_job.job1: Creating...
grid5000_job.job1: Still creating... [10s elapsed]
grid5000_job.job1: Still creating... [20s elapsed]
grid5000_job.job1: Creation complete after 23s [id=1846629]
data.grid5000_ipv6_nodelist.ipv6list: Reading...
data.grid5000_node.node: Reading...
data.grid5000_ipv6_nodelist.ipv6list: Read complete after 0s [id=2021-11-24 15:16:06.163297829 +0000 UTC]
grid5000_deployment.debian: Creating...
data.grid5000_node.node: Read complete after 0s [id=paravance-71]
grid5000_deployment.debian: Still creating... [10s elapsed]
grid5000_deployment.debian: Still creating... [20s elapsed]
# ...
grid5000_deployment.debian: Creation complete after 3m11s [id=D-a00c6e2b-a458-4070-9a1e-929aa6bf6bb4]
null_resource.nginx_install[0]: Creating...
null_resource.nginx_install[0]: Provisioning with 'remote-exec'...
# ...
null_resource.nginx_install[0]: Still creating... [10s elapsed]
null_resource.nginx_install[0]: Creation complete after 15s [id=4718359704867177795]
grid5000_firewall.f1: Creating...
grid5000_firewall.f1: Creation complete after 2s [id=1846629]

Apply complete! Resources: 4 added, 0 changed, 0 destroyed.
```

```sh
❯ terraform state show grid5000_firewall.f1             
# grid5000_firewall.f1:
resource "grid5000_firewall" "f1" {
    id     = "1846629"
    job_id = 1846629
    site   = "rennes"

    rule {
        dest     = [
            "paravance-71-ipv6.rennes.grid5000.fr",
        ]
        ports    = [
            80,
        ]
        protocol = "tcp+udp"
    }
    rule {
        dest     = [
            "paravance-71-ipv6.rennes.grid5000.fr",
        ]
        ports    = [
            443,
        ]
        protocol = "tcp+udp"
    }
}
```

```sh
❯ curl -XGET https://api.grid5000.fr/stable/sites/rennes/firewall/1846629 
[
  {
    "addr": [
      "2001:660:4406:700:1::47/128"
    ],
    "port": [
      80
    ],
    "proto": "tcp+udp"
  },
  {
    "addr": [
      "2001:660:4406:700:1::47/128"
    ],
    "port": [
      443
    ],
    "proto": "tcp+udp"
  }
]
```

### Destroy

```sh
❯ terraform destroy
# ...
grid5000_firewall.f1: Destroying... [id=1846629]
grid5000_firewall.f1: Destruction complete after 4s
null_resource.nginx_install[0]: Destroying... [id=4718359704867177795]
null_resource.nginx_install[0]: Destruction complete after 0s
grid5000_deployment.debian: Destroying... [id=D-a00c6e2b-a458-4070-9a1e-929aa6bf6bb4]
grid5000_deployment.debian: Destruction complete after 0s
grid5000_job.job1: Destroying... [id=1846629]
grid5000_job.job1: Destruction complete after 1s

Destroy complete! Resources: 4 destroyed.
```