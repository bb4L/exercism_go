package yacht

import (
	"sort"
)

const (
	YACHT           = "yacht"
	LITTLE_STRAIGHT = "little straight"
	BIG_STRAIGHT    = "big straight"
	FULL_HOUSE      = "full house"
	CHOICE          = "choice"
	FOUR_OF_A_KIND  = "four of a kind"
	ONES            = "ones"
	TWOS            = "twos"
	THREES          = "threes"
	FOURS           = "fours"
	FIVES           = "fives"
	SIXES           = "sixes"
)

var NUMBERS = map[string]int{ONES: 1, TWOS: 2, THREES: 3, FOURS: 4, FIVES: 5, SIXES: 6}

type dices []int

func (slice dices) Len() int {
	return len(slice)
}

func (slice dices) Less(i, j int) bool {
	return slice[i] < slice[j]
}
func (slice dices) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func Score(diceSlice []int, category string) (result int) {
	switch category {
	case CHOICE:
		return sum(diceSlice)

	case ONES, TWOS, THREES, FOURS, FIVES, SIXES:
		return countN(diceSlice, NUMBERS[category])

	case FOUR_OF_A_KIND:
		sort.Ints(diceSlice)
		if diceSlice[0] != diceSlice[3] && diceSlice[1] != diceSlice[4] {
			return
		}
		result = 4 * diceSlice[1]

	case FULL_HOUSE:
		sort.Ints(diceSlice)

		if !(diceSlice[0] == diceSlice[1] && diceSlice[2] == diceSlice[4] && diceSlice[0] != diceSlice[4]) && !(diceSlice[0] == diceSlice[2] && diceSlice[3] == diceSlice[4] && diceSlice[0] != diceSlice[4]) {
			return
		}

		result = sum(diceSlice)

	case YACHT:
		val := diceSlice[0]
		for _, dice := range diceSlice[1:] {
			if dice != val {
				return
			}
		}
		result = 50

	case LITTLE_STRAIGHT:
		sort.Ints(diceSlice)
		if !checkStraight(diceSlice, 1) {
			return
		}
		result = 30

	case BIG_STRAIGHT:
		sort.Sort(dices(diceSlice))
		if !checkStraight(diceSlice, 2) {
			return
		}
		result = 30
	}

	return
}

func countN(dices []int, n int) (result int) {
	for _, v := range dices {
		if v == n {
			result += n
		}
	}
	return
}

func checkStraight(diceSlice []int, startingValue int) bool {
	if diceSlice[4] != startingValue+4 || diceSlice[0] != startingValue {
		return false
	}

	for i, v := range diceSlice {
		if i+startingValue != v {
			return false
		}
	}
	return true
}

func sum(diceSlice []int) (result int) {
	for _, v := range diceSlice {
		result += v
	}
	return
}
