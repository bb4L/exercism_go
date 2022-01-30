package house

import "strings"

var steps = []string{"house that Jack built", "malt", "rat", "cat", "dog", "cow with the crumpled horn",
	"maiden all forlorn", "man all tattered and torn", "priest all shaven and shorn",
	"rooster that crowed in the morn", "farmer sowing his corn", "horse and the hound and the horn"}

var addendum = []string{"lay in", "ate", "killed", "worried", "tossed", "milked", "kissed",
	"married", "woke", "kept", "belonged to"}

// Verse return a verse
func Verse(v int) string {
	var sb strings.Builder
	sb.WriteString("This is the ")
	sb.WriteString(steps[v-1])
	if v > 1 {
		for i := v - 1; i > 0; i-- {
			sb.WriteString("\nthat ")
			sb.WriteString(addendum[i-1])
			sb.WriteString(" the ")
			sb.WriteString(steps[i-1])

		}
	}

	sb.WriteString(".")
	return sb.String()
}

// Song gives the whole song
func Song() string {
	var sb strings.Builder
	for i := 1; i < 13; i++ {
		sb.WriteString(Verse(i))
		if i < 12 {
			sb.WriteString("\n\n")
		}
	}
	return sb.String()
}
