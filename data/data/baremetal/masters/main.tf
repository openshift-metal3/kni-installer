resource "ironic_node_v1" "openshift-master-node" {
  count          = "${length(keys(var.master_nodes))}"
  name           = "${lookup(var.master_nodes[format("openshift-master-%d", count.index)], "name")}"
  resource_class = "baremetal"

  inspect   = true
  clean     = true
  available = true

  ports = [
    {
      address     = "${lookup(var.master_nodes[format("openshift-master-%d", count.index)], "port_address")}"
      pxe_enabled = "true"
    },
  ]

  properties  = "${var.properties[format("openshift-master-%d", count.index)]}"
  root_device = "${var.root_devices[format("openshift-master-%d", count.index)]}"

  driver      = "${lookup(var.master_nodes[format("openshift-master-%d", count.index)], "driver")}"
  driver_info = "${var.driver_infos[format("openshift-master-%d", count.index)]}"

  management_interface = "${lookup(var.master_nodes[format("openshift-master-%d", count.index)], "management_interface")}"
  power_interface      = "${lookup(var.master_nodes[format("openshift-master-%d", count.index)], "power_interface")}"
  vendor_interface     = "${lookup(var.master_nodes[format("openshift-master-%d", count.index)], "vendor_interface")}"
}

resource "ironic_allocation_v1" "openshift-master-allocation" {
  name           = "master-${count.index}"
  count          = 3
  resource_class = "baremetal"

  candidate_nodes = [
    "${ironic_node_v1.openshift-master-node.*.id}",
  ]
}

resource "ironic_deployment" "openshift-master-deployment" {
  count     = 3
  node_uuid = "${element(ironic_allocation_v1.openshift-master-allocation.*.node_uuid, count.index)}"

  instance_info = {
    image_source   = "${var.image_source}"
    image_checksum = "${var.image_checksum}"
    root_gb        = "${var.root_gb}"
  }

  user_data = "${var.ignition}"
}
