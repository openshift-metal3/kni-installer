resource "ironic_node_v1" "openshift-master" {
  count = "${length(keys(var.master_nodes))}"
  name = "${lookup(var.master_nodes[format("openshift-master-%d", count.index)], "name")}"

  target_provision_state = "active"
  user_data = "${var.ignition}"

  root_device = "${var.root_devices[format("openshift-master-%d", count.index)]}"
  driver = "${lookup(var.master_nodes[format("openshift-master-%d", count.index)], "driver")}"
  driver_info = "${var.driver_infos[format("openshift-master-%d", count.index)]}"

  ports = [
    {
      address = "${lookup(var.master_nodes[format("openshift-master-%d", count.index)], "port_address")}"
      pxe_enabled = "true"
    },
  ]

  properties = "${var.properties[format("openshift-master-%d", count.index)]}"

  instance_info = {
    image_source = "${var.image_source}"
    image_checksum = "${var.image_checksum}"
    root_gb = "${var.root_gb}"
  }

  management_interface = "${lookup(var.master_nodes[format("openshift-master-%d", count.index)], "management_interface")}"
  power_interface = "${lookup(var.master_nodes[format("openshift-master-%d", count.index)], "power_interface")}"
  vendor_interface = "${lookup(var.master_nodes[format("openshift-master-%d", count.index)], "vendor_interface")}"

}
