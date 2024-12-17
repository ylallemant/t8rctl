package utils

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func ResourceHash(resourceName string) string {
	hasher := sha1.New()
	io.WriteString(hasher, resourceName)
	hash := fmt.Sprintf("%x", hasher.Sum(nil))

	return hash[:5]
}
