package grid5000

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

// Schemas

func grid5000JobDiskFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"hostname": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"device": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
	return s
}
