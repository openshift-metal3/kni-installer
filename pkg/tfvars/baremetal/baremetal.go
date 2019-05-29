// Package baremetal contains bare metal specific Terraform-variable logic.
package baremetal

import (
	"encoding/json"
	libvirttfvars "github.com/openshift-metalkube/kni-installer/pkg/tfvars/libvirt"
	"github.com/pkg/errors"
)

type config struct {
	LibvirtURI         string `json:"libvirt_uri,omitempty"`
	IronicURI          string `json:"ironic_uri,omitempty"`
	Image              string `json:"os_image,omitempty"`
	ExternalBridge     string `json:"external_bridge,omitempty"`
	ProvisioningBridge string `json:"provisioning_bridge,omitempty"`

	// Data required for masters deployment - several maps per master, because of terraform's
	// limitation that maps cannot be strings
	MasterNodes interface{} `json:"master_nodes"`
	Properties  interface{} `json:"properties"`
	RootDevices interface{} `json:"root_devices"`
	DriverInfos interface{} `json:"driver_infos"`

	MasterConfiguration map[string]interface{} `json:"master_configuration"`
}

// TFVars generates bare metal specific Terraform variables.
func TFVars(libvirtURI, ironicURI, osImage, externalBridge, provisioningBridge string, nodes map[string]interface{}, configuration map[string]interface{}) ([]byte, error) {
	osImage, err := libvirttfvars.CachedImage(osImage)
	if err != nil {
		return nil, errors.Wrap(err, "failed to use cached libvirt image")
	}

	cfg := &config{
		LibvirtURI:          libvirtURI,
		IronicURI:           ironicURI,
		Image:               osImage,
		ExternalBridge:      externalBridge,
		ProvisioningBridge:  provisioningBridge,
		MasterNodes:         nodes["master_nodes"],
		Properties:          nodes["properties"],
		RootDevices:         nodes["root_devices"],
		DriverInfos:         nodes["driver_infos"],
		MasterConfiguration: configuration,
	}

	return json.MarshalIndent(cfg, "", "  ")
}
