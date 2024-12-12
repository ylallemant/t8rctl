package api

import "regexp"

var (
	ResourceNameAllowedChars         = regexp.MustCompile(`[A-Za-z0-9-]`)
	ResourceShortNameNonAllowedChars = regexp.MustCompile(`[^a-z0-9]`)
)

type ResourceInfo struct {
	Description string
	Prefix      string
	ShortFormat bool
}

var ResouceTypePrefix = map[string]ResourceInfo{
	"vault": {
		Description: "secret, key, certificate storage as KeyVaults in Azure cloud",
		Prefix:      "kv",
		ShortFormat: true,
	},
}
