package change

import (
	"errors"
)

type state struct {
	coins          []int
	remainingValue int
	change         []int
}

// Change returns the change for given coins and initialValue
func Change(coins []int, initialValue int) ([]int, error) {
	result := []int{}
	if initialValue == 0 {
		return result, nil
	}
	queue := []state{{coins, initialValue, []int{}}}
	var s state

	for len(queue) > 0 {
		s, queue = queue[0], queue[1:]

		if s.remainingValue == 0 {
			if len(result) == 0 || len(result) > len(s.change) {
				result = s.change
			}
			continue
		}
		if len(s.coins) <= 0 || s.remainingValue <= 0 || (len(result) > 0 && len(s.change) >= len(result)) {
			continue
		}

		lastCoin := s.coins[len(s.coins)-1]

		if lastCoin <= s.remainingValue {
			queue = append(queue, state{s.coins, s.remainingValue - lastCoin, append([]int{lastCoin}, s.change...)})
		}

		queue = append(queue, state{s.coins[:len(s.coins)-1], s.remainingValue, s.change})
	}

	if len(result) == 0 {
		return []int{}, errors.New("can not give change")
	}

	return result, nil
}
