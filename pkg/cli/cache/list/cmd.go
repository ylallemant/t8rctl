package list

import (
	"os"
	"time"

	humanize "github.com/dustin/go-humanize"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/ylallemant/t8rctl/pkg/cache"
)

var rootCmd = &cobra.Command{
	Use:   "list",
	Short: "list all caches or only specific ones",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		list := cache.CurrentManager.List()

		t := table.NewWriter()
		t.SetStyle(table.StyleLight)
		t.SetOutputMirror(os.Stdout)

		t.SetTitle("Cache List")
		t.AppendHeader(table.Row{"Path", "TTL", "Age", "Expires", "Size"})

		for _, cache := range list {
			t.AppendRow([]interface{}{
				cache.Path(),
				cache.TTL(),
				cache.Age().Round(time.Second),
				cache.Expires().Format("2006-01-02 15:04:05"),
				//time.Until(cache.Expires()),
				humanize.IBytes(uint64(cache.Size())),
			})
		}

		t.Render()

		return nil
	},
}

func init() {
	//rootCmd.PersistentFlags().BoolVar(&options.Current.Specific, "specific", options.Current.Specific, "uses only cluster id specific contexts (example \"workload-staging-green\")")
}

func Command() *cobra.Command {
	pflag.CommandLine.AddFlagSet(rootCmd.Flags())
	return rootCmd
}
