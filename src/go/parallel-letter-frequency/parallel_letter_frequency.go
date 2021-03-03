package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency calculates the frequencies concurrent
func ConcurrentFrequency(strings []string) (result FreqMap) {
	results := make(chan FreqMap)

	for _, s := range strings {
		go func(s string) {
			results <- Frequency(s)
		}(s)
	}

	result = <-results

	for i := 1; i < len(strings); i++ {
		for letter, freq := range <-results {
			result[letter] += freq
		}
	}

	return result
}
