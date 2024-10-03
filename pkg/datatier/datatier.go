package datatier

import "fmt"

var (
	Production  = "production"
	Staging     = "staging"
	Development = "development"
	instances   = []string{
		Production,
		Staging,
		Development,
	}
)

func New(name string) (*datatier, error) {
	instance := new(datatier)

	instance.name = name

	return instance, nil
}

type datatier struct {
	name  string
	alias string
}

func (i *datatier) Name() string {
	return i.name
}

func (i *datatier) Alias() string {
	if i.alias == "" {
		return i.Name()
	}

	return i.alias
}

func (i *datatier) SetAlias(alias string) error {
	if i.alias == "" {
		return fmt.Errorf("emply alias provided")
	}

	i.alias = alias

	return nil
}
