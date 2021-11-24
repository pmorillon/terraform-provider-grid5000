# grid5000_nodes_role Data Source

Manage roles based on subset of node list.

## Example Usage

```hcl
data "grid5000_nodes_role" "frontend" {
  name = "frontend"
  from_list = grid5000_job.frontend.assigned_nodes
  size = 1
}
```

## Argument Reference

* `name` - (Required)
* `from_list` - (Required)
* `size` - (Required)
* `pattern` - (Optional) Selection pattern regexp.

## Attribute Reference

The following attribute is exported :

* `nodes` - Resulting subset of nodes.
* `free_nodes` - Subset of nodes not selected.