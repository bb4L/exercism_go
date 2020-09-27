package twelve

import (
	"fmt"
	"strings"
)

var words = [12]string{"twelve Drummers Drumming", "eleven Pipers Piping", "ten Lords-a-Leaping", "nine Ladies Dancing", "eight Maids-a-Milking", "seven Swans-a-Swimming", "six Geese-a-Laying", "five Gold Rings", "four Calling Birds", "three French Hens", "two Turtle Doves", "and a Partridge in a Pear Tree"}
var numberish = [12]string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}

// Song returns the whole song
func Song() string {
	result := ""
	for i := 1; i < 13; i++ {
		result += Verse(i)
		if i < 12 {
			result += "\n"
		}
	}
	return result
}

// Verse returns the verse
func Verse(verseNumber int) string {
	var result = []string{}
	result = append(result, fmt.Sprintf("On the %s day of Christmas my true love gave to me:", numberish[verseNumber-1]))

	if verseNumber > 1 {
		for i := 0; i < verseNumber; i++ {

			if i < verseNumber-1 {
				result = append(result, words[12-verseNumber+i]+",")
			} else {
				result = append(result, words[12-verseNumber+i])
			}
		}
	} else {
		result = append(result, strings.Replace(words[11], "and ", "", -1))
	}

	return strings.Join(result[:], " ") + "."
}
