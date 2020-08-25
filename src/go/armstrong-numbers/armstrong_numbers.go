package armstrong

import (
	"strconv"
)

func pow(a, b int) (result int) {
	result = 1
	for i := 1; i <= b; i++ {
		result *= a
	}
	return
}

// IsNumber checks if the given number is a armstrong number
func IsNumber(number int) bool {
	length := len(strconv.Itoa(number))
	summedNumber := 0
	for i := 0; i < length; i++ {
		summedNumber += pow((number/pow(10, i))%10, length)
	}
	return number == int(summedNumber)
}
