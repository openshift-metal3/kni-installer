// Package baremetal contains utilities that help gather Baremetal specific
// information from terraform state.
package baremetal

import (
	"fmt"
	"github.com/libvirt/libvirt-go"
	"github.com/openshift-metalkube/kni-installer/pkg/types"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"strings"

	"github.com/openshift-metalkube/kni-installer/pkg/terraform"
)

// BootstrapIP returns the ip address for bootstrap host. Baremetal relies on a libvirt bootstrap node, that gets it's
// IP from DHCP.
func BootstrapIP(tfs *terraform.State, config *types.InstallConfig) (string, error) {
	br, err := terraform.LookupResource(tfs, "module.bootstrap", "libvirt_domain", "bootstrap")
	if err != nil {
		return "", errors.Wrap(err, "failed to lookup bootstrap libvirt domain")
	}

	if len(br.Instances) == 0 {
		return "", errors.New("no bootstrap instance found")
	}

	mac, err := getMACFromInstance(&br.Instances[0], config.BareMetal.ExternalBridge)
	if err != nil {
		return "", errors.Wrap(err, "could not fetch bootstrap mac address")
	}

	ip, err := getLeaseForMac(config.BareMetal.LibvirtURI, config.BareMetal.ExternalBridge, mac)
	if err != nil {
		return "", errors.Wrap(err, "no bootstrap dhcp lease found")
	}

	return ip, nil
}

// ControlPlaneIPs returns the ip addresses for control plane hosts, retrieved via Ironic introspection.
func ControlPlaneIPs(tfs *terraform.State) ([]string, error) {
	mrs, err := terraform.LookupResource(tfs, "module.masters", "ironic_introspection", "openshift-master-introspection")
	if err != nil {
		return nil, errors.Wrap(err, "failed to lookup masters introspection data")
	}

	var errs []error
	var masters []string
	for idx, inst := range mrs.Instances {
		interfaces, _, err := unstructured.NestedSlice(inst.Attributes, "interfaces")
		if err != nil {
			errs = append(errs, errors.Wrapf(err, "could not get interfaces for master-%d", idx))
		}
		ip, _, err := unstructured.NestedString(interfaces[0].(map[string]interface{}), "ip")
		masters = append(masters, ip)
	}
	return masters, utilerrors.NewAggregate(errs)
}

// getMACFromInstance finds the MAC address from the terraform state data
func getMACFromInstance(inst *terraform.StateResourceInstance, network string) (string, error) {
	interfaces, _, err := unstructured.NestedSlice(inst.Attributes, "network_interface")
	if err != nil {
		return "", err
	}

	for _, iface := range interfaces {
		bridge, _, err := unstructured.NestedString(iface.(map[string]interface{}), "bridge")
		if err != nil {
			return "", err
		}
		if bridge == network {
			mac, _, err := unstructured.NestedString(iface.(map[string]interface{}), "mac")
			return mac, err
		}
	}

	return "", fmt.Errorf("could not find mac")
}

// getLeaseForMac finds the DHCP lease for the given MAC address on a particular bridge.
func getLeaseForMac(libvirtURI, bridge, mac string) (string, error) {
	conn, err := libvirt.NewConnect(libvirtURI)
	if err != nil {
		return "", errors.Wrap(err, "could not connect to libvirt")
	}
	defer conn.Close()

	network, err := conn.LookupNetworkByName(bridge)
	if err != nil {
		return "", errors.Wrap(err, "could not find libvirt network")
	}

	leases, err := network.GetDHCPLeases()
	if err != nil {
		return "", errors.Wrap(err, "could not fetch dhcp leases")
	}

	for _, lease := range leases {
		if strings.ToUpper(lease.Mac) == strings.ToUpper(mac) {
			return lease.IPaddr, nil
		}
	}

	return "", fmt.Errorf("could not fetch dhcp lease")
}
