// Package baremetal collects bare metal specific configuration.
package baremetal

import (
	"github.com/metalkube/kni-installer/pkg/types/baremetal"
)

// Platform collects bare metal specific configuration.
func Platform() (*baremetal.Platform, error) {
	return &baremetal.Platform{}, nil
}
