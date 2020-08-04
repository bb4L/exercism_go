package raindrops

import "strconv"

// Pair mapping int to sound
type Pair struct {
	Key   int
	Value string
}

var mapping = [3]Pair{{3, "Pling"}, {5, "Plang"}, {7, "Plong"}}

// Convert a number to a string following the mapping rules given
func Convert(n int) string {
	result := ""

	for _, s := range mapping {
		if n%s.Key == 0 {
			result += s.Value
		}
	}

	if result == "" {
		return strconv.Itoa(n)
	}

	return result
}
