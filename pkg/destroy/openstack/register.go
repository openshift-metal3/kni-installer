// Package openstack provides a cluster-destroyer for openstack clusters.
package openstack

import (
	"github.com/openshift-metalkube/kni-installer/pkg/destroy/providers"
)

func init() {
	providers.Registry["openstack"] = New
}
