package gcp

import "github.com/openshift-metalkube/kni-installer/pkg/destroy/providers"

func init() {
	providers.Registry["gcp"] = New
}
