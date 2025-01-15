package name

import (
	"crypto/sha1"
	"fmt"
	"io"
	"strings"

	"github.com/ylallemant/t8rctl/pkg/api"
	"github.com/ylallemant/t8rctl/pkg/cli/resource/name/options"
	globalOptions "github.com/ylallemant/t8rctl/pkg/cli/resource/options"
)

func Name(env api.EnvironmentContext) (string, error) {
	return CafName(env), nil
}

func CoreName(env api.EnvironmentContext) string {
	if options.Current.Static != "" {
		return strings.ToLower(options.Current.Static)
	}

	tenant := ""

	if options.Current.Tenant != "" && options.Current.Tenant != api.None && options.Current.Tenant != api.DefaultTenant {
		tenant = fmt.Sprintf("-%s", options.Current.Tenant)
	}

	if tenant == "" && globalOptions.Current.CellTenant != "" && globalOptions.Current.CellTenant != api.None && globalOptions.Current.CellTenant != api.DefaultTenant {
		tenant = fmt.Sprintf("-%s", globalOptions.Current.CellTenant)
	}

	region := ""

	if options.Current.Region != "" && options.Current.Region != api.None && options.Current.Region != api.DefaultRegion {
		region = fmt.Sprintf("-%s", options.Current.Region)
	}

	if region == "" && globalOptions.Current.CellRegion != "" && globalOptions.Current.CellRegion != api.None && globalOptions.Current.CellRegion != api.DefaultRegion {
		region = fmt.Sprintf("-%s", globalOptions.Current.CellRegion)
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

func CoreHash(env api.EnvironmentContext) string {
	hasher := sha1.New()
	io.WriteString(hasher, CoreName(env))
	hash := fmt.Sprintf("%x", hasher.Sum(nil))

	return hash[:5]
}

func ShortName(env api.EnvironmentContext) string {
	core := CoreName(env)
	info := typeInfo(options.Current.Type)
	hash := CoreHash(env)

	maxCoreLengh := 24 - len(info.Prefix) - len(globalOptions.Current.CellId)
	maxCoreLenghShortened := maxCoreLengh - len(hash)

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
		sanitisedCore[:maxCoreLenghShortened],
		hash,
		globalOptions.Current.CellId,
	)
}

func CafName(env api.EnvironmentContext) string {
	core := CoreName(env)
	info := typeInfo(options.Current.Type)

	if info.ShortFormat {
		return ShortName(env)
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
