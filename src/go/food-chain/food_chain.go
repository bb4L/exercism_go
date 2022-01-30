package foodchain

import "strings"

var animals = []string{"fly", "spider", "bird", "cat", "dog", "goat", "cow", "horse"}
var description = []string{"", "It wriggled and jiggled and tickled inside her.",
	"How absurd to swallow a bird!", "Imagine that, to swallow a cat!", "What a hog, to swallow a dog!",
	"Just opened her throat and swallowed a goat!", "I don't know how she swallowed a cow!", ""}

// Verse returns the verse
func Verse(v int) string {
	var builder strings.Builder
	internalVerse(&builder, v)
	return builder.String()
}

func internalVerse(builder *strings.Builder, v int) {
	defer builder.WriteString(getEnding(v))
	builder.WriteString("I know an old lady who swallowed a ")
	builder.WriteString(animals[v-1])
	builder.WriteString(".\n")

	if v <= 1 || v >= 8 {
		return
	}

	if description[v-1] != "" {
		builder.WriteString(description[v-1])
		builder.WriteString("\n")
	}

	for i := v - 1; i > 0; i-- {
		builder.WriteString("She swallowed the ")
		builder.WriteString(animals[i])
		builder.WriteString(" to catch the ")
		builder.WriteString(animals[i-1])
		if i == 2 && v > 2 {
			builder.WriteString(" that wriggled and jiggled and tickled inside her")
		}
		builder.WriteString(".\n")
	}

}

func getEnding(v int) string {
	if v == 8 {
		return "She's dead, of course!"
	}
	return "I don't know why she swallowed the fly. Perhaps she'll die."
}

// Verses returns a range of verses
func Verses(start, end int) string {
	var builder strings.Builder
	for i := start; i <= end; i++ {
		internalVerse(&builder, i)
		if i < end {
			builder.WriteString("\n\n")
		}
	}
	return builder.String()
}

// Song gives the whole song
func Song() string {
	return Verses(1, 8)
}
