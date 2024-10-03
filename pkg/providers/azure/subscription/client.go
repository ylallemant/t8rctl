package subscription

import (
	"context"
	"fmt"
	"path/filepath"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
	"github.com/pkg/errors"
	"github.com/ylallemant/t8rctl/pkg/api"
	"github.com/ylallemant/t8rctl/pkg/cache"
	"github.com/ylallemant/t8rctl/pkg/global"
	"github.com/ylallemant/t8rctl/pkg/providers/azure/credentials"
	"gopkg.in/yaml.v3"
)

var (
	Current *SubscriptionClient
	_       api.AccountManager = &SubscriptionClient{}
)

func New() (*SubscriptionClient, error) {
	if Current != nil {
		return Current, nil
	}

	Current = new(SubscriptionClient)

	aksClient, err := armsubscriptions.NewClient(credentials.Current, &arm.ClientOptions{})
	if err != nil {
		return nil, err
	}

	Current.client = aksClient
	Current.cache = make([]api.Account, 0)

	Current.fsCache, err = cache.New(Current.CacheFile(), cache.DefaultTTL)
	if err != nil {
		return nil, err
	}

	if Current.fsCache.Valid() && !global.Current.DisableCache {
		err := Current.cacheRead()
		if err != nil {
			return nil, errors.Wrapf(err, "could perform initial cache file read %s", Current.fsCache.Path())
		}
	}

	return Current, nil
}

type SubscriptionClient struct {
	client  *armsubscriptions.Client
	cache   []api.Account
	fsCache api.Cache
	mux     sync.RWMutex
}

func (i *SubscriptionClient) CacheFile() string {
	return filepath.Join(cache.BasePath(), api.Azure, "subscriptions.yaml")
}

func (i *SubscriptionClient) cacheWrite(subscriptions []Subscription) error {
	content, err := yaml.Marshal(subscriptions)
	if err != nil {
		return errors.Wrapf(err, "could generate yaml content for cache file")
	}

	err = i.fsCache.Write(content)
	if err != nil {
		return errors.Wrapf(err, "could not write content yaml to cache file")
	}

	return nil
}

func (i *SubscriptionClient) cacheRead() error {
	clusters := make([]Subscription, 0)

	content, err := i.fsCache.Read()
	if err != nil {
		return errors.Wrapf(err, "could not read yaml content from cache file %s", i.fsCache.Path())
	}

	err = yaml.Unmarshal(content, &clusters)
	if err != nil {
		return errors.Wrapf(err, "could not unmarshal yaml content from cache file %s", i.fsCache.Path())
	}

	err = i.updateCache(clusters)
	if err != nil {
		return errors.Wrapf(err, "could not update in memory cache")
	}

	return nil
}

func (i *SubscriptionClient) updateCache(subscriptions []Subscription) error {
	i.mux.Lock()
	defer i.mux.Unlock()

	i.cache = make([]api.Account, 0)

	for _, subscription := range subscriptions {
		c, err := subscription.convert()
		if err != nil {
			return errors.Wrapf(err, "could not convert Azure subscription struct to generic account struct")
		}

		i.cache = append(i.cache, c)
	}

	return nil
}

func (i *SubscriptionClient) Provider() string {
	return api.Azure
}

func (i *SubscriptionClient) List() ([]api.Account, error) {
	if len(i.cache) > 0 {
		i.mux.RLock()
		defer i.mux.RUnlock()

		return i.cache, nil
	}

	if i.fsCache.Valid() && !global.Current.DisableCache {
		err := i.cacheRead()
		if err != nil {
			return i.cache, errors.Wrapf(err, "could not read cache file %s", i.fsCache.Path())
		}

		return i.cache, nil
	}

	subscriptions := make([]Subscription, 0)

	pager := i.client.NewListPager(&armsubscriptions.ClientListOptions{})

	for pager.More() {
		response, err := pager.NextPage(context.TODO())
		if err != nil {
			return i.cache, err
		}

		for _, subscription := range response.Value {
			subscriptions = append(subscriptions, NewSubscription(subscription))
		}
	}

	if len(subscriptions) > 0 {
		err := i.cacheWrite(subscriptions)
		if err != nil {
			return i.cache, errors.Wrapf(err, "could not write cache file %s", i.fsCache.Path())
		}

		err = i.updateCache(subscriptions)
		if err != nil {
			return i.cache, errors.Wrapf(err, "could not update in memory cache")
		}
	}

	return i.cache, nil
}

func (i *SubscriptionClient) Fetch(subscriptionId string) (api.Account, error) {
	raw, err := i.client.Get(context.TODO(), subscriptionId, &armsubscriptions.ClientGetOptions{})
	if err != nil {
		return nil, err
	}

	subscription := NewSubscription(&raw.Subscription)

	return subscription.convert()
}

func (i *SubscriptionClient) FromId(subscriptionId string) (api.Account, error) {
	if len(i.cache) == 0 {
		i.List()
	}

	i.mux.RLock()
	defer i.mux.RUnlock()

	for _, subscription := range i.cache {
		if subscription.Id() == subscriptionId {
			return subscription, nil
		}
	}

	return nil, fmt.Errorf("subscription with id %s not found", subscriptionId)
}
