package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

func normalize(data string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return ""
	}
	return strings.ToLower(reg.ReplaceAllString(data, ""))
}

// Encode the given data
func Encode(data string) string {
	normalizedRunes := []rune(normalize(data))

	length := len(normalizedRunes)

	if length == 0 {
		return ""
	}

	r := int(math.Round(math.Sqrt(float64(length))))
	c := r

	if r*r < length {
		c = r + 1
	}

	var sb strings.Builder
	sb.Grow(c * r)
	for j := 0; j < c; j++ {
		for i := 0; i < r; i++ {
			if i*c+j >= length {
				sb.WriteRune(' ')
			} else {
				sb.WriteRune(normalizedRunes[i*c+j])
			}
		}
		if j == c-1 {
			continue
		}
		sb.WriteString(" ")
	}

	return sb.String()
}
