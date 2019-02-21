// Package baremetal contains bare metal specific Terraform-variable logic.
package baremetal

import (
	"encoding/json"
)

type config struct {
}

// TFVars generates bare metal specific Terraform variables.
func TFVars() ([]byte, error) {
	cfg := &config{}
	return json.MarshalIndent(cfg, "", "  ")
}
