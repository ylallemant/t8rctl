package options

var (
	Domain  = "t8rctl"
	Current = NewOptions()
)

func NewOptions() *Options {
	options := new(Options)

	return options
}

type Options struct {
	Semver bool
	Commit bool
}
