package account

import "github.com/ylallemant/t8rctl/pkg/api"

var _ api.Account = &account{}

func New(provider, id, name string, properties map[string]string) (*account, error) {
	instance := new(account)

	instance.SProvider = provider
	instance.SId = id
	instance.SName = name
	instance.SProperties = properties

	return instance, nil
}

type account struct {
	SProvider   string
	SId         string
	SName       string
	SProperties map[string]string
}

func (i *account) Provider() string {
	return api.Azure
}

func (i *account) Id() string {
	return i.SId
}

func (i *account) Name() string {
	return i.SName
}

func (i *account) Property(property string) string {
	if value, found := i.SProperties[property]; found {
		return value
	}

	return ""
}

func (i *account) HasProperty(property string) bool {
	_, found := i.SProperties[property]
	return found
}
