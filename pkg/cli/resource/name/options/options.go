package options

var (
	Domain  = "t8rctl"
	Current = NewOptions()
)

func NewOptions() *Options {
	options := new(Options)

	options.Tenant = "none"
	options.Region = "none"
	options.Short = false
	options.Hash = false
	options.Core = false
	options.Caf = false

	return options
}

type Options struct {
	Name   string
	Static string
	Type   string
	Tenant string
	Region string
	Short  bool
	Caf    bool
	Hash   bool
	Core   bool
}
