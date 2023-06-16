terraform {
    required_providers {
        cbs = {
            source = "PureStorage-OpenConnect/cbs"
            version = "~> 0.9.0"
        }
    }
}
provider "cbs" {
    aws {
        region = var.region
    }
}

resource "cbs_array_aws" "aws_instance" {

    array_name = var.array_name

    deployment_template_url = var.template_url
    deployment_role_arn = var.deployment_role_arn

    log_sender_domain = var.log_sender_domain
    alert_recipients = var.alert_recipients
    array_model = var.purity_instance_type
    license_key = var.license_key

    pureuser_key_pair_name = var.key_name
    pureuser_private_key_path = var.key_file_path

    system_subnet = var.system_subnet
    replication_subnet = var.rep_subnet
    iscsi_subnet = var.iscsi_subnet
    management_subnet = var.mgmt_subnet

    replication_security_group = var.rep_security_group
    iscsi_security_group = var.iscsi_security_group
    management_security_group = var.mgmt_security_group
}