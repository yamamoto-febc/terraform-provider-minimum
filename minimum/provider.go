package minimum

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},
		ResourcesMap: map[string]*schema.Resource{
			"minimum_bool":        resourceMinimumBool(),
			"minimum_int":         resourceMinimumInt(),
			"minimum_float":       resourceMinimumFloat(),
			"minimum_string":      resourceMinimumString(),
			"minimum_list":        resourceMinimumList(),
			"minimum_nested_list": resourceMinimumNestedList(),
			"minimum_map":         resourceMinimumMap(),
			"minimum_invalid_map": resourceMinimumInvalidMap(),
			"minimum_set":         resourceMinimumSet(),
		},
	}
}
