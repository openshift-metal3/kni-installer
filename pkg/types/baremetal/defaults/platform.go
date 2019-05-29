package defaults

import (
	"github.com/openshift-metalkube/kni-installer/pkg/types/baremetal"
)

// Defaults for the baremetal platform.
const (
	LibvirtURI         = "qemu:///system"
	IronicURI          = "http://localhost:6385/v1"
	ExternalBridge     = "baremetal"
	ProvisioningBridge = "provisioning"
)

// SetPlatformDefaults sets the defaults for the platform.
func SetPlatformDefaults(p *baremetal.Platform) {
	if p.LibvirtURI == "" {
		p.LibvirtURI = LibvirtURI
	}

	if p.IronicURI == "" {
		p.IronicURI = IronicURI
	}

	if p.ExternalBridge == "" {
		p.ExternalBridge = ExternalBridge
	}

	if p.ProvisioningBridge == "" {
		p.ProvisioningBridge = ProvisioningBridge
	}
}
