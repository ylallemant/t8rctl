package update

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/ylallemant/t8rctl/pkg/cli/update/options"
)

var rootCmd = &cobra.Command{
	Use:   "update",
	Short: "update the binary",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("not implemented")

		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&options.Current.DryRun, "dry-run", options.Current.DryRun, "does not replace the binary")
	rootCmd.PersistentFlags().BoolVar(&options.Current.Force, "force", options.Current.Force, "force the replacement of the binary")
}

func Command() *cobra.Command {
	pflag.CommandLine.AddFlagSet(rootCmd.Flags())
	return rootCmd
}
