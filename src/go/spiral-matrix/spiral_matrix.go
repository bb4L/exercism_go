package spiralmatrix

import "math"

func SpiralMatrix(size int) (result [][]int) {

	result = make([][]int, size)
	for i := 0; i < size; i++ {
		result[i] = make([]int, size)
	}

	count := 1
	pos := [2]int{0, 0}
	if size > 0 {
		result[0][0] = 1
	}

	for i := size; i > 0; i -= 2 {
		for direction := 0; direction < 4; direction++ {
			for k := 0; k < i-1; k++ {
				if count > size*size {
					break
				}

				result[pos[0]][pos[1]] = count

				if direction == 3 && math.Abs(float64(i-k-1)) == 1 {
					pos[1] += 1
					if size%2 != 0 {
						result[pos[0]][pos[1]] = 1 + count
					}
					count++
					break
				}
				switch direction {
				case 0:
					pos[1] += 1
				case 1:
					pos[0] += 1
				case 2:
					pos[1] -= 1
				case 3:
					pos[0] -= 1
				}
				count++
			}
		}
	}
	return
}
