package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ylallemant/t8rctl/pkg/cli/cache"
	"github.com/ylallemant/t8rctl/pkg/cli/cluster"
	"github.com/ylallemant/t8rctl/pkg/cli/update"
	"github.com/ylallemant/t8rctl/pkg/cli/version"
)

var rootCmd = &cobra.Command{
	Use:   "t8rctl",
	Short: "t8rctl is a CLI tool to interact with the Transistor (t8r) Internal Developer Platform",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("please use a subcommand...")
		cmd.Usage()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(cluster.Command())
	rootCmd.AddCommand(cache.Command())
	rootCmd.AddCommand(update.Command())
	rootCmd.AddCommand(version.Command())
}

func Command() *cobra.Command {
	return rootCmd
}
