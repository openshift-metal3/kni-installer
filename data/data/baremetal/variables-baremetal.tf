variable "ironic_uri" {
  type = "string"
  description = "ironic connection URI"
}

variable "libvirt_uri" {
  type = "string"
  description = "libvirt connection URI"
}

variable "os_image" {
  type = "string"
  description = "The URL of the OS disk image"
}

variable "baremetal_bridge" {
  type = "string"
  description = "The name of the baremetal bridge"
}

variable "overcloud_bridge" {
  type = "string"
  description = "The name of the overcloud bridge"
}

variable "master_configuration" {
  type = "map"
  description = "Configuration information for masters such as image location"
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


