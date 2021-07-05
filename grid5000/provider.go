package grid5000

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Provider provider schema
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The username to use for HTTP basic authentication when accessing Grid5000 API.",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The password to use for HTTP basic authentication when accessing Grid5000 API.",
			},
			"restfully_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The restfully config file path to use for HTTP basic authentication when accessing Grid5000 API.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"grid5000_job":        resourceJob(),
			"grid5000_deployment": resourceDeployment(),
			"grid5000_ceph_pool":  resourceCephPool(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"grid5000_site":          dataSourceGrid5000Site(),
			"grid5000_node":          dataSourceGrid5000Node(),
			"grid5000_vlan_nodelist": dataSourceGrid5000VlanNodelist(),
			"grid5000_ceph_auth":     dataSourceGrid5000CephAuth(),
			"grid5000_nodes_role":    dataSourceGrid5000NodesRole(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Username:      d.Get("username").(string),
		Password:      d.Get("password").(string),
		RestfullyFile: d.Get("restfully_file").(string),
	}

	return config.Client(), nil
}
