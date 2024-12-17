package namespace

import (
	"github.com/pkg/errors"
	"github.com/ylallemant/t8rctl/pkg/cli/resource/namespace/options"
	globalOptions "github.com/ylallemant/t8rctl/pkg/cli/resource/options"
)

func validate() error {
	if globalOptions.Current.CellStage == "" {
		return errors.Errorf("providing the cell stage is mandatory")
	}

	if globalOptions.Current.Project == "" {
		return errors.Errorf("providing the project name is mandatory")
	}

	if options.Current.CurrentGitBranch == "" {
		return errors.Errorf("providing the current git branch is mandatory")
	}

	if options.Current.BaseGitBranch == "" {
		return errors.Errorf("providing the base git branch is mandatory")
	}

	return nil
}
