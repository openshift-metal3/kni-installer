package validation

import (
	"github.com/openshift-metalkube/kni-installer/pkg/types/baremetal"
	"github.com/openshift-metalkube/kni-installer/pkg/validate"
	"k8s.io/apimachinery/pkg/util/validation/field"
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

	if err := validate.Interface(p.ExternalBridge); err != nil {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("external_bridge"), p.ExternalBridge, err.Error()))
	}

	if err := validate.Interface(p.ProvisioningBridge); err != nil {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("provisioning_bridge"), p.ProvisioningBridge, err.Error()))
	}

	if p.Hosts == nil {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("nodes"), p.Hosts, "bare metal hosts are missing"))
	}

	if p.DefaultMachinePlatform != nil {
		allErrs = append(allErrs, ValidateMachinePool(p.DefaultMachinePlatform, fldPath.Child("defaultMachinePlatform"))...)
	}

	if err := validate.IP(p.APIVIP); err != nil {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("api_vip"), p.APIVIP, err.Error()))
	}

	if err := validate.IP(p.IngressVIP); err != nil {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("ingress_vip"), p.IngressVIP, err.Error()))
	}
	return allErrs
}
