package namespace

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"github.com/ylallemant/t8rctl/pkg/api"
	"github.com/ylallemant/t8rctl/pkg/resource/utils"
)

const (
	ErrorNameGeneration     = "failed to generage namespace name"
	ErrorVarianceGeneration = "failed to generage variance"
)

var (
	branchPrefixRegexp     = regexp.MustCompile(`^([a-zA-Z0-9-_]+)\s*[/:].*`)
	taskBranchRegexp       = regexp.MustCompile(`^([a-zA-Z]{0,5})[-]{0,1}(\d+)-.*`)
	namespaceNameMaxLength = 25
)

func Name(env api.EnvironmentContext) (string, error) {
	variance, err := Variance(env)
	if err != nil {
		return "", errors.Wrap(err, ErrorNameGeneration)
	}

	err = validateNameOptions(env)
	if err != nil {
		return "", errors.Wrap(err, ErrorNameGeneration)
	}

	output := fmt.Sprintf(
		"%s%s%s%s%s",
		env.Stack(),
		regionPart(env.StackRegion()),
		tenantPart(env.StackTenant()),
		utils.Separator,
		variance,
	)

	return utils.Sanitise(output), nil
}

func Variance(env api.EnvironmentContext) (string, error) {
	err := validateVarianceOptions(env)
	if err != nil {
		return "", errors.Wrap(err, ErrorVarianceGeneration)
	}

	variancePrefix := "f"
	branchName := env.GitBranchName()

	prefix, found := branchPrefix(branchName)
	if found {
		branchName = removePrefix(branchName, prefix)
		variancePrefix = prefix[:1]
	}

	branchName = utils.Sanitise(branchName)

	output := ""

	if env.IsGitBaseBranch() {
		output = env.StackDatatier()
	} else {
		output = fmt.Sprintf(
			"%s-%s",
			variancePrefix,
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

	return utils.Sanitise(output), nil
}

func varianceMaxLength(env api.EnvironmentContext) int {
	return namespaceNameMaxLength -
		len(env.Stack()) -
		len(regionPart(env.StackRegion())) -
		len(tenantPart(env.StackTenant()))
}

func regionPart(region string) string {
	if region == "" || region == api.None || region == api.DefaultEnvironmentRegion {
		return ""
	}

	return fmt.Sprintf(
		"%s%s",
		utils.Separator,
		region,
	)
}

func tenantPart(tenant string) string {
	if tenant == "" || tenant == api.None || tenant == api.DefaultEnvironmentTenant {
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
		return fmt.Sprintf(
			"%s%s%s",
			matches[1],
			utils.Separator,
			matches[2],
		), true
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
