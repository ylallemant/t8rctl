package namespace

import (
	"github.com/pkg/errors"
	"github.com/ylallemant/t8rctl/pkg/api"
)

const (
	ErrorMissingGitBranch    = "you have to specify the git branch"
	ErrorMissingStackName    = "you have to specify the stack name"
	ErrorMissingDatatierName = "you have to specify the datatier name"
)

func validateVarianceOptions(env api.EnvironmentContext) error {

	if env.GitBranchName() == "" {
		return errors.New(ErrorMissingGitBranch)
	}

	return nil
}

func validateNameOptions(env api.EnvironmentContext) error {

	if env.Stack() == "" {
		return errors.New(ErrorMissingStackName)
	}

	return nil
}
