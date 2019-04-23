// +build libvirt

package libvirt

import (
	"github.com/openshift-metal3/kni-installer/pkg/destroy"
)

func init() {
	destroy.Registry["libvirt"] = New
}
