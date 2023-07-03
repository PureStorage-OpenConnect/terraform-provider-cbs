#Variables
array_name             = "terraform-example-instance"
location               = "location_xxxx"
resource_group_name    = "resource_xxxx"
license_key            = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx"
log_sender_domain      = "example-company.org"
alert_recipients       = ["admin1@example-company.org", "admin2@example-company.org"]
array_model            = "V10MUR1"
zone                   = 1
virtual_network_id     = "/subscriptions/xxxxxxxxxxxxxx/resourceGroups/xxxxxxxxxxxxxx/providers/Microsoft.Network/virtualNetworks/xxxxxxxxxxxxxx"
key_vault_id           = "/subscriptions/xxxxxxxxxxxxxx/resourceGroups/xxxxxxxxxxxxxx/providers/Microsoft.KeyVault/vaults/xxxxxxxxxxxxxx"
management_subnet      = "SN-xxxxxxxxxxxxxx"
system_subnet          = "SN-xxxxxxxxxxxxxx"
iscsi_subnet           = "SN-xxxxxxxxxxxxxx"
replication_subnet     = "SN-xxxxxxxxxxxxxx"
jit_group_ids          = ["xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"]
key_file_path          = "example.pem"
user_assigned_identity = "/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx/resourcegroups/mock_resource_group_name/providers/Microsoft.ManagedIdentity/userAssignedIdentities/xxxxxxx"
