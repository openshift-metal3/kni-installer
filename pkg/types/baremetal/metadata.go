package baremetal

// Metadata contains baremetal metadata (e.g. for uninstalling the cluster).
type Metadata struct {
	LibvirtURI string `json:"libvirt_uri"`
	IronicURI  string `json:"ironic_uri"`
	APIVIP     string `json:"api_vip"`
	IngressVIP string `json:"ingress_vip"`
}
