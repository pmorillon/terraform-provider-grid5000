# grid5000_job Resource

Manage [OAR](https://oar.imag.fr) job submission on Grid'5000.

## Example Usage

```hcl
resource "grid5000_job" "my_job" {
  name      = "terraform"
  site      = "rennes"
  command   = "sleep 1h"
  resources = "/nodes=2"
  types     = ["deploy"]
}
```

## Argument Reference

* `site` - (Required) A grid'5000 site.
* `command` - (Required) Script or command to execute.
* `name` - (Optional) Name of OAR job.
* `resources` - (Optional) Specify wanted resources with [OAR resources request expression](http://oar.imag.fr/docs/2.5/user/usecases.html#using-the-resource-hierarchy).
* `types` - (Optional) OAR job type.
* `properties` - (Optional) OAR job properties.
* `scheduled_at_limit` - (Optional, default: 5m) Set the limit for scheduling time. Suffix may be 's' for seconds, 'm' for minutes, 'h' for hours. Ignored when reservation argument is used.
* `reservation` - (Optional) Request a job start time reservation, instead of a submission. The date format is "YYYY-MM-DD HH:MM:SS". Warning : the grid5000_job resource will be tainted automatically. When job is started, you need to untaint the resource before apply.

## Attribute Reference

* `state` - OAR job state.
* `assigned_nodes` - List of nodes hostnames assigned to this OAR job.
* `disks_resources` - List of disks resources assigned to this job.
* `subnets_resources` - List of subnets resources assigned to this job.
* `vlans_resources` - List of vlans resources assigned to this job.

## Nested blocks

### `disks_resources`

#### Attributes

* `hostname` - Node hostname
* `device` - Disk device name