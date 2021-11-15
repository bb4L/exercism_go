package railfence

import (
	"strings"
	"unicode/utf8"
)

func Encode(message string, rails int) string {
	var builders = make([]string, rails)
	line := 0
	change := 1
	for i := 0; i < utf8.RuneCountInString(message); i++ {
		builders[line] = builders[line] + string(message[i])
		switch line {
		case rails - 1:
			change = -1
		case 0:
			change = 1
		}
		line += change
	}

	return strings.Join(builders, "")
}

func Decode(message string, rails int) string {
	line := 0
	change := 1
	var counts = make([]int, rails)

	for i := 0; i < len(message); i++ {
		counts[line] = counts[line] + 1
		if line == rails-1 {
			change = -1
		} else if line == 0 {
			change = 1
		}
		line += change
	}

	var storedLines = []string{}
	for i := 0; i < rails; i++ {
		storedLines = append(storedLines, "")
	}

	start := 0

	for i := 0; i < rails; i++ {
		if i < rails-1 {
			storedLines[i] = message[start : start+counts[i]+1]
		} else {
			storedLines[i] = message[start:]
		}
		start = start + counts[i]
	}

	result := ""
	line = 0
	change = 1

	for i := 0; i < len(message); i++ {
		if len(storedLines[line]) > 0 {
			result += string(storedLines[line][0])
			if len(storedLines[line]) > 1 {
				storedLines[line] = storedLines[line][1:]
			} else {
				storedLines[line] = ""
			}
		}

		if line == rails-1 {
			change = -1
		} else if line == 0 {
			change = 1
		}
		line += change
	}

	return result
}
