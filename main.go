package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/yamamoto-febc/terraform-provider-minimum/minimum"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: minimum.Provider,
	})
}
