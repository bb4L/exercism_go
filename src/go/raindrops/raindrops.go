package raindrops

import "strconv"

var mapping = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

// Convert a number to a string following the mapping rules given
func Convert(n int) string {
	result := ""
	for i, s := range mapping {
		if n%i == 0 {
			result += s
		}
	}
	if result == "" {
		return strconv.FormatInt(int64(n), 10)
	}

	return result
}
