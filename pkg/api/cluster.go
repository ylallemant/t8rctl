package api

type Cluster interface {
	Provider() string
	Name() string
	Group() string
	Datatier() string
	Id() string
	Region() string
	Account() Account
	Section() string
	Version() string
	Managed() bool
	Active() bool
	Tags() map[string]string
}

type ClusterCache interface {
	Marshal([]Cluster) error
	Unmarshal([]Cluster, error)
}

type ClusterManager interface {
	Provider() string
	List(AccountManager) ([]Cluster, error)
	Connect(Cluster, string) error
	CheckActivity(id, group, datatier string) (bool, error)
}

type ClusterFilter struct {
	Provider     string
	Group        string
	Datatier     string
	Id           string
	All          bool
	ShowInactive bool
}

func ClusterManaged(tags map[string]string) bool {
	if _, found := tags[TAG_CLUSTER_GROUP]; !found {
		return false
	}

	if _, found := tags[TAG_CLUSTER_ID]; !found {
		return false
	}

	if _, found := tags[TAG_DATATIER]; !found {
		return false
	}

	return true
}
