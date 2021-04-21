variable "nodes_count" {
    description = "Cluster size"
    type = number
    default = 4
}

variable "walltime" {
    description = "OAR job walltime"
    type = string
    default = "1"
}

variable "username" {
    description = "Grid'5000 username"
    type = string
}

variable "nodes_location" {
    description = "Grid'5000 site where cluster will be deployed"
    type = string
    default = "rennes"
}

variable "data_location" {
    description = "Grid'5000 site where Ceph pool will be created"
    type = string
    default = "rennes"

    validation {
        condition     = contains(["rennes", "nantes"], var.data_location)
        error_message = "Valid values for data_location are (rennes, nantes)."
    }
}

variable "ceph_pool_quota" {
    description = "Quota for ceph pool"
    type = string
    default = "100G"
}

variable "nodes_selector" {
    description = "Nodes selector (OAR SQL notation surrounded by curly brackets)"
    type = string
    default = ""
}
