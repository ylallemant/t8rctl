package group

import (
	"fmt"
)

var (
	_ ClusterGroup = &clusterGroup{}
)

func FromString(name string) (ClusterGroup, error) {
	switch name {
	case Workload:
		return New(name, descriptions[Workload])
	case Concierge:
		return New(name, descriptions[Concierge])

	default:
		return nil, fmt.Errorf("cluster group \"%s\" unknown", name)
	}
}

func List() ([]ClusterGroup, error) {
	list := make([]ClusterGroup, 0)

	for _, name := range instances {
		element, _ := New(name, descriptions[name])
		list = append(list, element)
	}

	return list, nil
}
