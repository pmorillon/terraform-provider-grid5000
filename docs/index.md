# Grid5000 Provider

Terraform provider plugin for [OAR](https://oar.imag.fr) jobs submission and [Kadeploy](http://kadeploy3.gforge.inria.fr) bare-metal deployments on [Grid'5000](https://www.grid5000.fr). Grid'5000 is a large-scale and flexible testbed for experiment-driven research in all areas of computer science, with a focus on parallel and distributed computing including Cloud, HPC, Big Data and AI.

This provider is not maintained by the Grid'5000 Team.

## Example Usage

### Internal usage (frontend of one Grid5000 site)

Nothing to configure.

### External usage

```hcl
provider "grid5000" {
  # API credentials for external usages
  username = "john"
  password = "xxx"
}
```

Or

```hcl
provider "grid5000" {
  # Restfully configuration file location
  restfully_file = "/home/user/.restfully/api.grid5000.fr.yml"
}
```

## Argument Reference

* `username` - (Optional) The username to use for HTTP basic authentication when accessing Grid5000 API.
* `password` - (Optional) The password to use for HTTP basic authentication when accessing Grid5000 API.
* `restfully_file` - (Optional) The restfully config file path to use for authentication when accessing Grid5000 API.