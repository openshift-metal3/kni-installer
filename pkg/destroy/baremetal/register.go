// Package baremetal provides a cluster-destroyer for bare metal clusters.
package baremetal

import (
	"github.com/openshift-metal3/kni-installer/pkg/destroy"
)

func init() {
	destroy.Registry["baremetal"] = New
}
