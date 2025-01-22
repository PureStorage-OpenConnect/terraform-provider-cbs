terraform {
    required_providers {
        cbs = {
            source = "PureStorage-OpenConnect/cbs"
            version = "~> 0.11.1"
        }
    }
}

provider "cbs" {}

data "cbs_plan_azure" "version_plan" {
    plan_version = var.plan_version
}

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

    jit_approval_group_object_ids = var.jit_group_ids
    user_assigned_identity = var.user_assigned_identity

    resource_tags {
        resource = "Microsoft.Compute/virtualMachines"
        tag {
            name = "foo"
            value = "bar"
        }
        tag {
            name = "baz"
            value = "qux"
        }
    }

    resource_tags {
        resource = "Microsoft.Network/networkInterfaces"
        tag {
            name = "foo"
            value = "bar"
        }
    }

    plan {
        name = data.cbs_plan_azure.version_plan.name
        product = data.cbs_plan_azure.version_plan.product
        publisher = data.cbs_plan_azure.version_plan.publisher
        version = data.cbs_plan_azure.version_plan.version
    }

    lifecycle {
        ignore_changes = [
            plan,
        ]
    }
}

output "cbs_azure_available_plans" {
    value = data.cbs_plan_azure.version_plan
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