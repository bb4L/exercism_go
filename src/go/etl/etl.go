package etl

import "strings"

type legacy map[int][]string

// NewShiny format for the data
type NewShiny map[string]int

// Transform the old format to the new shiny format
func Transform(legacyData legacy) NewShiny {
	result := NewShiny{}

	for k, v := range legacyData {
		for _, value := range v {
			result[strings.ToLower(value)] = k
		}
	}

	return result
}
