package name

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/ylallemant/t8rctl/pkg/cli/resource/name/options"
	globalOptions "github.com/ylallemant/t8rctl/pkg/cli/resource/options"
	"github.com/ylallemant/t8rctl/pkg/environment"
	resourceName "github.com/ylallemant/t8rctl/pkg/resource/name"
	"github.com/ylallemant/t8rctl/pkg/runtime"
)

var rootCmd = &cobra.Command{
	Use:   "name",
	Short: "generates a resource name from environment information",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		provider := runtime.Providers.Get(globalOptions.Current.Provider)
		if provider == nil {
			return fmt.Errorf("provider \"%s\" not existing", provider.Type())
		}

		err := resourceName.Validate()
		if err != nil {
			return errors.Wrapf(err, "bad input")
		}

		envOptions := environment.OptionsFromEnvvars()

		env := environment.FromOptions(envOptions, false)

		output := resourceName.CoreName(env)

		if options.Current.Short {
			output = resourceName.ShortName(env)
		}

		if options.Current.Hash {
			output = resourceName.CoreHash(env)
		}

		if options.Current.Caf {
			output = resourceName.CafName(env)
		}

		fmt.Println(output)

		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&options.Current.Name, "name", options.Current.Name, "base resource name")
	rootCmd.PersistentFlags().StringVar(&options.Current.Static, "static", options.Current.Static, "bypass core name generation with a static value")
	rootCmd.PersistentFlags().StringVar(&options.Current.Type, "type", options.Current.Type, "base resource type")
	rootCmd.PersistentFlags().StringVar(&options.Current.Tenant, "tenant", options.Current.Tenant, "resource tenant name")
	rootCmd.PersistentFlags().StringVar(&options.Current.Region, "region", options.Current.Region, "resource region name")
	rootCmd.PersistentFlags().BoolVar(&options.Current.Short, "short", options.Current.Short, "output name limited to 24 chars")
	rootCmd.PersistentFlags().BoolVar(&options.Current.Core, "core", options.Current.Core, "output core name")
	rootCmd.PersistentFlags().BoolVar(&options.Current.Hash, "hash", options.Current.Hash, "output ony the core name 5 char hash")
	rootCmd.PersistentFlags().BoolVar(&options.Current.Caf, "caf", options.Current.Caf, "output core name with resource short prefix and cell-id suffix")
}

func Command() *cobra.Command {
	pflag.CommandLine.AddFlagSet(rootCmd.Flags())
	return rootCmd
}
