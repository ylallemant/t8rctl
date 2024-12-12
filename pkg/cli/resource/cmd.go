package resource

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/ylallemant/t8rctl/pkg/cli/resource/name"
	"github.com/ylallemant/t8rctl/pkg/cli/resource/options"
)

var rootCmd = &cobra.Command{
	Use:   "resource",
	Short: "used interact with resources",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("please use a subcommand...")
		cmd.Usage()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(name.Command())
}

func Command() *cobra.Command {
	rootCmd.PersistentFlags().StringVar(&options.Current.Provider, "provider", options.Current.Provider, "target provider")
	rootCmd.PersistentFlags().StringVar(&options.Current.Project, "project", options.Current.Project, "project short name")
	rootCmd.PersistentFlags().StringVar(&options.Current.CellStage, "cell-stage", options.Current.CellStage, "cell stage name")
	rootCmd.PersistentFlags().StringVar(&options.Current.CellTenant, "cell-tenant", options.Current.CellTenant, "cell tenant name")
	rootCmd.PersistentFlags().StringVar(&options.Current.CellRegion, "cell-region", options.Current.CellRegion, "cell region name")
	rootCmd.PersistentFlags().StringVar(&options.Current.CellId, "cell-id", options.Current.CellId, "cell id")

	pflag.CommandLine.AddFlagSet(rootCmd.Flags())
	return rootCmd
}
