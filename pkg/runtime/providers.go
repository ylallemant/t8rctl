package runtime

import (
	"sync"

	"github.com/ylallemant/t8rctl/pkg/api"
	"github.com/ylallemant/t8rctl/pkg/providers/azure"
)

var (
	_         api.ProviderManager = &providerManager{}
	Providers api.ProviderManager = nil
)

func init() {
	manager := new(providerManager)
	manager.cache = make(map[string]api.Provider)

	manager.mux.Lock()
	defer manager.mux.Unlock()

	azureProvider, err := azure.New()
	if err != nil {
		panic(err)
	}

	manager.cache[api.Azure] = azureProvider

	Providers = manager
}

type providerManager struct {
	cache map[string]api.Provider
	mux   sync.RWMutex
}

func (i *providerManager) Get(name string) api.Provider {
	i.mux.RLock()
	defer i.mux.RUnlock()

	if provider, exists := i.cache[name]; exists {
		return provider
	}

	return nil
}
