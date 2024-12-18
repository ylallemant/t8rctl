package environment

import (
	"github.com/ylallemant/t8rctl/pkg/api"
)

func FromOptions(options *api.EnvironmentContextOptions, enforceTopDown bool) api.EnvironmentContext {
	env := new(environmentContext)

	ensureRelatedValues(options, enforceTopDown)

	env.project = options.Project
	env.cellProvider = options.CellProvider
	env.cellId = options.CellId
	env.cellStage = options.CellStage
	env.cellRegion = options.CellRegion
	env.cellTenant = options.CellTenant
	env.clusterId = options.ClusterId
	env.clusterStage = options.ClusterStage
	env.clusterGroup = options.ClusterGroup
	env.clusterRegion = options.ClusterRegion
	env.clusterTenant = options.ClusterTenant
	env.stack = options.Stack
	env.stackDatatier = options.StackDatatier
	env.stackRegion = options.StackRegion
	env.stackTenant = options.StackTenant
	env.gitRepository = options.GitRepository
	env.gitBranchName = options.GitBranchName
	env.gitBranchHash = options.GitBranchHash
	env.gitBranchTag = options.GitBranchTag
	env.isGitBaseBranch = options.GitBranchName == options.GitBaseBranchName

	return env
}

func ensureRelatedValues(options *api.EnvironmentContextOptions, enforceTopDown bool) {
	ensureStages(options, enforceTopDown)
	ensureRegions(options, enforceTopDown)
	ensureTenants(options, enforceTopDown)

	if options.GitBaseBranchName == "" {
		options.GitBaseBranchName = api.DefaultGitBaseBranchName
	}
}

func ensureStages(options *api.EnvironmentContextOptions, enforceTopDown bool) {
	defaultValue := api.DefaultEnvironmentStage

	if options.CellStage != "" && options.CellStage != api.DefaultEnvironmentNotSet && options.CellStage != api.DefaultEnvironmentStage {
		defaultValue = options.CellStage
	}

	if options.ClusterStage != "" && options.ClusterStage != api.DefaultEnvironmentNotSet && options.ClusterStage != api.DefaultEnvironmentStage && !enforceTopDown {
		defaultValue = options.ClusterStage
	}

	if options.CellStage == "" {
		options.CellStage = defaultValue
	}

	if options.ClusterStage == "" {
		options.ClusterStage = defaultValue
	}

	if options.StackDatatier == "" {
		options.StackDatatier = defaultValue
	}
}

func ensureRegions(options *api.EnvironmentContextOptions, enforceTopDown bool) {
	defaultValue := api.DefaultEnvironmentRegion

	if options.CellRegion != "" && options.CellRegion != api.DefaultEnvironmentNotSet && options.CellRegion != api.DefaultEnvironmentRegion {
		defaultValue = options.CellRegion
	}

	if options.ClusterRegion != "" && options.ClusterRegion != api.DefaultEnvironmentNotSet && options.ClusterRegion != api.DefaultEnvironmentRegion && !enforceTopDown {
		defaultValue = options.CellStage
	}

	if options.CellRegion == "" {
		options.CellRegion = defaultValue
	}

	if options.ClusterRegion == "" {
		options.ClusterRegion = defaultValue
	}

	if options.StackRegion == "" {
		options.StackRegion = defaultValue
	}
}

func ensureTenants(options *api.EnvironmentContextOptions, enforceTopDown bool) {
	defaultValue := api.DefaultEnvironmentTenant

	if options.CellTenant != "" && options.CellTenant != api.DefaultEnvironmentNotSet && options.CellTenant != api.DefaultEnvironmentTenant {
		defaultValue = options.CellTenant
	}

	if options.ClusterTenant != "" && options.ClusterTenant != api.DefaultEnvironmentNotSet && options.ClusterTenant != api.DefaultEnvironmentTenant && !enforceTopDown {
		defaultValue = options.ClusterTenant
	}

	if options.CellTenant == "" {
		options.CellTenant = defaultValue
	}

	if options.ClusterTenant == "" {
		options.ClusterTenant = defaultValue
	}

	if options.StackTenant == "" {
		options.StackTenant = defaultValue
	}
}
