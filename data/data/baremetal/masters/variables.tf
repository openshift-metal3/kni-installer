variable "image_source" {
  description = "The URL of the OS disk image"
  type = "string"
}

variable "image_checksum" {
  type = "string"
  description = "The URL or checksum value of the image"
}

variable "root_gb" {
  type = "string"
  description = "Size of the root disk"
}

variable "root_disk" {
  type = "string"
  description = "Location of the root disk"
}

variable "ignition" {
  type = "string"
  description = "The content of the master ignition file"
}

variable "master_nodes" {
  type = "map"
  description = "Master bare metal node details"
}

variable "properties" {
  type = "map"
  description = "Master bare metal properties"
}

variable "root_devices" {
  type = "map"
  description = "Master root device configuration"
}

variable "driver_infos" {
  type = "map"
  description = "Master driver info"
}

