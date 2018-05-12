package acronym

import "strings"

func Abbreviate(s string) string {
	acronym := ""
	for _, word := range SplitAny(s, []string{" ", "-"}) {
		acronym += strings.ToUpper(string(word[0]))
	}
	return acronym
}

func SplitAny(s string, seps []string) []string {
	var result []string
    word := ""
	for _, char := range s {
        add := true
        for _, sep := range seps {
            if sep == string(char) {
                if word != "" {
                    result = append(result, word)
                    word = ""
                    add = false
                    break
                }
            }
        }
        if add {
            word += string(char)
        }
	}
    if word != "" {
        result = append(result, word)
    }
	return result
}
