package perfect

import (
	"errors"
)

// Classification type for classifications
type Classification int

// Classifications
const (
	ClassificationPerfect   Classification = iota
	ClassificationAbundant  Classification = iota
	ClassificationDeficient Classification = iota
)

// ErrOnlyPositive error for the not positive numbers
var ErrOnlyPositive = errors.New("number has to be greater than 0")

// Classify the given number
func Classify(number int64) (Classification, error) {
	if number <= 0 {
		return -1, ErrOnlyPositive
	}

	if number == 1 {
		return ClassificationDeficient, nil
	}

	sum := int64(1)
	for i := int64(2); i*i < number; i++ {
		if number%i == 0 {
			sum += i + number/i
		}
	}

	switch s := sum; {
	case s == number:
		return ClassificationPerfect, nil
	case s < number:
		return ClassificationDeficient, nil
	default:
		return ClassificationAbundant, nil
	}
}
