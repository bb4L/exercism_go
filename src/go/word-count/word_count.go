package wordcount

import (
	"regexp"
	"strings"
)

// Frequency represents a frequency map
type Frequency map[string]int

// WordCount returns a Frequency for the input given
func WordCount(input string) Frequency {
	frequency := make(map[string]int)
	reg, _ := regexp.Compile("[^A-Za-z0-9']+")

	for _, word := range strings.Fields(reg.ReplaceAllString(input, " ")) {
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
