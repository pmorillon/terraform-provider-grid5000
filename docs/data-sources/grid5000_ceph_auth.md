# grid5000_ceph_auth Data Source

Get Grid'5000 Ceph auth

## Exemple Usage

```hcl
data "grid5000_ceph_auth" "auth" {
  site = "rennes"
  username = "username"
}

output "key" {
  value = data.grid5000_ceph_auth.auth.key
}
```

## Argument Reference

The following arguments are supported:

* `site` - (Required) The name of Grid'5000 site.
* `username` - (Optional) Grid'5000 username (Required for external access).

## Attribute Reference

The following attributes are exported:

* `key` - Cephx key.