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

variable "master_0" {
  type = "map"
  description = "Master 0 bare metal node details"
}

variable "properties_0" {
  type = "map"
  description = "Master 0 bare metal properties"
}

variable "root_device_0" {
  type = "map"
  description = "Master 0 root device configuration"
}

variable "driver_info_0" {
  type = "map"
  description = "Master 0 driver info"
}

variable "master_1" {
  type = "map"
  description = "Master 1 bare metal node details"
}

variable "properties_1" {
  type = "map"
  description = "Master 1 bare metal properties"
}

variable "root_device_1" {
  type = "map"
  description = "Master 1 root device configuration"
}

variable "driver_info_1" {
  type = "map"
  description = "Master 1 driver info"
}

variable "master_2" {
  type = "map"
  description = "Master 2 bare metal node details"
}

variable "properties_2" {
  type = "map"
  description = "Master 2 bare metal properties"
}

variable "root_device_2" {
  type = "map"
  description = "Master 2 root device configuration"
}

variable "driver_info_2" {
  type = "map"
  description = "Master 2 driver info"
}


