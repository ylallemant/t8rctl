package list

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/ylallemant/t8rctl/pkg/cli/cluster/group/list/options"
	"github.com/ylallemant/t8rctl/pkg/cluster/group"
)

var rootCmd = &cobra.Command{
	Use:   "list",
	Short: "list cluster groups",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		groups, _ := group.List()

		switch options.Current.Output {

		case "table":
			t := table.NewWriter()
			t.SetStyle(table.StyleLight)
			t.SetOutputMirror(os.Stdout)

			t.SetTitle("Cluster Groups")

			t.AppendHeader(table.Row{"Name", "Description"})

			for _, group := range groups {
				t.AppendRow([]interface{}{group.Name(), group.Description()})
			}

			t.Render()

		case "list":
			for _, group := range groups {
				fmt.Println(group.Name())
			}

		default:
			return errors.Errorf("unknown ouput option \"%s\"", options.Current.Output)
		}

		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&options.Current.Output, "output", options.Current.Output, "set output format (\"table\" or \"list\")")
}

func Command() *cobra.Command {
	pflag.CommandLine.AddFlagSet(rootCmd.Flags())
	return rootCmd
}
