provider "libvirt" {
  uri = "${var.libvirt_uri}"
}

module "bootstrap" {
  source = "./bootstrap"

  cluster_id       = "${var.cluster_id}"
  image            = "${var.os_image}"
  ignition         = "${var.ignition_bootstrap}"
  baremetal_bridge = "${var.baremetal_bridge}"
  overcloud_bridge = "${var.overcloud_bridge}"
}
