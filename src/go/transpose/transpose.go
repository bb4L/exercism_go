package transpose

import "strings"

// Transpose transposes the given data
func Transpose(data []string) []string {
	rows := []*strings.Builder{}

	result := []string{}

	maxLine := 0
	for i, row := range data {
		if len(data) > maxLine {
			maxLine = len(data)
		}

		for k := 0; k < maxLine; k++ {
			var builder *strings.Builder

			if k < len(rows) {
				builder = rows[k]
			} else {
				builder = new(strings.Builder)
				builder.WriteString(strings.Repeat(" ", i))
			}

			if k < len(row) {
				builder.WriteByte(row[k])
			} else {
				builder.WriteString(" ")
			}

			if k >= len(rows) {
				rows = append(rows, builder)
			}
		}
	}

	for _, r := range rows {
		result = append(result, r.String())
	}

	return result
}
