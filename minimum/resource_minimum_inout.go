package minimum

import (
	"encoding/json"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/yamamoto-febc/terraform-provider-minimum/state"
)

var inOutDriver = state.NewDriver("inout")

func resourceMinimumInOut() *schema.Resource {
	return &schema.Resource{
		Create: resourceMinimumInOutCreate,
		Read:   resourceMinimumInOutRead,
		Update: resourceMinimumInOutUpdate,
		Delete: resourceMinimumInOutDelete,
		Schema: map[string]*schema.Schema{
			"required": {
				Type:     schema.TypeString,
				Required: true,
			},
			"optional": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"optional_computed": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"computed": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"input_only": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceMinimumInOutCreate(d *schema.ResourceData, meta interface{}) error {
	value := state.NewInOutValue(
		d.Get("required").(string),
		d.Get("optional").(string),
		d.Get("optional_computed").(string),
		d.Get("input_only").(string),
	)

	id, err := inOutDriver.Create(value)
	if err != nil {
		return err
	}

	d.SetId(id)
	return resourceMinimumInOutRead(d, meta)
}

func resourceMinimumInOutRead(d *schema.ResourceData, _ interface{}) error {
	id := d.Id()
	data, err := inOutDriver.Read(id)
	if err != nil {
		if state.IsStateNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return err
	}

	value := &state.InOutValue{}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	d.Set("required", value.Required)                  // nolint
	d.Set("optional", value.Optional)                  // nolint
	d.Set("optional_computed", value.OptionalComputed) // nolint
	d.Set("computed", value.Computed())                // nolint

	// d.Set("input_only", value.InputOnly)

	return nil
}

func resourceMinimumInOutUpdate(d *schema.ResourceData, meta interface{}) error {
	id := d.Id()

	// read current state
	data, err := inOutDriver.Read(id)
	if err != nil {
		if state.IsStateNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return err
	}

	value := &state.InOutValue{}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	// set updated values
	if d.HasChange("required") {
		value.Required = d.Get("required").(string)
	}
	if d.HasChange("optional") {
		value.Optional = d.Get("optional").(string)
	}
	if d.HasChange("optional_computed") {
		value.OptionalComputed = d.Get("optional_computed").(string)
	}
	if d.HasChange("input_only") {
		value.InputOnly = d.Get("input_only").(string)
	}

	// save values
	if err := inOutDriver.Update(id, value); err != nil {
		return err
	}

	return resourceMinimumInOutRead(d, meta)
}

func resourceMinimumInOutDelete(d *schema.ResourceData, _ interface{}) error {
	id := d.Id()
	if err := inOutDriver.Delete(id); err != nil {
		if !state.IsStateNotFoundError(err) {
			return err
		}
	}
	d.SetId("")
	return nil
}
