package util

import (
	"unicode"
)

// capitalize the first letter of a string
func CapitalizeFirst(s string) string {
	if s == "" {
		return s
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
