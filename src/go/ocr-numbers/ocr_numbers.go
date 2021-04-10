package ocr

import (
	"strings"
)

func recognizeDigit(data []string) string {
	if len(data) != 4 {
		return "?"
	}

	switch data[0] {
	case " _ ":
		// can be 2,3,5,6,7,8,9,0
		switch data[1] {
		case " _|":
			// 2,3
			if data[2] == "|_ " {
				return "2"
			}
			if data[2] == " _|" {
				return "3"
			}
			return "?"

		case "|_ ":
			// 5,6
			if data[2] == " _|" {
				return "5"
			}

			if data[2] == "|_|" {
				return "6"
			}
			return "?"

		case "  |":
			// 7
			if data[2] != "  |" {
				return "?"
			}
			return "7"

		case "|_|":
			// 8,9
			if data[2] == "|_|" {
				return "8"
			}

			if data[2] == " _|" {
				return "9"
			}
			return "?"

		case "| |":
			// 0
			if data[2] != "|_|" {
				return "?"
			}
			return "0"

		default:
			return "?"
		}

	case "   ":
		// can be 1,4
		switch data[1] {
		case "  |":
			// 1
			if data[2] != "  |" {
				return "?"
			}
			return "1"

		case "|_|":
			//4
			if data[2] != "  |" {
				return "?"

			}
			return "4"

		default:
			return "?"
		}

	}

	return "0"
}

func Recognize(data string) []string {
	splitted := strings.Split(data, "\n")

	if (len(splitted)-1)%4 != 0 {
		return []string{"?"}
	}

	dataToCheck := splitted[1:]

	result := []string{}

	for line := 0; line <= len(splitted)-4; line += 4 {
		lineResult := ""
		for i := 0; i <= len(dataToCheck[line])-3; i += 3 {
			numberData := []string{}

			for j := 0; j <= 3; j++ {
				if len(dataToCheck[j+line]) < (i)+3 {
					result = append(result, "?")
					continue
				}
				numberData = append(numberData, dataToCheck[j+line][i:i+3])
			}

			lineResult += recognizeDigit(numberData)
		}

		result = append(result, lineResult)
	}

	return result
}
