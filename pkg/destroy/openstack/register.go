// Package openstack provides a cluster-destroyer for openstack clusters.
package openstack

import (
	"github.com/metalkube/kni-installer/pkg/destroy"
)

func init() {
	destroy.Registry["openstack"] = New
}
