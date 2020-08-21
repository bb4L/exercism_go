package secret

// Relation defines the mapping of number to action
type Relation struct {
	number uint
	action string
}

var mapping = []Relation{Relation{1, "wink"}, Relation{2, "double blink"}, Relation{4, "close your eyes"}, Relation{8, "jump"}}

// Handshake returns the handshake for a given code
func Handshake(code uint) (result []string) {
	for _, rel := range mapping {
		if code&rel.number == rel.number {
			result = append(result, rel.action)
		}
	}

	if code&16 == 16 {
		for left, right := 0, len(result)-1; left < right; left, right = left+1, right-1 {
			result[left], result[right] = result[right], result[left]
		}
	}

	return result

}
