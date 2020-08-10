package pangram

import "unicode"

// IsPangram tests if the input is a pangram
func IsPangram(input string) bool {
	set := make(map[rune]struct{})

	for _, ru := range input {
		if !unicode.IsLetter(ru) {
			continue
		}

		lower := unicode.ToLower(ru)
		_, existing := set[lower]

		if existing {
			continue
		}

		set[lower] = struct{}{}
	}

	if len(set) != 26 {
		return false
	}

	return true
}
