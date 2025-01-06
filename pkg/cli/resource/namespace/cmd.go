package namespace

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/ylallemant/t8rctl/pkg/cli/resource/namespace/options"
	globalOptions "github.com/ylallemant/t8rctl/pkg/cli/resource/options"
	"github.com/ylallemant/t8rctl/pkg/environment"
	"github.com/ylallemant/t8rctl/pkg/resource/namespace"
	"github.com/ylallemant/t8rctl/pkg/runtime"
)

var rootCmd = &cobra.Command{
	Use:   "namespace",
	Short: "generates a namespace name from environment information",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		provider := runtime.Providers.Get(globalOptions.Current.Provider)
		if provider == nil {
			return fmt.Errorf("provider \"%s\" not existing", provider.Type())
		}

		err := validate()
		if err != nil {
			return errors.Wrapf(err, "bad input")
		}

		envOptions := environment.OptionsFromEnvvars()

		envOptions.GitBranchName = options.Current.CurrentGitBranch
		envOptions.GitBaseBranchName = options.Current.BaseGitBranch

		envOptions.Project = globalOptions.Current.Project
		envOptions.CellStage = globalOptions.Current.CellStage
		envOptions.CellRegion = options.Current.Region
		envOptions.CellTenant = options.Current.Tenant

		env := environment.FromOptions(envOptions, false)

		if options.Current.OnlyVariance {
			fmt.Println(namespace.Variance(env))
			return nil
		}

		fmt.Println(namespace.Name(env))
		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&options.Current.OnlyVariance, "variance", options.Current.OnlyVariance, "only returns the variable part of the name")
	rootCmd.PersistentFlags().StringVar(&options.Current.CurrentGitBranch, "branch-current", options.Current.CurrentGitBranch, "current git branch")
	rootCmd.PersistentFlags().StringVar(&options.Current.BaseGitBranch, "branch-base", options.Current.BaseGitBranch, "git branch used for base environments")
	rootCmd.PersistentFlags().StringVar(&options.Current.Tenant, "tenant", options.Current.Tenant, "resource tenant name")
	rootCmd.PersistentFlags().StringVar(&options.Current.Region, "region", options.Current.Region, "resource region name")
}

func Command() *cobra.Command {
	pflag.CommandLine.AddFlagSet(rootCmd.Flags())
	return rootCmd
}
