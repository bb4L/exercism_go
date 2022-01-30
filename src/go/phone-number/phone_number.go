package phonenumber

import (
	"errors"
	"unicode"
)

// Number returns the phonenumber formatted or a error
func Number(phoneNumber string) (string, error) {
	digits := []rune{}
	for _, val := range phoneNumber {
		if unicode.IsDigit(val) {
			digits = append(digits, val)
		} else {
			if val == '(' || val == ')' || val == '-' || val == '.' || val == '+' || unicode.IsSpace(val) {
				continue
			}

			if unicode.IsLetter(val) {
				return "", errors.New("leters not permitted")
			}

			return "", errors.New("punctuations not permitted")
		}
	}

	if len(digits) < 10 {
		return "", errors.New("incorrect number of digits")
	}

	if len(digits) > 11 {
		return "", errors.New("more than 11 digits")
	}

	if len(digits) == 11 {
		if digits[0] != '1' {
			return "", errors.New("11 digits must start with 1")
		}
		digits = digits[1:]
	}

	areaCode := digits[0]
	if areaCode == '0' {
		return "", errors.New("area code cannot start with zero")
	}
	if areaCode == '1' {
		return "", errors.New("area code cannot start with one")
	}

	exchangeCode := digits[3]
	if exchangeCode == '0' {
		return "", errors.New("exchange code cannot start with zero")
	}
	if exchangeCode == '1' {
		return "", errors.New("exchange code cannot start with one")
	}
	return string(digits), nil
}

// AreaCode returns the areacode of a phone number
func AreaCode(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return phoneNumber[0:3], nil
}

// Format returns the formatted phone number
func Format(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return "(" + phoneNumber[0:3] + ") " + phoneNumber[3:6] + "-" + phoneNumber[6:10], nil
}
