variable "location" {
    type = string
}

variable "resource_group_name" {
    type = string
}

variable "jit_group_ids" {
    type = list(string)
}

variable "zone" {
    type = number
}

variable "fusion_sec_name" {
    type = string
}


variable "load_balancer_network_rg" {
    type = string
}

variable "load_balancer_network_name" {
    type = string
}

variable "load_balancer_subnet" {
    type = string
}
