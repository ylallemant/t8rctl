package vault

import (
	"context"
	"fmt"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
	"github.com/ylallemant/t8rctl/pkg/api"
	"github.com/ylallemant/t8rctl/pkg/providers/azure/credentials"
)

var (
	Current *VaultClient
	_       api.VaultManager = &VaultClient{}
)

func New() (*VaultClient, error) {
	if Current != nil {
		return Current, nil
	}

	Current = new(VaultClient)

	Current.clients = make(map[string]*azsecrets.Client)
	Current.cache = make(map[string]map[string]string)

	return Current, nil
}

type VaultClient struct {
	clients map[string]*azsecrets.Client
	cache   map[string]map[string]string
	mux     sync.RWMutex
}

func (i *VaultClient) Provider() string {
	return api.Azure
}

func (i *VaultClient) List(vaultName string) (map[string]string, error) {
	keyvalue := make(map[string]string)

	if !i.initialized(vaultName) {
		err := i.connect(vaultName)
		if err != nil {
			return keyvalue, err
		}

		pager := i.clients[vaultName].NewListSecretsPager(nil)
		for pager.More() {
			page, err := pager.NextPage(context.TODO())
			if err != nil {
				return keyvalue, err
			}

			for _, secret := range page.Value {
				value, err := i.Get(vaultName, secret.ID.Name())
				if err != nil {
					return keyvalue, err
				}

				keyvalue[secret.ID.Name()] = value
			}
		}

		i.mux.Lock()
		i.cache[vaultName] = keyvalue
		i.mux.Unlock()
	}

	return i.cache[vaultName], nil
}

func (i *VaultClient) Get(vaultName, secret string) (string, error) {
	if !i.initialized(vaultName) {
		err := i.connect(vaultName)
		if err != nil {
			return "", err
		}
	}

	if i.cached(vaultName, secret) {
		i.mux.RLock()
		defer i.mux.RUnlock()

		return i.cache[vaultName][secret], nil
	}

	resp, err := i.clients[vaultName].GetSecret(context.TODO(), secret, "", nil)
	if err != nil {
		return "", err
	}

	i.mux.Lock()
	defer i.mux.Unlock()
	i.cache[vaultName][secret] = *resp.Value

	return i.cache[vaultName][secret], nil
}

func (i *VaultClient) cached(vaultName, secret string) bool {
	if !i.initialized(vaultName) {
		return false
	}

	i.mux.RLock()
	_, exists := i.cache[vaultName][secret]
	i.mux.RUnlock()

	return exists
}

func (i *VaultClient) initialized(vaultName string) bool {
	i.mux.RLock()
	_, exists := i.clients[vaultName]
	i.mux.RUnlock()

	return exists
}

func (i *VaultClient) connect(vaultName string) error {
	i.mux.Lock()
	defer i.mux.Unlock()

	url := fmt.Sprintf("https://%s.vault.azure.net", vaultName)

	client := azsecrets.NewClient(url, credentials.Current, &azsecrets.ClientOptions{})

	i.clients[vaultName] = client
	i.cache[vaultName] = make(map[string]string)

	return nil
}
