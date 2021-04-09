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
			if (b == ')' && brackets[len(brackets)-1] == '(') || (b == ']' && brackets[len(brackets)-1] == '[') || (b == '}' && brackets[len(brackets)-1] == '{') {
				if len(brackets) == 1 {
					brackets = []rune{}
				} else {
					brackets = brackets[:len(brackets)-1]
				}
			} else {
				return false
			}
		}
	}

	return len(brackets) == 0
}
