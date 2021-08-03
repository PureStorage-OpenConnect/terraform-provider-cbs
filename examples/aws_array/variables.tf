variable "region" {
    type = string
}

variable "template_url" {
    type = string
}

variable "deployment_role_arn" {
    type = string
}

variable "log_sender_domain" {
    type = string
}

variable "alert_recipients" {
    type = list(string)
}

variable "purity_instance_type" {
    type = string
}

variable "license_key" {
    type = string
}

variable "system_subnet" {
    type = string
}
variable "rep_subnet" {
    type = string
}

variable "iscsi_subnet" {
    type = string
}

variable "mgmt_subnet" {
    type = string
}

variable "key_name" {
    type = string
}

variable "key_file_path" {
    type = string
}

variable "rep_security_group" {
    type = string
}

variable "iscsi_security_group" {
    type = string
}

variable "mgmt_security_group" {
    type = string
}

variable "array_name" {
    type = string
}