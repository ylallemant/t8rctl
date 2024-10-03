package name

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ylallemant/t8rctl/pkg/api"
	"github.com/ylallemant/t8rctl/pkg/cli/cluster/name/options"
	"github.com/ylallemant/t8rctl/pkg/cluster"
	"github.com/ylallemant/t8rctl/pkg/runtime"
)

var rootCmd = &cobra.Command{
	Use:   "name",
	Short: "returns the active cluster for the specified cluster group and datatier",
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

		if options.Current.Id != "" {
			// if an ID was specified it may be from an inactive cluster
			filter.ShowInactive = true
		}

		clusters, err := provider.Clusters().List(provider.Accounts())
		if err != nil {
			return err
		}

		// context are only created for managed clusters
		clusters = cluster.FilterManaged(clusters)

		filteredClusters := cluster.Filter(clusters, filter)

		if len(filteredClusters) == 0 && options.Current.FallbackDatatier != "" {
			filter.Datatier = options.Current.FallbackDatatier
			filteredClusters = cluster.Filter(clusters, filter)
		}

		if len(filteredClusters) == 0 {
			return fmt.Errorf("no cluster was found with filter: %s => %s-%s", api.Azure, options.Current.Group, options.Current.Datatier)
		}

		if len(filteredClusters) > 1 {
			return fmt.Errorf("filter returned multiple clusters (%d): %s => %s-%s", len(filteredClusters), api.Azure, options.Current.Group, options.Current.Datatier)
		}

		fmt.Println(filteredClusters[0].Name())

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

	if options.Current.Datatier == "" {
		return fmt.Errorf("datatier was not specified")
	}

	return nil
}

func init() {
	rootCmd.Flags().StringVar(&options.Current.Provider, "provider", options.Current.Provider, "filter by cloud provider")
	rootCmd.Flags().StringVar(&options.Current.Datatier, "datatier", options.Current.Datatier, "filter by datatier")
	rootCmd.Flags().StringVar(&options.Current.Group, "group", options.Current.Group, "filter by cluster group")
	rootCmd.Flags().StringVar(&options.Current.Id, "id", options.Current.Id, "filter by cluster id")
	rootCmd.Flags().StringVar(&options.Current.FallbackDatatier, "fallback-datatier", options.Current.FallbackDatatier, "datatier to use if the provided one returns no matches")
}

func Command() *cobra.Command {
	return rootCmd
}

func filterFromFlags() api.ClusterFilter {
	return api.ClusterFilter{
		Provider: options.Current.Provider,
		Datatier: options.Current.Datatier,
		Id:       options.Current.Id,
		Group:    options.Current.Group,
	}
}
