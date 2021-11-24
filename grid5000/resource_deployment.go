package grid5000

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gitlab.inria.fr/pmorillo/gog5k"
)

func resourceDeployment() *schema.Resource {
	return &schema.Resource{
		Create: resourceDeploymentCreate,
		Read:   resourceDeploymentRead,
		Update: resourceDeploymentUpdate,
		Delete: resourceDeploymentDelete,

		Schema: grid5000DeploymentFields(),
	}
}

func resourceDeploymentCreate(d *schema.ResourceData, m interface{}) error {
	site := d.Get("site").(string)

	client := m.(*gog5k.Client)
	ctx := context.Background()

	nodes := expandNodes(d.Get("nodes").(*schema.Set).List())
	createRequest := &gog5k.DeploymentCreateRequest{
		Environment: d.Get("environment").(string),
		Key:         d.Get("key").(string),
		Nodes:       nodes,
	}

	if p := d.Get("partition_number").(int); p != 0 {
		createRequest.PartitionNumber = p
	}

	if v := d.Get("vlan").(int); v != 0 {
		createRequest.Vlan = int16(v)
	}

	deployment, _, err := client.Deployments.Create(ctx, site, createRequest)
	if err != nil {
		return err
	}

	id := deployment.ID

	_, err = client.Deployments.WaitForState(ctx, site, id, gog5k.DeploymentTerminatedState)
	if err != nil {
		return err
	}

	d.SetId(id)

	return resourceDeploymentRead(d, m)
}

func resourceDeploymentRead(d *schema.ResourceData, m interface{}) error {
	if d.Id() != "" {
		client := m.(*gog5k.Client)
		ctx := context.Background()

		deployment, _, err := client.Deployments.Get(ctx, d.Get("site").(string), d.Id())
		if err != nil {
			return err
		}
		d.Set("state", deployment.Status)
	}

	return nil
}

func resourceDeploymentUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceDeploymentDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}

func expandNodes(nodes []interface{}) []string {
	expandedNodes := make([]string, len(nodes))
	for i, v := range nodes {
		expandedNodes[i] = v.(string)
	}

	return expandedNodes
}
