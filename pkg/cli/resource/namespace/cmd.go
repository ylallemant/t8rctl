package namespace

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/ylallemant/t8rctl/pkg/api"
	"github.com/ylallemant/t8rctl/pkg/cli/resource/namespace/options"
	globalOptions "github.com/ylallemant/t8rctl/pkg/cli/resource/options"
	"github.com/ylallemant/t8rctl/pkg/resource/namespace"
	"github.com/ylallemant/t8rctl/pkg/runtime"
)

var rootCmd = &cobra.Command{
	Use:   "namespace",
	Short: "generates a namespace name from environment information",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		provider := runtime.Providers.Get(api.Azure)
		if provider == nil {
			return fmt.Errorf("provider \"%s\" not existing", api.Azure)
		}

		err := validate()
		if err != nil {
			return errors.Wrapf(err, "bad input")
		}

		env := &api.EnvironmentContext{
			BranchName: options.Current.CurrentGitBranch,
			IsBase:     options.Current.BaseGitBranch == options.Current.CurrentGitBranch,
			Project:    globalOptions.Current.Project,
			Datatier:   globalOptions.Current.CellStage,
			Region:     options.Current.Region,
			Tenant:     options.Current.Tenant,
		}

		fmt.Println(namespace.Name(env))

		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&options.Current.CurrentGitBranch, "branch-current", options.Current.CurrentGitBranch, "current git branch")
	rootCmd.PersistentFlags().StringVar(&options.Current.BaseGitBranch, "branch-base", options.Current.BaseGitBranch, "git branch used for base environments")
	rootCmd.PersistentFlags().StringVar(&options.Current.Tenant, "tenant", options.Current.Tenant, "resource tenant name")
	rootCmd.PersistentFlags().StringVar(&options.Current.Region, "region", options.Current.Region, "resource region name")
}

func Command() *cobra.Command {
	pflag.CommandLine.AddFlagSet(rootCmd.Flags())
	return rootCmd
}
