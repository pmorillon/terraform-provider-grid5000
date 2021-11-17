package grid5000

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

// Schemas

func grid5000FirewallFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"site": {
			Type:     schema.TypeString,
			Required: true,
		},
		"address": {
			Type:     schema.TypeSet,
			Required: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"port": {
			Type:     schema.TypeString,
			Required: true,
		},
		"protocol": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "all",
		},
		"source": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"job_id": {
			Type:     schema.TypeInt,
			Required: true,
		},
	}
}
