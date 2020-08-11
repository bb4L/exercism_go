package luhn

import (
	"strconv"
	"strings"
)

// Valid tests if a string contains a per Luhn valid number
func Valid(input string) bool {
	noSpaces := strings.ReplaceAll(input, " ", "")

	if len([]rune(noSpaces)) <= 1 {
		return false
	}

	digits := []rune(noSpaces)
	mod2 := len(digits)%2 == 0
	sum := 0

	for i, digit := range digits {

		var n, err = strconv.Atoi(string(digit))

		if err != nil {
			return false
		}

		if (i%2 == 0) == mod2 {
			n = 2 * n
			if n > 9 {
				n -= 9
			}
		}

		sum += n

	}

	return sum%10 == 0
}
