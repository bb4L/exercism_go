package ocr

import (
	"strings"
)

var digits = map[string]string{
	" _ | ||_|   ": "0",
	"     |  |   ": "1",
	" _  _||_    ": "2",
	" _  _| _|   ": "3",
	"   |_|  |   ": "4",
	" _ |_  _|   ": "5",
	" _ |_ |_|   ": "6",
	" _   |  |   ": "7",
	" _ |_||_|   ": "8",
	" _ |_| _|   ": "9",
}

func recognizeDigit(data []string) string {
	val, ok := digits[strings.Join(data, "")]
	if !ok {
		return "?"
	}
	return val
}

// Recognize parses the ocr numbers to normal numbers
func Recognize(data string) []string {
	splitted := strings.Split(data, "\n")

	if (len(splitted)-1)%4 != 0 {
		return []string{"?"}
	}

	dataToCheck := splitted[1:]
	result := []string{}

	for line := 0; line <= len(splitted)-4; line += 4 {
		var lineResult strings.Builder
		for i := 0; i <= len(dataToCheck[line])-3; i += 3 {
			numberData := []string{}

			for j := 0; j <= 3; j++ {
				if len(dataToCheck[j+line]) < (i)+3 {
					result = append(result, "?")
					continue
				}
				numberData = append(numberData, dataToCheck[j+line][i:i+3])
			}

			lineResult.WriteString(recognizeDigit(numberData))
		}

		result = append(result, lineResult.String())
	}

	return result
}
