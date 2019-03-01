variable "base_volume_id" {
  type        = "string"
  description = "The ID of the base volume for the bootstrap node."
}

variable "cluster_id" {
  type        = "string"
  description = "The identifier for the cluster."
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
