package isbn

import (
	"strings"
	"unicode"
)

// IsValidISBN checks if a string is a valid ISBN
func IsValidISBN(isbn string) bool {
	firstX := strings.Index(isbn, "X")
	if firstX != len(isbn)-1 && firstX != -1 {
		return false
	}

	runes := []rune(strings.Replace(isbn, "-", "", -1))

	if len(runes) != 10 {
		return false
	}

	sum := 0

	for i, r := range runes {
		if !unicode.IsDigit(r) {
			if i != 9 || r != 'X' {
				return false
			}

			sum += 10
			continue
		}

		sum += (10 - i) * int(r-'0')
	}

	return sum%11 == 0
}
