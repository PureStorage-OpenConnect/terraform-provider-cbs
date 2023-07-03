package acceptance

type AccTestCbsAwsParams struct {
	ArrayName              string `json:"array_name"`
	ArrayModel             string `json:"array_model"`
	DeploymentTemplateUrl  string `json:"deployment_template_url"`
	DeploymentRoleArn      string `json:"deployment_role_arn"`
	LicenseKey             string `json:"license_key"`
	PureuserKeyPairName    string `json:"pureuser_key_pair_name"`
	PureuserPrivateKeyPath string `json:"pureuser_private_key_path"`
	PureuserPrivateKey     string `json:"pureuser_private_key"`
	Subnet                 string `json:"subnet"`
	SecurityGroup          string `json:"security_group"`
}

type AccTestCbsAzureParams struct {
	ArrayName              string `json:"array_name"`
	ArrayModel             string `json:"array_model"`
	AppDefinitionId        string `json:"app_definition_id"`
	PlanName               string `json:"plan_name"`
	PlanProduct            string `json:"plan_product"`
	PlanPublisher          string `json:"plan_publisher"`
	PlanVersion            string `json:"plan_version"`
	ResourceGroupName      string `json:"resource_group_name"`
	Location               string `json:"location"`
	LicenseKey             string `json:"license_key"`
	PureuserPrivateKeyPath string `json:"pureuser_private_key_path"`
	PureuserPrivateKey     string `json:"pureuser_private_key"`
	KeyvaultId             string `json:"keyvault_id"`
	ManagementSubnet       string `json:"management_subnet"`
	ISCSISubnet            string `json:"iscsi_subnet"`
	ReplicationSubnet      string `json:"replication_subnet"`
	SystemSubnet           string `json:"system_subnet"`
	VirtualNetworkId       string `json:"virtual_network_id"`
	JitGroup               string `json:"jit_group"`
	JitGroupID             string `json:"jit_group_id"`
	FusionSECIdentity      string `json:"fusion_sec_identity"`
	UserAssignedIdentity   string `json:"user_assigned_identity"`
}

type AccTestCbsFusionSECAzureParams struct {
	FusionSECName           string `json:"fusion_sec_name"`
	PlanName                string `json:"plan_name"`
	PlanProduct             string `json:"plan_product"`
	PlanPublisher           string `json:"plan_publisher"`
	PlanVersion             string `json:"plan_version"`
	ResourceGroupName       string `json:"resource_group_name"`
	Location                string `json:"location"`
	LoadBalancerNetworkRg   string `json:"load_balancer_network_rg"`
	LoadBalancerNetworkName string `json:"load_balancer_network_name"`
	LoadBalancerSubnet      string `json:"load_balancer_subnet"`
	JitGroupID              string `json:"jit_group_id"`
}
