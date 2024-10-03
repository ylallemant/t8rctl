package azure

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/ylallemant/t8rctl/pkg/api"
	"github.com/ylallemant/t8rctl/pkg/cache"
	"github.com/ylallemant/t8rctl/pkg/providers/azure/aks"
	"github.com/ylallemant/t8rctl/pkg/providers/azure/subscription"
	"github.com/ylallemant/t8rctl/pkg/providers/azure/vault"
)

var _ api.Provider = &azure{}

func New() (*azure, error) {
	instance := new(azure)

	keyvaultClient, err := vault.New()
	if err != nil {
		return nil, err
	}

	SubscriptionClient, err := subscription.New()
	if err != nil {
		return nil, err
	}
	instance.subscription = SubscriptionClient
	instance.subscription.List()
	if err != nil {
		return nil, err
	}

	aksClient, err := aks.New()
	if err != nil {
		return nil, err
	}

	instance.aks = aksClient

	instance.accounts = SubscriptionClient
	instance.clusters = aksClient
	instance.vaults = keyvaultClient

	return instance, nil
}

type azure struct {
	cred         *azidentity.DefaultAzureCredential
	aks          *aks.AksClient
	subscription *subscription.SubscriptionClient
	accounts     api.AccountManager
	clusters     api.ClusterManager
	vaults       api.VaultManager
}

func (i *azure) Type() string {
	return api.Azure
}

func (i *azure) Accounts() api.AccountManager {
	return i.accounts
}

func (i *azure) Clusters() api.ClusterManager {
	return i.clusters
}

func (i *azure) Vaults() api.VaultManager {
	return i.vaults
}

func (i *azure) PurgeCaches() error {
	return cache.CurrentManager.Purge(api.Azure)
}
