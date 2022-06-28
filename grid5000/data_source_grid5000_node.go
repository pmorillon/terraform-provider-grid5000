package grid5000

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gitlab.inria.fr/pmorillo/gog5k"
)

func dataSourceGrid5000Node() *schema.Resource {
	return &schema.Resource{
		Read: datasourceGrid5000NodeRead,

		Schema: grid5000NodeFields(),
	}
}

func datasourceGrid5000NodeRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*gog5k.Client)
	ctx := context.Background()

	node, _, err := client.Nodes.Get(ctx, d.Get("site").(string), shortHostname(d.Get("name").(string)))
	if err != nil {
		return fmt.Errorf("failed to get node : %v", err)
	}

	d.SetId(node.UID)
	// Find main network adapter
	for _, n := range node.NetworkAdapters {
		if n.Mounted {
			d.Set("ip", n.IP)
			d.Set("ip6", n.IP6)
			d.Set("primary_network_interface", n.Name)
			break
		}
	}

	network_adapters := make([]interface{}, 0)
	for _, netif := range node.NetworkAdapters {
		n := make(map[string]interface{})
		n["name"] = netif.Name
		n["device"] = netif.Device
		n["mounted"] = netif.Mounted
		n["mountable"] = netif.Mountable
		n["enabled"] = netif.Enabled
		n["ip"] = netif.IP
		n["ip6"] = netif.IP6
		n["driver"] = netif.Driver
		n["vendor"] = netif.Vendor
		n["switch"] = netif.Switch
		n["switch_port"] = netif.SwitchPort
		n["model"] = netif.Model
		n["interface"] = netif.Interface
		n["mac"] = netif.MAC
		n["rate"] = netif.Rate
		network_adapters = append(network_adapters, n)
	}
	if err := d.Set("network_adapters", network_adapters); err != nil {
		log.Printf("[ERROR] %v\n", err)
	}

	return nil
}

func shortHostname(hostname string) string {
	return strings.Split(hostname, ".")[0]
}
