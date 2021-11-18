package brackets

func Bracket(data string) bool {
	brackets := []rune{}

	for _, b := range data {
		switch b {
		case '(', '[', '{':
			brackets = append(brackets, b)
		case ')', ']', '}':
			if len(brackets) == 0 {
				return false
			}
			if areMatching(brackets[len(brackets)-1], b) {
				brackets = brackets[:len(brackets)-1]
			} else {
				return false
			}
		}
	}

	return len(brackets) == 0
}

func areMatching(a, b rune) bool {
	switch a {
	case '(':
		return b == ')'
	case '{':
		return b == '}'
	case '[':
		return b == ']'
	}
	return false
}
