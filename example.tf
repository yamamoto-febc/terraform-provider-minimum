resource minimum_inout "inout" {
  required          = "required"
  optional          = ""
  optional_computed = ""

  // this field cannot be set
  #computed          = ""

  input_only        = "input_only"
}
