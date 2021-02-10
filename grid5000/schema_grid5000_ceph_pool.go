package grid5000

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Schemas
func grid5000CephPoolFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"pool_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"site": {
			Type:     schema.TypeString,
			Required: true,
		},
		"username": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"quota": {
			Type:     schema.TypeString,
			Required: true,
			ValidateFunc: func(i interface{}, k string) (ws []string, errors []error) {
				if _, _, err := quotaParse(i.(string)); err != nil {
					errors = append(errors, fmt.Errorf("%q: invalid value, must use format '<int>M|G|T'", k))
				}
				return
			},
		},
		"real_pool_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
	return s
}
