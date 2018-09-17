resource minimum_schema "schema_example" {
  default               = "value"
  default_func          = "value"
  default_func_required = "value"
  diff_suppress_func    = "Value"
  force_new             = "value"
  state_func            = "value"
  sensitive             = "value"
  validate_func         = "value"
  #conflicts_with1       = "value"
  #conflicts_with2       = "value"
  #deprecated            = "value"
  #removed               = "value"
}

output sensitive {
  value = "${minimum_schema.schema_example.sensitive}"
}
