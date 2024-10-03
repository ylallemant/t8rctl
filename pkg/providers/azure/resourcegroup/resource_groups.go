package resourcegroup

import "strings"

// TODO: remove this for a query against an API
func InfoFromId(azureResourceId string) (subscriptionId string, resourceGroup string) {
	if len(azureResourceId) > 0 {
		parts := strings.Split(azureResourceId, "/")

		if len(parts) > 6 {
			return parts[2], parts[4]
		}
	}

	return "", ""
}
