package series

func All(n int, s string) (result []string) {
	if n > len(s) {
		return
	}
	for i := 0; i <= len(s)-n; i++ {
		result = append(result, s[i:n+i])
	}

	return
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}

func First(n int, s string) (string, bool) {
	if n > len(s) {
		return "", false
	}

	return s[:n], true
}
