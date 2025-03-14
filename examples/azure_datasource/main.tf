terraform {
    required_providers {
        cbs = {
            source = "PureStorage-OpenConnect/cbs"
            version = "~> 0.11.1"
        }
    }
}

data "cbs_azure_plans" "plans" {
}

data "cbs_plan_azure" "version_plan" {
    plan_version = "6.8.x"
}


output "cbs_azure_available_plans" {
    value = data.cbs_azure_plans.plans
}

output "cbs_filtered_plan" {
    value = data.cbs_plan_azure.version_plan
}
