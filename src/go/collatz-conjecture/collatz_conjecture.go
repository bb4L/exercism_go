package collatzconjecture

import "math"
import "errors"

// CollatzConjecture return the number of iterations needed
func CollatzConjecture(num int) (int, error) {
	actNum := float64(num)
	iterations := 0

	if actNum <= 0 {
		return iterations, errors.New("number must be > 0")
	}

	for {
		if actNum == 1 {
			break
		}

		if math.Mod(actNum, 2.0) == 0 {
			actNum = actNum / 2
		} else {
			actNum = 3*actNum + 1
		}

		iterations++
	}
	return iterations, nil
}
