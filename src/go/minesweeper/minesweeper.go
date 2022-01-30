package minesweeper

import (
	"errors"
	"strings"
)

var errBadFormat = errors.New("bad format")

// Count adds the bomb counts for all the fields
func (board Board) Count() error {
	lenghtLine := len(board[0])

	for _, border := range []string{string(board[0]), string(board[len(board)-1])} {
		if border != "+"+strings.Repeat("-", len(board[0])-2)+"+" {
			return errBadFormat
		}
	}

	for i, line := range board {

		if len(line) != lenghtLine {
			return errBadFormat
		}

		for _, border := range []byte{line[0], line[len(line)-1]} {

			switch border {
			case '|':
				// do nothing
			case '+':
				if i != 0 && i != len(board)-1 {
					return errBadFormat
				}
			default:
				return errBadFormat
			}
		}

		for j, entry := range line {
			bombCount := byte('0')

			switch entry {
			case '*', '|', '+', '-':
				continue

			case ' ':
				for a := -1; a < 2; a++ {
					if (j+a) < 0 || (j+a) > lenghtLine-1 {
						continue
					}

					for b := -1; b < 2; b++ {
						if (i+b) < 0 || (i+b) > len(board)-1 {
							continue
						}

						if board[i+b][j+a] != '*' {
							continue
						}

						bombCount++
					}
				}

			default:
				return errors.New("invalid value")
			}

			if bombCount != '0' {
				board[i][j] = bombCount
			}

		}

	}

	return nil
}
