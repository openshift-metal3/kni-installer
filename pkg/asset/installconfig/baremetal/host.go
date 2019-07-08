package baremetal

import (
	"github.com/openshift-metalkube/kni-installer/pkg/types/baremetal"
	"gopkg.in/AlecAivazis/survey.v1"
)

// Host prompts the user for hardware details about a baremetal host.
func Host() (*baremetal.Host, error) {
	var host baremetal.Host

	err := survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Input{
				Message: "Name",
			},
			Validate: survey.ComposeValidators(survey.Required),
		},
	}, &host.Name)
	if err != nil {
		return nil, err
	}

	err = survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Input{
				Message: "BMC Address",
				Help:    "The address for the BMC, e.g. ipmi://192.168.0.1",
			},
			Validate: survey.ComposeValidators(survey.Required, uriValidator),
		},
	}, &host.BMC.Address)
	if err != nil {
		return nil, err
	}

	err = survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Input{
				Message: "BMC Username",
			},
			Validate: survey.ComposeValidators(survey.Required),
		},
	}, &host.BMC.Username)
	if err != nil {
		return nil, err
	}

	err = survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Password{
				Message: "BMC Password",
			},
			Validate: survey.ComposeValidators(survey.Required),
		},
	}, &host.BMC.Password)
	if err != nil {
		return nil, err
	}

	err = survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Input{
				Message: "Boot MAC Address",
			},
			Validate: survey.ComposeValidators(survey.Required, macValidator),
		},
	}, &host.BootMACAddress)
	if err != nil {
		return nil, err
	}

	err = survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Select{
				Message: "Hardware profile",
				Options: []string{"default", "libvirt", "dell", "dell-raid"},
				Default: "default",
			},
		},
	}, &host.HardwareProfile)
	if err != nil {
		return nil, err
	}

	return &host, nil
}
