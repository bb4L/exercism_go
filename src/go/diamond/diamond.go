package diamond

import (
	"errors"
	"strings"
)

// Gen generates a diamond
func Gen(b byte) (string, error) {
	if b < byte('A') || b > byte('Z') {
		return "", errors.New("invalid character")
	}

	difference := int(b - 'A')

	if difference == 0 {
		return "A\n", nil
	}
	result := []string{getLine(difference, difference)}

	for i := difference - 1; i >= 0; i-- {
		line := getLine(difference, i)
		result = append(append([]string{line}, result...), line)
	}

	return strings.Join(result, "\n") + "\n", nil
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
