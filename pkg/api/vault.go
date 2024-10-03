package api

type Vault interface {
	Provider() string
	Id() string
	Name() string
}

type VaultManager interface {
	Provider() string
	List(string) (map[string]string, error)
	Get(string, string) (string, error)
}
