package grid5000

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gitlab.inria.fr/pmorillo/gog5k"
)

func resourceGrid5000JobImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	if len(d.Id()) == 0 {
		return nil, fmt.Errorf("import ID (<job_id>@<site>) is nil")
	}

	args := strings.Split(d.Id(), "@")
	if len(args) != 2 {
		return nil, fmt.Errorf("import ID bad format (<job_id>@<site>)")
	}

	jobID, _ := strconv.Atoi(args[0])
	site := args[1]

	client := m.(*gog5k.Client)
	ctx := context.Background()

	job, _, err := client.OARJobs.Get(ctx, site, int32(jobID))
	if err != nil {
		return nil, fmt.Errorf("failed to get job : %v", err)
	}

	if isJobAvailable(*job) {
		d.SetId(fmt.Sprint(job.ID))
		d.Set("site", site)
		d.Set("scheduled_at_limit", "5m")
		d.Set("command", job.Command)
		d.Set("state", job.State)
		d.Set("need_state", job.State)
		d.Set("assigned_nodes", job.AssignedNodes)
		d.Set("vlans_resources", job.ResourcesByType.Vlans)
		d.Set("subnets_resources", job.ResourcesByType.Subnets)
		disks := make([]interface{}, 0, len(job.ResourcesByType.Disks))
		for _, d := range job.ResourcesByType.Disks {
			s := strings.Split(d, ".")
			device := make(map[string]interface{}, 2)
			device["hostname"] = strings.Join(s[1:], ".")
			device["device"] = s[0]
			disks = append(disks, device)
		}
		if err := d.Set("disks_resources", disks); err != nil {
			log.Printf("Error : %v", err)
		}
	} else {
		return nil, fmt.Errorf("job is not running")
	}
	return []*schema.ResourceData{d}, nil
}
