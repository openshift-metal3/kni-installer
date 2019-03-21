# FIXME: This Terraform HCL file defines the 3 master nodes.  It uses the original ironic_nodes.json format,
# flattened because Terraform v0.11 does not support nested data structures.  Maps may only be key/value
# pairs.  We could use terraform's resource `count` provider to have just one resource declaration, but
# the data would have to be structured differently.

resource "ironic_node_v1" "openshift-master-0" {
  name = "${var.master_0["name"]}"

  target_provision_state = "active"
  user_data = "${var.ignition}"
  root_device = "${var.root_device_0}"
  driver = "${var.master_0["driver"]}"
  driver_info = "${var.driver_info_0}"

  ports = [
    {
      address = "${var.master_0["port_address"]}"
      pxe_enabled = "true"
    },
  ]

  properties = "${var.properties_0}"

  instance_info = {
    image_source = "${var.image_source}"
    image_checksum = "${var.image_checksum}"
    root_gb = "${var.root_gb}"
  }

  management_interface = "${var.master_0["management_interface"]}"
  power_interface = "${var.master_0["power_interface"]}"
  vendor_interface = "${var.master_0["vendor_interface"]}"
}

resource "ironic_node_v1" "openshift-master-1" {
  name = "${var.master_1["name"]}"

  target_provision_state = "active"
  user_data = "${var.ignition}"
  root_device = "${var.root_device_1}"
  driver = "${var.master_1["driver"]}"
  driver_info = "${var.driver_info_1}"

  ports = [
    {
      address = "${var.master_1["port_address"]}"
      pxe_enabled = "true"
    },
  ]

  properties = "${var.properties_1}"

  instance_info = {
    image_source = "${var.image_source}"
    image_checksum = "${var.image_checksum}"
    root_gb = "${var.root_gb}"
  }

  management_interface = "${var.master_1["management_interface"]}"
  power_interface = "${var.master_1["power_interface"]}"
  vendor_interface = "${var.master_1["vendor_interface"]}"
}

resource "ironic_node_v1" "openshift-master-2" {
  name = "${var.master_2["name"]}"

  target_provision_state = "active"
  user_data = "${var.ignition}"
  root_device = "${var.root_device_2}"
  driver = "${var.master_2["driver"]}"
  driver_info = "${var.driver_info_2}"

  ports = [
    {
      address = "${var.master_2["port_address"]}"
      pxe_enabled = "true"
    },
  ]

  properties = "${var.properties_2}"

  instance_info = {
    image_source = "${var.image_source}"
    image_checksum = "${var.image_checksum}"
    root_gb = "${var.root_gb}"
  }

  management_interface = "${var.master_2["management_interface"]}"
  power_interface = "${var.master_2["power_interface"]}"
  vendor_interface = "${var.master_2["vendor_interface"]}"
}
