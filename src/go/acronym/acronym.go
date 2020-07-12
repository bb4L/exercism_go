package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate creates abbreviations from input string.
func Abbreviate(sInitial string) string {
	s := strings.ReplaceAll(sInitial, "'", "")
	re := regexp.MustCompile("[^a-zA-Z]+")
	t := re.ReplaceAllLiteralString(s, " ")

	result := ""
	for _, word := range strings.Fields(t) {
		result = result + strings.ToUpper(string([]rune(word)[0]))
	}

	return result
}
