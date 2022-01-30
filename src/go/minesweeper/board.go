package minesweeper

import "bytes"

// Board for minesweeper
type Board [][]byte

func (b Board) String() string {
	return "\n" + string(bytes.Join(b, []byte{'\n'}))
}
