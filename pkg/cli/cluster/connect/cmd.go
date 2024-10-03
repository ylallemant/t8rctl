package connect

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/ylallemant/t8rctl/pkg/api"
	"github.com/ylallemant/t8rctl/pkg/cli/cluster/connect/options"
	globalOptions "github.com/ylallemant/t8rctl/pkg/cli/cluster/options"
	"github.com/ylallemant/t8rctl/pkg/cluster"
	"github.com/ylallemant/t8rctl/pkg/runtime"
)

var rootCmd = &cobra.Command{
	Use:   "connect",
	Short: "connects to current running managed clusters using meta contexts (without cluster id: \"workload-staging\", without \"green\")",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		provider := runtime.Providers.Get(api.Azure)
		if provider == nil {
			return fmt.Errorf("provider \"%s\" not existing", api.Azure)
		}

		err := provider.PurgeCaches()
		if err != nil {
			return errors.Wrapf(err, "provider \"%s\" could not purge its caches")
		}

		filter := filterFromFlags()

		clusters, err := provider.Clusters().List(provider.Accounts())
		if err != nil {
			return err
		}

		// context are only created for managed clusters
		managed := cluster.FilterManaged(clusters)

		filteredClusters := cluster.Filter(managed, filter)

		if len(filteredClusters) == 0 && options.Current.FallbackDatatier != "" {
			filter.Datatier = options.Current.FallbackDatatier
			filteredClusters = cluster.Filter(clusters, filter)
		}

		if !options.Current.Specific {
			fmt.Println("updating kubernetes meta-contexts for active clusters:")
			for _, cluster := range filteredClusters {
				provider := runtime.Providers.Get(cluster.Provider())
				context := fmt.Sprintf("%s-%s", cluster.Group(), cluster.Datatier())

				if cluster.Active() {
					err := provider.Clusters().Connect(cluster, context)
					if err != nil {
						return errors.Wrapf(err, "could not retrieve context for %s cluster %s", provider.Type(), context)
					}

					fmt.Println("  - updated context for ", context)
				}
			}
		}

		if options.Current.All || options.Current.Specific {
			fmt.Println("updating kubernetes specific contexts for all managed clusters")
			for _, cluster := range filteredClusters {
				provider := runtime.Providers.Get(cluster.Provider())
				context := fmt.Sprintf("%s-%s-%s", cluster.Group(), cluster.Datatier(), cluster.Id())

				err := provider.Clusters().Connect(cluster, context)
				if err != nil {
					return errors.Wrapf(err, "could not retrieve context for %s cluster %s", provider.Type(), context)
				}

				fmt.Println("  - updated context for ", context)
			}
		}

		if options.Current.Unmanaged {
			fmt.Println("updating kubernetes specific contexts for all non-managed clusters")
			for _, cluster := range clusters {
				if !cluster.Managed() {
					provider := runtime.Providers.Get(cluster.Provider())
					context := cluster.Name()

					err := provider.Clusters().Connect(cluster, context)
					if err != nil {
						return errors.Wrapf(err, "could not retrieve context for %s cluster %s", provider.Type(), context)
					}

					fmt.Println("  - updated context for ", context)
				}
			}
		}

		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&options.Current.All, "all", "a", options.Current.All, "adds also cluster id specific contexts (example \"workload-staging-green\")")
	rootCmd.PersistentFlags().BoolVar(&options.Current.Specific, "specific", options.Current.Specific, "uses only cluster id specific contexts (example \"workload-staging-green\")")
	rootCmd.PersistentFlags().BoolVarP(&options.Current.Unmanaged, "unmanaged", "u", options.Current.Unmanaged, "create contexts for unmanaged clusters")
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
