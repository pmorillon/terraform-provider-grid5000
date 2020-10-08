# grid5000_vlan_nodelist Data Source

Tranform a node list by adding kavlan node list

## Example Usage

```hcl
data "grid5000_vlan_nodelist" "nodes" {
  nodelist = flatten(grid5000_job.nodes[*].assigned_nodes)
  vlan     = grid5000_job.vlan.vlans_resources[0]
}

output "grid5000_kavlan_nodelist" {
  value = data.grid5000_vlan_nodelist.nodes.result
}
```

Result :

```
Outputs:

grid5000_kavlan_nodelist = [
  "chetemi-1-kavlan-12.lille.grid5000.fr",
  "chetemi-8-kavlan-12.lille.grid5000.fr",
  "econome-12-kavlan-12.nantes.grid5000.fr",
  "econome-2-kavlan-12.nantes.grid5000.fr",
  "parapide-13-kavlan-12.rennes.grid5000.fr",
  "parapide-8-kavlan-12.rennes.grid5000.fr",
]
```

## Argument Reference

The following arguments are supported :

* `nodelist` - (Require - List) A node list.
* `vlan` - (Required - Int) Vlan ID.

## Attribute Reference

The following attribute is exported :

* `result` - Node list with kavlan adresses.