package allyourbase

import (
	"errors"
	"math"
)

// ConvertToBase convert the base of a representation
func ConvertToBase(base int, digits []int, outpuBase int) (result []int, err error) {
	if base < 2 {
		return nil, errors.New("input base must be >= 2")
	}
	if outpuBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}
	result = []int{}
	var sum = 0
	length := len(digits) - 1
	for i := length; i >= 0; i-- {
		if digits[i] < 0 || digits[i] >= base {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
		sum += digits[i] * int(math.Pow(float64(base), float64(length-i)))
	}

	var exp = 1
	for {
		if sum < outpuBase {
			result = append([]int{sum}, result...)

			break
		}
		remainder := sum % outpuBase
		sum = sum / outpuBase
		result = append([]int{remainder}, result...)

		exp++
	}
	return
}
