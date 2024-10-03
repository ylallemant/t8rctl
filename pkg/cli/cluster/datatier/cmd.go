package datatier

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ylallemant/t8rctl/pkg/api"
	"github.com/ylallemant/t8rctl/pkg/cli/cluster/datatier/options"
	"github.com/ylallemant/t8rctl/pkg/cluster"
	"github.com/ylallemant/t8rctl/pkg/runtime"
)

var rootCmd = &cobra.Command{
	Use:   "datatier",
	Short: "returns the corresponding cluster datatier for any provided stack datatier",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := validate()
		if err != nil {
			return err
		}

		provider := runtime.Providers.Get(api.Azure)
		if provider == nil {
			return fmt.Errorf("provider \"%s\" not existing", api.Azure)
		}

		filter := filterFromFlags()

		clusters, err := provider.Clusters().List(provider.Accounts())
		if err != nil {
			return err
		}

		// context are only created for managed clusters
		clusters = cluster.FilterManaged(clusters)

		filteredClusters := cluster.Filter(clusters, filter)

		if len(filteredClusters) == 0 && options.Current.DefaultClusterDatatier != "" {
			filter.Datatier = options.Current.DefaultClusterDatatier
			filteredClusters = cluster.Filter(clusters, filter)
		}

		if len(filteredClusters) == 0 {
			return fmt.Errorf("no cluster was found with filter: %s => %s-%s", api.Azure, options.Current.Group, options.Current.StackDatatier)
		}

		if len(filteredClusters) > 1 {
			return fmt.Errorf("filter returned multiple clusters (%d): %s => %s-%s", len(filteredClusters), api.Azure, options.Current.Group, options.Current.StackDatatier)
		}

		fmt.Println(filteredClusters[0].Datatier())

		return nil
	},
}

func validate() error {
	if api.Azure == "" {
		return fmt.Errorf("cloud provider was not specified")
	}

	if options.Current.Group == "" {
		return fmt.Errorf("cluster group was not specified")
	}

	if options.Current.StackDatatier == "" {
		return fmt.Errorf("datatier was not specified")
	}

	return nil
}

func init() {
	rootCmd.Flags().StringVar(&options.Current.Provider, "provider", options.Current.Provider, "filter by cloud provider")
	rootCmd.Flags().StringVarP(&options.Current.StackDatatier, "stack-datatier", "s", options.Current.StackDatatier, "stack datatier")
	rootCmd.Flags().StringVar(&options.Current.Group, "group", options.Current.Group, "filter by cluster group")
	rootCmd.Flags().StringVar(&options.Current.DefaultClusterDatatier, "default-cluster-datatier", options.Current.DefaultClusterDatatier, "cluster datatier to use if the provided one returns no matches")
}

func Command() *cobra.Command {
	return rootCmd
}

func filterFromFlags() api.ClusterFilter {
	return api.ClusterFilter{
		Provider: options.Current.Provider,
		Datatier: options.Current.StackDatatier,
		Group:    options.Current.Group,
	}
}
