package encode

import (
	"strconv"
	"strings"
	"unicode"
)

func writeRune(sb strings.Builder, r rune, n int) strings.Builder {
	if n > 0 {
		sb.WriteString(string(n))
	}
	sb.WriteRune(r)

	return sb
}

// RunLengthEncode encode the data
func RunLengthEncode(data string) string {
	runes := []rune(data)
	if len(runes) < 1 {
		return ""
	}
	actualRune := runes[0]
	actualCount := 1
	var sb strings.Builder
	sb.Grow(2 * len(runes))
	for _, r := range runes[1:] {
		if actualRune == r {
			actualCount++
		} else {
			if actualCount > 1 {
				sb.WriteString(strconv.Itoa(actualCount))
				actualCount = 1
			}
			sb.WriteRune(actualRune)
			actualRune = r
		}
	}

	if actualCount > 1 {
		sb.WriteString(strconv.Itoa(actualCount))
		actualCount = 1
	}
	sb.WriteRune(actualRune)

	return sb.String()
}

// RunLengthDecode decode the data
func RunLengthDecode(data string) string {
	var sb strings.Builder

	count := 0
	for _, r := range []rune(data) {
		if unicode.IsDigit(r) {
			if count > 0 {
				count = count * 10
			}
			count += int(r - '0')
		} else {
			if count > 0 {
				sb.WriteString(strings.Repeat(string(r), count))
				count = 0
			} else {
				sb.WriteRune(r)
			}
		}
	}
	return sb.String()
}
