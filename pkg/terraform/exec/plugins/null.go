// +build baremetal

package plugins

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-null/null"
)

func init() {
	exec := func() {
		plugin.Serve(&plugin.ServeOpts{
			ProviderFunc: null.Provider,
		})
	}
	KnownPlugins["terraform-provider-null"] = exec
}
