package environment

import "strings"

func sanitiseValue(value string) string {
	value = strings.TrimSpace(value)
	value = strings.ToLower(value)
	return value
}
