package api

var (
	Azure = "azure"
)

type Provider interface {
	Type() string
	Vaults() VaultManager
	Clusters() ClusterManager
	Accounts() AccountManager
	PurgeCaches() error
}

type ProviderManager interface {
	Get(string) Provider
}
