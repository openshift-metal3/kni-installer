package validation

import (
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/openshift-metalkube/kni-installer/pkg/types/openstack"
)

// ValidateMachinePool checks that the specified machine pool is valid.
func ValidateMachinePool(p *openstack.MachinePool, fldPath *field.Path) field.ErrorList {
	return field.ErrorList{}
}
