package namespace

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ylallemant/t8rctl/pkg/api"
	"github.com/ylallemant/t8rctl/pkg/environment"
)

func TestVariance(t *testing.T) {
	cases := []struct {
		name         string
		input        *api.EnvironmentContextOptions
		expected     string
		throwsError  bool
		errorMessage string
	}{
		{
			name:        "error missing git branch",
			input:       &api.EnvironmentContextOptions{},
			expected:    "",
			throwsError: true,
			errorMessage: fmt.Sprintf(
				"%s: %s",
				ErrorVarianceGeneration,
				ErrorMissingGitBranch,
			),
		},
		{
			name: "default vaiance",
			input: &api.EnvironmentContextOptions{
				GitBranchName: api.DefaultGitBaseBranchName,
			},
			expected: "all",
		},
		{
			name: "vaiance with custom datatier",
			input: &api.EnvironmentContextOptions{
				GitBranchName: api.DefaultGitBaseBranchName,
				StackDatatier: "dev",
			},
			expected: "dev",
		},
		{
			name: "vaiance for feature branch without prefix",
			input: &api.EnvironmentContextOptions{
				GitBranchName: "some improvement",
			},
			expected: "f-some-improvement",
		},
		{
			name: "vaiance for long feature branch without prefix",
			input: &api.EnvironmentContextOptions{
				GitBranchName: "some improvement for our best ecom system on the planet",
			},
			expected: "f-some-improvement-f-ec7e5",
		},
		{
			name: "vaiance for feature branch with prefix",
			input: &api.EnvironmentContextOptions{
				GitBranchName: "feature/ some improvement ",
			},
			expected: "f-some-improvement",
		},
		{
			name: "vaiance for branch with custom prefix",
			input: &api.EnvironmentContextOptions{
				GitBranchName: "issue / fix this problem",
			},
			expected: "i-fix-this-problem",
		},
		{
			name: "vaiance for branch with custom prefix using double points",
			input: &api.EnvironmentContextOptions{
				GitBranchName: "issue: do everthing better",
			},
			expected: "i-do-everthing-better",
		},
		{
			name: "vaiance remove forbbiden characters and consecutive separators",
			input: &api.EnvironmentContextOptions{
				GitBranchName: "bugfix/fix--this 💩 problem-",
			},
			expected: "b-fix-this-problem",
		},
		{
			name: "vaiance for feature branch with task reference",
			input: &api.EnvironmentContextOptions{
				GitBranchName: "feature/jira239_some improvement",
			},
			expected: "f-jira-239",
		},
		{
			name: "vaiance for long feature branch with task reference",
			input: &api.EnvironmentContextOptions{
				GitBranchName: "issue/jira-2453 = some improvement for our best ecom system on the planet",
			},
			expected: "i-jira-2453",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			env := environment.FromOptions(c.input, false)

			variance, err := Variance(env)

			assert.Equal(tt, c.expected, variance, "wrong variance")

			if c.throwsError {
				assert.NotNil(tt, err)
				assert.Equal(tt, c.errorMessage, err.Error(), "wrong error massage")
			} else {
				assert.Nil(tt, err)
			}

		})
	}
}

func TestName(t *testing.T) {
	cases := []struct {
		name         string
		input        *api.EnvironmentContextOptions
		expected     string
		throwsError  bool
		errorMessage string
	}{
		{
			name:        "error missing git branch",
			input:       &api.EnvironmentContextOptions{},
			expected:    "",
			throwsError: true,
			errorMessage: fmt.Sprintf(
				"%s: %s: %s",
				ErrorNameGeneration,
				ErrorVarianceGeneration,
				ErrorMissingGitBranch,
			),
		},
		{
			name: "error missing stack name",
			input: &api.EnvironmentContextOptions{
				GitBranchName: api.DefaultGitBaseBranchName,
			},
			expected:    "",
			throwsError: true,
			errorMessage: fmt.Sprintf(
				"%s: %s",
				ErrorNameGeneration,
				ErrorMissingStackName,
			),
		},
		{
			name: "default name with custom stack name",
			input: &api.EnvironmentContextOptions{
				GitBranchName: api.DefaultGitBaseBranchName,
				Stack:         "ecom",
			},
			expected: "ecom-all",
		},
		{
			name: "name with custom stack and datatier name",
			input: &api.EnvironmentContextOptions{
				GitBranchName: api.DefaultGitBaseBranchName,
				Stack:         "ecom",
				StackDatatier: "dev",
			},
			expected: "ecom-dev",
		},
		{
			name: "name with custom region",
			input: &api.EnvironmentContextOptions{
				GitBranchName: api.DefaultGitBaseBranchName,
				StackRegion:   "emea",
				Stack:         "ecom",
				StackDatatier: "dev",
			},
			expected: "ecom-emea-dev",
		},
		{
			name: "name with custom tenat",
			input: &api.EnvironmentContextOptions{
				GitBranchName: api.DefaultGitBaseBranchName,
				StackTenant:   "myshop",
				Stack:         "ecom",
				StackDatatier: "dev",
			},
			expected: "ecom-myshop-dev",
		},
		{
			name: "name with custom region and tenant",
			input: &api.EnvironmentContextOptions{
				GitBranchName: api.DefaultGitBaseBranchName,
				StackTenant:   "myshop",
				StackRegion:   "emea",
				Stack:         "ecom",
				StackDatatier: "dev",
			},
			expected: "ecom-emea-myshop-dev",
		},
		{
			name: "name from feature branch",
			input: &api.EnvironmentContextOptions{
				GitBranchName: "feature/JIRA-2435: user bonus",
				Stack:         "ecom",
				StackDatatier: "dev",
			},
			expected: "ecom-f-jira-2435",
		},
		{
			name: "name from long feature branch",
			input: &api.EnvironmentContextOptions{
				GitBranchName: "feature/our best shop improvement to date",
				Stack:         "ecom",
				StackDatatier: "dev",
			},
			expected: "ecom-f-our-best-shop-f50a1",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			env := environment.FromOptions(c.input, false)

			variance, err := Name(env)

			assert.Equal(tt, c.expected, variance, "wrong variance")

			if c.throwsError {
				assert.NotNil(tt, err)
				assert.Equal(tt, c.errorMessage, err.Error(), "wrong error massage")
			} else {
				assert.Nil(tt, err)
			}

		})
	}
}
