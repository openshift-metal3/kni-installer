// Package baremetal extracts bare metal metadata from install
// configurations.
package baremetal

import (
	"github.com/openshift-metal3/kni-installer/pkg/types"
	"github.com/openshift-metal3/kni-installer/pkg/types/baremetal"
)

// Metadata converts an install configuration to bare metal metadata.
func Metadata(infraID string, config *types.InstallConfig) *baremetal.Metadata {
	return &baremetal.Metadata{
		LibvirtURI: config.Platform.BareMetal.LibvirtURI,
		IronicURI: config.Platform.BareMetal.IronicURI,
	}
}
