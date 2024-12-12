package options

import (
	"github.com/ylallemant/t8rctl/pkg/api"
)

var (
	Current = NewOptions()
)

func NewOptions() *Options {
	options := new(Options)

	options.Provider = api.Azure
	options.CellRegion = "global"
	options.CellTenant = "shared"
	options.CellId = ""

	return options
}

type Options struct {
	Provider   string
	Project    string
	CellStage  string
	CellTenant string
	CellRegion string
	CellId     string
}
