variable "site" {
  description = "Grid'5000 site where cluster will be deployed"
  type        = string
  default     = "rennes"
}

variable "nodes_count" {
  description = "Cluster size"
  type        = number
  default     = 4
}

variable "walltime" {
  description = "Experiment duration (OAR notation)"
  type        = string
  default     = "1"
}

variable "cluster_selector" {
  description = "Grid'5000 cluster name"
  type        = string
  default     = "parasilo"
}

variable "ceph_metadata_device" {
  description = "Use a device for Bluestore metadata"
  type        = string
  default     = ""
}

variable "rook_ceph_version" {
  description = "Rook Ceph version"
  type        = string
  default     = "v1.9.0"
}

variable "ceph_version" {
    description = "Ceph version"
    type = string
    default = "v16.2.7"
}
