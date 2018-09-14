package minimum

import (
	"encoding/json"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/yamamoto-febc/terraform-provider-minimum/state"
)

var client = state.NewDriver("basic")

func resourceMinimumBasic() *schema.Resource {
	return &schema.Resource{
		Create: resourceMinimumBasicCreate,
		Read:   resourceMinimumBasicRead,
		Update: resourceMinimumBasicUpdate,
		Delete: resourceMinimumBasicDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceMinimumBasicCreate(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)
	value := &state.BasicValue{
		Name: name,
	}

	id, err := client.Create(value)
	if err != nil {
		return err
	}

	d.SetId(id)
	return resourceMinimumBasicRead(d, meta)
}

func resourceMinimumBasicRead(d *schema.ResourceData, _ interface{}) error {
	id := d.Id()
	data, err := client.Read(id)
	if err != nil {
		if state.IsStateNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return err
	}

	value := &state.BasicValue{}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	d.Set("name", value.Name) // nolint

	return nil
}

func resourceMinimumBasicUpdate(d *schema.ResourceData, meta interface{}) error {
	id := d.Id()

	// read current state
	data, err := client.Read(id)
	if err != nil {
		if state.IsStateNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return err
	}

	value := &state.BasicValue{}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	// set updated values
	if d.HasChange("name") {
		value.Name = d.Get("name").(string)
	}

	// save values
	if err := client.Update(id, value); err != nil {
		return err
	}

	return resourceMinimumBasicRead(d, meta)
}

func resourceMinimumBasicDelete(d *schema.ResourceData, _ interface{}) error {
	id := d.Id()
	if err := client.Delete(id); err != nil {
		if !state.IsStateNotFoundError(err) {
			return err
		}
	}
	d.SetId("")
	return nil
}
