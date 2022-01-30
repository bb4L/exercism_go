package binarysearch

// SearchInts searches a int in the slice
func SearchInts(list []int, key int) int {
	startIdx := 0
	endIdx := len(list) - 1

	for startIdx <= endIdx {
		middleIdx := (startIdx + endIdx) / 2
		middle := list[middleIdx]

		if key > middle {
			startIdx = middleIdx + 1
		} else if key < middle {
			endIdx = middleIdx - 1
		} else {
			return middleIdx
		}
	}

	return -1
}
