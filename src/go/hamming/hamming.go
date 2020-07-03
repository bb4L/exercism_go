package hamming

import "errors"

// Distance computes the hamming distance of a and b
func Distance(a, b string) (int, error) {
	distance := 0
	if len(a) == len(b) {
		for k := 0; k < len(a); k++ {
			if a[k] != b[k] {
				distance++
			}
		}
		return distance, nil
	}
	return distance, errors.New("can not work with strings with different size")
}
