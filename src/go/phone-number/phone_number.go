package phonenumber

import (
	"errors"
	"regexp"
	"strings"
)

func Number(phoneNumber string) (string, error) {
	if len(phoneNumber) != len(regexp.MustCompile("[a-zA-Z]").ReplaceAllString(phoneNumber, "")) {
		return "", errors.New("leters not permitted")
	}

	if len(strings.ReplaceAll(regexp.MustCompile("[0-9()-.]").ReplaceAllString(phoneNumber, ""), " ", "")) > 0 {
		return "", errors.New("punctuations not permitted")
	}

	phoneNumber = regexp.MustCompile("[^0-9]").ReplaceAllString(phoneNumber, "")
	numberLength := len(phoneNumber)
	if numberLength < 10 {
		return "", errors.New("incorrect number of digits")
	}
	if numberLength > 11 {
		return "", errors.New("more than 11 digits")
	}
	if numberLength == 11 {
		if string(phoneNumber[0]) != "1" {
			return "", errors.New("11 digits must start with 1")
		}
		phoneNumber = phoneNumber[1:]
	}

	areaCode := phoneNumber[0]
	if areaCode == '0' {
		return "", errors.New("area code cannot start with zero")
	}
	if areaCode == '1' {
		return "", errors.New("area code cannot start with one")
	}

	exchangeCode := phoneNumber[3]
	if exchangeCode == '0' {
		return "", errors.New("exchange code cannot start with zero")
	}
	if exchangeCode == '1' {
		return "", errors.New("exchange code cannot start with one")
	}
	return phoneNumber, nil
}

func AreaCode(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return string(phoneNumber[0:3]), nil
}

func Format(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return "(" + phoneNumber[0:3] + ") " + phoneNumber[3:6] + "-" + phoneNumber[6:10], nil
}
