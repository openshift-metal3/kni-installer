variable "cluster_id" {
  type        = "string"
  description = "The identifier for the cluster."
}

variable "image" {
  description = "The URL of the OS disk image"
  type        = "string"
}

variable "ignition" {
  type        = "string"
  description = "The content of the bootstrap ignition file."
}

variable "baremetal_bridge" {
  type        = "string"
  description = "The name of the baremetal bridge"
}

variable "overcloud_bridge" {
  type        = "string"
  description = "The name of the overcloud bridge"
}
