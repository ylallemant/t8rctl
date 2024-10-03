package options

var (
	Domain  = "t8rctl"
	Current = NewOptions()
)

func NewOptions() *Options {
	options := new(Options)

	options.Id = ""

	return options
}

type Options struct {
	Provider         string
	Datatier         string
	Group            string
	Id               string
	FallbackDatatier string
}
