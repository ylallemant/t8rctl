package options

import "github.com/ylallemant/t8rctl/pkg/api"

var (
	Domain  = "t8rctl"
	Current = NewOptions()
)

func NewOptions() *Options {
	options := new(Options)

	options.Tenant = api.DefaultTenant
	options.Region = api.DefaultRegion

	return options
}

type Options struct {
	CurrentGitBranch string
	BaseGitBranch    string
	Tenant           string
	Region           string
	OnlyVariance     bool
}
