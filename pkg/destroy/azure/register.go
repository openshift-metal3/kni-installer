package azure

import (
	"github.com/openshift-metalkube/kni-installer/pkg/destroy/providers"
)

func init() {
	providers.Registry["azure"] = New
}
