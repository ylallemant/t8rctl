package name

import (
	"crypto/sha1"
	"fmt"
	"io"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/ylallemant/t8rctl/pkg/api"
	"github.com/ylallemant/t8rctl/pkg/cli/resource/name/options"
	globalOptions "github.com/ylallemant/t8rctl/pkg/cli/resource/options"
	"github.com/ylallemant/t8rctl/pkg/runtime"
)

var rootCmd = &cobra.Command{
	Use:   "name",
	Short: "generates a resource name from environment information",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		provider := runtime.Providers.Get(api.Azure)
		if provider == nil {
			return fmt.Errorf("provider \"%s\" not existing", api.Azure)
		}

		err := validate()
		if err != nil {
			return errors.Wrapf(err, "bad input")
		}

		output := coreName()

		if options.Current.Short {
			output = shortName()
		}

		if options.Current.Hash {
			output = coreHash()
		}

		if options.Current.Caf {
			output = cafName()
		}

		fmt.Println(output)

		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&options.Current.Name, "name", options.Current.Name, "base resource name")
	rootCmd.PersistentFlags().StringVar(&options.Current.Static, "static", options.Current.Static, "bypass core name generation with a static value")
	rootCmd.PersistentFlags().StringVar(&options.Current.Type, "type", options.Current.Type, "base resource type")
	rootCmd.PersistentFlags().StringVar(&options.Current.Tenant, "tenant", options.Current.Tenant, "resource tenant name")
	rootCmd.PersistentFlags().StringVar(&options.Current.Region, "region", options.Current.Region, "resource region name")
	rootCmd.PersistentFlags().BoolVar(&options.Current.Short, "short", options.Current.Short, "output name limited to 24 chars")
	rootCmd.PersistentFlags().BoolVar(&options.Current.Core, "core", options.Current.Core, "output core name")
	rootCmd.PersistentFlags().BoolVar(&options.Current.Hash, "hash", options.Current.Hash, "output ony the core name 5 char hash")
	rootCmd.PersistentFlags().BoolVar(&options.Current.Caf, "caf", options.Current.Caf, "output core name with resource short prefix and cell-id suffix")
}

func Command() *cobra.Command {
	pflag.CommandLine.AddFlagSet(rootCmd.Flags())
	return rootCmd
}

func coreName() string {
	if options.Current.Static != "" {
		return strings.ToLower(options.Current.Static)
	}

	tenant := ""

	if globalOptions.Current.CellTenant != "" && globalOptions.Current.CellTenant != "shared" {
		tenant = fmt.Sprintf("-%s", globalOptions.Current.CellTenant)
	}

	if tenant == "" && options.Current.Tenant != "" && options.Current.Tenant != "none" {
		tenant = fmt.Sprintf("-%s", options.Current.Tenant)
	}

	region := ""

	if globalOptions.Current.CellRegion != "" && globalOptions.Current.CellRegion != "global" {
		region = fmt.Sprintf("-%s", globalOptions.Current.CellRegion)
	}

	if region == "" && options.Current.Region != "" && options.Current.Region != "none" {
		region = fmt.Sprintf("-%s", options.Current.Region)
	}

	core := strings.ToLower(
		fmt.Sprintf(
			"%s%s-%s%s-%s",
			globalOptions.Current.Project,
			region,
			globalOptions.Current.CellStage,
			tenant,
			options.Current.Name,
		),
	)

	return core
}

func coreHash() string {
	hasher := sha1.New()
	io.WriteString(hasher, coreName())
	hash := fmt.Sprintf("%x", hasher.Sum(nil))

	return hash[:5]
}

func shortName() string {
	core := coreName()
	info := typeInfo(options.Current.Type)
	hash := coreHash()

	maxCoreLengh := 24 - len(info.Prefix) - len(globalOptions.Current.CellId) - len(hash)

	sanitisedCore := sanitiseCore(core)

	if len(sanitisedCore) < maxCoreLengh {
		return fmt.Sprintf(
			"%s%s%s",
			info.Prefix,
			sanitisedCore,
			globalOptions.Current.CellId,
		)
	}

	return fmt.Sprintf(
		"%s%s%s%s",
		info.Prefix,
		sanitisedCore[:maxCoreLengh],
		hash,
		globalOptions.Current.CellId,
	)
}

func cafName() string {
	core := coreName()
	info := typeInfo(options.Current.Type)

	if info.ShortFormat {
		return shortName()
	}

	return fmt.Sprintf(
		"%s-%s-%s",
		info.Prefix,
		core,
		globalOptions.Current.CellId,
	)
}

func typeInfo(typeName string) api.ResourceInfo {
	return api.ResouceTypePrefix[typeName]
}

func sanitiseCore(core string) string {
	return api.ResourceShortNameNonAllowedChars.ReplaceAllString(core, "")
}
