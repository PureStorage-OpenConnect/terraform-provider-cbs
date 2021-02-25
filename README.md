# CBS Terraform Provider
Terraform Provider for Pure Storage Cloud Block Store deployments

## Prerequisites

### Tools
- Install Terraform 0.13 (or higher) from: https://terraform.io
- Install Golang from: https://golang.org/dl

### Install Go Modules

```shell
go mod download
```

### Build and Install Provider

```shell
make
```
Note: The provider has only been tested on Linux and might not work on other platforms.

### AWS Credentials

The provider is built using AWS SDK for Go, and will require the AWS credentials configured
in the sytem. for more information about setting up the SDK, refer to this page:
https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html

### Azure Credentials

The provider can be authenticated to Azure using either the Azure CLI or an Azure Service Principal. For Azure CLI,
Terraform only supports authenticating using the `az` CLI (and this must be available on your PATH). Authenticating
using the older `azure` CLI or PowerShell Cmdlets are not supported.
Login to the Azure CLI:
```shell
az login
```

And list the Subscriptions associated with the account:
```shell
az account list
```

You can specify the Subscription to be default if you have more than one Subscription:
```shell
az account set --subscription="SUBSCRIPTION_ID"
```

Terraform can be authenticated with a Service Principal in the provider configuration:
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
Alternatively, those values can also be supplied to the provider through the following environment variables:
```shell
export ARM_CLIENT_ID="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
export ARM_CLIENT_SECRET="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
export ARM_SUBSCRIPTION_ID="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
export ARM_TENANT_ID="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
```
## Using the Provider

[Provider Documentation](https://registry.terraform.io/providers/PureStorage-OpenConnect/cbs/latest/docs)

Example terraform configuration templates can be found in the examples/ subdirectory.

## Testing

To run acceptance tests, a JSON config file must be specified to supply the input parameters to
the test CBS instances. An example config file for AWS can be found at `testing/example_aws_config`.
The config file is passed to the test using the `TEST_ACC_AWS_PARAMS_PATH` environment variable.
An example config file for Azure can be found at `testing/example_azure_config`.
The config file is passed to the test using the `TEST_ACC_AZURE_PARAMS_PATH` environment variable.

Acceptance tests can be run with `make testacc`:
```shell
export TEST_ACC_AWS_PARAMS_PATH=<path to config file>
export TEST_ACC_AZURE_PARAMS_PATH=<path to config file>
make testacc
```

By default, test resources are created in the `us-west-2` AWS region. The `AWS_TEST_REGION`
can be used to specify an alternate region.

**Note**: Running acceptance tests requires a valid Cloud Block Store license and will create
actual CBS instances. CBS and AWS billing costs will apply.

For assistance, please either open GitHub issues or go to the Pure Code() Slack channel here:

#terraform
https://code-purestorage.slack.com/archives/CFYKBGNMP

Self-register for our Slack team at https://codeinvite.purestorage.com/
