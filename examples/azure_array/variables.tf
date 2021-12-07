variable "location" {
    type = string
}

variable "resource_group_name" {
    type = string
}

variable "virtual_network_id" {
    type = string
}

variable "log_sender_domain" {
    type = string
}

variable "alert_recipients" {
    type = list(string)
}

variable "groups" {
    type = list(string)
}

variable "array_model" {
    type = string
}

variable "license_key" {
    type = string
}

variable "management_subnet" {
    type = string
}

variable "system_subnet" {
    type = string
}

variable "iscsi_subnet" {
    type = string
}

variable "replication_subnet" {
    type = string
}

variable "zone" {
    type = number
}

variable "array_name" {
    type = string
}

variable "key_vault_id" {
    type = string
}

variable "key_file_path" {
    type = string
}
variable "plan_name" {
    type = string
}

variable "plan_product" {
    type = string
}

variable "plan_publisher" {
    type = string
}

variable "plan_version" {
    type = string
}