package utils

import (
	"regexp"
	"strings"
)

var (
	Separator             = "-"
	forbiddenChars        = regexp.MustCompile(`[^a-zA-Z0-9-]`)
	forbiddenStartChar    = regexp.MustCompile(`^[^a-zA-Z0-9]`)
	forbiddenEndChar      = regexp.MustCompile(`[^a-zA-Z0-9]$`)
	consecutiveSeparators = regexp.MustCompile(`-+`)
)

func Sanitise(resourceName string) string {
	resourceName = replaceForbiddenCharsWithSeparator(resourceName)
	resourceName = removeMultipleConsecutiveSeparators(resourceName)
	resourceName = trim(resourceName)
	return strings.ToLower(resourceName)
}

func replaceForbiddenCharsWithSeparator(resourceName string) string {
	return forbiddenChars.ReplaceAllString(resourceName, Separator)
}

func removeMultipleConsecutiveSeparators(resourceName string) string {
	return consecutiveSeparators.ReplaceAllString(resourceName, Separator)
}

func trim(resourceName string) string {
	resourceName = strings.TrimSpace(resourceName)
	resourceName = forbiddenStartChar.ReplaceAllString(resourceName, "")
	return forbiddenEndChar.ReplaceAllString(resourceName, "")
}
