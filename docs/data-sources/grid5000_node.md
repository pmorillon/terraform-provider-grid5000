# grid5000_node Data Source

Get Grid'5000 node information

## Example Usage

```hcl
data "grid5000_node" "parasilo-1" {
  site = "rennes"
  name = "parasilo-1"
}

output "grid5000_node_parasilo-1" {
  value = data.grid5000_node_parasilo-1.ip
}
```

## Argument Reference

The following arguments are supported:

* `site` - (Required) The name of Grid'5000 site.
* `name` - (Required) The node hostname. FQDN is supported.

## Attribute Reference

The following attributes are exported:

* `ip` - IPv4 address of main network adapter.
* `ip6` - IPv6 address of main network adapter.