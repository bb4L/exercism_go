package rotationalcipher

import (
	"strings"
	"unicode"
)

// RotationalCipher encode with rotational cipher
func RotationalCipher(input string, shift int) string {
	var sb strings.Builder
	sb.Grow(len(input))
	maxR := 'z'
	minR := 'a'
	for _, r := range input {
		if unicode.IsLetter(r) {
			if unicode.IsUpper(r) {
				maxR = 'Z'
				minR = 'A'
			} else {
				maxR = 'z'
				minR = 'a'
			}

			newR := r + rune(shift)

			if newR > maxR {
				newR = minR + (newR - (maxR + 1))
			} else if newR < minR {
				newR = maxR - minR + newR + 1
			}
			sb.WriteRune(newR)
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}
