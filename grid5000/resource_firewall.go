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

	rules := d.Get("rule").([]interface{})
	req := []*gog5k.Firewall{}

	if len(rules) > 0 {
		for _, r := range rules {
			firewallRuleRequest := &gog5k.Firewall{}
			rule := r.(map[string]interface{})

			firewallRuleRequest.Dest = expandNodes(rule["dest"].(*schema.Set).List())
			firewallRuleRequest.Src = expandNodes(rule["src"].(*schema.Set).List())

			if rule["protocol"].(string) != "tcp+udp" {
				firewallRuleRequest.Protocol = rule["protocol"].(string)
			}

			if rule["protocol"].(string) != "all" {
				firewallRuleRequest.Ports = expandInt(rule["ports"].(*schema.Set).List())
			}

			req = append(req, firewallRuleRequest)
		}
	}

	_, err := client.Firewall.Create(ctx, site, int32(jobID), req)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%d", jobID))

	return resourceFirewallRead(d, m)
}

func resourceFirewallRead(d *schema.ResourceData, m interface{}) error {
	// if d.Id() != "" {
	// 	client := m.(*gog5k.Client)
	// 	ctx := context.Background()
	// 	jobID, _ := strconv.Atoi(d.Id())

	// 	firewall, _, err := client.Firewall.Get(ctx, d.Get("site").(string), int32(jobID))
	// 	if err != nil {
	// 		return err
	// 	}

	// 	rules := make([]interface{}, len(firewall))

	// 	for _, rule := range firewall {
	// 		r := make(map[string]interface{})
	// 		r["dest"] = rule.Dest
	// 		r["protocol"] = rule.Protocol
	// 		r["src"] = rule.Src
	// 		r["ports"] = rule.Ports
	// 		rules = append(rules, r)
	// 	}
	// 	if err := d.Set("rule", rules); err != nil {
	// 		log.Printf("[ERROR] %v", err)
	// 	}
	// }

	return nil
}

func resourceFirewallUpdate(d *schema.ResourceData, m interface{}) error {
	resourceFirewallDelete(d, m)
	resourceFirewallCreate(d, m)

	return resourceFirewallRead(d, m)
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

func expandInt(nodes []interface{}) []int {
	expandedInt := make([]int, len(nodes))
	for i, v := range nodes {
		expandedInt[i] = v.(int)
	}

	return expandedInt
}
