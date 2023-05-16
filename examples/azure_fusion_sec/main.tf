terraform {
    required_providers {
        cbs = {
            source = "PureStorage-OpenConnect/cbs"
            version = "~> 0.8.0"
        }
    }
}

provider "cbs" {}
resource "cbs_fusion_sec_azure" "fusion_sec_azure_instance" {

    fusion_sec_name = var.fusion_sec_name

    location = var.location
    resource_group_name = var.resource_group_name
    load_balancer_network_rg = var.load_balancer_network_rg
    load_balancer_network_name = var.load_balancer_network_name
    load_balancer_subnet = var.load_balancer_subnet


    jit_approval_group_object_ids = var.jit_group_ids

    plan {
        name = "xxxx"
        product = "xxxx"
        publisher = "xxxx"
        version = "xxxx"
    }
}


output "fusion_sec_hmvip0" {
    value = fusion_sec_azure_instance.hmvip0
}

output "fusion_sec_hmvip0" {
    value = fusion_sec_azure_instance.hmvip1
}

output "fusion_sec_load_balancer_full_identity_id" {
    value = fusion_sec_azure_instance.load_balancer_full_identity_id
}