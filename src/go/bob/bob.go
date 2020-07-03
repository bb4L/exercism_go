// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"

// Hey returns minimal answers
func Hey(remark string) string {
	remarkTrimmed := strings.Trim(remark, " \t\n\r")

	if remarkTrimmed == "" {
		return "Fine. Be that way!"
	}

	isQuestion := len(remarkTrimmed) > 0 && remark[len(remarkTrimmed)-1] == '?'
	isShouted := remarkTrimmed == strings.ToUpper(remarkTrimmed) && strings.ContainsAny(strings.ToLower(remarkTrimmed), "abcdefghijklmnopqrtsufvwxy")

	if isQuestion {
		if isShouted {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	}

	if isShouted {
		return "Whoa, chill out!"
	}
	return "Whatever."

}
