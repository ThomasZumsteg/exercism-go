package bob

import (
	"strings"
	"unicode"
)

//TestVersion is the test suite that this will pass
const TestVersion = 1

/*Hey responds to a greeting like a lackadaisical teenager.*/
func Hey(greeting string) string {
	greeting = strings.TrimSpace(greeting)

	switch {
	case greeting == "":
		return "Fine. Be that way!"
	case any(greeting, unicode.IsUpper) && !any(greeting, unicode.IsLower):
		return "Whoa, chill out!"
	case greeting[len(greeting)-1] == '?':
		return "Sure."
	default:
		return "Whatever."
	}
}

/*any determines if any items in a string pass some test*/
func any(items string, test func(rune) bool) bool {
	for _, item := range items {
		if test(item) {
			return true
		}
	}
	return false
}
