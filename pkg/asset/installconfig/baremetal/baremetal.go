// Package baremetal collects bare metal specific configuration.
package baremetal

import (
	"encoding/json"
	"gopkg.in/AlecAivazis/survey.v1"

	"github.com/openshift-metalkube/kni-installer/pkg/types/baremetal"
	baremetaldefaults "github.com/openshift-metalkube/kni-installer/pkg/types/baremetal/defaults"
	"github.com/openshift-metalkube/kni-installer/pkg/validate"
)

// Platform collects bare metal specific configuration.
func Platform() (*baremetal.Platform, error) {
	var libvirtURI, ironicURI, externalBridge, provisioningBridge, apiVIP, nodesJSON string
	err := survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Input{
				Message: "Libvirt Connection URI",
				Help:    "The libvirt connection URI to be used.",
				Default: baremetaldefaults.LibvirtURI,
			},
			Validate: survey.ComposeValidators(survey.Required, uriValidator),
		},
	}, &libvirtURI)
	if err != nil {
		return nil, err
	}

	err = survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Input{
				Message: "Ironic Connection URI",
				Help:    "The ironic connection URI to be used.",
				Default: baremetaldefaults.IronicURI,
			},
			Validate: survey.ComposeValidators(survey.Required, uriValidator),
		},
	}, &libvirtURI)
	if err != nil {
		return nil, err
	}

	err = survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Input{
				Message: "External bridge",
				Help:    "External bridge is used for external communication.",
				Default: baremetaldefaults.ExternalBridge,
			},
			Validate: survey.ComposeValidators(survey.Required, interfaceValidator),
		},
	}, &externalBridge)
	if err != nil {
		return nil, err
	}

	err = survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Input{
				Message: "Provisioning bridge",
				Help:    "Provisioning bridge is used to provision machines.",
				Default: baremetaldefaults.ProvisioningBridge,
			},
			Validate: survey.ComposeValidators(survey.Required, interfaceValidator),
		},
	}, &provisioningBridge)
	if err != nil {
		return nil, err
	}

	err = survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Input{
				Message: "Master node definition JSON",
				Help:    "JSON data containing information about the baremetal nodes for use by Ironic.",
			},
		},
	}, &nodesJSON)
	if err != nil {
		return nil, err
	}

	var nodes map[string]interface{}
	if err = json.Unmarshal([]byte(nodesJSON), &nodes); err != nil {
		return nil, err
	}

	err = survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Input{
				Message: "API VIP",
				Help:    "The VIP to be used for internal API communication.",
			},
			Validate: survey.ComposeValidators(survey.Required, ipValidator),
		},
	}, &apiVIP)
	if err != nil {
		return nil, err
	}

	return &baremetal.Platform{
		LibvirtURI:         libvirtURI,
		IronicURI:          ironicURI,
		Nodes:              nodes,
		ApiVIP:             apiVIP,
		ExternalBridge:     externalBridge,
		ProvisioningBridge: provisioningBridge,
	}, nil
}

// uriValidator validates if the answer provided in prompt is a valid
// url and has non-empty scheme.
func uriValidator(ans interface{}) error {
	return validate.URI(ans.(string))
}

func ipValidator(ans interface{}) error {
	return validate.IP(ans.(string))
}

// interfaceValidator validates if the answer provided is a valid network
// interface.  net.Interfaces in Go does not let us know if it's a bridge, but
// we can at least make sure an interface by that name exists.
func interfaceValidator(ans interface{}) error {
	return validate.Interface(ans.(string))
}
