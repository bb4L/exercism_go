package matrix

import (
	"strconv"
	"strings"
)

// Define the Matrix and Pair types here.
type Matrix struct {
	data [][]int
}
type Pair [2]int

func New(s string) (*Matrix, error) {
	m := Matrix{}

	for i, row := range strings.Split(s, "\n") {
		m.data = append(m.data, []int{})

		for _, val := range strings.Split(row, " ") {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				return &m, err
			}

			m.data[i] = append(m.data[i], intVal)
		}
	}

	return &m, nil
}

func (m *Matrix) Saddle() []Pair {
	maxVals := [][]int{}
	saddlePoints := []Pair{}

	for _, row := range m.data {
		var rowMax int
		var rowMaxs []int

		for j, val := range row {
			if j == 0 || val > rowMax {
				rowMax = val
				rowMaxs = []int{j}
			} else if val == rowMax {
				rowMaxs = append(rowMaxs, j)
			}
		}
		maxVals = append(maxVals, rowMaxs)
	}

	for i, maxIdxs := range maxVals {
		maxVal := m.data[i][maxIdxs[0]]

		for _, maxIdx := range maxIdxs {
			isSaddle := true

			for j, r := range m.data {
				if j == i {
					continue
				}

				if r[maxIdx] < maxVal {
					isSaddle = false
					break
				}
			}

			if isSaddle {
				saddlePoints = append(saddlePoints, Pair{i, maxIdx})
			}
		}
	}

	return saddlePoints
}
