package grid5000

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceGrid5000VlanNodelist() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGrid5000VlanNodelistRead,

		Schema: map[string]*schema.Schema{
			// Filters
			"nodelist": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"vlan": {
				Type:     schema.TypeInt,
				Required: true,
			},
			// Out parameters
			"result": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceGrid5000VlanNodelistRead(d *schema.ResourceData, m interface{}) error {
	nodeList := expandNodes(d.Get("nodelist").(*schema.Set).List())
	vlanNumber := d.Get("vlan").(int)

	result := make([]string, 0, len(nodeList))
	for _, n := range nodeList {
		s := strings.Split(n, ".")
		r := fmt.Sprintf("%s-kavlan-%d", s[0], vlanNumber)
		s[0] = r
		result = append(result, strings.Join(s, "."))
	}

	d.SetId(time.Now().UTC().String())
	if err := d.Set("result", result); err != nil {
		log.Printf("Error : %v", err)
	}
	return nil
}
