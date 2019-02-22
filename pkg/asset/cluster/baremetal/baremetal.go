// Package baremetal extracts bare metal metadata from install
// configurations.
package baremetal

import (
	"github.com/metalkube/kni-installer/pkg/types"
	"github.com/metalkube/kni-installer/pkg/types/baremetal"
)

// Metadata converts an install configuration to bare metal metadata.
func Metadata(clusterID string, config *types.InstallConfig) *baremetal.Metadata {
	return &baremetal.Metadata{
		URI: config.Platform.BareMetal.URI,
	}
}
