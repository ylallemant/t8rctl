package name

import (
	"github.com/pkg/errors"
	"github.com/ylallemant/t8rctl/pkg/api"
	"github.com/ylallemant/t8rctl/pkg/cli/resource/name/options"
	globalOptions "github.com/ylallemant/t8rctl/pkg/cli/resource/options"
)

func Validate() error {
	bools := 0

	if options.Current.Core {
		bools += 1
	}

	if options.Current.Short {
		bools += 1
	}

	if options.Current.Hash {
		bools += 1
	}

	if options.Current.Caf {
		bools += 1
	}

	if bools > 1 {
		return errors.Errorf("only one output flag is allowed")
	}

	if options.Current.Caf || options.Current.Short {
		if options.Current.Type != "" && !typeExists(options.Current.Type) {
			return errors.Errorf("unknown resource type \"%s\"", options.Current.Type)
		}

		if globalOptions.Current.CellId == "" && !typeExists(options.Current.Type) {
			return errors.Errorf("cell-id is missing")
		}
	}

	if options.Current.Name == "" && options.Current.Static == "" {
		return errors.Errorf("one of following inputs is manatory: \"name\" or \"static\"")
	}

	if options.Current.Name != "" && (globalOptions.Current.CellStage == "" || globalOptions.Current.Project == "") {
		return errors.Errorf("using dynamic naming makes following values mandatory: \"project\" and \"cell-stage\"")
	}

	if options.Current.Caf && (globalOptions.Current.CellId == "") {
		return errors.Errorf("requesting a CAF name makes some values mandatory: \"cell-id\"")
	}

	return nil
}

func typeExists(typeName string) bool {
	_, found := api.ResouceTypePrefix[typeName]

	return found
}
