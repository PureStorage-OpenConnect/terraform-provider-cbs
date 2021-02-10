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

variable "subnet" {
    type = string
}

variable "key_name" {
    type = string
}

variable "security_group" {
    type = string
}