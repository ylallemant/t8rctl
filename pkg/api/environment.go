package api

const (
	ENVAR_PLATFORM_PROJECT        = "PLATFORM_PROJECT"
	ENVAR_PLATFORM_CELL_ID        = "PLATFORM_CELL_ID"
	ENVAR_PLATFORM_CELL_STAGE     = "PLATFORM_CELL_STAGE"
	ENVAR_PLATFORM_CELL_REGION    = "PLATFORM_CELL_REGION"
	ENVAR_PLATFORM_CELL_TENANT    = "PLATFORM_CELL_TENANT"
	ENVAR_PLATFORM_CLUSTER_ID     = "PLATFORM_CLUSTER_ID"
	ENVAR_PLATFORM_CLUSTER_STAGE  = "PLATFORM_CLUSTER_STAGE"
	ENVAR_PLATFORM_CLUSTER_GROUP  = "PLATFORM_CLUSTER_GROUP"
	ENVAR_PLATFORM_CLUSTER_REGION = "PLATFORM_CLUSTER_REGION"
	ENVAR_PLATFORM_CLUSTER_TENANT = "PLATFORM_CLUSTER_TENANT"
	ENVAR_PLATFORM_STACK          = "PLATFORM_STACK"
	ENVAR_PLATFORM_STACK_DATATIER = "PLATFORM_STACK_DATATIER"
	ENVAR_PLATFORM_STACK_REGION   = "PLATFORM_STACK_REGION"
	ENVAR_PLATFORM_STACK_TENANT   = "PLATFORM_STACK_TENANT"
	ENVAR_GIT_BRANCH_NAME         = "GIT_BRANCH_NAME"
	ENVAR_GIT_BRANCH_HASH         = "GIT_BRANCH_HASH"
	ENVAR_GIT_BRANCH_TAG          = "GIT_BRANCH_TAG"
	ENVAR_GIT_BASE_BRANCH_NAME    = "GIT_BASE_BRANCH_NAME"

	DefaultEnvironmentStage  = "all"
	DefaultEnvironmentRegion = "global"
	DefaultEnvironmentTenant = "shared"
	DefaultEnvironmentNotSet = "none"

	DefaultGitBaseBranchName = "main"
)

type EnvironmentContextOptions struct {
	Project           string
	CellId            string
	CellProvider      string
	CellStage         string
	CellRegion        string
	CellTenant        string
	ClusterId         string
	ClusterStage      string
	ClusterGroup      string
	ClusterRegion     string
	ClusterTenant     string
	Stack             string
	StackDatatier     string
	StackRegion       string
	StackTenant       string
	GitRepository     string
	GitBranchName     string
	GitBranchHash     string
	GitBranchTag      string
	GitBaseBranchName string
}

type EnvironmentContext interface {
	Project() string
	CellProvider() string
	CellId() string
	CellStage() string
	CellRegion() string
	CellTenant() string
	ClusterId() string
	ClusterStage() string
	ClusterGroup() string
	ClusterRegion() string
	ClusterTenant() string
	Stack() string
	StackDatatier() string
	StackRegion() string
	StackTenant() string
	GitRepository() string
	GitBranchName() string
	GitBranchHash() string
	GitBranchTag() string
	IsGitBaseBranch() bool
}
