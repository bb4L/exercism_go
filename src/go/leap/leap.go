// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

import "math"

// IsLeapYear returns true if the passed integer is a leap year
func IsLeapYear(year int) bool {
	return math.Mod(float64(year), 4) == 0 && ((math.Mod(float64(year), 100) == 0 && math.Mod(float64(year), 400) == 0) || math.Mod(float64(year), 100) != 0)
}
