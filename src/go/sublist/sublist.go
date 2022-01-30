package sublist

// Relation of two lists
type Relation string

// Sublist returns relation of list1 and list2
func Sublist(list1, list2 []int) Relation {
	if len(list1) == 0 && len(list2) > 0 {
		return "sublist"
	}
	if len(list2) == 0 && len(list1) > 0 {
		return "superlist"
	}

	if len(list1) == len(list2) {
		for i, k := range list1 {
			if k != list2[i] {
				return "unequal"
			}
		}
		return "equal"
	}

	reverse := false

	if len(list1) < len(list2) {
		reverse = true
		list1, list2 = list2, list1
	}

	for j, k := range list1 {
		if k == list2[0] && len(list2) <= len(list1)-j {
			isSub := true
			for i, n := range list2[1:] {
				if list1[j+i+1] != n {
					isSub = false
					break
				}
			}

			if isSub {
				if reverse {
					return "sublist"
				}
				return "superlist"
			}
		}

	}

	return "unequal"
}
