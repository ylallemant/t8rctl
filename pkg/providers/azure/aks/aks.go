package aks

import (
	"context"
	"fmt"
	"path/filepath"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice"
	"github.com/pkg/errors"
	"github.com/ylallemant/t8rctl/pkg/api"
	"github.com/ylallemant/t8rctl/pkg/cache"
	"github.com/ylallemant/t8rctl/pkg/command"
	"github.com/ylallemant/t8rctl/pkg/global"
	"github.com/ylallemant/t8rctl/pkg/providers/azure/credentials"
	"github.com/ylallemant/t8rctl/pkg/providers/azure/vault"
	"gopkg.in/yaml.v3"
)

var (
	Current *AksClient
	_       api.ClusterManager = &AksClient{}
)

func New() (*AksClient, error) {
	if Current != nil {
		return Current, nil
	}

	Current = new(AksClient)

	Current.activityCache = map[string]map[string]map[string]bool{}
	Current.clients = make([]*armcontainerservice.ManagedClustersClient, 0)

	hmgServicesClient, err := armcontainerservice.NewManagedClustersClient("15bc7278-fb36-46fc-9f9c-eea2b20bf9c9", credentials.Current, &arm.ClientOptions{})
	if err != nil {
		return nil, err
	}

	hmgInfraClient, err := armcontainerservice.NewManagedClustersClient("609aece9-8cbe-48d3-8ef5-510ad67699fa", credentials.Current, &arm.ClientOptions{})
	if err != nil {
		return nil, err
	}

	/*
		hmgLivingdocsClient, err := armcontainerservice.NewManagedClustersClient("7d66e7f3-cfa2-4d28-9aa8-18a1e56572fe", credentials.Current, &arm.ClientOptions{})
		if err != nil {
			return nil, err
		}
	*/

	Current.clients = append(Current.clients, hmgInfraClient)
	Current.clients = append(Current.clients, hmgServicesClient)
	// Current.clients = append(Current.clients, hmgLivingdocsClient)

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

type AksClient struct {
	clients       []*armcontainerservice.ManagedClustersClient
	cache         []api.Cluster
	fsCache       api.Cache
	activityCache map[string]map[string]map[string]bool
	mux           sync.RWMutex
}

func (i *AksClient) CacheFile() string {
	return filepath.Join(cache.BasePath(), api.Azure, "clusters.yaml")
}

func (i *AksClient) cacheWrite(clusters []Cluster) error {
	content, err := yaml.Marshal(clusters)
	if err != nil {
		return errors.Wrapf(err, "could generate yaml content for cache file")
	}

	err = i.fsCache.Write(content)
	if err != nil {
		return errors.Wrapf(err, "could not write content yaml to cache file")
	}

	return nil
}

func (i *AksClient) cacheRead() error {
	clusters := make([]Cluster, 0)

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

func (i *AksClient) updateCache(clusters []Cluster) error {
	i.mux.Lock()
	defer i.mux.Unlock()

	i.cache = make([]api.Cluster, 0)

	for _, aks := range clusters {
		c, err := aks.convert()
		if err != nil {
			return errors.Wrapf(err, "could not convert Azure cluster struct to generic cluster struct")
		}

		i.cache = append(i.cache, c)
	}

	return nil
}

func (i *AksClient) Provider() string {
	return api.Azure
}

func (i *AksClient) PurgeCache() error {
	return nil
}

func (i *AksClient) List(subs api.AccountManager) ([]api.Cluster, error) {
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

	clusters := make([]Cluster, 0)

	for _, client := range i.clients {

		pager := client.NewListPager(&armcontainerservice.ManagedClustersClientListOptions{})

		for pager.More() {
			response, err := pager.NextPage(context.TODO())
			if err != nil {
				return i.cache, err
			}

			for _, azCluster := range response.Value {
				azureCluster, err := NewCluster(azCluster)
				if err != nil {
					return nil, errors.Wrapf(err, "could not convert Azure Dataset to Azure cluster struct")
				}

				managed := api.ClusterManaged(azureCluster.Tags)

				if managed {
					azureCluster.TagId = azureCluster.Tags[api.TAG_CLUSTER_ID]
					azureCluster.TagDatatier = azureCluster.Tags[api.TAG_DATATIER]
					azureCluster.TagGroup = azureCluster.Tags[api.TAG_CLUSTER_GROUP]
					azureCluster.Active, err = i.CheckActivity(azureCluster.TagId, azureCluster.TagGroup, azureCluster.TagDatatier)
					if err != nil {
						return nil, errors.Wrapf(err, "failed to check activity status for cluster %s", azureCluster.Name)
					}
				}

				clusters = append(clusters, azureCluster)
			}
		}
	}

	if len(clusters) > 0 {
		err := i.cacheWrite(clusters)
		if err != nil {
			return i.cache, errors.Wrapf(err, "could not write cache file %s", i.fsCache.Path())
		}

		err = i.updateCache(clusters)
		if err != nil {
			return i.cache, errors.Wrapf(err, "could not update in memory cache")
		}
	}

	return i.cache, nil
}

func (i *AksClient) Connect(cluster api.Cluster, context string) error {
	cmd := command.New("az")
	cmd.AddArg("aks")
	cmd.AddArg("get-credentials")
	cmd.AddArg("--overwrite-existing")
	cmd.AddArg("--name")
	cmd.AddArg(cluster.Name())
	cmd.AddArg("--resource-group")
	cmd.AddArg(cluster.Section())
	cmd.AddArg("--subscription")
	cmd.AddArg(cluster.Account().Id())

	if context != "" {
		cmd.AddArg("--context")
		cmd.AddArg(context)
	}

	_, err := cmd.Execute()

	return err
}

func (i *AksClient) CheckActivity(id, group, datatier string) (bool, error) {
	if _, found := i.activityCache[group]; found {
		if _, found := i.activityCache[group][datatier]; found {
			if status, found := i.activityCache[group][datatier][id]; found {
				return status, nil
			}
		} else {
			i.activityCache[group][datatier] = map[string]bool{}
		}
	} else {
		i.activityCache[group] = map[string]map[string]bool{}
		i.activityCache[group][datatier] = map[string]bool{}
	}

	activeId, err := vault.Current.Get("hmg-shared", fmt.Sprintf("cluster-%s-active-id-%s", group, datatier))
	if err != nil {
		return false, err
	}

	i.activityCache[group][datatier][id] = (id == activeId)
	//fmt.Println("is debug logging enabled:", log.Debug().Enabled())
	//log.Debug().Msgf("is %s-%s-%s active (%s): %v", group, datatier, id, activeId, i.activityCache[group][datatier][id])

	return i.activityCache[group][datatier][id], nil
}
