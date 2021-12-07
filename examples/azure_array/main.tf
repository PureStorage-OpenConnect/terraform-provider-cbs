terraform {
    required_providers {
        cbs = {
            source = "PureStorage-OpenConnect/cbs"
            version = "~> 0.6.0"
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
    virtual_network_id = var.virtual_network_id

    key_vault_id = var.key_vault_id
    pureuser_private_key_path = var.key_file_path

    management_subnet = var.management_subnet
    system_subnet = var.system_subnet
    iscsi_subnet = var.iscsi_subnet
    replication_subnet = var.replication_subnet

    jit_approval {
        approvers {
              groups = var.groups
        }
    }
    plan {
        name = var.plan_name
        product = var.plan_product
        publisher = var.plan_publisher
        version = var.plan_version
    }
}

output "cbs_mgmt_endpoint" {
    value = cbs_array_azure.azure_instance.management_endpoint
}
output "cbs_mgmt_endpoint_ct0" {
    value = cbs_array_azure.azure_instance.management_endpoint_ct0
}
output "cbs_mgmt_endpoint_ct1" {
    value = cbs_array_azure.azure_instance.management_endpoint_ct1
}
output "cbs_repl_endpoint_ct0" {
    value = cbs_array_azure.azure_instance.replication_endpoint_ct0
}
output "cbs_repl_endpoint_ct1" {
    value = cbs_array_azure.azure_instance.replication_endpoint_ct1
}
output "cbs_iscsi_endpoint_ct0" {
    value = cbs_array_azure.azure_instance.iscsi_endpoint_ct0
}
output "cbs_iscsi_endpoint_ct1" {
    value = cbs_array_azure.azure_instance.iscsi_endpoint_ct1
}