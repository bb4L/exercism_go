package series

// All returns a list of all substrings of s with length n
func All(n int, s string) (result []string) {
	if n > len(s) {
		return
	}
	for i := 0; i <= len(s)-n; i++ {
		result = append(result, s[i:n+i])
	}

	return
}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	return s[:n]
}

// First is the safe variant of UnsafeFirst
func First(n int, s string) (string, bool) {
	if n > len(s) {
		return "", false
	}

	return s[:n], true
}
