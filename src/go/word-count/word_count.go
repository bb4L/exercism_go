package wordcount

import (
	"strings"
	"unicode"
)

// Frequency represents a frequency map
type Frequency map[string]int

// WordCount returns a Frequency for the input given
func WordCount(input string) Frequency {
	frequency := make(map[string]int)

	inline := func(r rune) bool { return !(unicode.IsLetter(r) || unicode.IsNumber(r) || r == '\'') }
	for _, word := range strings.FieldsFunc(input, inline) {
		lowerWord := strings.Trim(strings.ToLower(word), "'")
		val, ok := frequency[lowerWord]
		if ok {
			frequency[lowerWord] = val + 1
		} else {
			frequency[lowerWord] = 1
		}
	}
	return frequency
}
