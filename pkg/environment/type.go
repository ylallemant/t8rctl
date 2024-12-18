package environment

import "github.com/ylallemant/t8rctl/pkg/api"

var _ api.EnvironmentContext = &environmentContext{}

func New() api.EnvironmentContext {
	return nil
}

type environmentContext struct {
	project         string
	cellProvider    string
	cellId          string
	cellStage       string
	cellRegion      string
	cellTenant      string
	clusterId       string
	clusterStage    string
	clusterGroup    string
	clusterRegion   string
	clusterTenant   string
	stack           string
	stackDatatier   string
	stackRegion     string
	stackTenant     string
	gitRepository   string
	gitBranchName   string
	gitBranchHash   string
	gitBranchTag    string
	isGitBaseBranch bool
}

func (i *environmentContext) Project() string {
	return i.project
}

func (i *environmentContext) CellProvider() string {
	return i.cellProvider
}

func (i *environmentContext) CellId() string {
	return i.cellId
}

func (i *environmentContext) CellStage() string {
	return i.cellStage
}

func (i *environmentContext) CellRegion() string {
	return i.cellRegion
}

func (i *environmentContext) CellTenant() string {
	return i.cellTenant
}

func (i *environmentContext) ClusterId() string {
	return i.clusterId
}

func (i *environmentContext) ClusterStage() string {
	return i.clusterStage
}

func (i *environmentContext) ClusterGroup() string {
	return i.clusterGroup
}

func (i *environmentContext) ClusterRegion() string {
	return i.clusterRegion
}

func (i *environmentContext) ClusterTenant() string {
	return i.clusterTenant
}

func (i *environmentContext) Stack() string {
	return i.stack
}

func (i *environmentContext) StackDatatier() string {
	return i.stackDatatier
}

func (i *environmentContext) StackTenant() string {
	return i.stackTenant
}

func (i *environmentContext) StackRegion() string {
	return i.stackRegion
}

func (i *environmentContext) GitRepository() string {
	return i.gitRepository
}

func (i *environmentContext) GitBranchName() string {
	return i.gitBranchName
}

func (i *environmentContext) GitBranchHash() string {
	return i.gitBranchHash
}

func (i *environmentContext) GitBranchTag() string {
	return i.gitBranchTag
}

func (i *environmentContext) IsGitBaseBranch() bool {
	return i.isGitBaseBranch
}
