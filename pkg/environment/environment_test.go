package environment

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ylallemant/t8rctl/pkg/api"
)

func TestFromOptions(t *testing.T) {
	cases := []struct {
		name            string
		input           *api.EnvironmentContextOptions
		expected        *api.EnvironmentContextOptions
		isGitBaseBranch bool
	}{
		{
			name:  "default values",
			input: &api.EnvironmentContextOptions{},
			expected: &api.EnvironmentContextOptions{
				CellStage:         api.DefaultEnvironmentStage,
				CellRegion:        api.DefaultEnvironmentRegion,
				CellTenant:        api.DefaultEnvironmentTenant,
				ClusterStage:      api.DefaultEnvironmentStage,
				ClusterRegion:     api.DefaultEnvironmentRegion,
				ClusterTenant:     api.DefaultEnvironmentTenant,
				StackDatatier:     api.DefaultEnvironmentStage,
				StackRegion:       api.DefaultEnvironmentRegion,
				StackTenant:       api.DefaultEnvironmentTenant,
				GitBaseBranchName: api.DefaultGitBaseBranchName,
			},
			isGitBaseBranch: false,
		},
		{
			name: "stage, region, tenant value propagation",
			input: &api.EnvironmentContextOptions{
				CellStage:  "DEV",
				CellRegion: " emea ",
				CellTenant: "sHop\n",
			},
			expected: &api.EnvironmentContextOptions{
				CellStage:         "dev",
				CellRegion:        "emea",
				CellTenant:        "shop",
				ClusterStage:      "dev",
				ClusterRegion:     "emea",
				ClusterTenant:     "shop",
				StackDatatier:     "dev",
				StackRegion:       "emea",
				StackTenant:       "shop",
				GitBaseBranchName: api.DefaultGitBaseBranchName,
			},
			isGitBaseBranch: false,
		},
		{
			name: "cluster overwrites cell propagation",
			input: &api.EnvironmentContextOptions{
				StackDatatier: "DEV",
				StackRegion:   " emea ",
				StackTenant:   "sHop\n",
			},
			expected: &api.EnvironmentContextOptions{
				CellStage:         api.DefaultEnvironmentStage,
				CellRegion:        api.DefaultEnvironmentRegion,
				CellTenant:        api.DefaultEnvironmentTenant,
				ClusterStage:      api.DefaultEnvironmentStage,
				ClusterRegion:     api.DefaultEnvironmentRegion,
				ClusterTenant:     api.DefaultEnvironmentTenant,
				StackDatatier:     "dev",
				StackRegion:       "emea",
				StackTenant:       "shop",
				GitBaseBranchName: api.DefaultGitBaseBranchName,
			},
			isGitBaseBranch: false,
		},
		{
			name: "stack overwrites cell and cluster propagation",
			input: &api.EnvironmentContextOptions{
				CellStage:     "DEV",
				ClusterStage:  "TEST",
				CellRegion:    " emea ",
				ClusterRegion: "amer",
				CellTenant:    "sHop\n",
				ClusterTenant: "ecom",
			},
			expected: &api.EnvironmentContextOptions{
				CellStage:         "dev",
				CellRegion:        "emea",
				CellTenant:        "shop",
				ClusterStage:      "test",
				ClusterRegion:     "amer",
				ClusterTenant:     "ecom",
				StackDatatier:     "test",
				StackRegion:       "amer",
				StackTenant:       "ecom",
				GitBaseBranchName: api.DefaultGitBaseBranchName,
			},
			isGitBaseBranch: false,
		},
		{
			name: "git base git branch = true",
			input: &api.EnvironmentContextOptions{
				GitBranchName: api.DefaultGitBaseBranchName,
			},
			expected: &api.EnvironmentContextOptions{
				CellStage:         api.DefaultEnvironmentStage,
				CellRegion:        api.DefaultEnvironmentRegion,
				CellTenant:        api.DefaultEnvironmentTenant,
				ClusterStage:      api.DefaultEnvironmentStage,
				ClusterRegion:     api.DefaultEnvironmentRegion,
				ClusterTenant:     api.DefaultEnvironmentTenant,
				StackDatatier:     api.DefaultEnvironmentStage,
				StackRegion:       api.DefaultEnvironmentRegion,
				StackTenant:       api.DefaultEnvironmentTenant,
				GitBranchName:     api.DefaultGitBaseBranchName,
				GitBaseBranchName: api.DefaultGitBaseBranchName,
			},
			isGitBaseBranch: true,
		},
		{
			name: "check all properties",
			input: &api.EnvironmentContextOptions{
				Project:           "project",
				CellProvider:      "cellprovider",
				CellId:            "cellid",
				ClusterId:         "ClusterId",
				ClusterGroup:      "ClusterGroup",
				Stack:             "Stack",
				GitRepository:     "GitRepository",
				GitBranchName:     "refs/heads/GitBranchName",
				GitBranchHash:     "GitBranchHash",
				GitBranchTag:      "GitBranchTag",
				GitBaseBranchName: "GitBaseBranchName",
			},
			expected: &api.EnvironmentContextOptions{
				Project:           "project",
				CellProvider:      "cellprovider",
				CellId:            "cellid",
				ClusterId:         "clusterid",
				ClusterGroup:      "clustergroup",
				Stack:             "stack",
				GitRepository:     "gitrepository",
				GitBranchName:     "gitbranchname",
				GitBranchHash:     "gitbranchhash",
				GitBranchTag:      "gitbranchtag",
				GitBaseBranchName: "gitbasebranchname",
				CellStage:         api.DefaultEnvironmentStage,
				CellRegion:        api.DefaultEnvironmentRegion,
				CellTenant:        api.DefaultEnvironmentTenant,
				ClusterStage:      api.DefaultEnvironmentStage,
				ClusterRegion:     api.DefaultEnvironmentRegion,
				ClusterTenant:     api.DefaultEnvironmentTenant,
				StackDatatier:     api.DefaultEnvironmentStage,
				StackRegion:       api.DefaultEnvironmentRegion,
				StackTenant:       api.DefaultEnvironmentTenant,
			},
			isGitBaseBranch: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			env := FromOptions(c.input, false)

			assert.Equal(tt, c.expected.Project, env.Project(), "wrong Project")
			assert.Equal(tt, c.expected.CellProvider, env.CellProvider(), "wrong CellProvider")
			assert.Equal(tt, c.expected.CellId, env.CellId(), "wrong CellId")
			assert.Equal(tt, c.expected.CellStage, env.CellStage(), "wrong CellStage")
			assert.Equal(tt, c.expected.CellRegion, env.CellRegion(), "wrong CellRegion")
			assert.Equal(tt, c.expected.CellTenant, env.CellTenant(), "wrong CellTenant")
			assert.Equal(tt, c.expected.ClusterId, env.ClusterId(), "wrong ClusterId")
			assert.Equal(tt, c.expected.ClusterStage, env.ClusterStage(), "wrong ClusterStage")
			assert.Equal(tt, c.expected.ClusterGroup, env.ClusterGroup(), "wrong ClusterGroup")
			assert.Equal(tt, c.expected.ClusterRegion, env.ClusterRegion(), "wrong ClusterRegion")
			assert.Equal(tt, c.expected.ClusterTenant, env.ClusterTenant(), "wrong ClusterTenant")
			assert.Equal(tt, c.expected.Stack, env.Stack(), "wrong Stack")
			assert.Equal(tt, c.expected.StackDatatier, env.StackDatatier(), "wrong StackDatatier")
			assert.Equal(tt, c.expected.StackRegion, env.StackRegion(), "wrong StackRegion")
			assert.Equal(tt, c.expected.StackTenant, env.StackTenant(), "wrong StackTenant")
			assert.Equal(tt, c.expected.GitRepository, env.GitRepository(), "wrong GitRepository")
			assert.Equal(tt, c.expected.GitBranchName, env.GitBranchName(), "wrong GitBranchName")
			assert.Equal(tt, c.expected.GitBaseBranchName, env.GitBaseBranchName(), "wrong GitBranchTag")
			assert.Equal(tt, c.expected.GitBranchHash, env.GitBranchHash(), "wrong GitBranchHash")
			assert.Equal(tt, c.expected.GitBranchTag, env.GitBranchTag(), "wrong GitBranchTag")
			assert.Equal(tt, c.isGitBaseBranch, env.IsGitBaseBranch(), "wrong IsGitBaseBranch")
		})
	}
}
