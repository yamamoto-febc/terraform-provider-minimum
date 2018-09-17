package minimum

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/yamamoto-febc/terraform-provider-minimum/state"
)

var schemaDriver = state.NewDriver("schema")

func resourceMinimumSchema() *schema.Resource {
	return &schema.Resource{
		Create: resourceMinimumSchemaCreate,
		Read:   resourceMinimumSchemaRead,
		Update: resourceMinimumSchemaUpdate,
		Delete: resourceMinimumSchemaDelete,
		Schema: map[string]*schema.Schema{
			"default": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
			},
			"default_func": {
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: func() (interface{}, error) {
					return "default_func", nil
				},
			},
			"default_func_required": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("MINIMUM_DEFAULT", nil),
			},
			"diff_suppress_func": {
				Type:     schema.TypeString,
				Optional: true,
				DiffSuppressFunc: func(_, old, new string, d *schema.ResourceData) bool {
					return strings.ToLower(old) == strings.ToLower(new)
				},
			},
			"force_new": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"state_func": {
				Type:     schema.TypeString,
				Optional: true,
				StateFunc: func(value interface{}) string {
					return hash(value.(string))
				},
			},
			"sensitive": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"validate_func": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: func(value interface{}, k string) (ws []string, es []error) {
					v, ok := value.(string)
					if !ok {
						es = append(es, fmt.Errorf("expected type of %s to be string", k))
						return
					}
					if !(1 <= len(v) && len(v) <= 5) {
						es = append(es, fmt.Errorf("expected length of %s to be in the range (%d - %d), got %s", k, 1, 5, v))
					}
					return
				},
			},
			"conflicts_with1": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"conflicts_with2"},
			},
			"conflicts_with2": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"conflicts_with1"},
			},
			"deprecated": {
				Type:       schema.TypeString,
				Optional:   true,
				Deprecated: "deprecated",
			},
			"removed": {
				Type:     schema.TypeString,
				Optional: true,
				Removed:  "removed",
			},
		},
	}
}

func hash(v string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(v)))
}

func resourceMinimumSchemaCreate(d *schema.ResourceData, meta interface{}) error {
	value := &state.SchemaValue{
		Default:             d.Get("default").(string),
		DefaultFunc:         d.Get("default_func").(string),
		DefaultFuncRequired: d.Get("default_func_required").(string),
		DiffSuppressFunc:    d.Get("diff_suppress_func").(string),
		ForceNew:            d.Get("force_new").(string),
		StateFunc:           d.Get("state_func").(string),
		Sensitive:           d.Get("sensitive").(string),
		ValidateFunc:        d.Get("validate_func").(string),
		ConflictsWith1:      d.Get("conflicts_with1").(string),
		ConflictsWith2:      d.Get("conflicts_with2").(string),
		Deprecated:          d.Get("deprecated").(string),
		Removed:             d.Get("removed").(string),
	}

	id, err := schemaDriver.Create(value)
	if err != nil {
		return err
	}

	d.SetId(id)
	return resourceMinimumSchemaRead(d, meta)
}

func resourceMinimumSchemaRead(d *schema.ResourceData, _ interface{}) error {
	id := d.Id()
	data, err := schemaDriver.Read(id)
	if err != nil {
		if state.IsStateNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return err
	}

	value := &state.SchemaValue{}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	d.Set("default", value.Default)                           // nolint
	d.Set("default_func", value.DefaultFunc)                  // nolint
	d.Set("default_func_required", value.DefaultFuncRequired) // nolint
	d.Set("diff_suppress_func", value.DiffSuppressFunc)       // nolint
	d.Set("force_new", value.ForceNew)                        // nolint
	d.Set("state_func", hash(value.StateFunc))                // nolint
	d.Set("sensitive", value.Sensitive)                       // nolint
	d.Set("validate_func", value.ValidateFunc)                // nolint
	d.Set("conflicts_with1", value.ConflictsWith1)            // nolint
	d.Set("conflicts_with2", value.ConflictsWith2)            // nolint
	d.Set("deprecated", value.Deprecated)                     // nolint
	d.Set("removed", value.Removed)                           // nolint

	return nil
}

func resourceMinimumSchemaUpdate(d *schema.ResourceData, meta interface{}) error {
	id := d.Id()

	// read current state
	data, err := schemaDriver.Read(id)
	if err != nil {
		if state.IsStateNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return err
	}

	value := &state.SchemaValue{}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	// set updated values
	if d.HasChange("default") {
		value.Default = d.Get("default").(string)
	}
	if d.HasChange("default_func") {
		value.DefaultFunc = d.Get("default_func").(string)
	}
	if d.HasChange("default_func_required") {
		value.DefaultFuncRequired = d.Get("default_func_required").(string)
	}
	if d.HasChange("diff_suppress_func") {
		value.DiffSuppressFunc = d.Get("diff_suppress_func").(string)
	}
	if d.HasChange("force_new") {
		value.ForceNew = d.Get("force_new").(string)
	}
	if d.HasChange("state_func") {
		value.StateFunc = d.Get("state_func").(string)
	}
	if d.HasChange("sensitive") {
		value.Sensitive = d.Get("sensitive").(string)
	}
	if d.HasChange("validate_func") {
		value.ValidateFunc = d.Get("validate_func").(string)
	}
	if d.HasChange("conflicts_with1") {
		value.ConflictsWith1 = d.Get("conflicts_with1").(string)
	}
	if d.HasChange("conflicts_with2") {
		value.ConflictsWith2 = d.Get("conflicts_with2").(string)
	}
	if d.HasChange("deprecated") {
		value.Deprecated = d.Get("deprecated").(string)
	}
	if d.HasChange("removed") {
		value.Removed = d.Get("removed").(string)
	}
	// save values
	if err := schemaDriver.Update(id, value); err != nil {
		return err
	}

	return resourceMinimumSchemaRead(d, meta)
}

func resourceMinimumSchemaDelete(d *schema.ResourceData, _ interface{}) error {
	id := d.Id()
	if err := schemaDriver.Delete(id); err != nil {
		if !state.IsStateNotFoundError(err) {
			return err
		}
	}
	d.SetId("")
	return nil
}
