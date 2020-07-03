package darts

import "math"

// Score returns the score as defined
func Score(x, y float64) int {
	distance := math.Sqrt(math.Pow(x, 2.0) + math.Pow(y, 2.0))

	if distance > 10 {
		return 0
	}
	if distance > 5 {
		return 1
	}

	if distance > 1 {
		return 5
	}

	return 10
}
