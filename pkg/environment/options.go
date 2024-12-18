package environment

import (
	"os"
	"strings"

	"github.com/ylallemant/t8rctl/pkg/api"
)

func OptionsFromEnvvars() *api.EnvironmentContextOptions {
	options := new(api.EnvironmentContextOptions)

	value, found := os.LookupEnv(api.ENVAR_PLATFORM_PROJECT)
	if found {
		options.Project = sanitiseValue(value)
	}

	value, found = os.LookupEnv(api.ENVAR_PLATFORM_CELL_ID)
	if found {
		options.CellId = sanitiseValue(value)
	}

	value, found = os.LookupEnv(api.ENVAR_PLATFORM_CELL_STAGE)
	if found {
		options.CellStage = sanitiseValue(value)
	}

	value, found = os.LookupEnv(api.ENVAR_PLATFORM_CELL_REGION)
	if found {
		options.CellRegion = sanitiseValue(value)
	}

	value, found = os.LookupEnv(api.ENVAR_PLATFORM_CELL_TENANT)
	if found {
		options.CellTenant = sanitiseValue(value)
	}

	value, found = os.LookupEnv(api.ENVAR_PLATFORM_CLUSTER_ID)
	if found {
		options.ClusterId = sanitiseValue(value)
	}

	value, found = os.LookupEnv(api.ENVAR_PLATFORM_CLUSTER_STAGE)
	if found {
		options.ClusterStage = sanitiseValue(value)
	}

	value, found = os.LookupEnv(api.ENVAR_PLATFORM_CLUSTER_GROUP)
	if found {
		options.ClusterGroup = sanitiseValue(value)
	}

	value, found = os.LookupEnv(api.ENVAR_PLATFORM_CLUSTER_REGION)
	if found {
		options.ClusterRegion = sanitiseValue(value)
	}

	value, found = os.LookupEnv(api.ENVAR_PLATFORM_CLUSTER_TENANT)
	if found {
		options.ClusterTenant = sanitiseValue(value)
	}

	value, found = os.LookupEnv(api.ENVAR_PLATFORM_STACK)
	if found {
		options.Stack = sanitiseValue(value)
	}

	value, found = os.LookupEnv(api.ENVAR_PLATFORM_STACK_DATATIER)
	if found {
		options.StackDatatier = sanitiseValue(value)
	}

	value, found = os.LookupEnv(api.ENVAR_PLATFORM_STACK_REGION)
	if found {
		options.StackRegion = sanitiseValue(value)
	}

	value, found = os.LookupEnv(api.ENVAR_PLATFORM_STACK_TENANT)
	if found {
		options.StackTenant = sanitiseValue(value)
	}

	value, found = os.LookupEnv(api.ENVAR_GIT_BRANCH_NAME)
	if found {
		options.GitBranchName = sanitiseValue(value)
	}

	value, found = os.LookupEnv(api.ENVAR_GIT_BRANCH_HASH)
	if found {
		options.GitBranchHash = sanitiseValue(value)
	}

	value, found = os.LookupEnv(api.ENVAR_GIT_BRANCH_TAG)
	if found {
		options.GitBranchTag = sanitiseValue(value)
	}

	value, found = os.LookupEnv(api.ENVAR_GIT_BASE_BRANCH_NAME)
	if found {
		options.GitBaseBranchName = sanitiseValue(value)
	}

	return options
}

func sanitiseValue(value string) string {
	value = strings.TrimSpace(value)
	value = strings.ToLower(value)
	return value
}
