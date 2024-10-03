package cache

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/ylallemant/t8rctl/pkg/cli/cache/list"
	"github.com/ylallemant/t8rctl/pkg/cli/cache/purge"
)

var rootCmd = &cobra.Command{
	Use:   "cache",
	Short: "used interact with local caches",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("please use a subcommand...")
		cmd.Usage()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(list.Command())
	rootCmd.AddCommand(purge.Command())
}

func Command() *cobra.Command {
	pflag.CommandLine.AddFlagSet(rootCmd.Flags())
	return rootCmd
}
