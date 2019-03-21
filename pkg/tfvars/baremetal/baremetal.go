// Package baremetal contains bare metal specific Terraform-variable logic.
package baremetal

import (
	"encoding/json"
	"github.com/openshift-metalkube/kni-installer/pkg/types/baremetal"

	libvirttfvars "github.com/openshift-metalkube/kni-installer/pkg/tfvars/libvirt"
	"github.com/pkg/errors"
)

type config struct {
	LibvirtURI      string `json:"libvirt_uri,omitempty"`
	IronicURI       string `json:"ironic_uri,omitempty"`
	Image           string `json:"os_image,omitempty"`
	BareMetalBridge string `json:"baremetal_bridge,omitempty"`
	OverCloudBridge string `json:"overcloud_bridge,omitempty"`

	// Data required for masters deployment - several maps per master, because of terraform's
	// limitation that maps cannot be strings
	Master0     interface{} `json:"master_0"`
	Properties0 interface{} `json:"properties_0"`
	RootDevice0 interface{} `json:"root_device_0"`
	DriverInfo0 interface{} `json:"driver_info_0"`

	Master1     interface{} `json:"master_1"`
	Properties1 interface{} `json:"properties_1"`
	RootDevice1 interface{} `json:"root_device_1"`
	DriverInfo1 interface{} `json:"driver_info_1"`

	Master2     interface{} `json:"master_2"`
	Properties2 interface{} `json:"properties_2"`
	RootDevice2 interface{} `json:"root_device_2"`
	DriverInfo2 interface{} `json:"driver_info_2"`

	MasterConfiguration baremetal.MasterConfiguration `json:"master_configuration"`
}

// TFVars generates bare metal specific Terraform variables.
func TFVars(libvirtURI, ironicURI, osImage, baremetalBridge, overcloudBridge string, nodes map[string]interface{}, configuration baremetal.MasterConfiguration) ([]byte, error) {
	osImage, err := libvirttfvars.CachedImage(osImage)
	if err != nil {
		return nil, errors.Wrap(err, "failed to use cached libvirt image")
	}

	cfg := &config{
		LibvirtURI:          libvirtURI,
		IronicURI:           ironicURI,
		Image:               osImage,
		BareMetalBridge:     baremetalBridge,
		OverCloudBridge:     overcloudBridge,
		Master0:             nodes["master_0"],
		Properties0:         nodes["properties_0"],
		RootDevice0:         nodes["root_device_0"],
		DriverInfo0:         nodes["driver_info_0"],
		Master1:             nodes["master_1"],
		Properties1:         nodes["properties_1"],
		RootDevice1:         nodes["root_device_1"],
		DriverInfo1:         nodes["driver_info_1"],
		Master2:             nodes["master_2"],
		Properties2:         nodes["properties_2"],
		RootDevice2:         nodes["root_device_2"],
		DriverInfo2:         nodes["driver_info_2"],
		MasterConfiguration: configuration,
	}

	return json.MarshalIndent(cfg, "", "  ")
}
