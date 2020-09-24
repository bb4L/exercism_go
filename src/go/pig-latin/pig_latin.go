package piglatin

import (
	"strings"
)

// Sentence returns the sentence in piglatin
func Sentence(sentence string) string {
	result := ""

	for _, w := range strings.Split(sentence, " ") {
		result = result + Word(w) + " "
	}

	return strings.Trim(result, " ")
}

// Word converts a word into piglatin
func Word(word string) string {

	index := strings.IndexAny(word, "aeiou")
	if index == 0 || strings.HasPrefix(word, "xr") || strings.HasPrefix(word, "yt") {
		return word + "ay"
	}

	if index > 0 {
		if word[index-1:index+1] == "qu" || (index == 1 && word[index-1:index] == "y") {
			index = strings.IndexAny(word, "aeio")
		}
		return word[index:len(word)] + word[:index] + "ay"
	}

	index = strings.Index(word, "y")
	if index > 0 {
		return word[index:len(word)] + word[:index] + "ay"

	}

	return word
}
