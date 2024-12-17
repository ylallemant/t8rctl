package options

var (
	Domain  = "t8rctl"
	Current = NewOptions()
)

func NewOptions() *Options {
	options := new(Options)

	options.Tenant = "none"
	options.Region = "none"

	return options
}

type Options struct {
	CurrentGitBranch string
	BaseGitBranch    string
	Tenant           string
	Region           string
}
