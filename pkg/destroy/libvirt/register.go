// +build libvirt

package libvirt

import (
	"github.com/openshift-metalkube/kni-installer/pkg/destroy"
)

func init() {
	destroy.Registry["libvirt"] = New
}
