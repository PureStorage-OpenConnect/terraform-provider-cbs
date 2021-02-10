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
Note: The provider has only been tested on Linux and may not work on other platforms

### AWS Credentials

The provider is built using AWS SDK for Go, and will require the AWS credentials configured
in the sytem. for more information about setting up the SDK, refer to this page:
https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html

## Using the Provider

Example terraform configuration templates can be found in the examples/ subdirectory.

## Testing

To run acceptance tests, a JSON config file must be specified to supply the input parameters to
the test CBS instances. An example config file can be found at `testing/example_aws_config`.
The config file is passed to the test using the `TEST_ACC_AWS_PARAMS_PATH` environment variable.

Acceptance tests can be run with `make testacc`:
```shell
export TEST_ACC_AWS_PARAMS_PATH=<path to config file>
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
