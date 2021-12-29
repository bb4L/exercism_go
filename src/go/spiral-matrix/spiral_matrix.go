package spiralmatrix

type Position struct {
	Row int
	Col int
}

type Direction int

const (
	Right Direction = iota
	Down
	Left
	Up
)

func SpiralMatrix(size int) (result [][]int) {

	result = make([][]int, size)
	for i := 0; i < size; i++ {
		result[i] = make([]int, size)
	}

	count := 1
	pos := Position{0, 0}
	if size > 0 {
		result[0][0] = 1
	}

	for i := size; i > 0; i -= 2 {
		for _, direction := range []Direction{Right, Down, Left, Up} {
			for k := 0; k < i-1; k++ {
				if count > size*size {
					break
				}

				result[pos.Row][pos.Col] = count

				if direction == Up && IntAbs(i-k-1) == 1 {
					pos.Col += 1
					if size%2 != 0 {
						result[pos.Row][pos.Col] = 1 + count
					}
					count++
					break
				}
				switch direction {
				case Right:
					pos.Col += 1
				case Down:
					pos.Row += 1
				case Left:
					pos.Col -= 1
				case Up:
					pos.Row -= 1
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
