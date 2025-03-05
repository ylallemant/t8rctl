package namespace

import (
	"github.com/pkg/errors"
	"github.com/ylallemant/t8rctl/pkg/cli/resource/namespace/options"
	globalOptions "github.com/ylallemant/t8rctl/pkg/cli/resource/options"
)

const (
	ErrorMissingCellProject      = "providing the project name is mandatory"
	ErrorMissingCellStage        = "providing the cell stage name is mandatory"
	ErrorMissingBaseGitBranch    = "providing the base git branch name is mandatory"
	ErrorMissingCurrentGitBranch = "providing the current git branch name is mandatory"
)

func validate(opts *options.Options) error {
	if globalOptions.Current.Project == "" {
		return errors.Errorf(ErrorMissingCellProject)
	}

	if globalOptions.Current.CellStage == "" {
		return errors.Errorf(ErrorMissingCellStage)
	}

	if opts.BaseGitBranch == "" {
		return errors.Errorf(ErrorMissingBaseGitBranch)
	}

	if opts.CurrentGitBranch == "" {
		return errors.Errorf(ErrorMissingCurrentGitBranch)
	}

	return nil
}
