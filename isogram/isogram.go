//Package isogram Go language track, lesson 6
package isogram

import (
	"strings"
)

// IsIsogram checks wether or not a given word is an Isogram
func IsIsogram(input string) bool {
	if strings.Compare(input, "") == 0 {
		return true
	}
	palavra := []byte(strings.ToLower(input))
	Alphabet := make(map[byte]bool)
	for _, l := range palavra {
		if Alphabet[l] {
			return false
		}
		if l != '-' && l != ' ' {
			Alphabet[l] = true
		}
	}
	return true
}
