// +build libvirt

package types

import (
	"sort"

	"github.com/metalkube/kni-installer/pkg/types/libvirt"
)

func init() {
	PlatformNames = append(PlatformNames, libvirt.Name)
	sort.Strings(PlatformNames)
}
