package grid5000

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

// Schemas

func grid5000DeploymentFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"site": {
			Type:     schema.TypeString,
			Required: true,
		},
		"environment": {
			Type:     schema.TypeString,
			Required: true,
		},
		"key": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"nodes": {
			Type:     schema.TypeSet,
			Required: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"partition_number": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"vlan": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"state": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}
