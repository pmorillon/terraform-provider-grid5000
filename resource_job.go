package main

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gitlab.inria.fr/pmorillo/gog5k"
)

func resourceJob() *schema.Resource {
	return &schema.Resource{
		Create: resourceJobCreate,
		Read:   resourceJobRead,
		Update: resourceJobUpdate,
		Delete: resourceJobDelete,

		Schema: map[string]*schema.Schema{
			"site": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"types": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"resources": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"command": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"assigned_nodes": &schema.Schema{
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"vlans_resources": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"subnets_resources": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"disks_resources": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceJobCreate(d *schema.ResourceData, m interface{}) error {
	site := d.Get("site").(string)

	client := m.(*gog5k.Client)
	ctx := context.Background()

	createRequest := &gog5k.OARJobCreateRequest{
		Name:      d.Get("name").(string),
		Command:   d.Get("command").(string),
		Resources: d.Get("resources").(string),
		Types:     expandTypes(d.Get("types").(*schema.Set).List()),
	}

	job, _, err := client.OARJobs.Create(ctx, site, createRequest)
	if err != nil {
		return err
	}

	jobID := job.ID

	job, err = client.OARJobs.WaitForState(ctx, site, jobID, gog5k.OARJobRunningState)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprint(job.ID))
	log.Printf("JOB ID : %d, resource ID : %s", job.ID, d.Id())

	return resourceJobRead(d, m)
}

func resourceJobRead(d *schema.ResourceData, m interface{}) error {
	if d.Id() != "" {
		client := m.(*gog5k.Client)
		ctx := context.Background()

		jobID, _ := strconv.Atoi(d.Id())
		job, _, err := client.OARJobs.Get(ctx, d.Get("site").(string), int32(jobID))
		if err != nil {
			return err
		}
		if isJobAvailable(*job) {
			d.SetId(fmt.Sprint(job.ID))
			d.Set("state", job.State)
			d.Set("assigned_nodes", job.AssignedNodes)
			d.Set("vlans_resources", job.ResourcesByType.Vlans)
			d.Set("subnets_resources", job.ResourcesByType.Subnets)
			d.Set("disks_resources", job.ResourcesByType.Disks)
		} else {
			d.SetId("")
		}
	}

	return nil
}

func resourceJobUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceJobRead(d, m)
}

func resourceJobDelete(d *schema.ResourceData, m interface{}) error {
	if d.Id() != "" {
		client := m.(*gog5k.Client)
		ctx := context.Background()

		jobID, _ := strconv.Atoi(d.Id())

		job, _, err := client.OARJobs.Get(ctx, d.Get("site").(string), int32(jobID))
		if err != nil {
			return err
		}
		var unvalidState = regexp.MustCompile(`(terminated|error)`)

		if !(unvalidState.MatchString(job.State)) {
			_, err := client.OARJobs.Delete(ctx, d.Get("site").(string), int32(jobID))
			if err != nil {
				return err
			}
		}
	}
	d.SetId("")

	return nil
}

func expandTypes(types []interface{}) []string {
	expandedTypes := make([]string, len(types))
	for i, v := range types {
		expandedTypes[i] = v.(string)
	}

	return expandedTypes
}

func isJobAvailable(job gog5k.OARJob) bool {
	if (job.State == gog5k.OARJobTerminatedState) || (job.State == gog5k.OARJobErrorState) {
		return false
	}
	return true
}
