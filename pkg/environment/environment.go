package environment

import (
	"github.com/ylallemant/t8rctl/pkg/api"
)

func FromOptions(options *api.EnvironmentContextOptions, enforceTopDown bool) api.EnvironmentContext {
	env := new(environmentContext)

	ensureRelatedValues(options, enforceTopDown)

	env.project = sanitiseValue(options.Project)
	env.cellProvider = sanitiseValue(options.CellProvider)
	env.cellId = sanitiseValue(options.CellId)
	env.cellStage = sanitiseValue(options.CellStage)
	env.cellRegion = sanitiseValue(options.CellRegion)
	env.cellTenant = sanitiseValue(options.CellTenant)
	env.clusterId = sanitiseValue(options.ClusterId)
	env.clusterStage = sanitiseValue(options.ClusterStage)
	env.clusterGroup = sanitiseValue(options.ClusterGroup)
	env.clusterRegion = sanitiseValue(options.ClusterRegion)
	env.clusterTenant = sanitiseValue(options.ClusterTenant)
	env.stack = sanitiseValue(options.Stack)
	env.stackDatatier = sanitiseValue(options.StackDatatier)
	env.stackRegion = sanitiseValue(options.StackRegion)
	env.stackTenant = sanitiseValue(options.StackTenant)
	env.gitRepository = sanitiseValue(options.GitRepository)
	env.gitBranchName = sanitiseValue(options.GitBranchName)
	env.gitBaseBranchName = sanitiseValue(options.GitBaseBranchName)
	env.gitBranchHash = sanitiseValue(options.GitBranchHash)
	env.gitBranchTag = sanitiseValue(options.GitBranchTag)
	env.isGitBaseBranch = env.gitBranchName == env.gitBaseBranchName

	return env
}

func ensureRelatedValues(options *api.EnvironmentContextOptions, enforceTopDown bool) {
	ensureStages(options, enforceTopDown)
	ensureRegions(options, enforceTopDown)
	ensureTenants(options, enforceTopDown)

	if options.Stack == "" {
		options.Stack = options.Project
	}

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
		defaultValue = options.ClusterRegion
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
