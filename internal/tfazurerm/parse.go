package tfazurerm

import "fmt"

// Ported from VaultID: https://github.com/terraform-providers/terraform-provider-azurerm/blob/64ea516c0be29868a6c1ab14b6999c547b6184cb/azurerm/internal/services/keyvault/parse/vault.go#L41
// and VirtualNetworkID: https://github.com/terraform-providers/terraform-provider-azurerm/blob/64ea516c0be29868a6c1ab14b6999c547b6184cb/azurerm/internal/services/network/parse/virtual_network.go

// Parse Name and ResourceGroup name from an Azure Resource Manager ID
func ParseNameRGFromID(azureId string, resourceType string) (string, string, error) {
	id, err := ParseAzureResourceID(azureId)
	if err != nil {
		return "", "", err
	}

	if id.SubscriptionID == "" {
		return "", "", fmt.Errorf("ID was missing the 'subscriptions' element")
	}

	if id.ResourceGroup == "" {
		return "", "", fmt.Errorf("ID was missing the 'resourceGroups' element")
	}

	idName, err := id.PopSegment(resourceType)
	if err != nil {
		return "", "", err
	}

	if err := id.ValidateNoEmptySegments(azureId); err != nil {
		return "", "", err
	}

	return idName, id.ResourceGroup, nil
}
