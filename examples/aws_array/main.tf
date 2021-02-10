provider "cbs" {
    aws{
        region = var.region
    }
}

resource "cbs_array_aws" "aws_instance" {

    array_name = "tf-array"

    deployment_template_url = var.template_url
    deployment_role_arn = var.deployment_role_arn

    log_sender_domain = var.log_sender_domain
    alert_recipients = var.alert_recipients
    array_model = var.purity_instance_type
    license_key = var.license_key

    pureuser_key_pair_name = var.key_name

    system_subnet = var.subnet
    replication_subnet = var.subnet
    iscsi_subnet = var.subnet
    management_subnet = var.subnet

    replication_security_group = var.security_group
    iscsi_security_group = var.security_group
    management_security_group = var.security_group
}