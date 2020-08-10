package anagram

import (
	"sort"
	"strings"
)

func equals(firstSlice, secondSlice []rune) bool {
	if len(firstSlice) != len(secondSlice) {
		return false
	}

	for i, r := range firstSlice {
		if r != secondSlice[i] {
			return false
		}
	}
	return true

}

// Detect anagram's of a given word in a list
func Detect(word string, toTest []string) []string {
	var result []string
	lowerWord := strings.ToLower(word)
	sortedSlice := []rune(lowerWord)
	sort.Slice(sortedSlice, func(i, j int) bool {
		return sortedSlice[i] < sortedSlice[j]
	})

	for _, val := range toTest {
		if len(lowerWord) != len(val) {
			continue
		}
		lowerVal := strings.ToLower(val)

		if lowerWord == lowerVal {
			continue
		}

		testSlice := []rune(lowerVal)

		sort.Slice(testSlice, func(i, j int) bool {
			return testSlice[i] < testSlice[j]
		})

		if !equals(sortedSlice, testSlice) {
			continue
		}

		result = append(result, val)
	}
	return result
}
