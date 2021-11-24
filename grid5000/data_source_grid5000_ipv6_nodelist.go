package grid5000

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceGrid5000IPv6Nodelist() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGrid5000IPv6NodelistRead,
		Schema: map[string]*schema.Schema{
			// Filter
			"nodelist": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

func dataSourceGrid5000IPv6NodelistRead(d *schema.ResourceData, m interface{}) error {
	nodeList := expandNodes(d.Get("nodelist").(*schema.Set).List())

	result := make([]string, 0, len(nodeList))
	for _, n := range nodeList {
		s := strings.Split(n, ".")
		r := fmt.Sprintf("%s-ipv6", s[0])
		s[0] = r
		result = append(result, strings.Join(s, "."))
	}

	d.SetId(time.Now().UTC().String())
	if err := d.Set("result", result); err != nil {
		log.Printf("Error : %v", err)
	}
	return nil
}
