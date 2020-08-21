package armstrong

import (
	"math"
	"strconv"
)

// IsNumber checks if the given number is a armstrong number
func IsNumber(number int) bool {
	length := len(strconv.Itoa(number))
	summedNumber := float64(0)
	for i := 0; i < length; i++ {
		summedNumber += math.Pow(float64(number/int(math.Pow(10, float64(i)))%10), float64(length))
	}
	return number == int(summedNumber)
}
