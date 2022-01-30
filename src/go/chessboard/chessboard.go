package chessboard

// Rank which stores if a square is occupied by a piece
type Rank []bool

// Chessboard contains a map of eight Ranks, accessed with values from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) (count int) {
	for _, val := range cb[rank] {
		if val {
			count++
		}
	}
	return
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) (count int) {
	for _, val := range cb {
		if file > len(val) {
			continue
		}
		if val[file-1] {
			count++
		}
	}
	return
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) (count int) {
	for _, file := range cb {
		count += len(file)
	}
	return
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) (count int) {
	for _, file := range cb {
		for _, val := range file {
			if val {
				count++
			}
		}
	}
	return
}
