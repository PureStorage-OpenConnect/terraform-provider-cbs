---
page_title: "cbs Provider"
subcategory: ""
description: |-

---

# CBS Provider

Use the provider to deploy and manage Pure Storage Cloud Block Store instances. Configure the provider
with the proper credentials for the desired cloud back using the authentication methods outlined below.


## Example Usage

**AWS:**
```hcl
provider "cbs" {
    aws {
        region  = "us-west-2"
    }
}
```

**Azure:**
```hcl
provider "cbs" {
    azure {
        client_id = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
        client_secret = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
        subscription_id = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
        tenant_id = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
    }
}
```

## Authentication

### AWS

The AWS component of the provider is built on top of the AWS SDK for Go, and delegates AWS
authentication to that SDK. There are a number of ways to configure the environment with AWS
credentials; to view a list of possible options, please see the [credentials](https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/#specifying-credentials) section of the SDK
setup guide.

### Azure

The provider supports Azure authentication through the Azure CLI or through an Azure Service Principal
with a Client Secret. You can authenticate with the Azure CLI by running `az login` on the same system you are running
Terraform on. The provider will automatically grab the credentials from the system. Alternatively, you
can authenticate through a Service Principal by passing the service principal information into the provider,
either by using the Terraform configuration block or by using environment variables.

## Argument Reference

- `aws` (Optional) - An AWS configuration block. See [below for nested schema](#nestedblock--aws))

- `azure` (Optional) - An Azure configuration block. See [below for nested schema](#nestedblock--azure))

<a id="nestedblock--aws"></a>
### Nested Schema for `aws`

- `region` (Optional) - The AWS region. This can also be specified with the `AWS_DEFAULT_REGION` environment variable.

<a id="nestedblock--azure"></a>
### Nested Schema for `azure`

- `client_id` (Optional) The Client ID to use for service principal authentication. This can also be specified with the `ARM_CLIENT_ID` environment variable.
- `client_secret` (Optional) The Client Secret to use for service principal authentication. This can also be specified with the `ARM_CLIENT_SECRET` environment variable.
- `subscription_id` (Optional) The Subscription to use for service principal authentication. This can also be specified with the `ARM_SUBSCRIPTION_ID` environment variable.
- `tenant_id` (Optional) The Tenant ID to use for service principal authentication. This can also be specified with the `ARM_TENANT_ID` environment variable.
