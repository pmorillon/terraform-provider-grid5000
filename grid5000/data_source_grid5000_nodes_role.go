package grid5000

import (
	"regexp"

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
	free_nodes := d.Get("from_list").(*schema.Set).List()
	size := d.Get("size").(int)
	pattern := d.Get("pattern").(string)
	new_list := make([]interface{}, 0)
	if pattern != "" {
		for _, v := range from_list {
			matched, err := regexp.MatchString(pattern, v.(string))
			if err != nil {
				return err
			}
			if matched {
				new_list = append(new_list, v.(string))
			}
		}
		new_list = new_list[0:size]
	} else {
		new_list = from_list[0:size]
	}

	for _, v := range new_list {
		for i2, v2 := range free_nodes {
			if v2.(string) == v.(string) {
				free_nodes = removeAtIndex(free_nodes, i2)
				break
			}
		}
	}

	d.Set("nodes", new_list)
	d.Set("free_nodes", free_nodes)
	d.SetId(d.Get("name").(string))

	return nil
}

func removeAtIndex(s []interface{}, index int) []interface{} {
	return append(s[:index], s[index+1:]...)
}
