package prime

import (
	"math"
)

var firstSixPrimes = []int{2, 3, 5, 7, 11, 13}

func isPrime(n int) bool {
	for _, i := range firstSixPrimes {
		if n%i == 0 {
			return n == i
		}
	}

	for i := 15; i <= int(math.Sqrt(float64(n))); i = i + 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Nth returns the n-th prime
func Nth(n int) (int, bool) {
	if n < 1 {
		return 0, false
	}
	if n < 7 {
		return firstSixPrimes[n-1], true
	}

	x := 13
	for count := 6; count < n; {
		x = x + 2
		if isPrime(x) {
			count++
		}
	}
	return x, true
}
