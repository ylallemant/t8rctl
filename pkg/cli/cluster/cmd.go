package cluster

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/ylallemant/t8rctl/pkg/cli/cluster/active"
	"github.com/ylallemant/t8rctl/pkg/cli/cluster/connect"
	"github.com/ylallemant/t8rctl/pkg/cli/cluster/datatier"
	"github.com/ylallemant/t8rctl/pkg/cli/cluster/group"
	"github.com/ylallemant/t8rctl/pkg/cli/cluster/id"
	"github.com/ylallemant/t8rctl/pkg/cli/cluster/name"
	"github.com/ylallemant/t8rctl/pkg/cli/cluster/options"
	"github.com/ylallemant/t8rctl/pkg/cli/cluster/ps"
)

var rootCmd = &cobra.Command{
	Use:   "cluster",
	Short: "used interact with accessible clusters",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("please use a subcommand...")
		cmd.Usage()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(group.Command())
	rootCmd.AddCommand(ps.Command())
	rootCmd.AddCommand(connect.Command())
	rootCmd.AddCommand(active.Command())
	rootCmd.AddCommand(name.Command())
	rootCmd.AddCommand(id.Command())
	rootCmd.AddCommand(datatier.Command())
}

func Command() *cobra.Command {
	rootCmd.PersistentFlags().StringVar(&options.Current.Provider, "provider", options.Current.Provider, "filter by cloud provider")
	rootCmd.PersistentFlags().StringVar(&options.Current.Datatier, "datatier", options.Current.Datatier, "filter by datatier")
	rootCmd.PersistentFlags().StringVar(&options.Current.Group, "group", options.Current.Group, "filter by cluster group")
	rootCmd.PersistentFlags().StringVar(&options.Current.Id, "id", options.Current.Id, "filter by cluster id")
	rootCmd.PersistentFlags().BoolVar(&options.Current.ShowInactive, "show-inactive", options.Current.ShowInactive, "show also inactive clusters")
	rootCmd.PersistentFlags().BoolVarP(&options.Current.Debug, "debug", "d", options.Current.Debug, "output debugging logs")
	rootCmd.PersistentFlags().BoolVarP(&options.Current.DisableCache, "disable-cache", "c", options.Current.DisableCache, "disables caching")

	pflag.CommandLine.AddFlagSet(rootCmd.Flags())
	return rootCmd
}
