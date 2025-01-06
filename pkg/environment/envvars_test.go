package environment

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ylallemant/t8rctl/pkg/api"
)

func TestOptionsFromEnvvars(t *testing.T) {
	cases := []struct {
		name            string
		input           map[string]string
		expected        *api.EnvironmentContextOptions
		isGitBaseBranch bool
	}{
		{
			name:            "empty envvars",
			input:           map[string]string{},
			expected:        &api.EnvironmentContextOptions{},
			isGitBaseBranch: false,
		},
		{
			name: "test all envvars",
			input: map[string]string{
				api.ENVAR_PLATFORM_PROJECT:        "Project",
				api.ENVAR_PLATFORM_CELL_PROVIDER:  "CellProvider",
				api.ENVAR_PLATFORM_CELL_ID:        "CellId",
				api.ENVAR_PLATFORM_CELL_STAGE:     "CellStage",
				api.ENVAR_PLATFORM_CELL_REGION:    "CellRegion",
				api.ENVAR_PLATFORM_CELL_TENANT:    "CellTenant",
				api.ENVAR_PLATFORM_CLUSTER_ID:     "ClusterId",
				api.ENVAR_PLATFORM_CLUSTER_STAGE:  "ClusterStage",
				api.ENVAR_PLATFORM_CLUSTER_GROUP:  "ClusterGroup",
				api.ENVAR_PLATFORM_CLUSTER_REGION: "ClusterRegion",
				api.ENVAR_PLATFORM_CLUSTER_TENANT: "ClusterTenant",
				api.ENVAR_PLATFORM_STACK:          "Stack",
				api.ENVAR_PLATFORM_STACK_DATATIER: "StackDatatier",
				api.ENVAR_PLATFORM_STACK_REGION:   "StackRegion",
				api.ENVAR_PLATFORM_STACK_TENANT:   "StackTenant",
				api.ENVAR_GIT_REPOSITORY:          "GitRepository",
				api.ENVAR_GIT_BRANCH_NAME:         "GitBranchName",
				api.ENVAR_GIT_BRANCH_HASH:         "GitBranchHash",
				api.ENVAR_GIT_BRANCH_TAG:          "GitBranchTag",
				api.ENVAR_GIT_BASE_BRANCH_NAME:    "GitBaseBranchName",
			},
			expected: &api.EnvironmentContextOptions{
				Project:           "project",
				CellProvider:      "cellprovider",
				CellId:            "cellid",
				CellStage:         "cellstage",
				CellRegion:        "cellregion",
				CellTenant:        "celltenant",
				ClusterId:         "clusterid",
				ClusterGroup:      "clustergroup",
				ClusterStage:      "clusterstage",
				ClusterRegion:     "clusterregion",
				ClusterTenant:     "clustertenant",
				Stack:             "stack",
				StackDatatier:     "stackdatatier",
				StackRegion:       "stackregion",
				StackTenant:       "stacktenant",
				GitRepository:     "gitrepository",
				GitBranchName:     "gitbranchname",
				GitBranchHash:     "gitbranchhash",
				GitBranchTag:      "gitbranchtag",
				GitBaseBranchName: "gitbasebranchname",
			},
			isGitBaseBranch: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			for envvar, value := range c.input {
				os.Setenv(envvar, value)
			}

			envvarOptions := OptionsFromEnvvars()

			assert.Equal(tt, c.expected.Project, envvarOptions.Project, "wrong Project")
			assert.Equal(tt, c.expected.CellProvider, envvarOptions.CellProvider, "wrong CellProvider")
			assert.Equal(tt, c.expected.CellId, envvarOptions.CellId, "wrong CellId")
			assert.Equal(tt, c.expected.CellStage, envvarOptions.CellStage, "wrong CellStage")
			assert.Equal(tt, c.expected.CellRegion, envvarOptions.CellRegion, "wrong CellRegion")
			assert.Equal(tt, c.expected.CellTenant, envvarOptions.CellTenant, "wrong CellTenant")
			assert.Equal(tt, c.expected.ClusterId, envvarOptions.ClusterId, "wrong ClusterId")
			assert.Equal(tt, c.expected.ClusterStage, envvarOptions.ClusterStage, "wrong ClusterStage")
			assert.Equal(tt, c.expected.ClusterGroup, envvarOptions.ClusterGroup, "wrong ClusterGroup")
			assert.Equal(tt, c.expected.ClusterRegion, envvarOptions.ClusterRegion, "wrong ClusterRegion")
			assert.Equal(tt, c.expected.ClusterTenant, envvarOptions.ClusterTenant, "wrong ClusterTenant")
			assert.Equal(tt, c.expected.Stack, envvarOptions.Stack, "wrong Stack")
			assert.Equal(tt, c.expected.StackDatatier, envvarOptions.StackDatatier, "wrong StackDatatier")
			assert.Equal(tt, c.expected.StackRegion, envvarOptions.StackRegion, "wrong StackRegion")
			assert.Equal(tt, c.expected.StackTenant, envvarOptions.StackTenant, "wrong StackTenant")
			assert.Equal(tt, c.expected.GitRepository, envvarOptions.GitRepository, "wrong GitRepository")
			assert.Equal(tt, c.expected.GitBranchName, envvarOptions.GitBranchName, "wrong GitBranchName")
			assert.Equal(tt, c.expected.GitBaseBranchName, envvarOptions.GitBaseBranchName, "wrong GitBranchTag")
			assert.Equal(tt, c.expected.GitBranchHash, envvarOptions.GitBranchHash, "wrong GitBranchHash")
			assert.Equal(tt, c.expected.GitBranchTag, envvarOptions.GitBranchTag, "wrong GitBranchTag")

			for envvar := range c.input {
				os.Unsetenv(envvar)
			}
		})
	}
}
