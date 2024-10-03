package purge

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/ylallemant/t8rctl/pkg/cache"
)

var rootCmd = &cobra.Command{
	Use:   "purge",
	Short: "purges all caches or only specific ones",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cache.CurrentManager.Purge("")
	},
}

func init() {
	//rootCmd.PersistentFlags().BoolVar(&options.Current.Specific, "specific", options.Current.Specific, "uses only cluster id specific contexts (example \"workload-staging-green\")")
}

func Command() *cobra.Command {
	pflag.CommandLine.AddFlagSet(rootCmd.Flags())
	return rootCmd
}
