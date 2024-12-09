package cluster

import (
	"github.com/ylallemant/t8rctl/pkg/api"
)

var _ api.Cluster = &cluster{}

func NewCluster(provider, rawName, id, group, datatier, region, section, version string, account api.Account, managed, active bool, tags map[string]string) (*cluster, error) {
	instance := new(cluster)

	instance.CName = rawName
	instance.managed = managed
	instance.tags = tags

	if managed {
		instance.CGroup = instance.Tags()[api.TAG_CLUSTER_GROUP]
		instance.CDatatier = instance.Tags()[api.TAG_CLUSTER_STAGE]
		instance.CID = instance.Tags()[api.TAG_CLUSTER_ID]
	} else {
		instance.CGroup = "none"
		instance.CDatatier = "none"
		instance.CID = "none"
	}

	instance.CProvider = provider
	instance.CRegion = region
	instance.CAccount = account
	instance.section = section
	instance.version = version
	instance.active = active

	return instance, nil
}

type cluster struct {
	CName     string
	CGroup    string
	CDatatier string
	CID       string
	CProvider string
	CRegion   string
	CAccount  api.Account
	section   string
	version   string
	managed   bool
	active    bool
	tags      map[string]string
}

func (c *cluster) Name() string {
	return c.CName
}

func (c *cluster) Group() string {
	return c.CGroup
}

func (c *cluster) Datatier() string {
	return c.CDatatier
}

func (c *cluster) Id() string {
	return c.CID
}

func (c *cluster) Provider() string {
	return c.CProvider
}

func (c *cluster) Region() string {
	return c.CRegion
}

func (c *cluster) Account() api.Account {
	return c.CAccount
}

func (c *cluster) Section() string {
	return c.section
}

func (c *cluster) Version() string {
	return c.version
}

func (c *cluster) Managed() bool {
	return c.managed
}

func (c *cluster) Active() bool {
	return c.active
}

func (c *cluster) Tags() map[string]string {
	return c.tags
}

func FilterManaged(clusters []api.Cluster) []api.Cluster {
	filtered := make([]api.Cluster, 0)

	for _, candidate := range clusters {
		if candidate.Managed() {
			filtered = append(filtered, candidate)
		}
	}

	return filtered
}

func Filter(clusters []api.Cluster, filter api.ClusterFilter) []api.Cluster {
	filtered := make([]api.Cluster, 0)

	for _, cluster := range clusters {
		if Ignore(cluster, filter) {
			continue
		}

		filtered = append(filtered, cluster)
	}

	return filtered
}

func Ignore(cluster api.Cluster, filter api.ClusterFilter) bool {

	if filter.All {
		return false
	}

	if !cluster.Active() && !filter.ShowInactive {
		return true
	}

	if filter.Provider != "" && cluster.Provider() != filter.Provider {
		return true
	}

	if filter.Datatier != "" && cluster.Datatier() != filter.Datatier {
		return true
	}

	if filter.Group != "" && cluster.Group() != filter.Group {
		return true
	}

	if filter.Id != "" && cluster.Id() != filter.Id {
		return true
	}

	return false
}
