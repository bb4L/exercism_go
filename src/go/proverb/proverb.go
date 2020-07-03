package proverb

import "fmt"

// Proverb create proverb from list of strings
func Proverb(rhyme []string) []string {
	inputSize := len(rhyme)
	result := make([]string, inputSize)

	if inputSize == 0 {
		return result
	}

	if inputSize > 1 {
		for i, s := range rhyme[1:] {
			result[i] = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], s)
		}
	}

	result[inputSize-1] = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
	return result
}
