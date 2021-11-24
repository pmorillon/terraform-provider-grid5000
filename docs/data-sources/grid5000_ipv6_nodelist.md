# grid5000_ipv6_nodelist Data Source

Tranform a node list to IPv6 FQDN node list.

## Example Usage

```hcl
data "grid5000_ipv6_nodelist" "ipv6list" {
  nodelist = grid5000_job.firewall.assigned_nodes 
}
```

Result :

```sh
❯ terraform state show data.grid5000_ipv6_nodelist.ipv6list 
# data.grid5000_ipv6_nodelist.ipv6list:
data "grid5000_ipv6_nodelist" "ipv6list" {
    id       = "2021-11-18 08:16:32.188333 +0000 UTC"
    nodelist = [
        "paravance-8.rennes.grid5000.fr",
    ]
    result   = [
        "paravance-8-ipv6.rennes.grid5000.fr",
    ]
}
```

## Argument Reference

The following arguments are supported :

* `nodelist` - (Require - List) A node list.

## Attribute Reference

The following attribute is exported :

* `result` - Node list with IPv6 adresses.