package isogram

import "unicode"

func IsIsogram(s string) bool {
	seen := map[rune]bool{}
	for _, char := range s {
		char = unicode.ToLower(char)
		if 'a' <= char && char <= 'z' {
			if _, ok := seen[char]; ok {
				return false
			}
			seen[char] = true
		}
	}
	return true
}
