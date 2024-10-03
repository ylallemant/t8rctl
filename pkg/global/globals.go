package global

var (
	Current = new(Globals)
)

type Globals struct {
	Debug        bool
	DisableCache bool
}
