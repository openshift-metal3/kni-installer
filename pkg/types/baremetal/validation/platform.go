package validation

import (
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/metalkube/kni-installer/pkg/types/baremetal"
	"github.com/metalkube/kni-installer/pkg/validate"
)

// ValidatePlatform checks that the specified platform is valid.
func ValidatePlatform(p *baremetal.Platform, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	if err := validate.URI(p.URI); err != nil {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("uri"), p.URI, err.Error()))
	}
	if p.DefaultMachinePlatform != nil {
		allErrs = append(allErrs, ValidateMachinePool(p.DefaultMachinePlatform, fldPath.Child("defaultMachinePlatform"))...)
	}
	return allErrs
}
