package datatier

import (
	"fmt"

	"github.com/ylallemant/t8rctl/pkg/api"
)

var (
	_ api.Datatier = &datatier{}
)

func FromString(name string) (api.Datatier, error) {
	switch name {
	case Production:
		return New(name)
	case Staging:
		return New(name)
	case Development:
		return New(name)

	default:
		return nil, fmt.Errorf("datatier \"%s\" unknown", name)
	}
}

func List() ([]api.Datatier, error) {
	list := make([]api.Datatier, 0)

	for _, name := range instances {
		element, _ := New(name)
		list = append(list, element)
	}

	return list, nil
}
