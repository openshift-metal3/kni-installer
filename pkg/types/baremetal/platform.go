package baremetal

type MasterConfiguration struct {
	ImageSource   string `json:"image_source,omitempty"`
	ImageChecksum string `json:"image_checksum,omitempty"`
	RootGb        string `json:"root_gb,omitempty"`
	RootDisk      string `json:"root_disk,omitempty"`
}

// Platform stores all the global configuration that all
// machinesets use.
type Platform struct {
	// LibvirtURI is the identifier for the libvirtd connection.  It must be
	// reachable from the host where the installer is run.
	// +optional
	// Default is qemu:///system
	LibvirtURI string `json:"libvirt_uri,omitempty"`

	// IronicURI is the identifier for the Ironic connection.  It must be
	// reachable from the host where the installer is run.
	// +optional
	IronicURI string `json:"ironic_uri,omitempty"`

	// Nodes is the information needed to create the master nodes in
	// Ironic.
	Nodes map[string]interface{} `json:"nodes"`

	// MasterConfiguration contains the information needed to provision
	// a master.
	MasterConfiguration MasterConfiguration `json:"master_configuration"`

	// DefaultMachinePlatform is the default configuration used when
	// installing on bare metal for machine pools which do not define their own
	// platform configuration.
	// +optional
	DefaultMachinePlatform *MachinePool `json:"defaultMachinePlatform,omitempty"`
}
