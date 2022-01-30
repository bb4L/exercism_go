package spiralmatrix

type position struct {
	Row int
	Col int
}

type direction int

const (
	right direction = iota
	down
	left
	up
)

// SpiralMatrix returns e spiral matrix with the given size
func SpiralMatrix(size int) (result [][]int) {

	result = make([][]int, size)
	for i := 0; i < size; i++ {
		result[i] = make([]int, size)
	}

	count := 1
	pos := position{0, 0}
	if size > 0 {
		result[0][0] = 1
	}

	for i := size; i > 0; i -= 2 {
		for _, direction := range []direction{right, down, left, up} {
			for k := 0; k < i-1; k++ {
				if count > size*size {
					break
				}

				result[pos.Row][pos.Col] = count

				if direction == up && IntAbs(i-k-1) == 1 {
					pos.Col++
					if size%2 != 0 {
						result[pos.Row][pos.Col] = 1 + count
					}
					count++
					break
				}
				switch direction {
				case right:
					pos.Col++
				case down:
					pos.Row++
				case left:
					pos.Col--
				case up:
					pos.Row--
				}
				count++
			}
		}
	}
	return result
}

// IntAbs returns the absolute value of an integer.
func IntAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
