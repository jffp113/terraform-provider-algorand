package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/jffp113/terraform-provider-algorand/algorand"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: algorand.Provider})
}
