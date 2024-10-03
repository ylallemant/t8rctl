package cache

import (
	"strings"
	"sync"

	"github.com/pkg/errors"
	"github.com/ylallemant/t8rctl/pkg/api"
)

var CurrentManager *cacheManager

func init() {
	CurrentManager = new(cacheManager)
	CurrentManager.caches = make([]*cache, 0)
}

type cacheManager struct {
	caches []*cache
	mux    sync.RWMutex
}

func (i *cacheManager) register(cache *cache) {
	i.mux.Lock()
	defer i.mux.Unlock()

	found := i.Filter(cache.path)

	if len(found) == 0 {
		i.caches = append(i.caches, cache)
	}
}

func (i *cacheManager) List() []api.Cache {
	i.mux.RLock()
	defer i.mux.RUnlock()

	casted := make([]api.Cache, 0)

	for _, cache := range i.caches {
		casted = append(casted, cache)
	}

	return casted
}

func (i *cacheManager) Purge(provider string) error {
	i.mux.RLock()
	defer i.mux.RUnlock()

	filtered := i.Filter(provider)

	for _, cache := range filtered {
		err := cache.Purge()
		if err != nil {
			return errors.Wrapf(err, "could not purge cache %s", cache.path)
		}
	}

	return nil
}

func (i *cacheManager) Filter(search string) []*cache {
	if search == "" {
		return i.caches
	}

	filtered := make([]*cache, 0)

	for _, cache := range i.caches {
		if strings.Contains(cache.path, search) {
			filtered = append(filtered, cache)
		}
	}

	return filtered
}
