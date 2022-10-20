#Variables
region                  = "us-west-2"
array_name              = "array-name"
template_url            = "https://s3.amazonaws.com/awsmp-fulfillment-cf-templates-prod/4ea2905b-7939-4ee0-a521-d5c2fcb41214/f7a53b584af54f7ba62ef555b11ed859.template"
deployment_role_arn     = "arn:aws:iam::xxxxxxxxxxxx:role/example_role"
log_sender_domain       = "example-company.org"
alert_recipients        = ["admin1@example-company.org", "admin2@example-company.org"]
purity_instance_type    = "V10AR1"
license_key             = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx"
key_name                = "cbs-key"
key_file_path           = "path/to/cbs-key"
rep_security_group      = "sg-xxxx"
iscsi_security_group    = "sg-xxxx"
mgmt_security_group     = "sg-xxxx"
system_subnet           = "subnet-xxxx"
rep_subnet              = "subnet-xxxx"
iscsi_subnet            = "subnet-xxxx"
mgmt_subnet             = "subnet-xxxx"

