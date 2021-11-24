package grid5000

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

// Schemas

func grid5000FirewallFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"site": {
			Type:     schema.TypeString,
			Required: true,
		},
		"job_id": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"rule": {
			Type:     schema.TypeList,
			MinItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: grid5000FirewallRule(),
			},
		},
	}
}

func grid5000FirewallRule() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"dest": {
			Type:     schema.TypeSet,
			Required: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"ports": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
		},
		"protocol": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "tcp+udp",
		},
		"src": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}
