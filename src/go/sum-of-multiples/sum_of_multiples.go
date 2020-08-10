package summultiples

// SumMultiples returns the sum of the multiples
func SumMultiples(limit int, divisors ...int) (result int) {
	// multiples := map[int]struct{}{}

	result = 0

	for i := 1; i < limit; i++ {
		for _, k := range divisors {
			if k == 0 {
				continue
			}
			if i%k == 0 {
				result += i
				break
			}
		}
	}

	// for _, k := range divisors {
	// 	if k <= 0 {
	// 		continue
	// 	}
	// 	for n := k; n < limit; {
	// 		multiples[n] = struct{}{}
	// 		n = n + k
	// 	}
	// }

	// for k := range multiples {
	// 	result += k
	// }
	return result
}
