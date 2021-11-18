package binarysearch

func SearchInts(list []int, key int) int {
	if len(list) == 0 {
		return -1
	}

	startIdx := 0

	if list[startIdx] == key {
		return startIdx
	}

	endIdx := len(list) - 1

	for startIdx <= endIdx {
		if list[startIdx] > key {
			return -1
		}

		middleIdx := (startIdx + endIdx) / 2
		middle := list[middleIdx]

		if middle == key {
			return middleIdx
		}

		if key > middle {
			startIdx = middleIdx + 1
		}

		if key < middle {
			endIdx = middleIdx - 1
		}

		if startIdx >= len(list) {
			return -1
		}

		if list[startIdx] == key {
			return startIdx
		}
	}

	return -1
}
