package hamming

import "errors"

// Distance computes the hamming distance of a and b
func Distance(a, b string) (int, error) {
	distance := 0
	ar := []rune(a)
	br := []rune(b)

	if len(ar) != len(br) {
		return distance, errors.New("can not work with strings with different size")
	}

	for index, runeValue := range ar {
		if runeValue != br[index] {
			distance++
		}
	}
	return distance, nil
}
