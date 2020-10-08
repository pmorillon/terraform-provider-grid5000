package grid5000

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Schemas

func grid5000JobFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"site": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"types": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"resources": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"reservation": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"command": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"scheduled_at_limit": &schema.Schema{
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
		"need_state": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "running",
		},
		"properties": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"state": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		"assigned_nodes": &schema.Schema{
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"vlans_resources": &schema.Schema{
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"subnets_resources": &schema.Schema{
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
