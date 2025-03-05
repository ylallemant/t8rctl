package namespace

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/ylallemant/t8rctl/pkg/api"
	"github.com/ylallemant/t8rctl/pkg/cli/resource/namespace/options"
	globalOptions "github.com/ylallemant/t8rctl/pkg/cli/resource/options"
	"github.com/ylallemant/t8rctl/pkg/environment"
	"github.com/ylallemant/t8rctl/pkg/resource/namespace"
	"github.com/ylallemant/t8rctl/pkg/runtime"
)

const (
	ErrorBadInput = "bad input"
)

var rootCmd = &cobra.Command{
	Use:   "namespace",
	Short: "generates a namespace name from environment information",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		opts := readFlags(cmd.Flags())
		provider := runtime.Providers.Get(globalOptions.Current.Provider)
		if provider == nil {
			return fmt.Errorf("provider \"%s\" not existing", provider.Type())
		}

		err := validate(opts)
		if err != nil {
			return errors.Wrapf(err, ErrorBadInput)
		}

		envOptions := environment.OptionsFromEnvvars()

		envOptions.GitBranchName = opts.CurrentGitBranch
		envOptions.GitBaseBranchName = opts.BaseGitBranch

		envOptions.Project = globalOptions.Current.Project
		envOptions.StackDatatier = globalOptions.Current.CellStage
		envOptions.StackRegion = opts.Region
		envOptions.StackTenant = opts.Tenant

		env := environment.FromOptions(envOptions, false)

		if opts.OnlyVariance {
			variance, err := namespace.Variance(env)
			if err != nil {
				return errors.Wrapf(err, "failed to generate namespace variance")
			}

			fmt.Fprint(cmd.OutOrStdout(), variance)
			return nil
		}

		name, err := namespace.Name(env)
		if err != nil {
			return errors.Wrapf(err, "failed to generate namespace name")
		}

		fmt.Fprint(cmd.OutOrStdout(), name)
		return nil
	},
}

func readFlags(flags *pflag.FlagSet) *options.Options {
	var err error
	opts := options.NewOptions()

	opts.BaseGitBranch, err = flags.GetString("branch-base")
	if err != nil {
		panic(err.Error())
	}
	opts.CurrentGitBranch, err = flags.GetString("branch-current")
	if err != nil {
		panic(err.Error())
	}
	opts.Region, err = flags.GetString("region")
	if err != nil {
		panic(err.Error())
	}
	opts.Tenant, err = flags.GetString("tenant")
	if err != nil {
		panic(err.Error())
	}
	opts.OnlyVariance, err = flags.GetBool("variance")
	if err != nil {
		panic(err.Error())
	}

	return opts
}

func init() {
	rootCmd.PersistentFlags().Bool("variance", false, "only returns the variable part of the name")
	rootCmd.PersistentFlags().String("branch-current", "", "current git branch")
	rootCmd.PersistentFlags().String("branch-base", api.DefaultGitBaseBranchName, "git branch used for base environments")
	rootCmd.PersistentFlags().String("tenant", api.DefaultTenant, "resource tenant name")
	rootCmd.PersistentFlags().String("region", api.DefaultRegion, "resource region name")
}

func Command() *cobra.Command {
	pflag.CommandLine.AddFlagSet(rootCmd.Flags())
	return rootCmd
}
