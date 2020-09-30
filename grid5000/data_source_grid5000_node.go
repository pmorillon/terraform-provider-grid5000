package grid5000

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gitlab.inria.fr/pmorillo/gog5k"
)

func dataSourceGrid5000Node() *schema.Resource {
	return &schema.Resource{
		Read: datasourceGrid5000NodeRead,

		Schema: map[string]*schema.Schema{
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
		},
	}
}

func datasourceGrid5000NodeRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*gog5k.Client)
	ctx := context.Background()

	node, _, err := client.Nodes.Get(ctx, d.Get("site").(string), shortHostname(d.Get("name").(string)))
	if err != nil {
		return fmt.Errorf("Failed to get node : %v", err)
	}

	d.SetId(node.UID)
	// Find main network adapter
	for _, n := range node.NetworkAdapters {
		if n.Mounted {
			d.Set("ip", n.IP)
			d.Set("ip6", n.IP6)
			break
		}
	}

	return nil
}

func shortHostname(hostname string) string {
	return strings.Split(hostname, ".")[0]
}
