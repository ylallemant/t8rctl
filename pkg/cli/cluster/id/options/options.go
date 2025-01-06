package options

import "github.com/ylallemant/t8rctl/pkg/api"

var (
	Domain  = "t8rctl"
	Current = NewOptions()
)

func NewOptions() *Options {
	options := new(Options)

	options.Provider = api.Azure

	return options
}

type Options struct {
	Provider         string
	Datatier         string
	Group            string
	FallbackDatatier string
}
