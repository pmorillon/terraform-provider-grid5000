package grid5000

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func grid5000NodeFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		// Filters
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"site": {
			Type:     schema.TypeString,
			Required: true,
		},

		// Out parameters
		"ip": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"ip6": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"primary_network_interface": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}
