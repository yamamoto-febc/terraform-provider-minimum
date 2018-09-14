package minimum

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/yamamoto-febc/terraform-provider-minimum/state"
)

func resourceMinimumBool() *schema.Resource {
	driver := state.NewDriver("types-bool")
	return &schema.Resource{
		Create: resourceMinimumResourceDataCreate(driver),
		Read:   resourceMinimumResourceDataRead(driver),
		Update: resourceMinimumResourceDataUpdate(driver),
		Delete: resourceMinimumResourceDataDelete(driver),
		Schema: map[string]*schema.Schema{
			"value": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func resourceMinimumInt() *schema.Resource {
	driver := state.NewDriver("types-int")
	return &schema.Resource{
		Create: resourceMinimumResourceDataCreate(driver),
		Read:   resourceMinimumResourceDataRead(driver),
		Update: resourceMinimumResourceDataUpdate(driver),
		Delete: resourceMinimumResourceDataDelete(driver),
		Schema: map[string]*schema.Schema{
			"value": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceMinimumFloat() *schema.Resource {
	driver := state.NewDriver("types-float")
	return &schema.Resource{
		Create: resourceMinimumResourceDataCreate(driver),
		Read:   resourceMinimumResourceDataRead(driver),
		Update: resourceMinimumResourceDataUpdate(driver),
		Delete: resourceMinimumResourceDataDelete(driver),
		Schema: map[string]*schema.Schema{
			"value": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}

func resourceMinimumString() *schema.Resource {
	driver := state.NewDriver("types-string")
	return &schema.Resource{
		Create: resourceMinimumResourceDataCreate(driver),
		Read:   resourceMinimumResourceDataRead(driver),
		Update: resourceMinimumResourceDataUpdate(driver),
		Delete: resourceMinimumResourceDataDelete(driver),
		Schema: map[string]*schema.Schema{
			"value": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceMinimumList() *schema.Resource {
	driver := state.NewDriver("types-list")
	return &schema.Resource{
		Create: resourceMinimumResourceDataCreate(driver),
		Read:   resourceMinimumResourceDataRead(driver),
		Update: resourceMinimumResourceDataUpdate(driver),
		Delete: resourceMinimumResourceDataDelete(driver),
		Schema: map[string]*schema.Schema{
			"value": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceMinimumNestedList() *schema.Resource {
	driver := state.NewDriver("types-nested-list")
	return &schema.Resource{
		Create: resourceMinimumResourceDataCreate(driver),
		Read:   resourceMinimumResourceDataRead(driver),
		Update: resourceMinimumResourceDataUpdate(driver),
		Delete: resourceMinimumResourceDataDelete(driver),
		Schema: map[string]*schema.Schema{
			"value": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceMinimumMap() *schema.Resource {
	driver := state.NewDriver("types-map")
	return &schema.Resource{
		Create: resourceMinimumResourceDataCreate(driver),
		Read:   resourceMinimumResourceDataRead(driver),
		Update: resourceMinimumResourceDataUpdate(driver),
		Delete: resourceMinimumResourceDataDelete(driver),
		Schema: map[string]*schema.Schema{
			"value": {
				Type:     schema.TypeMap,
				Optional: true,
				// Elem:     &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}

func resourceMinimumInvalidMap() *schema.Resource {
	driver := state.NewDriver("types-invalid-map")
	return &schema.Resource{
		Create: resourceMinimumResourceDataCreate(driver),
		Read:   resourceMinimumResourceDataRead(driver),
		Update: resourceMinimumResourceDataUpdate(driver),
		Delete: resourceMinimumResourceDataDelete(driver),
		Schema: map[string]*schema.Schema{
			"value": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value1": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value2": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"value3": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceMinimumSet() *schema.Resource {
	driver := state.NewDriver("types-set")
	return &schema.Resource{
		Create: resourceMinimumResourceDataCreate(driver),
		Read:   resourceMinimumResourceDataRead(driver),
		Update: resourceMinimumResourceDataUpdate(driver),
		Delete: resourceMinimumResourceDataDelete(driver),
		Schema: map[string]*schema.Schema{
			"value": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
				Set: func(v interface{}) int {
					var buf bytes.Buffer
					m := v.(map[string]interface{})
					keys := []string{"name", "value"}
					for _, key := range keys {
						if v, ok := m[key]; ok {
							buf.WriteString(fmt.Sprintf("%s-", v.(string)))
						}
					}
					return hashcode.String(buf.String())
				},
			},
		},
	}
}

func resourceMinimumResourceDataCreate(driver *state.Driver) func(*schema.ResourceData, interface{}) error {
	return func(d *schema.ResourceData, meta interface{}) error {
		raw := d.Get("value")
		if v, ok := raw.(*schema.Set); ok {
			raw = v.List()
		}

		value := &state.Types{
			Value: raw,
		}

		id, err := driver.Create(value)
		if err != nil {
			return err
		}

		d.SetId(id)
		return resourceMinimumResourceDataRead(driver)(d, meta)
	}
}

func resourceMinimumResourceDataRead(driver *state.Driver) func(*schema.ResourceData, interface{}) error {
	return func(d *schema.ResourceData, _ interface{}) error {
		id := d.Id()
		data, err := driver.Read(id)
		if err != nil {
			if state.IsStateNotFoundError(err) {
				d.SetId("")
				return nil
			}
			return err
		}

		value := &state.Types{}
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}

		return nil
	}
}

func resourceMinimumResourceDataUpdate(driver *state.Driver) func(*schema.ResourceData, interface{}) error {
	return func(d *schema.ResourceData, meta interface{}) error {
		id := d.Id()

		// read current state
		data, err := driver.Read(id)
		if err != nil {
			if state.IsStateNotFoundError(err) {
				d.SetId("")
				return nil
			}
			return err
		}

		value := &state.Types{}
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}

		if d.HasChange("value") {
			raw := d.Get("value")
			if v, ok := raw.(*schema.Set); ok {
				raw = v.List()
			}
			value.Value = raw
		}

		// save values
		if err := driver.Update(id, value); err != nil {
			return err
		}

		return resourceMinimumResourceDataRead(driver)(d, meta)
	}
}

func resourceMinimumResourceDataDelete(driver *state.Driver) func(*schema.ResourceData, interface{}) error {
	return func(d *schema.ResourceData, _ interface{}) error {
		id := d.Id()
		if err := driver.Delete(id); err != nil {
			if !state.IsStateNotFoundError(err) {
				return err
			}
		}
		d.SetId("")
		return nil
	}
}
