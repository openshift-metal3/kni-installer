package validation

import (
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/metalkube/kni-installer/pkg/types/baremetal"
)

// ValidatePlatform checks that the specified platform is valid.
func ValidatePlatform(p *baremetal.Platform, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	return allErrs
}
