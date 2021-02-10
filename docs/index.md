---
page_title: "cbs Provider"
subcategory: ""
description: |-

---

# cbs Provider

Use the provider to deploy and manage Pure Storage Cloud Block Store instances. The provider
must be configured with the proper credentials for the desired cloud back end before it can be
used.


## Example Usage

```hcl
provider "cbs" {
    aws {
        region  = "us-west-2"
    }
}
```

## Authentication

### AWS

The AWS component of the provider is built on top of the AWS SDK for Go, and delegates AWS
authentication to that SDK. There are a number of ways to configure the environment with AWS
credentials; to view a list of possible options, please see the [credentials](https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/#specifying-credentials) section of the SDK
setup guide.

## Argument Reference

- `aws` (Optional) - A configuration block for the provider's AWS configuration. See
[below for nested schema](#nestedblock--aws))

<a id="nestedblock--aws"></a>
### Nested Schema for `aws`

- `region` (Optional) - The AWS region. This must be specified either in the Terraform template or
with the `AWS_DEFAULT_REGION` environment variable.
