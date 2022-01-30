package yacht

import (
	"sort"
)

const (
	yacht          = "yacht"
	littleStraight = "little straight"
	bigStraight    = "big straight"
	fullHouse      = "full house"
	choice         = "choice"
	fourOfAKind    = "four of a kind"
	ones           = "ones"
	twos           = "twos"
	threes         = "threes"
	fours          = "fours"
	fives          = "fives"
	sixes          = "sixes"
)

var numbers = map[string]int{ones: 1, twos: 2, threes: 3, fours: 4, fives: 5, sixes: 6}

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

// Score returns the score for a slice of dices and the chosen category
func Score(diceSlice []int, category string) (result int) {
	switch category {
	case choice:
		return sum(diceSlice)

	case ones, twos, threes, fours, fives, sixes:
		return countN(diceSlice, numbers[category])

	case fourOfAKind:
		sort.Ints(diceSlice)
		if diceSlice[0] != diceSlice[3] && diceSlice[1] != diceSlice[4] {
			return
		}
		result = 4 * diceSlice[1]

	case fullHouse:
		sort.Ints(diceSlice)

		if !(diceSlice[0] == diceSlice[1] && diceSlice[2] == diceSlice[4] && diceSlice[0] != diceSlice[4]) && !(diceSlice[0] == diceSlice[2] && diceSlice[3] == diceSlice[4] && diceSlice[0] != diceSlice[4]) {
			return
		}

		result = sum(diceSlice)

	case yacht:
		val := diceSlice[0]
		for _, dice := range diceSlice[1:] {
			if dice != val {
				return
			}
		}
		result = 50

	case littleStraight:
		sort.Ints(diceSlice)
		if !checkStraight(diceSlice, 1) {
			return
		}
		result = 30

	case bigStraight:
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
