package options

var (
	Domain  = "t8rctl"
	Current = NewOptions()
)

func NewOptions() *Options {
	options := new(Options)

	options.Output = "table"

	return options
}

type Options struct {
	All              bool
	FallbackDatatier string
	Output           string
}
