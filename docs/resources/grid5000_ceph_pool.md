# grid5000_ceph_pool

Manage on demand Ceph pools.

## Exemple Usage

```hcl
resource "grid5000_ceph_pool" "k8s" {
  site = "rennes"
  username = "username"
  pool_name = "k8s"
  quota = "200G"
}
```

## Argument Reference

* `site` - (Required) A grid'5000 site.
* `pool_name` - (Required) Ceph pool name.
* `quota` - (Required) Ceph pool quota (format '<int>M|G|T').
* `username` - (Optional) Grid'5000 username (Required for external access).

## Attribute Reference

* `real_pool_name` - namespaced pool name.