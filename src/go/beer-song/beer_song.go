package beer

import (
	"errors"
	"fmt"
	"strings"
)

// Song returns the full song
func Song() string {
	res, _ := Verses(99, 0)
	return res
}

// Verses gives the verses from start to stop
func Verses(start, stop int) (string, error) {
	if start < stop {
		return "", errors.New("start needs to be higher than stop")
	}

	result := ""
	for i := start; i >= stop; i-- {
		newVerse, err := Verse(i)
		if err != nil {
			return "", err
		}
		result += newVerse + "\n"
	}

	return result, nil
}

var beerStrings = []string{
	"No more bottles",
	"%d bottle",
	"%d bottles",
	"%d bottles",
}

var secondStrings = []string{
	"Go to the store and buy some more, 99 bottles of beer on the wall.",
	"Take it down and pass it around, no more bottles of beer on the wall.",
	"Take one down and pass it around, %d bottle of beer on the wall.",
	"Take one down and pass it around, %d bottles of beer on the wall.",
}

// Verse returns one verse
func Verse(n int) (string, error) {
	if n > 100 || n < 0 {
		return "", errors.New("invalid verse number")
	}
	idx := n
	if n > 3 {
		idx = 3
	}

	firstString := strings.Split(fmt.Sprintf(beerStrings[idx], n), "%")[0]
	secondString := strings.Split(fmt.Sprintf(secondStrings[idx], n-1), "%")[0]

	return fmt.Sprintf("%s of beer on the wall, %s of beer.", firstString, strings.ToLower(firstString)) + "\n" + secondString + "\n", nil
}
