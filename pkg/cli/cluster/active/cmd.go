package active

import (
	"fmt"

	"github.com/ylallemant/t8rctl/pkg/api"
	"github.com/ylallemant/t8rctl/pkg/cli/cluster/active/options"
	"github.com/ylallemant/t8rctl/pkg/cluster"
	"github.com/ylallemant/t8rctl/pkg/global"
	"github.com/ylallemant/t8rctl/pkg/logger"
	"github.com/ylallemant/t8rctl/pkg/runtime"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "active",
	Short: "checks whether or not the specified cluster is active and receives network traffic (returns \"true\" or \"false\")",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		logger.Ensure(options.Current.Debug)

		global.Current.Debug = options.Current.Debug
		global.Current.DisableCache = options.Current.DisableCache
		log.Debug().Msgf("used global options: %#+v", global.Current)

		err := validate()
		if err != nil {
			return err
		}

		provider := runtime.Providers.Get(api.Azure)
		if provider == nil {
			return fmt.Errorf("provider \"%s\" not existing", api.Azure)
		}

		filter := filterFromFlags()
		filter.ShowInactive = true
		log.Debug().Msgf("used filters: %#+v", filter)

		clusters, err := provider.Clusters().List(provider.Accounts())
		if err != nil {
			return err
		}
		log.Debug().Msgf("found clusters: %#+v", clusters)

		// context are only created for managed clusters
		clusters = cluster.FilterManaged(clusters)

		filteredClusters := cluster.Filter(clusters, filter)

		if len(filteredClusters) == 0 {
			if options.Current.IgnoreUnknown {
				fmt.Println("false")
				return nil
			} else {
				return fmt.Errorf("specified cluster was not found: %s => %s-%s-%s", api.Azure, options.Current.Group, options.Current.Datatier, options.Current.Id)
			}
		}

		if len(filteredClusters) > 1 {
			return fmt.Errorf("filter returned multiple clusters (%d): %s => %s-%s-%s", len(filteredClusters), api.Azure, options.Current.Group, options.Current.Datatier, options.Current.Id)
		}

		fmt.Println(filteredClusters[0].Active())

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

	if options.Current.Id == "" {
		return fmt.Errorf("cluster id was not specified")
	}

	return nil
}

func init() {
	rootCmd.PersistentFlags().StringVar(&options.Current.Provider, "provider", options.Current.Provider, "filter by cloud provider")
	rootCmd.PersistentFlags().StringVar(&options.Current.Datatier, "datatier", options.Current.Datatier, "filter by datatier")
	rootCmd.PersistentFlags().StringVar(&options.Current.Group, "group", options.Current.Group, "filter by cluster group")
	rootCmd.PersistentFlags().StringVar(&options.Current.Id, "id", options.Current.Id, "filter by cluster id")
	rootCmd.PersistentFlags().BoolVar(&options.Current.IgnoreUnknown, "ignore-unknown", options.Current.IgnoreUnknown, "return \"false\" if cluster is unknown")
	rootCmd.PersistentFlags().BoolVarP(&options.Current.Debug, "debug", "d", options.Current.Debug, "output debugging logs")
	rootCmd.PersistentFlags().BoolVarP(&options.Current.DisableCache, "disable-cache", "c", options.Current.DisableCache, "disables caching")
}

func Command() *cobra.Command {
	pflag.CommandLine.AddFlagSet(rootCmd.Flags())
	return rootCmd
}

func filterFromFlags() api.ClusterFilter {
	return api.ClusterFilter{
		Provider: options.Current.Provider,
		Datatier: options.Current.Datatier,
		Group:    options.Current.Group,
		Id:       options.Current.Id,
	}
}
