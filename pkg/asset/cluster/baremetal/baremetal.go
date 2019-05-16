// Package baremetal extracts bare metal metadata from install
// configurations.
package baremetal

import (
	"github.com/openshift-metalkube/kni-installer/pkg/types"
	"github.com/openshift-metalkube/kni-installer/pkg/types/baremetal"
)

// Metadata converts an install configuration to bare metal metadata.
func Metadata(infraID string, config *types.InstallConfig) *baremetal.Metadata {
	return &baremetal.Metadata{
		LibvirtURI: config.Platform.BareMetal.LibvirtURI,
		IronicURI: config.Platform.BareMetal.IronicURI,
		ApiVIP: config.Platform.BareMetal.ApiVIP,
	}
}
