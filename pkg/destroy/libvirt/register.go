// +build libvirt

package libvirt

import (
	"github.com/openshift-metalkube/kni-installer/pkg/destroy/providers"
)

func init() {
	providers.Registry["libvirt"] = New
}
