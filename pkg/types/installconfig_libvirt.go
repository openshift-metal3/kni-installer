// +build libvirt

package types

import (
	"sort"

	"github.com/openshift-metal3/kni-installer/pkg/types/libvirt"
)

func init() {
	PlatformNames = append(PlatformNames, libvirt.Name)
	sort.Strings(PlatformNames)
}
