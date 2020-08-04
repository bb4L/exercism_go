package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram tests if a given string is a isogram
func IsIsogram(text string) bool {
	var lowerCaseText = strings.ToLower(text)
	for _, ru := range []rune(text) {
		if !unicode.IsLetter(ru) {
			continue
		}

		if strings.IndexAny(lowerCaseText, string(ru)) != strings.LastIndexAny(lowerCaseText, string(ru)) {
			return false
		}
	}
	return true
}
