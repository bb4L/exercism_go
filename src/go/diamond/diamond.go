package diamond

import (
	"errors"
	"strings"
)

func Gen(b byte) (string, error) {
	if b < byte('A') || b > byte('Z') {
		return "", errors.New("invalid character")
	}

	difference := int(b - 'A')

	if difference == 0 {
		return "A\n", nil
	}
	reversedResult := []string{}
	result := []string{}

	for i := 0; i < difference; i++ {
		result = append(result, getLine(difference, i))
		reversedResult = append([]string{getLine(difference, i)}, reversedResult...)
	}

	result = append(result, getLine(difference, difference))

	return strings.Join(result, "\n") + "\n" + strings.Join(reversedResult, "\n") + "\n", nil
}

func getLine(totalDifference int, iteration int) string {
	filler := " "
	currentChar := string(byte('A' + iteration))
	startEndSpace := strings.Repeat(filler, totalDifference-iteration)

	if iteration == 0 {
		startEndSpace = strings.Repeat(filler, totalDifference-iteration/2-(iteration%2))
		return startEndSpace + currentChar + startEndSpace
	}

	return startEndSpace + currentChar + strings.Repeat(filler, iteration*2-1) + currentChar + startEndSpace

}
