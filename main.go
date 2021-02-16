package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	archive "github.com/terraform-providers/terraform-provider-archive/internal/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: archive.Provider})
}
