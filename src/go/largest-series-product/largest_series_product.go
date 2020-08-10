package lsproduct

import (
	"errors"
	"unicode"
)

func checkRune(r rune) (int64, error) {
	if !unicode.IsDigit(r) {
		return -1, errors.New("digits input must only contain digits")
	}
	return int64(r - '0'), nil
}

func findMaxStartingFrom(runes []rune, startIdx int, span int) (int64, error) {
	current := int64(1)
	if startIdx > len(runes)-span {
		return 0, nil
	}

	for i := startIdx; i < span+startIdx; i++ {
		r, err := checkRune(runes[i])

		if err != nil {
			return -1, err
		}
		if r == 0 {
			n, err := findMaxStartingFrom(runes, i+1, span)
			if err == nil {
				return n, err
			}
		}
		current *= r
	}
	max := current

	for i := span + startIdx; i < len(runes); i++ {
		rOld, _ := checkRune(runes[i-span])
		r, err := checkRune(runes[i])

		if err != nil {
			return -1, err
		}

		if r == 0 {
			n, err2 := findMaxStartingFrom(runes, i+1, span)
			if err2 != nil {
				return n, err2
			}
			if max > n {
				return max, nil
			}
			return n, nil
		}

		current = (current * r / rOld)

		if current > max {
			max = current
		}
	}
	return max, nil
}

// LargestSeriesProduct returns the largest product of a series of length span
func LargestSeriesProduct(digits string, span int) (int64, error) {
	runes := []rune(digits)
	if len(runes) == 0 && span == 0 {
		return 1, nil
	}

	if span == 0 {
		return 1, nil
	}

	if span < 0 {
		return -1, errors.New("span must be greater than zero")
	}

	if len(runes) < span {
		return -1, errors.New("span must be smaller than string length")
	}

	return findMaxStartingFrom(runes, 0, span)
}
