package grid5000

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gitlab.inria.fr/pmorillo/gog5k"
)

func dataSourceGrid5000CephAuth() *schema.Resource {
	return &schema.Resource{
		Read: datasourceGrid5000CephAuthRead,

		Schema: map[string]*schema.Schema{
			// Filters
			"site": {
				Type:     schema.TypeString,
				Required: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
			},

			// Out parameters
			"key": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
		},
	}
}

func datasourceGrid5000CephAuthRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*gog5k.Client)
	ctx := context.Background()

	if d.Get("username").(string) == "" {
		d.Set("username", os.Getenv("USER"))
	}

	auth, _, err := client.CSOD.GetAuth(ctx, d.Get("site").(string), d.Get("username").(string))
	if err != nil {
		return fmt.Errorf("failed to get node : %v", err)
	}

	d.Set("key", auth.Key)
	d.SetId(d.Get("username").(string))

	return nil
}
