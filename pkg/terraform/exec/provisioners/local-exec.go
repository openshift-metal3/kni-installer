// +build baremetal

package provisioners

import (
	"github.com/hashicorp/terraform/builtin/provisioners/local-exec"
	"github.com/hashicorp/terraform/plugin"
)

func init() {
	localExecProvisioner := func() {
		plugin.Serve(&plugin.ServeOpts{
			ProvisionerFunc: localexec.Provisioner,
		})
	}
	KnownProvisioners["local-exec"] = localExecProvisioner
}
