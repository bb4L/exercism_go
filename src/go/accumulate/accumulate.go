package accumulate

// Accumulate the given function on the input
func Accumulate(input []string, action func(string) string) []string {
	result := make([]string, len(input))
	for i, v := range input {
		result[i] = action(v)
	}
	return result
}
