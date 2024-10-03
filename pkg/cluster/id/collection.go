package id

import (
	"fmt"
)

var (
	_ ClusterId = &clusterId{}
)

func FromString(name string) (ClusterId, error) {
	switch name {
	case Blue:
		return New(name)
	case Green:
		return New(name)

	default:
		return nil, fmt.Errorf("cluster id \"%s\" unknown", name)
	}
}

func List() ([]ClusterId, error) {
	list := make([]ClusterId, 0)

	for _, name := range instances {
		element, _ := New(name)
		list = append(list, element)
	}

	return list, nil
}
