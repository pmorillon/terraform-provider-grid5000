package grid5000

import (
	"context"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gitlab.inria.fr/pmorillo/gog5k"
)

func resourceCephPool() *schema.Resource {
	return &schema.Resource{
		Read:   resourceCephPoolRead,
		Create: resourceCephPoolCreate,
		Delete: resourceCephPoolDelete,
		Update: resourceCephPoolUpdate,

		Schema: grid5000CephPoolFields(),
	}
}

func resourceCephPoolRead(d *schema.ResourceData, m interface{}) error {
	if d.Id() != "" {
		client := m.(*gog5k.Client)
		ctx := context.Background()

		if d.Get("username").(string) == "" {
			d.Set("username", os.Getenv("USER"))
		}

		d.Set("real_pool_name", fmt.Sprintf("%s_%s", d.Get("username").(string), d.Get("pool_name").(string)))
		pool, _, err := client.CSOD.GetPool(ctx, d.Get("site").(string), d.Get("real_pool_name").(string))
		if err != nil {
			return err
		}
		if pool != nil {
			d.SetId(d.Get("pool_name").(string))
			d.Set("real_pool_name", pool.PoolName)
			//d.Set("quota", pool.QuotaMaxBytes)
		}
	}
	return nil
}

func resourceCephPoolCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*gog5k.Client)
	ctx := context.Background()

	if d.Get("username").(string) == "" {
		d.Set("username", os.Getenv("USER"))
	}

	d.Set("real_pool_name", fmt.Sprintf("%s_%s", d.Get("username").(string), d.Get("pool_name").(string)))
	pool, _, err := client.CSOD.GetPool(ctx, d.Get("site").(string), d.Get("real_pool_name").(string))
	if err != nil {
		return err
	}
	if pool != nil {
		d.SetId(d.Get("pool_name").(string))
	} else {
		quota, err := quotaToBytes(d.Get("quota").(string))
		if err != nil {
			return err
		}
		createPoolRequest := &gog5k.CSODPoolCreateRequest{
			PoolName:       d.Get("pool_name").(string),
			QuotaMaxBytes:  quota,
			ExpirationDate: time.Now().AddDate(0, 1, 0),
			Size:           3,
		}
		_, err = client.CSOD.CreatePool(ctx, d.Get("site").(string), createPoolRequest)
		if err != nil {
			return err
		}
		d.SetId(d.Get("pool_name").(string))
	}
	return nil
}

func resourceCephPoolDelete(d *schema.ResourceData, m interface{}) error {
	if d.Id() != "" {
		client := m.(*gog5k.Client)
		ctx := context.Background()

		_, err := client.CSOD.DeletePool(ctx, d.Get("site").(string), d.Get("real_pool_name").(string))
		if err != nil {
			return err
		}
	}
	d.SetId("")
	return nil
}

func resourceCephPoolUpdate(d *schema.ResourceData, m interface{}) error {
	// client := m.(*gog5k.Client)
	// ctx := context.Background()

	// quota, err := quotaToBytes(d.Get("quota").(string))
	// if err != nil {
	// 	return err
	// }
	// createPoolRequest := &gog5k.CSODPool{
	// 	PoolName:       d.Get("real_pool_name").(string),
	// 	QuotaMaxBytes:  quota,
	// 	ExpirationDate: time.Now().AddDate(0, 1, 0),
	// 	Size:           3,
	// }
	// _, err = client.CSOD.UpdatePool(ctx, d.Get("site").(string), createPoolRequest)
	// if err != nil {
	// 	return err
	// }
	// d.SetId(d.Get("pool_name").(string))
	return nil
}

func quotaParse(quota string) (value string, unit string, err error) {
	validFormat := regexp.MustCompile(`^(\d+)(M|G|T)$`)
	if validFormat.MatchString(quota) {
		res := validFormat.FindAllStringSubmatch(quota, -1)
		return res[0][1], res[0][2], nil
	}
	return "", "", errors.New("quota parsing error")
}

func quotaToBytes(quota string) (bytes int64, err error) {
	value, unit, err := quotaParse(quota)
	if err != nil {
		return 0, err
	}
	valueInt64, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, err
	}
	switch unit {
	case "M":
		bytes = valueInt64 * 1024 * 1024
	case "G":
		bytes = valueInt64 * 1024 * 1024 * 1024
	case "T":
		bytes = valueInt64 * 1024 * 1024 * 1024 * 1024
	}
	return bytes, nil
}
