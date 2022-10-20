## 0.8.0 (May 15, 2023)

* Added support for Cloud Block Store on Azure deployment for Pure Fusion using the `fusion_sec_identity` parameter, refer to the [documentation](docs/resources/array_azure.md)
* Added cbs_azure_plans datasource, refer to the [documentation](docs/data-sources/azure_plans.md)
* Added cbs_fusion_sec_azure resource, refer to the [documentation](docs/resources/fusion_sec_azure.md)
* Update Purity version of cbs_array_azure to fetch the latest azure release using cbs_azure_plans
* Update Purity version of cbs_array_aws to 6.4.0

## 0.7.0 (Oct 19, 2022)

* Update Purity version of cbs_array_azure to 6.3.5
* Update Purity version of cbs_array_aws to 6.2.13
* Add support for JIT approvals using `jit_approval_group_object_ids` parameter, refer to the [documentation](docs/resources/array_azure.md)
* Marking `jit_approval` as Deprecated, will be removed in future major release

## 0.6.0 (Aug 3, 2021)

* Update Purity version of cbs_array_azure from 6.1.7 to 6.1.8
* Deactivation now supported for Azure
* Better handling of SecureString fields, fixes issue #3
* Simplified Azure networking parameters
* Multiple Azure parameters were changed as part of deactivation and networking work, as always refer to the [documentation](docs/resources/array_azure.md)

## 0.5.0 (July 20, 2021)

* Update Purity version of cbs_array_azure from 6.1.6 to 6.1.7

## 0.4.0 (May 13, 2021)

* Update Purity version of cbs_array_azure from 6.1.4 to 6.1.6

## 0.3.0 (April 13, 2021)

* Update Purity version of cbs_array_azure from 6.1.0 to 6.1.4

## 0.2.0 (March 9, 2021)

* Release of CBS on Azure support through the cbs_array_azure resource
* Various docs updates and typo fixes

## 0.1.0 (February 9, 2021)

* Initial Release. Support for AWS CBS instances only.
* Create/Destroy, no updates yet.