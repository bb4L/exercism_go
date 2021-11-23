package matrix

import (
	"strconv"
	"strings"
)

// Define the Matrix and Pair types here.
type Matrix struct {
	data         [][]int
	saddlePoints []Pair
}
type Pair [2]int

func New(s string) (*Matrix, error) {
	m := Matrix{}
	maxVals := [][]int{}

	for i, row := range strings.Split(s, "\n") {
		m.data = append(m.data, []int{})

		var rowMax int
		var rowMaxs []int

		for j, val := range strings.Split(row, " ") {
			intVal, _ := strconv.Atoi(val)

			if j == 0 || intVal > rowMax {
				rowMax = intVal
				rowMaxs = []int{j}
			} else if intVal == rowMax {
				rowMaxs = append(rowMaxs, j)
			}

			m.data[i] = append(m.data[i], intVal)
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
				m.saddlePoints = append(m.saddlePoints, Pair{i, maxIdx})
			}
		}

	}
	return &m, nil
}

func (m *Matrix) Saddle() []Pair {
	return m.saddlePoints
}
