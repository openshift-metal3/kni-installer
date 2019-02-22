locals {
  cluster_domain = "${var.cluster_name}.${var.base_domain}"
}

provider "libvirt" {
  uri = "${var.libvirt_uri}"
}

module "volume" {
  source = "./volume"

  cluster_name = "${var.cluster_name}"
  image        = "${var.os_image}"
}

module "bootstrap" {
  source = "./bootstrap"

  base_volume_id   = "${module.volume.coreos_base_volume_id}"
  cluster_name     = "${var.cluster_name}"
  ignition         = "${var.ignition_bootstrap}"
  baremetal_bridge = "${var.baremetal_bridge}"
  overcloud_bridge = "${var.overcloud_bridge}"
}
