package atbash

import (
	"strings"
	"unicode"
)

// Atbash return altbash encoded input
func Atbash(input string) string {
	var sb strings.Builder
	sb.Grow(len(input) + len(input)/5 + 1)

	count := 0
	for _, r := range input {
		isDigit := unicode.IsDigit(r)
		if !isDigit && !unicode.IsLetter(r) {
			continue
		}

		if isDigit {
			sb.WriteRune(r)
		} else {
			sb.WriteRune(rune('z' + 'a' - unicode.ToLower(r)))
		}
		count++

		if count%5 == 0 {
			sb.WriteRune(' ')
		}
	}
	result := sb.String()
	return strings.TrimRight(result, " ")
}
