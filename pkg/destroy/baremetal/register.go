// Package baremetal provides a cluster-destroyer for bare metal clusters.
package baremetal

import (
	"github.com/metalkube/kni-installer/pkg/destroy"
)

func init() {
	destroy.Registry["baremetal"] = New
}
