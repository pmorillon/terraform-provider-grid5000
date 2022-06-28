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
		"network_adapters": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: grid5000NodeNetworkAdapterFields(),
			},
		},
	}
}

func grid5000NodeNetworkAdapterFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"device": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"mounted": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"mountable": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"enabled": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"ip": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"ip6": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"driver": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"vendor": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"switch": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"switch_port": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"model": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"mac": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"interface": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"rate": {
			Type:     schema.TypeInt,
			Computed: true,
		},
	}
}
