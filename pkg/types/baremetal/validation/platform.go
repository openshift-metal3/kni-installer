package validation

import (
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/openshift-metalkube/kni-installer/pkg/types/baremetal"
	"github.com/openshift-metalkube/kni-installer/pkg/validate"
)

// ValidatePlatform checks that the specified platform is valid.
func ValidatePlatform(p *baremetal.Platform, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	if err := validate.URI(p.LibvirtURI); err != nil {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("libvirt_uri"), p.LibvirtURI, err.Error()))
	}

	if err := validate.URI(p.IronicURI); err != nil {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("ironic_uri"), p.LibvirtURI, err.Error()))
	}

	if p.Nodes == nil {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("nodes"), p.Nodes, "nodes is missing"))
	}

	if p.DefaultMachinePlatform != nil {
		allErrs = append(allErrs, ValidateMachinePool(p.DefaultMachinePlatform, fldPath.Child("defaultMachinePlatform"))...)
	}
	return allErrs
}
