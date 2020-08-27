package grid5000

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

// Schemas

func grid5000DeploymentFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"site": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"environment": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"key": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"nodes": &schema.Schema{
			Type:     schema.TypeSet,
			Required: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"state": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}
