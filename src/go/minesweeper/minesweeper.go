package minesweeper

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func (board Board) Count() error {
	lenghtLine := len(board[0])

	if string(board[0]) != "+"+strings.Repeat("-", len(board[0])-2)+"+" {
		return errors.New("bad format")
	}

	for i, val := range board {
		if len(val) != lenghtLine {
			return errors.New("bad format")
		}

		switch val[0] {
		case []byte("|")[0], []byte("+")[0], []byte("-")[0]:
			// do nothing
		default:
			return errors.New("bad format")
		}
		switch val[len(val)-1] {
		case []byte("|")[0], []byte("+")[0], []byte("-")[0]:
			// do nothing
		default:
			return errors.New("bad format")
		}

		for j, val2 := range val {
			switch val2 {
			case []byte("*")[0], []byte("|")[0], []byte("+")[0], []byte("-")[0]:
				continue

			case []byte(" ")[0]:
				for a := -1; a < 2; a++ {
					if (j+a) < 0 || (j+a) > lenghtLine-1 {
						continue
					}
					for b := -1; b < 2; b++ {
						if (i+b) < 0 || (i+b) > len(board)-1 {
							continue
						}

						if board[i][j] == []byte("*")[0] || board[i][j] == []byte("|")[0] || board[i][j] == []byte("+")[0] || board[i][j] == []byte("-")[0] {
							continue
						}

						if board[i+b][j+a] != []byte("*")[0] {
							continue
						}

						val, _ := strconv.Atoi(string(board[i][j]))
						board[i][j] = []byte(fmt.Sprint(val + 1))[0]
					}
				}
			default:
				return errors.New("invalid value")
			}
		}
	}

	return nil
}
