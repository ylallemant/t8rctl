package options

var (
	Domain  = "datatier"
	Current = NewOptions()
)

func NewOptions() *Options {
	options := new(Options)
	return options
}

type Options struct {
	Provider               string
	StackDatatier          string
	Group                  string
	DefaultClusterDatatier string
}
