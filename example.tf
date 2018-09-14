resource minimum_bool "bool" {
  value = true
}

resource minimum_int "int" {
  value = 1
}

resource minimum_float "float" {
  value = 1.1
}

resource minimum_string "string" {
  value = "foobar"
}

resource minimum_list "list" {
  value = ["list0", "list1", "list2"]
}

resource minimum_nested_list "nested-list" {
  value = [
    {
      name        = "name1"
      value       = "value1"
      description = "description1"
    },
    {
      name        = "name2"
      value       = "value2"
      description = "description2"
    },
  ]
}

resource minimum_map "map" {
  value = {
    key1 = "value1"
    key2 = true
    key3 = 1

    # can't use nested value
    #key4 = {
    #  nest1 = {
    #    nest2 = "value"
    #  }
    #}
  }
}

resource minimum_invalid_map "invalid" {
  value = {
    value1 = "value1"
    value2 = 2
    value3 = true
  }
}

resource minimum_invalid_map "invalid2" {
  value = {
    # value1 = "value1" 
    value2 = "foo"      
    value3 = "bar"      
    value4 = "disapper" 
  }
}

resource minimum_set "set" {
  value = [
    {
      name        = "name1"
      value       = "value1"
      description = "description1"
    },
    {
      name        = "name2"
      value       = "value2"
      description = "description2"
    },
  ]
}
