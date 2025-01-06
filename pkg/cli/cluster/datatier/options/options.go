package options

import "github.com/ylallemant/t8rctl/pkg/api"

var (
	Domain  = "datatier"
	Current = NewOptions()
)

func NewOptions() *Options {
	options := new(Options)

	options.Provider = api.Azure

	return options
}

type Options struct {
	Provider               string
	StackDatatier          string
	Group                  string
	DefaultClusterDatatier string
}
