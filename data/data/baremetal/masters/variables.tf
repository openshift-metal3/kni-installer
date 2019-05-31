variable "image_source" {
  description = "The URL of the OS disk image"
  type        = string
}

variable "image_checksum" {
  type        = string
  description = "The URL or checksum value of the image"
}

variable "root_gb" {
  type        = string
  description = "Size of the root disk"
}

variable "ignition" {
  type        = string
  description = "The content of the master ignition file"
}

variable "master_nodes" {
  type        = map(map(string))
  description = "Master bare metal node details"
}

variable "properties" {
  type        = map(map(string))
  description = "Master bare metal properties"
}

variable "root_devices" {
  type        = map(map(string))
  description = "Master root device configuration"
}

variable "driver_infos" {
  type        = map(map(string))
  description = "Master driver info"
}

