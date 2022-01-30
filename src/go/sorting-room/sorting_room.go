package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

// NumberBox a box holding a number
type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %d.0", nb.Number())
}

// FancyNumber reprsent number as sting
type FancyNumber struct {
	n string
}

// Value returns the value of the FancyNumber
func (i FancyNumber) Value() string {
	return i.n
}

// FancyNumberBox box for a FancyNumber
type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	val, err := strconv.Atoi(fnb.Value())
	if err != nil {
		return 0
	}
	return val
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	return fmt.Sprintf("This is a fancy box containing the number %d.0", ExtractFancyNumber(fnb))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch i := i.(type) {
	case int:
		return DescribeNumber(float64(i))
	case float64:
		return DescribeNumber(i)
	case NumberBox:
		return DescribeNumberBox(i)
	case FancyNumberBox:
		return DescribeFancyNumberBox(i)
	default:
		return "Return to sender"
	}
}
