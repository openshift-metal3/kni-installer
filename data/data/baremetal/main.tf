provider "libvirt" {
  uri = "${var.libvirt_uri}"
}

module "volume" {
  source = "./volume"

  cluster_id = "${var.cluster_id}"
  image      = "${var.os_image}"
}

module "bootstrap" {
  source = "./bootstrap"

  base_volume_id   = "${module.volume.coreos_base_volume_id}"
  cluster_id       = "${var.cluster_id}"
  ignition         = "${var.ignition_bootstrap}"
  baremetal_bridge = "${var.baremetal_bridge}"
  overcloud_bridge = "${var.overcloud_bridge}"
}
