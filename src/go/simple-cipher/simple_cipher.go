package cipher

import (
	"strings"
	"unicode"
)

// Shift new version of cesar
type Shift []int

// NewShift create new shift instance
func NewShift(n int) Shift {
	if n <= -26 || n == 0 || n >= 26 {
		return nil
	}
	return Shift{n}
}

// NewCaesar create new cesar object
func NewCaesar() Shift {
	return Shift{3}
}

func (s Shift) shift(text string, direction int) string {
	var sb strings.Builder
	sb.Grow(len(text))
	sLen := len(s)
	i := 0
	for _, r1 := range text {
		r := r1
		if !unicode.IsLetter(r) {
			continue
		}
		if unicode.IsUpper(r) {
			r = unicode.ToLower(r)
		}
		newR := r + rune((s[i%sLen]%26)*direction)

		if newR > 'z' {
			newR = 'a' + (newR - ('z' + 1))
		} else if newR < 'a' {
			newR = 'z' - 'a' + newR + 1
		}

		i++
		sb.WriteRune(newR)
	}

	return sb.String()
}

// Encode the given text
func (s Shift) Encode(text string) string {
	return s.shift(text, 1)
}

// Decode the given text
func (s Shift) Decode(text string) string {
	return s.shift(text, -1)
}

// NewVigenere create new cesar object
func NewVigenere(key string) Shift {
	allA := true

	shifts := make(Shift, len(key))
	for i, c := range key {
		if c != 'a' {
			allA = false
		}
		if !unicode.IsLetter(c) || unicode.IsUpper(c) {
			return nil
		}
		shifts[i] = int(c - 'a')
	}
	if allA {
		return nil
	}
	return shifts
}
