package grid5000

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gitlab.inria.fr/pmorillo/gog5k"
)

func resourceFirewall() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallCreate,
		Read:   resourceFirewallRead,
		Update: resourceFirewallUpdate,
		Delete: resourceFirewallDelete,
		Schema: grid5000FirewallFields(),
	}
}

func resourceFirewallCreate(d *schema.ResourceData, m interface{}) error {
	site := d.Get("site").(string)
	jobID := d.Get("job_id").(int)

	client := m.(*gog5k.Client)
	ctx := context.Background()

	addresses := expandNodes(d.Get("address").(*schema.Set).List())

	firewallRequest := &gog5k.FirewallCreateRequest{
		Address: addresses,
		Port:    d.Get("port").(string),
		//Protocol: d.Get("protocol").(string),
	}

	req := []*gog5k.FirewallCreateRequest{firewallRequest}

	_, err := client.Firewall.Create(ctx, site, int32(jobID), req)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%d", jobID))

	return nil
}

func resourceFirewallRead(d *schema.ResourceData, m interface{}) error {

	return nil
}

func resourceFirewallUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFirewallDelete(d *schema.ResourceData, m interface{}) error {
	if d.Id() != "" {
		client := m.(*gog5k.Client)
		ctx := context.Background()

		_, err := client.Firewall.Delete(ctx, d.Get("site").(string), int32(d.Get("job_id").(int)))
		if err != nil {
			return err
		}
	}
	d.SetId("")
	return nil
}
