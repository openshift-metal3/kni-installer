// +build libvirt

package libvirt

import (
	"github.com/metalkube/kni-installer/pkg/destroy"
)

func init() {
	destroy.Registry["libvirt"] = New
}
