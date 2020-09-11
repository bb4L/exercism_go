package prime

import "math"

// Factors returns the primefactors of the given number
func Factors(number int64) (result []int64) {
	var remainingNumber = number
	result = []int64{}

	for {
		if remainingNumber%2 != 0 {
			break
		}
		result = append(result, 2)
		remainingNumber /= 2
	}

	if remainingNumber <= 2 {
		return
	}

	sqrt := int64(math.Sqrt(float64(remainingNumber)))

	for i := int64(3); i <= sqrt; i = i + 2 {
		if remainingNumber < i {
			return
		}
		for {
			if remainingNumber%i != 0 || remainingNumber <= 1 {
				break
			}
			if i == remainingNumber {
				result = append(result, i)
				return
			}
			remainingNumber /= i
			result = append(result, i)
		}
	}

	if remainingNumber > 1 {
		result = append(result, remainingNumber)
	}
	return

}
