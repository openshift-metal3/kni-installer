// Package openstack provides a cluster-destroyer for openstack clusters.
package openstack

import (
	"github.com/openshift-metal3/kni-installer/pkg/destroy"
)

func init() {
	destroy.Registry["openstack"] = New
}
