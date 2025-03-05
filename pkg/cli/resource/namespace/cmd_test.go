package namespace

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ylallemant/t8rctl/pkg/cli/resource/namespace/options"
	globalOptions "github.com/ylallemant/t8rctl/pkg/cli/resource/options"
	"github.com/ylallemant/t8rctl/pkg/testutils"
)

func setGlobals(project, cellStage, cellTenant, cellRegion string) {
	globalOptions.Current.Project = project
	globalOptions.Current.CellStage = cellStage

	if cellTenant != "" {
		globalOptions.Current.CellTenant = cellTenant
	}

	if cellRegion != "" {
		globalOptions.Current.CellRegion = cellRegion
	}
}

func TestCommand(t *testing.T) {
	cases := []struct {
		name         string
		args         []string
		project      string
		cellStage    string
		cellTenant   string
		cellRegion   string
		onlyVariance bool
		expected     string
		expectError  bool
		errorMessage string
	}{
		{
			name:         "input error project",
			args:         []string{},
			expectError:  true,
			errorMessage: fmt.Sprintf("%s: %s", ErrorBadInput, ErrorMissingCellProject),
		},
		{
			name:         "input error stage",
			args:         []string{},
			project:      "ecom",
			expectError:  true,
			errorMessage: fmt.Sprintf("%s: %s", ErrorBadInput, ErrorMissingCellStage),
		},
		{
			name:         "input error base git branch name",
			args:         []string{},
			project:      "ecom",
			cellStage:    "dev",
			expectError:  true,
			errorMessage: fmt.Sprintf("%s: %s", ErrorBadInput, ErrorMissingBaseGitBranch),
		},
		{
			name: "input error current git branch name",
			args: []string{
				"--branch-base", "main",
			},
			project:      "ecom",
			cellStage:    "dev",
			expectError:  true,
			errorMessage: fmt.Sprintf("%s: %s", ErrorBadInput, ErrorMissingCurrentGitBranch),
		},
		{
			name: "base environment",
			args: []string{
				"--branch-current", "main",
				"--branch-base", "main",
			},
			project:     "ecom",
			cellStage:   "dev",
			expected:    "ecom-dev",
			expectError: false,
		},
		{
			name: "feature environment",
			args: []string{
				"--branch-current", "some-branch",
				"--branch-base", "main",
			},
			project:     "ecom",
			cellStage:   "dev",
			expected:    "ecom-f-some-branch",
			expectError: false,
		},
		{
			name: "with tenant",
			args: []string{
				"--branch-current", "some-branch-with-very-long-name",
				"--branch-base", "main",
				"--tenant", "first",
			},
			project:     "ecom",
			cellStage:   "dev",
			expected:    "ecom-first-f-some-bra-2ba2e",
			expectError: false,
		},
		{
			name: "with region",
			args: []string{
				"--branch-current", "some-branch-with-very-long-name",
				"--branch-base", "main",
				"--region", "emea",
			},
			project:     "ecom",
			cellStage:   "dev",
			expected:    "ecom-emea-f-some-bran-2ba2e",
			expectError: false,
		},
		{
			name: "with region and tenant",
			args: []string{
				"--branch-current", "some-branch-with-very-long-name",
				"--branch-base", "main",
				"--tenant", "first",
				"--region", "emea",
			},
			project:     "ecom",
			cellStage:   "dev",
			expected:    "ecom-emea-first-f-som-2ba2e",
			expectError: false,
		},
		{
			name: "from shortened branch name",
			args: []string{
				"--branch-current", "some-branch-with-very-long-name",
				"--branch-base", "main",
			},
			project:     "ecom",
			cellStage:   "dev",
			expected:    "ecom-f-some-branch-wi-2ba2e",
			expectError: false,
		},
		{
			name: "ignore git path imformation in branch name",
			args: []string{
				"--branch-current", "refs/heads/feature/support-new-cluster-connection",
				"--branch-base", "main",
			},
			project:     "ecom",
			cellStage:   "dev",
			expected:    "ecom-f-support-new-cl-e0655",
			expectError: false,
		},
		{
			name: "branch name with hyphen JIRA reference",
			args: []string{
				"--branch-current", "refs/heads/feature/TEST-456-some-new-feature",
				"--branch-base", "main",
			},
			project:     "ecom",
			cellStage:   "dev",
			expected:    "ecom-f-test-456",
			expectError: false,
		},
		{
			name: "branch name with underscore JIRA reference",
			args: []string{
				"--branch-current", "refs/heads/feature/TEST_456-some-new-feature",
				"--branch-base", "main",
			},
			project:     "ecom",
			cellStage:   "dev",
			expected:    "ecom-f-test-456",
			expectError: false,
		},
		{
			name: "branch name with underscore JIRA reference",
			args: []string{
				"--branch-current", "refs/heads/feature/TEST456-some-new-feature",
				"--branch-base", "main",
			},
			project:     "ecom",
			cellStage:   "dev",
			expected:    "ecom-f-test-456",
			expectError: false,
		},
		{
			name: "variance for base environment",
			args: []string{
				"--branch-current", "main",
				"--branch-base", "main",
			},
			project:      "ecom",
			cellStage:    "dev",
			onlyVariance: true,
			expected:     "dev",
			expectError:  false,
		},
		{
			name: "variance for feature environment",
			args: []string{
				"--branch-current", "refs/heads/feature/support-new-cluster-connection",
				"--branch-base", "main",
			},
			project:      "ecom",
			cellStage:    "dev",
			onlyVariance: true,
			expected:     "f-support-new-cl-e0655",
			expectError:  false,
		},
	}

	for _, c := range cases {
		testutils.ResetCobraFlags(rootCmd)

		t.Run(c.name, func(tt *testing.T) {
			options.Current = options.NewOptions()
			globalOptions.Current = globalOptions.NewOptions()

			setGlobals(c.project, c.cellStage, c.cellTenant, c.cellRegion)

			buf := bytes.NewBufferString("")
			rootCmd.SetOut(buf)

			if c.onlyVariance {
				c.args = append(c.args, "--variance")
			}

			rootCmd.SetArgs(c.args)
			cmdErr := rootCmd.Execute()

			out, err := io.ReadAll(buf)
			assert.Nil(tt, err)

			if c.expectError {
				assert.NotNil(tt, cmdErr)
				if err != nil {
					assert.Equal(tt, c.errorMessage, cmdErr.Error(), "wrong error massage")
				}
			} else {
				assert.Nil(tt, cmdErr)
				assert.Equal(tt, c.expected, string(out), "wrong result")
			}
		})
	}
}
