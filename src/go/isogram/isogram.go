package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram tests if a given string is a isogram
func IsIsogram(text string) bool {
	var lowerCaseText = strings.ToLower(text)
	var seenRunes = make(map[rune]bool)
	for _, ru := range lowerCaseText {
		if !unicode.IsLetter(ru) {
			continue
		}

		if seenRunes[ru] {
			return false
		}

		seenRunes[ru] = true
	}
	return true
}
