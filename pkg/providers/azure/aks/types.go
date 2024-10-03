package aks

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice"
	"github.com/ylallemant/t8rctl/pkg/api"
	"github.com/ylallemant/t8rctl/pkg/cluster"
	"github.com/ylallemant/t8rctl/pkg/providers/azure/resourcegroup"
	"github.com/ylallemant/t8rctl/pkg/providers/azure/subscription"
)

func NewCluster(aksCluster *armcontainerservice.ManagedCluster) (Cluster, error) {
	c := Cluster{}

	subscriptionId, resourceGroup := resourcegroup.InfoFromId(*aksCluster.ID)

	c.Id = *aksCluster.ID
	c.Name = *aksCluster.Name
	c.ResourceGroup = resourceGroup
	c.SubscriptionId = subscriptionId
	c.Location = *aksCluster.Location
	c.KubernetesVersion = *aksCluster.Properties.KubernetesVersion

	c.Tags = map[string]string{}
	for key, value := range aksCluster.Tags {
		c.Tags[key] = *value
	}

	return c, nil
}

type Cluster struct {
	Id                string
	Name              string
	TagId             string
	TagDatatier       string
	TagGroup          string
	Location          string
	ResourceGroup     string
	KubernetesVersion string
	SubscriptionId    string
	Active            bool
	Tags              map[string]string
}

func (c *Cluster) convert() (api.Cluster, error) {
	subscriptionAccount, err := subscription.Current.FromId(c.SubscriptionId)
	if err != nil {
		return nil, err
	}

	managed := api.ClusterManaged(c.Tags)

	return cluster.NewCluster(
		api.Azure,
		c.Name,
		c.TagId,
		c.TagGroup,
		c.TagDatatier,
		c.Location,
		c.ResourceGroup,
		c.KubernetesVersion,
		subscriptionAccount,
		managed,
		c.Active,
	)
}
