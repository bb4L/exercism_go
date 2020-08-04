package romannumerals

import (
	"errors"
	"strings"
)

// Triplet object to store the three values needed per digit
type Triplet struct {
	single string
	five   string
	ten    string
}

var digitTriplets = []Triplet{
	{"I", "V", "X"},
	{"X", "L", "C"},
	{"C", "D", "M"},
}

func handleDigit(digit int, triplet Triplet) string {
	result := ""

	switch digit {
	case 1, 2, 3:
		result = strings.Repeat(triplet.single, digit)

	case 4, 5, 6, 7, 8:
		if digit < 5 {
			result = strings.Repeat(triplet.single, 5-digit) + triplet.five
		} else {
			result = triplet.five + strings.Repeat(triplet.single, digit-5)
		}

	default:
		result = strings.Repeat(triplet.single, 10-digit) + triplet.ten
	}
	return result
}

// ToRomanNumeral convert a given number to the roman string
func ToRomanNumeral(arabic int) (string, error) {
	if arabic <= 0 {
		return "", errors.New("number must be > 0")
	}

	if arabic > 3000 {
		return "", errors.New("number must be < 3000")
	}

	result := ""
	digit := 0
	remaining := arabic
	count := -1

	for (remaining > 0 || digit > 0) && count < 2 {
		digit = remaining % 10
		remaining = remaining / 10
		count++

		if digit <= 0 {
			continue
		}

		digitResult := handleDigit(digit, digitTriplets[count])
		result = digitResult + result
	}

	if count == 2 && remaining > 0 {
		result = strings.Repeat("M", remaining) + result
	}

	return result, nil
}
