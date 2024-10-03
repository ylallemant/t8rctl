package ps

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/ylallemant/t8rctl/pkg/api"
	globalOptions "github.com/ylallemant/t8rctl/pkg/cli/cluster/options"
	"github.com/ylallemant/t8rctl/pkg/cli/cluster/ps/options"
	"github.com/ylallemant/t8rctl/pkg/cluster"
	"github.com/ylallemant/t8rctl/pkg/runtime"
)

var rootCmd = &cobra.Command{
	Use:   "ps",
	Short: "list running clusters",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		provider := runtime.Providers.Get(api.Azure)
		if provider == nil {
			return fmt.Errorf("provider \"%s\" not existing", api.Azure)
		}

		filter := filterFromFlags()

		clusters, err := provider.Clusters().List(provider.Accounts())
		if err != nil {
			return err
		}

		if !options.Current.All {
			// remove non managed clusters
			clusters = cluster.FilterManaged(clusters)
		}

		filteredClusters := cluster.Filter(clusters, filter)

		if len(filteredClusters) == 0 && options.Current.FallbackDatatier != "" {
			filter.Datatier = options.Current.FallbackDatatier
			filteredClusters = cluster.Filter(clusters, filter)
		}

		switch options.Current.Output {

		case "table":
			t := table.NewWriter()
			t.SetStyle(table.StyleLight)
			t.SetOutputMirror(os.Stdout)

			t.SetTitle("Cluster List")

			if options.Current.All {
				t.SetTitle("Cluster List (all)")
				t.AppendHeader(table.Row{"Name", "Is Managed", "Is Active", "Group", "Datatier", "ID", "Provider", "Region", "Account", "Section", "Kubernetes Version"})
			} else {
				t.SetTitle("Cluster List (managed)")
				t.AppendHeader(table.Row{"Name", "Is Active", "Group", "Datatier", "ID", "Provider", "Region", "Account", "Section", "Kubernetes Version"})
			}

			for _, cluster := range filteredClusters {
				if options.Current.All {
					t.AppendRow([]interface{}{
						cluster.Name(),
						cluster.Managed(),
						cluster.Active(),
						cluster.Group(),
						cluster.Datatier(),
						cluster.Id(),
						cluster.Provider(),
						cluster.Region(),
						cluster.Account().Name(),
						cluster.Section(),
						cluster.Version(),
					})
				} else {
					t.AppendRow([]interface{}{
						cluster.Name(),
						cluster.Active(),
						cluster.Group(),
						cluster.Datatier(),
						cluster.Id(),
						cluster.Provider(),
						cluster.Region(),
						cluster.Account().Name(),
						cluster.Section(),
						cluster.Version(),
					})
				}
			}

			t.Render()

		case "list":
			for _, cluster := range filteredClusters {
				fmt.Println(cluster.Name())
			}

		default:
			return errors.Errorf("unknown ouput option \"%s\"", options.Current.Output)
		}

		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&options.Current.All, "all", options.Current.All, "list also non managed clusters")
	rootCmd.PersistentFlags().StringVar(&options.Current.Output, "output", options.Current.Output, "set output format (\"table\" or \"list\")")
	rootCmd.PersistentFlags().StringVar(&options.Current.FallbackDatatier, "fallback-datatier", options.Current.FallbackDatatier, "datatier to use if the provided one returns no matches")
}

func Command() *cobra.Command {
	pflag.CommandLine.AddFlagSet(rootCmd.Flags())
	return rootCmd
}

func filterFromFlags() api.ClusterFilter {
	return api.ClusterFilter{
		Provider:     globalOptions.Current.Provider,
		Datatier:     globalOptions.Current.Datatier,
		Group:        globalOptions.Current.Group,
		Id:           globalOptions.Current.Id,
		All:          options.Current.All,
		ShowInactive: globalOptions.Current.ShowInactive,
	}
}
