package binarysearch

func SearchInts(list []int, key int) int {
	if len(list) == 0 {
		return -1
	}

	idx := 0

	if list[idx] == key {
		return idx
	}

	searchRange := len(list) - 1

	for searchRange > 0 {
		if list[idx] > key {
			return -1
		}

		middle := list[idx+(searchRange/2)]
		if middle == key {
			return idx + searchRange/2
		}

		if key > middle {
			idx += searchRange / 2
			if searchRange%2 == 1 && searchRange > 1 {
				idx++
				searchRange++
			}
		}

		searchRange /= 2

		if idx >= len(list) {
			return -1
		}
		if list[idx] == key {
			return idx
		}
	}

	return -1
}
