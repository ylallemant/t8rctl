package namespace

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ylallemant/t8rctl/pkg/api"
	"github.com/ylallemant/t8rctl/pkg/resource/utils"
)

var (
	branchPrefixRegexp = regexp.MustCompile(`^([a-zA-Z0-9-_]+)[/:].*`)
	taskBranchRegexp   = regexp.MustCompile(`^([a-zA-Z0-9]{0,5}-\d+)[-_/].*`)
	maxLength          = 25
)

func Name(env *api.EnvironmentContext) string {
	branchName := env.BranchName
	prefix, found := branchPrefix(branchName)

	if found {
		branchName = removePrefix(branchName, prefix)
	}

	branchName = utils.Sanitise(branchName)

	output := ""

	if env.IsBase {
		output = fmt.Sprintf(
			"%s%s%s",
			env.Project,
			utils.Separator,
			env.Datatier,
		)
	} else {
		output = fmt.Sprintf(
			"%s%sf%s%s",
			env.Project,
			utils.Separator,
			utils.Separator,
			fromBranch(branchName),
		)
	}

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
