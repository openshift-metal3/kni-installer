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

  master_nodes = "${var.master_nodes}"
  properties = "${var.properties}"
  root_devices = "${var.root_devices}"
  driver_infos = "${var.driver_infos}"
}

