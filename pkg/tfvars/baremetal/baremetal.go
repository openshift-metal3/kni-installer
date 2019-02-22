// Package baremetal contains bare metal specific Terraform-variable logic.
package baremetal

import (
	"encoding/json"

	libvirttfvars "github.com/metalkube/kni-installer/pkg/tfvars/libvirt"
	"github.com/pkg/errors"
)

type config struct {
	URI             string `json:"libvirt_uri,omitempty"`
	Image           string `json:"os_image,omitempty"`
	BareMetalBridge string `json:"baremetal_bridge,omitempty"`
	OverCloudBridge string `json:"overcloud_bridge,omitempty"`
}

// TFVars generates bare metal specific Terraform variables.
func TFVars(libvirtURI, osImage, baremetalBridge, overcloudBridge string) ([]byte, error) {
	osImage, err := libvirttfvars.CachedImage(osImage)
	if err != nil {
		return nil, errors.Wrap(err, "failed to use cached libvirt image")
	}

	cfg := &config{
		URI:             libvirtURI,
		Image:           osImage,
		BareMetalBridge: baremetalBridge,
		OverCloudBridge: overcloudBridge,
	}

	return json.MarshalIndent(cfg, "", "  ")
}
