---
page_title: "cbs_array_aws Resource - terraform-provider-cbs"
subcategory: ""
description: |-

---

# cbs_array_aws Resource

Allows the deployment and management of a Cloud Block Store instance on AWS. The instance is deployed as a CloudFormation stack.

Refer to the [deployment guide](https://support.purestorage.com/FlashArray/PurityFA/Cloud_Block_Store/Cloud_Block_Store_Deployment_and_Configuration_Guide_for_AWS) for information on how to configure the AWS environment for the CBS instance.

!>Currently, destroying the Terraform resource **will not** deactivate the CBS instance. Due to this,
we recommend that you deactivate the instance from inside the array itself. For information on how
to do this, refer to the [relevant section](https://support.purestorage.com/FlashArray/PurityFA/Cloud_Block_Store/Cloud_Block_Store_Deployment_and_Configuration_Guide_for_AWS#Removing_Cloud_Block_Store) of the guide.
This method is not available for instances at Purity version 5.3.0.aws2 or earlier. If the instance is destroyed by another
method than the one outlined in the guide, including `terraform destroy`, or if the instance is at version 5.3.0.aws2 or earlier,
then you must contact Pure Storage Support in order to deactivate the instance.

~>Updates are currently not supported for this resource.

## Example Usage

```hcl
resource "cbs_array_aws" "cbs_example" {

    array_name = "terraform-example-instance"

    deployment_template_url = "https://s3.amazonaws.com/awsmp-fulfillment-cf-templates-prod/4ea2905b-7939-4ee0-a521-d5c2fcb41214/f7a53b584af54f7ba62ef555b11ed859.template"
    deployment_role_arn = "arn:aws:iam::xxxxxxxxxxxx:role/example_role"

    log_sender_domain = "example-company.org"
    alert_recipients = ["admin1@example-company.org", "admin2@example-company.org"]
    array_model = "V10AR1"
    license_key = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx"

    pureuser_key_pair_name = "cbs_key"

    system_subnet = "subnet-xxxxxxxxxxxxxxxxx"
    replication_subnet = "subnet-xxxxxxxxxxxxxxxxx"
    iscsi_subnet = "subnet-xxxxxxxxxxxxxxxxx"
    management_subnet = "subnet-xxxxxxxxxxxxxxxxx"

    replication_security_group = "sg-xxxxxxxxxxxxxxxxx"
    iscsi_security_group = "sg-xxxxxxxxxxxxxxxxx"
    management_security_group = "sg-xxxxxxxxxxxxxxxxx"

    tags = {
        foo = "bar"
        test = "value"
    }
}
```

## Argument Reference

- `array_name` (Required) - Name of the array, and the name of the CloudFormation stack.
- `array_model` (Required) - CBS array size to launch. The possible values are `V10AR1` or `V20AR1`
- `alert_recipients` (Optional) - List of email addresses to receive alerts.
- `deployment_role_arn` (Required) - The ARN of an IAM role that AWS CloudFormation assumes to create the stack
- `deployment_template_url` (Required) -  AWS S3 URL Containing the CloudFormation template.
- `iscsi_security_group` (Required) - Security Group to apply to iSCSI interfaces. Must allow inbound TCP traffic on port 3260.
- `iscsi_subnet` (Required) - Subnet in which to launch iSCSI interfaces.
- `license_key` (Required) - Pure Storage-provided license key.
- `log_sender_domain` (Required) - Domain name used to determine how CBS logs are parsed and treated by Pure Storage Support and Escalations.
- `management_security_group` (Required) - Security Group to apply to management interfaces. Must allow inbound TCP traffic on ports 22, 80, 8084, and inbound/outbound on port 443.
- `management_subnet` (Required) - Subnet in which to launch management interfaces.
- `pureuser_key_pair_name` (Required) - Name of an existing EC2 KeyPair to enable SSH access to the controllers.
- `replication_security_group` (Required) - Security Group to apply to replication interfaces. Must allow inbound and outbound TCP traffic on ports 8117.
- `replication_subnet` (Required) - Subnet in which to launch replication interfaces.
- `system_subnet` (Required) - Subnet in which to launch system interfaces (internal to the array).
- `tags` (Optional) - A list of tags to apply to the CloudFormation stack.

## Attribute Reference

- `gui_endpoint` - URL of the Purity GUI
- `iscsi_endpoint_ct0` - iSCSI IP address and port of the ct0 instance.
- `iscsi_endpoint_ct1` - iSCSI IP address and port of the ct1 instance.
- `management_endpoint` - Management IP address of the CBS instance.
- `replication_endpoint_ct0` - Replication IP address of the ct0 instance.
- `replication_endpoint_ct1` - Replication IP address of the ct1 instance.
- `resume_lambda` - Lambda function used to resume the array. Only available when deploying arrays
at version 6.1.0 or higher.
- `stack_name` - The name of the CloudFormation stack.



