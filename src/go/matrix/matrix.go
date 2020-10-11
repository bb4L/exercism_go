package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix representing a matrix
type Matrix [][]int

// New create new matrix
func New(input string) (Matrix, error) {
	rowlen := 0
	rows := [][]int{}
	for _, row := range strings.Split(input, "\n") {
		newRow := []int{}
		vals := strings.Split(strings.TrimSpace(row), " ")
		if rowlen > 0 && len(vals) != rowlen {
			return nil, errors.New("Rows don't have the same length")
		}
		rowlen = len(vals)
		for _, val := range vals {
			newVal, err := strconv.Atoi(val)
			if err != nil {
				return nil, errors.New("Could not parse all values")
			}
			newRow = append(newRow, newVal)
		}
		rows = append(rows, newRow)
	}

	return rows, nil
}

// Rows returns all the rows
func (m Matrix) Rows() [][]int {
	result := [][]int{}
	for _, row := range m {
		newRow := []int{}
		for _, val := range row {
			newRow = append(newRow, val)
		}
		result = append(result, newRow)
	}
	return result
}

// Cols returns all the rows
func (m Matrix) Cols() [][]int {
	cols := [][]int{}
	for i := 0; i < len(m[0]); i++ {
		newCol := []int{}
		for _, row := range m {
			newCol = append(newCol, row[i])
		}
		cols = append(cols, newCol)
	}
	return cols
}

// Set sets a value
func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(m) || col < 0 || col >= len(m[0]) {
		return false
	}
	m[row][col] = val
	return true
}
