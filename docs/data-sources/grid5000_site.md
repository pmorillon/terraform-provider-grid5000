# grid5000_site Data Source

Get Grid'5000 site information

## Example Usage

```hcl
data "grid5000_site" "rennes" {
  name = "rennes"
}

output "grid5000_site_rennes_frontend_ip" {
  value = data.grid5000_site.rennes.frontend_ip
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of Grid'5000 site.

## Attribute Reference

The following attributes are exported:

* `frontend_ip` - IP address of the frontend.
* `email_contact` - Contact email for support.
* `production` - Is site supporting production jobs.
* `renater_ip` - RENATER IP.
* `longitude` - Longitude position.
* `latitude` - Latitude position.
* `description` - Site description.
* `location` - Site location.