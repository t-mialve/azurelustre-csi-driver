# Driver Parameters

For more information, see the [Azure Managed Lustre Filesystem (AMLFS) service documentation](https://learn.microsoft.com/en-us/azure/azure-managed-lustre/) and the [AMLFS CSI documentation](https://learn.microsoft.com/en-us/azure/azure-managed-lustre/use-csi-driver-kubernetes).

## Dynamic Provisioning

Name | Meaning | Available Value | Mandatory | Default value
--- | --- | --- | --- | ---
amlfilesystem-name | The vanity name of the Azure Managed Lustre Filesystem (AMLFS) cluster. This is not the name of the file system used in mount commands. | This field can begin and end with an alphanumeric character. The field can only contain alphanumeric characters, hyphens, and underscores. It can also interpret metadata such as `"${pvc.metadata.name}"`, `"${pvc.metadata.namespace}"`, and `"${pv.metadata.name}"`. | Yes | This value must be provided.
location | Azure region in which the AMLFS cluster will be created. The region name should only have lower-case letters or numbers. | `eastus2`, `westus`, etc. | No | If empty, the driver will use the same region name as the current AKS cluster.
vnet-resource-group | The name of the resource group containing the virtual network to be connected to the AMLFS cluster. This resource group must already exist. | Resource group names can only include alphanumeric characters, underscores, parentheses, hyphens, periods (except at the end), and Unicode characters that match the allowed characters. | Yes | This value must be provided.
vnet-name | The name of the virtual network to be connected to the AMLFS cluster. This virtual network must already exist. | The name must begin with a letter or number, end with a letter, number, or underscore, and may contain only letters, numbers, underscores, periods, or hyphens. | Yes | This value must be provided.
subnet-name | The name of the subnet within the virtual network to be connected to the AMLFS cluster. This subnet must already exist. | The name must begin with a letter or number, end with a letter, number, or underscore, and may contain only letters, numbers, underscores, periods, or hyphens. | Yes | This value must be provided.
maintenance-day-of-week | The day of the week for maintenance to be performed on the AMLFS cluster. | `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday` | Yes | This value must be provided.
time-of-day-utc | The time (in UTC) when the maintenance window can begin on the AMLFS cluster. | Time value can only be in 24-hour format i.e., HH:MM | Yes | This value must be provided.
sku-name | SKU name for the Azure Managed Lustre file system. The SKU determines the throughput of the AMLFS cluster. | The SKU value must be one of the following: `AMLFS-Durable-Premium-40`, `AMLFS-Durable-Premium-125`, `AMLFS-Durable-Premium-250`, `AMLFS-Durable-Premium-500`. | Yes | This value must be provided.
zones | The availability zone is where your resource will be created. For the best performance, locate your AMLFS cluster in the same region and availability zone that houses your AKS cluster and other compute clients. | The zones must be either `"1"`, `"2"`, `"3"`, or a comma-separated list of multiple zones i.e., `"1,2"`. | Yes | This value must be provided.
resource-group-name | The name of the resource group in which to create the AMLFS cluster. This resource group must already exist. | Resource group names can only include alphanumeric characters, underscores, parentheses, hyphens, periods (except at the end), and Unicode characters that match the allowed characters. | No | If empty, the driver will use the AKS infrastructure resource group.
identities | User-assigned identity to assign to the AMLFS cluster. This identity must already exist. | This must be the resource identifier for the identity i.e., `"/subscriptions/12345678-1234-1234-1234-123456789abc/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myManagedIdentity"` | No | None
tags | Tags to apply to the AMLFS cluster resource. These tags do not affect AMLFS cluster functionality. | The tag name has a limit of 512 characters and the tag value has a limit of 256 characters. Tag names can't contain these characters: `<, >, %, &, \, ?, /`. | No | None
sub-dir | Mount into a subdirectory of the AMLFS cluster. This subdirectory does not need to exist. | This must be a valid Linux file path. | No | None

## Static Provisioning

Name | Meaning | Available Value | Mandatory | Default value
--- | --- | --- | --- | ---
mgs-ip-address | The IP address of the Lustre MGS, see AMLFS cluster details. | Must be a valid IP address i.e., `x.x.x.x` | Yes | This value must be provided.
sub-dir | Mount into a subdirectory of the AMLFS cluster. This subdirectory does not need to exist. | This must be a valid Linux file path. | No | None, will default to mounting the root directory of the AMLFS cluster.