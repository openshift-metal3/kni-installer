package azure

import (
	"github.com/openshift-metalkube/kni-installer/pkg/destroy"
)

func init() {
	destroy.Registry["azure"] = New
}
