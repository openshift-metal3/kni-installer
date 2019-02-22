package defaults

import (
	"github.com/openshift-metalkube/kni-installer/pkg/types/baremetal"
)

// Defaults for the baremetal platform.
const (
	// DefaultURI is the default URI of the libvirtd connection.
	DefaultURI = "qemu:///system"
)

// SetPlatformDefaults sets the defaults for the platform.
func SetPlatformDefaults(p *baremetal.Platform) {
	if p.URI == "" {
		p.URI = DefaultURI
	}
}
