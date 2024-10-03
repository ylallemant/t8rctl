package api

type Account interface {
	Provider() string
	Id() string
	Name() string
	Property(string) string
	HasProperty(string) bool
}

type AccountManager interface {
	Provider() string
	List() ([]Account, error)
	FromId(string) (Account, error)
}
