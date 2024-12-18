package namespace

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ylallemant/t8rctl/pkg/api"
	"github.com/ylallemant/t8rctl/pkg/resource/utils"
)

var (
	branchPrefixRegexp     = regexp.MustCompile(`^([a-zA-Z0-9-_]+)[/:].*`)
	taskBranchRegexp       = regexp.MustCompile(`^([a-zA-Z0-9]{0,5}-\d+)[-_/].*`)
	namespaceNameMaxLength = 25
)

func Name(env api.EnvironmentContext) string {
	variance := Variance(env)
	output := fmt.Sprintf(
		"%s%s%s%s%s",
		env.Project(),
		regionPart(env.StackRegion()),
		tenantPart(env.StackTenant()),
		utils.Separator,
		variance,
	)

	return utils.Sanitise(output)
}

func Variance(env api.EnvironmentContext) string {
	branchName := env.GitBranchName()
	prefix, found := branchPrefix(branchName)

	if found {
		branchName = removePrefix(branchName, prefix)
	}

	branchName = utils.Sanitise(branchName)

	output := ""

	if env.IsGitBaseBranch() {
		output = env.StackDatatier()
	} else {
		output = fmt.Sprintf(
			"f-%s",
			fromBranch(branchName),
		)
	}

	maxLength := varianceMaxLength(env)

	if len(output) > maxLength {
		hash := utils.ResourceHash(output)
		output = fmt.Sprintf(
			"%s-%s",
			output[0:maxLength-len(hash)],
			hash,
		)
	}

	return utils.Sanitise(output)
}

func varianceMaxLength(env api.EnvironmentContext) int {
	return namespaceNameMaxLength -
		len(env.Project()) -
		len(regionPart(env.StackRegion())) -
		len(tenantPart(env.StackTenant()))
}

func regionPart(region string) string {
	if region == "" || region == api.DefaultEnvironmentRegion {
		return ""
	}

	return fmt.Sprintf(
		"%s%s",
		utils.Separator,
		region,
	)
}

func tenantPart(tenant string) string {
	if tenant == "" || tenant == api.DefaultEnvironmentTenant {
		return ""
	}

	return fmt.Sprintf(
		"%s%s",
		utils.Separator,
		tenant,
	)
}

func branchPrefix(branchName string) (string, bool) {
	matches := branchPrefixRegexp.FindStringSubmatch(branchName)
	if len(matches) > 0 {
		return matches[1], true
	}

	return "", false
}

func removePrefix(branchName, prefix string) string {
	// remove prefix string
	branchName = strings.ReplaceAll(
		branchName,
		prefix,
		"",
	)

	// remove separator
	return branchName[1:]
}

func taskReference(branchName string) (string, bool) {
	matches := taskBranchRegexp.FindStringSubmatch(branchName)
	if len(matches) > 0 {
		return matches[1], true
	}

	return "", false
}

func fromBranch(branchName string) string {
	taslRef, found := taskReference(branchName)

	if found {
		return taslRef
	}

	return utils.Sanitise(branchName)
}
