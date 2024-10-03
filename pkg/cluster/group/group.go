package group

var (
	Workload     = "workload"
	Concierge    = "concierge"
	descriptions = map[string]string{
		Workload:  "main group used to run the team stacks",
		Concierge: "group used to run infrastructure tooling",
	}
	instances = []string{
		Workload,
		Concierge,
	}
)

func New(name, description string) (*clusterGroup, error) {
	instance := new(clusterGroup)

	instance.name = name
	instance.description = description

	return instance, nil
}

type clusterGroup struct {
	name        string
	description string
}

func (i *clusterGroup) Name() string {
	return i.name
}

func (i *clusterGroup) Description() string {
	return i.description
}
