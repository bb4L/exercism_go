package strain

// Ints is an integer array
type Ints []int

// Strings array of strings
type Strings []string

// Lists list of integer lists
type Lists [][]int

// Keep Ints
func (a Ints) Keep(f func(int) bool) Ints {
	if len(a) == 0 {
		return nil
	}
	result := Ints{}
	for _, val := range a {
		if !f(val) {
			continue
		}

		result = append(result, val)
	}
	return result
}

// Discard Ints
func (a Ints) Discard(f func(int) bool) Ints {
	return a.Keep(func(i int) bool { return !f(i) })
}

// Keep Strings
func (a Strings) Keep(f func(string) bool) Strings {
	result := Strings{}

	if len(a) == 0 {
		return nil
	}

	for _, val := range a {
		if !f(val) {
			continue
		}
		result = append(result, val)

	}
	return result
}

// Keep Lists
func (a Lists) Keep(f func([]int) bool) Lists {
	result := Lists{}

	if len(a) == 0 {
		return nil
	}

	for _, val := range a {
		if !f(val) {
			continue
		}
		result = append(result, val)

	}
	return result
}
