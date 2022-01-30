package poker

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

const (
	high = iota
	pair
	twoPair
	threeOfAKind
	straight
	flush
	fullHouse
	fourOfAKind
	straightFlush
)

var colors = []string{"♢", "♡", "♧", "♤"}

var pictureCards = map[string]int{
	"J": 11, "Q": 12, "K": 13, "A": 14,
}

type hand struct {
	handScore    int
	highCard     int
	originalHand string
	numbers      map[int]int
	fourOfAKind  int
	threeOfAKind int
	pairs        []int
}

func createHand(rawHand string) (result *hand, err error) {
	result = &hand{}
	result.originalHand = rawHand

	cards := strings.Split(rawHand, " ")
	if len(cards) != 5 {
		return result, errors.New("invalid number of cards")
	}

	minNumber := 14
	maxNumber := 0
	maxMultipleNumbers := 0
	handColors := map[string]int{}
	numbers := map[int]int{}

	for _, card := range cards {
		r := []rune(card)

		length := len(r)

		if length < 2 || length > 3 {
			return nil, errors.New("invalid card")
		}

		rawNumber := string(r[0 : len(r)-1])
		color := string(r[len(r)-1])

		foundColor := false
		for _, v := range colors {
			if v == color {
				foundColor = true
				break
			}
		}

		if !foundColor {
			return nil, errors.New("invalid color")
		}

		handColors[color] = handColors[color] + 1

		number := -1
		n, ok := pictureCards[rawNumber]
		if ok {
			number = n
		} else {
			number, err = strconv.Atoi(rawNumber)
			if err != nil || number > 10 || number < 2 {
				return nil, errors.New("invalid card value for number")
			}
		}

		if number > maxNumber {
			maxNumber = number
		}
		if number < minNumber {
			minNumber = number
		}

		numbers[number] = numbers[number] + 1

		if numbers[number] > maxMultipleNumbers {
			maxMultipleNumbers = numbers[number]
		}
	}

	result.numbers = numbers
	result.highCard = -1
	result.handScore = -1

	isFlush := len(handColors) == 1
	isStraight := getIsStraight(numbers, minNumber, maxNumber)

	if isStraight {
		if maxNumber-minNumber == 4 {
			result.highCard = maxNumber
		} else {
			result.highCard = 4
		}
	}

	if isFlush && isStraight {
		result.handScore = straightFlush
	} else {
		if isFlush {
			result.handScore = flush
		}

		if isStraight {
			result.handScore = straight
		}
	}

	switch maxMultipleNumbers {
	case 4:
		result.handScore = fourOfAKind
		for k, v := range numbers {
			if v == 4 {
				result.fourOfAKind = k
			} else {
				result.highCard = k
			}
		}
	case 3:
		if len(numbers) == 2 {
			result.handScore = fullHouse
		} else {
			result.handScore = threeOfAKind
		}
		for k, v := range numbers {
			if v == 3 {
				result.highCard = k
				result.threeOfAKind = k
			}
			if v == 2 {
				result.pairs = append(result.pairs, k)
			}
		}
	case 2:
		if len(numbers) == 3 {
			result.handScore = twoPair
		} else {
			result.handScore = pair
		}

		higherPair := 0
		for k, v := range numbers {

			if v == 2 {
				if k > higherPair {
					higherPair = k
				}
				result.pairs = append(result.pairs, k)
			}
		}
		result.highCard = higherPair
	default:
		if !(isStraight || isFlush) {
			result.handScore = high
		}
		if result.highCard == -1 {
			for key := range numbers {
				if key > result.highCard {
					result.highCard = key
				}
			}
		}
	}

	return result, nil
}

func getIsStraight(numbers map[int]int, minNumber, maxNumber int) bool {
	if len(numbers) != 5 {
		return false
	}
	if maxNumber-minNumber == 4 {
		return true
	}

	if minNumber == 2 && maxNumber == 14 {
		for _, i := range []int{3, 4, 5} {
			if numbers[i] != 1 {
				return false
			}
		}
		return true
	}
	return false
}

// BestHand retruns the best hand
func BestHand(rawHands []string) ([]string, error) {
	var hands []*hand

	for _, h := range rawHands {
		hand, err := createHand(h)

		if err != nil {
			continue
		}
		hands = append(hands, hand)
	}

	if len(hands) == 0 {
		return nil, errors.New("no valid hand")
	}

	if len(hands) == 1 {
		return []string{hands[0].originalHand}, nil
	}

	sort.Slice(hands, func(i, j int) bool {
		return compareHands(hands[i], hands[j]) == 1
	})

	result := []string{hands[0].originalHand}

	for _, k := range hands[1:] {
		if compareHands(hands[0], k) == 0 {
			result = append(result, k.originalHand)
		} else {
			break
		}
	}
	return result, nil
}

func compareHands(h1, h2 *hand) int {
	handComparison := comparingNumbers([]int{h1.handScore}, []int{h2.handScore})

	if handComparison == 0 {

		if h1.handScore == fullHouse {
			return comparingNumbers([]int{h1.threeOfAKind, h1.pairs[0]}, []int{h2.threeOfAKind, h2.pairs[0]})
		}

		if h1.handScore == fourOfAKind {
			return comparingNumbers([]int{h1.fourOfAKind, h1.highCard}, []int{h2.fourOfAKind, h2.highCard})
		}

		highCardComparison := comparingNumbers([]int{h1.highCard}, []int{h2.highCard})
		if highCardComparison == 0 {
			k := 1
			for {
				highCard := h1.highCard - k
				if highCard < 2 {
					return 0
				}

				comparation := comparingNumbers([]int{h1.numbers[highCard]}, []int{h2.numbers[highCard]})
				if comparation != 0 {
					return comparation
				}
				k++
			}
		}
		return highCardComparison
	}

	return handComparison
}

func comparingNumbers(a, b []int) int {
	for i, k := range a {
		if k == b[i] {
			continue
		}

		if k > b[i] {
			return 1
		}
		return -1
	}
	return 0
}
