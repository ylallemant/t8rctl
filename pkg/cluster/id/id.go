package id

var (
	Blue      = "blue"
	Green     = "green"
	instances = []string{
		Blue,
		Green,
	}
)

func New(name string) (*clusterId, error) {
	instance := new(clusterId)

	instance.name = name

	return instance, nil
}

type clusterId struct {
	name string
}

func (i *clusterId) Name() string {
	return i.name
}
