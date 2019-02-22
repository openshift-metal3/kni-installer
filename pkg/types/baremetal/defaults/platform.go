package defaults

import (
	"github.com/metalkube/kni-installer/pkg/types/baremetal"
)

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
