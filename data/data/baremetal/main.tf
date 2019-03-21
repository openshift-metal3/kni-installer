provider "libvirt" {
  uri = "${var.libvirt_uri}"
}

provider "ironic" {
  url          = "${var.ironic_uri}"
  microversion = "1.50"
}

module "bootstrap" {
  source = "./bootstrap"

  cluster_id       = "${var.cluster_id}"
  image            = "${var.os_image}"
  ignition         = "${var.ignition_bootstrap}"
  baremetal_bridge = "${var.baremetal_bridge}"
  overcloud_bridge = "${var.overcloud_bridge}"
}

module "masters" {
  source = "./masters"

  ignition       = "${var.ignition_master}"
  image_source   = "${var.master_configuration["image_source"]}"
  image_checksum = "${var.master_configuration["image_checksum"]}"
  root_gb        = "${var.master_configuration["root_gb"]}"
  root_disk      = "${var.master_configuration["root_disk"]}"

  master_0 = "${var.master_0}"
  properties_0 = "${var.properties_0}"
  root_device_0 = "${var.root_device_0}"
  driver_info_0 = "${var.driver_info_0}"

  master_1 = "${var.master_1}"
  properties_1 = "${var.properties_1}"
  root_device_1 = "${var.root_device_1}"
  driver_info_1 = "${var.driver_info_1}"

  master_2 = "${var.master_2}"
  properties_2 = "${var.properties_2}"
  root_device_2 = "${var.root_device_2}"
  driver_info_2 = "${var.driver_info_2}"
}

