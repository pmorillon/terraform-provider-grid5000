# grid5000_firewall

Manage Grid'5000 reconfigurable firewall, see [Grid'5000 service documentation](https://www.grid5000.fr/w/Reconfigurable_Firewall). 

## Example usage

```hcl
data "grid5000_ipv6_nodelist" "ipv6list" {
  nodelist = grid5000_job.firewall.assigned_nodes 
}

resource "grid5000_firewall" "f1" {
    site = local.site
    job_id = grid5000_job.firewall.id
    dest = data.grid5000_ipv6_nodelist.ipv6list.result
    ports = [
      80
    ]
}
```

```sh
❯ terraform state show grid5000_firewall.f1
# grid5000_firewall.f1:
resource "grid5000_firewall" "f1" {
    dest     = [
        "paravance-8-ipv6.rennes.grid5000.fr",
    ]
    id       = "1828028"
    job_id   = 1828028
    ports    = [
        80,
    ]
    protocol = "tcp+udp"
    site     = "rennes"
}
```

## Argument Reference

* `site` - (Required) A grid'5000 site.
* `job_id` - (Required) OAR job ID.
* `dest` - (Required) Set of IPv6 destination addresses.
* `src` - (Optional) Set of IPv6 source addresses.
* `ports` - (Optional) Set of opened ports. Not used if protocal argument is __all__.
* `protocol` - (Optional) __tcp+udp__ by default. 
