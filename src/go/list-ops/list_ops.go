package listops

// IntList own implementation of a list
type IntList []int

type binFunc func(x, y int) int
type unaryFunc func(x int) int
type predFunc func(x int) bool

// Foldl applying funciton from left to right
func (integers IntList) Foldl(f binFunc, initial int) int {
	acc := initial

	for _, val := range integers {
		acc = f(acc, val)
	}

	return acc
}

// Foldr reverse applying function
func (integers IntList) Foldr(f binFunc, initial int) int {
	acc := initial

	for i := integers.Length() - 1; i >= 0; i-- {
		acc = f(integers[i], acc)
	}

	return acc
}

// Append list to a existing list
func (integers IntList) Append(toAppend IntList) IntList {
	result := make(IntList, integers.Length()+toAppend.Length())

	i := 0
	for _, val := range integers {
		result[i] = val
		i++
	}
	for _, val := range toAppend {
		result[i] = val
		i++
	}

	return result
}

// Concat a list with a list of lists
func (integers IntList) Concat(toConcat []IntList) IntList {
	result := integers

	for _, list := range toConcat {
		result = result.Append(list)
	}
	return result
}

// Reverse a list
func (integers IntList) Reverse() IntList {
	result := IntList{}

	for i := integers.Length() - 1; i >= 0; i-- {
		result = result.Append(IntList{integers[i]})
	}

	return result
}

// Map apply a function to each element
func (integers IntList) Map(f unaryFunc) IntList {
	result := IntList{}

	for _, i := range integers {
		result = result.Append(IntList{f(i)})
	}

	return result
}

// Filter the elements with a given function
func (integers IntList) Filter(f predFunc) IntList {
	result := IntList{}

	for _, i := range integers {
		if f(i) {
			result = result.Append(IntList{i})
		}
	}

	return result
}

// Length get the length of a list
func (integers IntList) Length() int {
	length := 0
	for i := range integers {
		length = i + 1
	}
	return length
}
