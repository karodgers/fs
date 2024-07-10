package handlers

import (
	"unicode"
)

// Checks if a string contains non-ASCII characters i.e characters >127 in the ascii table.
func ContainsNonASCII(s string) bool {
	for _, char := range s {
		if char > unicode.MaxASCII {
			return true
		}
	}
	return false
}
