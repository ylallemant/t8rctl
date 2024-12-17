package api

type EnvironmentContext struct {
	Project    string
	Region     string
	Tenant     string
	Datatier   string
	BranchName string
	IsBase     bool
}
