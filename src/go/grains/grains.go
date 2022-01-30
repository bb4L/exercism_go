package grains

import (
	"errors"
)

// Square return value on field n
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return uint64(0), errors.New("wrong input should be in range 1 to 64 including")
	}
	return uint64(1 << (n - 1)), nil
}

// Total return the total of the square
func Total() uint64 {
	k := uint64(1 << 32)
	return k*k - 1
}
