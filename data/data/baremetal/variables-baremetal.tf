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

variable "external_bridge" {
  type = "string"
  description = "The name of the external bridge"
}

variable "provisioning_bridge" {
  type = "string"
  description = "The name of the provisioning bridge"
}

variable "master_configuration" {
  type = "map"
  description = "Configuration information for masters such as image location"
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
  description = "Master root device configurations"
}

variable "driver_infos" {
  type = "map"
  description = "Master driver infos"
}
