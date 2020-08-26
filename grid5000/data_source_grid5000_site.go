package grid5000

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gitlab.inria.fr/pmorillo/gog5k"
)

func dataSourceGrid5000Site() *schema.Resource {
	return &schema.Resource{
		Read: datasourceGrid5000SiteRead,

		Schema: map[string]*schema.Schema{
			// Filters
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Out parameters
			"frontend_ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"email_contact": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"production": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"renater_ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"longitude": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"latitude": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"location": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func datasourceGrid5000SiteRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*gog5k.Client)
	ctx := context.Background()

	site, _, err := client.Sites.GetByName(ctx, d.Get("name").(string))
	if err != nil {
		return fmt.Errorf("Failed to get site : %v", err)
	}

	d.SetId(site.UID)
	d.Set("frontend_ip", site.FrontendIP)
	d.Set("email_contact", site.EmailContact)
	d.Set("production", site.Production)
	d.Set("renater_ip", site.RenaterIP)
	d.Set("longitude", site.Longitude)
	d.Set("latitude", site.Latitude)
	d.Set("description", site.Description)
	d.Set("location", site.Location)

	return nil
}
