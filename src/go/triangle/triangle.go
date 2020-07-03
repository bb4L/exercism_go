// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import (
	"math"
)

// Kind will be returned
type Kind string

const (
	// NaT =  not a triangle
	NaT = "NaT"
	// Equ = equilateral
	Equ = "Equ"
	// Iso = isosceles
	Iso = "Iso"
	// Sca = scalene
	Sca = "Sca"
)

// KindFromSides returns the value for the provided sizes
func KindFromSides(a, b, c float64) Kind {
	// check for positive lenghts
	if a <= 0 || b <= 0 || c <= 0 ||
		// check for Triangle Inequality
		a+b < c || a+c < b || b+c < a ||
		//
		math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) ||
		math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return NaT
	}

	iso1 := a == b
	iso2 := b == c
	iso3 := a == c

	if iso1 && iso2 {
		return Equ
	}

	if iso1 || iso2 || iso3 {
		return Iso
	}

	return Sca
}
