package grid5000

import (
	"context"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gitlab.inria.fr/pmorillo/gog5k"
)

func resourceJob() *schema.Resource {
	return &schema.Resource{
		Create: resourceJobCreate,
		Read:   resourceJobRead,
		Update: resourceJobUpdate,
		Delete: resourceJobDelete,

		Schema: grid5000JobFields(),
	}
}

func resourceJobCreate(d *schema.ResourceData, m interface{}) error {
	site := d.Get("site").(string)

	client := m.(*gog5k.Client)
	ctx := context.Background()

	createRequest := &gog5k.OARJobCreateRequest{
		Name:       d.Get("name").(string),
		Command:    d.Get("command").(string),
		Resources:  d.Get("resources").(string),
		Properties: d.Get("properties").(string),
		Types:      expandTypes(d.Get("types").(*schema.Set).List()),
	}

	advanceReservation := d.Get("reservation").(string)

	if advanceReservation != "" {
		createRequest.Reservation = advanceReservation
	}

	job, _, err := client.OARJobs.Create(ctx, site, createRequest)
	if err != nil {
		return err
	}

	jobID := job.ID

	for {
		job, _, err = client.OARJobs.Get(ctx, site, jobID)
		if job.ScheduledAt != 0 {
			break
		}
		time.Sleep(5 * time.Second)
	}

	if advanceReservation == "" {
		isValid, err := validateScheduling(job.ScheduledAt, d.Get("scheduled_at_limit").(string))
		if err != nil {
			return err
		}

		if !isValid {
			_, err = client.OARJobs.Delete(ctx, site, jobID)
			if err != nil {
				return err
			}
			return errors.New("OAR job will not be scheduled at time, job is deleted")
		}

		job, err = client.OARJobs.WaitForState(ctx, site, jobID, gog5k.OARJobRunningState)
		if err != nil {
			return err
		}
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
			switch job.State {
			case gog5k.OARJobErrorState, gog5k.OARJobTerminatedState:
				d.SetId("")
				return fmt.Errorf("oar job %d is in %s state", job.ID, job.State)
			case gog5k.OARJobWaitingState:
				d.SetId(fmt.Sprint(job.ID))
				d.Set("need_state", "waiting")
				return fmt.Errorf("oar job %d is in %s state, it will be scheduled at %d, restart terraform apply later", job.ID, job.State, job.ScheduledAt)
			}
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
	if (job.State == gog5k.OARJobTerminatedState) || (job.State == gog5k.OARJobErrorState) || (job.State == gog5k.OARJobWaitingState) {
		return false
	}
	return true
}

func durationParse(duration string) (value string, unit string, err error) {
	validFormat := regexp.MustCompile(`^(\d+)(s|m|h)$`)
	if validFormat.MatchString(duration) {
		res := validFormat.FindAllStringSubmatch(duration, -1)
		return res[0][1], res[0][2], nil
	}
	return "", "", errors.New("duration parsing error")
}

func validateScheduling(scheduledAt int32, limit string) (bool, error) {
	value, unit, err := durationParse(limit)
	if err != nil {
		return false, err
	}
	valueInt64, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return false, err
	}
	var limitInSeconds int64 = 0
	switch unit {
	case "s":
		limitInSeconds = valueInt64
	case "m":
		limitInSeconds = valueInt64 * 60
	case "h":
		limitInSeconds = valueInt64 * 60 * 60
	default:
		log.Printf("Unrecognize unit : %s", unit)
	}
	now := time.Now()
	sec := now.Unix()

	return int64(scheduledAt) < (sec + limitInSeconds), nil
}
