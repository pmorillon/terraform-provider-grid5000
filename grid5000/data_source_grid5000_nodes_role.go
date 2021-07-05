package grid5000

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceGrid5000NodesRole() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			// Filters
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"from_list": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"size": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"pattern": {
				Type:     schema.TypeString,
				Optional: true,
			},

			// Out parameters
			"nodes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"free_nodes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},

		Read: dataSourceGrid5000NodesRoleRead,
	}
}

func dataSourceGrid5000NodesRoleRead(d *schema.ResourceData, m interface{}) error {
	from_list := d.Get("from_list").(*schema.Set).List()
	size := d.Get("size").(int)
	d.Set("nodes", from_list[0:size])
	d.Set("free_nodes", from_list[size:])
	d.SetId(d.Get("name").(string))
	return nil
}
