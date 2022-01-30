package queenattack

import (
	"errors"
)

// CanQueenAttack returns a bool if two queens can attack
func CanQueenAttack(q1, q2 string) (bool, error) {
	if q1 == q2 {
		return false, errors.New("same position")
	}
	if invalidValue(q1) || invalidValue(q2) {
		return false, errors.New("invalid position")
	}
	return q1[0] == q2[0] || q1[1] == q2[1] || (q1[0]-q2[0] == q1[1]-q2[1]) || (q1[0]-q2[0] == q2[1]-q1[1]), nil
}

func invalidValue(val string) bool {
	return len(val) != 2 || (val[0] < 'a' || val[0] > 'h') || (val[1] < '1' || val[1] > '8')
}
