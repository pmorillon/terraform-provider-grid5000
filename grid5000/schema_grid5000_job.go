package grid5000

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Schemas

func grid5000JobFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"site": {
			Type:     schema.TypeString,
			Required: true,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"types": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"resources": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"reservation": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"command": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"scheduled_at_limit": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "5m",
			ValidateFunc: func(i interface{}, k string) (ws []string, errors []error) {
				if _, _, err := durationParse(i.(string)); err != nil {
					errors = append(errors, fmt.Errorf("%q: invalid value, must use the format '5m', suffix may be 's' for seconds, 'm' for minutes, 'h' for hours", k))
				}
				return
			},
		},
		"need_state": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "running",
		},
		"properties": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"state": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"assigned_nodes": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"vlans_resources": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"subnets_resources": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"disks_resources": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: grid5000JobDiskFields(),
			},
		},
	}
}
