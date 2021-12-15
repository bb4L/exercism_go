package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.

type SillyNephewError struct {
	msg string
}

func (err SillyNephewError) Error() string {
	return err.msg
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	if cows < 0 {
		return 0, SillyNephewError{fmt.Sprintf("silly nephew, there cannot be %d cows", cows)}
	}

	amount, err := weightFodder.FodderAmount()
	if amount < 0 {
		return 0, errors.New("Negative fodder")
	}

	if err != nil {
		if err != ErrScaleMalfunction {
			return 0, err
		}
		amount *= 2

	}
	if cows == 0 {
		return 0, errors.New("Division by zero")
	}
	return amount / float64(cows), nil
}
