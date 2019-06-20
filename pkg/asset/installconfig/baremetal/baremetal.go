// Package baremetal collects bare metal specific configuration.
package baremetal

import (
	"fmt"
	"gopkg.in/AlecAivazis/survey.v1"

	"github.com/openshift-metalkube/kni-installer/pkg/types/baremetal"
	baremetaldefaults "github.com/openshift-metalkube/kni-installer/pkg/types/baremetal/defaults"
	"github.com/openshift-metalkube/kni-installer/pkg/validate"
)

// Platform collects bare metal specific configuration.
func Platform() (*baremetal.Platform, error) {
	var libvirtURI, ironicURI, externalBridge, provisioningBridge, apiVIP string
	var hosts []*baremetal.Host

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
				Message: "API VIP",
				Help:    "The VIP to be used for internal API communication. If blank, the address will be looked up from DNS.",
				Default: baremetaldefaults.ApiVIP,
			},
			Validate: ipValidator,
		},
	}, &apiVIP)
	if err != nil {
		return nil, err
	}

	// Keep prompting for hosts
	for {
		var host *baremetal.Host
		host, err = Host()
		host.Role = "master" // FIXME(stbenjam): Support workers

		if err != nil {
			fmt.Printf("Invalid host - please try again")
			continue
		}
		hosts = append(hosts, host)

		more := false
		prompt := &survey.Confirm{
			Message: "Add another?",
		}
		survey.AskOne(prompt, &more, nil)
		if !more {
			break
		}
	}

	return &baremetal.Platform{
		LibvirtURI:         libvirtURI,
		IronicURI:          ironicURI,
		APIVIP:             apiVIP,
		ExternalBridge:     externalBridge,
		ProvisioningBridge: provisioningBridge,
		Hosts:              hosts,
	}, nil
}

// uriValidator validates if the answer provided in prompt is a valid
// url and has non-empty scheme.
func uriValidator(ans interface{}) error {
	return validate.URI(ans.(string))
}

func ipValidator(ans interface{}) error {
	if (ans.(string) != baremetaldefaults.ApiVIP) {
		return validate.IP(ans.(string))
	}
	return nil
}

// interfaceValidator validates if the answer provided is a valid network
// interface.  net.Interfaces in Go does not let us know if it's a bridge, but
// we can at least make sure an interface by that name exists.
func interfaceValidator(ans interface{}) error {
	return validate.Interface(ans.(string))
}

// macValidator validates if the answer provided is a valid mac address
func macValidator(ans interface{}) error {
	return validate.MAC(ans.(string))
}
