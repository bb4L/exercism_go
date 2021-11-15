package bookstore

import (
	"sort"
)

func Cost(books []int) int {
	booksByIndex := make([]int, 5)
	differentBooks := 0

	for _, val := range books {
		booksByIndex[val-1] += 1
	}

	for _, val := range booksByIndex {
		if val > 0 {
			differentBooks++
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(booksByIndex)))
	return getMinCosts(booksByIndex, differentBooks)
}

func getMinCosts(books []int, maxSet int) int {
	sum := 0
	differentBooks := 0

	for _, val := range books {
		sum += val
		if val > 0 {
			differentBooks++
		}
	}

	if sum == 0 {
		return 0
	}

	if differentBooks == 0 {
		return 0
	}

	if maxSet == 1 {
		return sum * bookPrice * 100
	}

	price := 0
	otherPrice := 0

	if differentBooks >= maxSet {
		newBooks := getRemoveAfterSet(books, maxSet)
		price = getSetPrice(maxSet)
		price += getMinCosts(newBooks, maxSet)
		otherPrice = getMinCosts(books, maxSet-1)
	} else {
		price += getMinCosts(books, differentBooks)
		otherPrice = price
	}

	if otherPrice > price {
		return price
	}
	return otherPrice
}

func getRemoveAfterSet(books []int, setMax int) []int {
	result := []int{}
	i := 0

	for _, val := range books {
		if i < setMax {
			result = append(result, val-1)
		} else {
			result = append(result, val)
		}
		i++
	}

	sort.Sort(sort.Reverse(sort.IntSlice(result)))
	return result

}

const bookPrice = 8

func getSetPrice(setSize int) int {
	reduction := 100

	switch setSize {
	case 2:
		reduction = 95
	case 3:
		reduction = 90
	case 4:
		reduction = 80
	case 5:
		reduction = 75
	}
	return setSize * reduction * bookPrice
}
