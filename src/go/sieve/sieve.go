package sieve

// Sieve returns all the primenumbers up to the limit
func Sieve(limit int) (result []int) {
	vals := make([]bool, limit-1)
	min := 2

	for i := 2; i <= limit; i++ {
		if !vals[i-min] {
			result = append(result, i)
		}
		if 2*i > limit {
			continue
		}
		for k := 2 * i; k <= limit; k = k + i {
			vals[k-min] = true
		}
	}

	return
}
