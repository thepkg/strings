// Package strings implements simple string functions.
package strings

import (
	"strings"
)

// ToUpperFirst returns a copy of the string s with first letter mapped to its upper case.
func ToUpperFirst(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}
