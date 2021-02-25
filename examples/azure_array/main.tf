terraform {
    required_providers {
        cbs = {
            source = "PureStorage-OpenConnect/cbs"
            version = "0.2.0"
        }
    }
}

provider "cbs" {}

resource "cbs_array_azure" "azure_instance" {

    array_name = var.array_name

    location = var.location
    resource_group_name = var.resource_group_name
    license_key = var.license_key
    log_sender_domain = var.log_sender_domain
    alert_recipients = var.alert_recipients
    array_model = var.array_model
    zone = var.zone
    virtual_network = var.virtual_network

    management_subnet = var.management_subnet
    system_subnet = var.system_subnet
    iscsi_subnet = var.iscsi_subnet
    replication_subnet = var.replication_subnet

    management_resource_group = var.management_resource_group
    system_resource_group = var.system_resource_group
    iscsi_resource_group = var.iscsi_resource_group
    replication_resource_group = var.replication_resource_group

    jit_approval {
        approvers {
              groups = var.groups
        }
    }
}