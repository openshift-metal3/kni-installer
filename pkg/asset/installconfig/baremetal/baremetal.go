// Package baremetal collects bare metal specific configuration.
package baremetal

import (
	survey "gopkg.in/AlecAivazis/survey.v1"

	"github.com/metalkube/kni-installer/pkg/types/baremetal"
	baremetaldefaults "github.com/metalkube/kni-installer/pkg/types/baremetal/defaults"
	"github.com/metalkube/kni-installer/pkg/validate"
)

// Platform collects bare metal specific configuration.
func Platform() (*baremetal.Platform, error) {
	var uri string
	err := survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Input{
				Message: "Libvirt Connection URI",
				Help:    "The libvirt connection URI to be used.",
				Default: baremetaldefaults.DefaultURI,
			},
			Validate: survey.ComposeValidators(survey.Required, uriValidator),
		},
	}, &uri)
	if err != nil {
		return nil, err
	}

	return &baremetal.Platform{
		URI: uri,
	}, nil
}

// uriValidator validates if the answer provided in prompt is a valid
// url and has non-empty scheme.
func uriValidator(ans interface{}) error {
	return validate.URI(ans.(string))
}
