package igpay

import (
	"regexp"
	"strings"
)

/*PigLatin converts a sentance to pig latin.*/
func PigLatin(engl string) string {
	var pigWords []string
	words := regexp.MustCompile("\\s+").Split(engl, -1)
	for _, word := range words {
		pigWords = append(pigWords, pigLatinWord(word))
	}
	return strings.Join(pigWords, " ")
}

/*pigLatinWord converts a single word to piglating.*/
func pigLatinWord(engl string) string {
	matches := []string{
		"^()(y[^aeiou].*)", // Leading y
		"(.*?[^q])(u.*)",   // u without qu
		"(.*?)([aeoi].*)",  // Default
	}
	for _, match := range matches {
		parts := regexp.MustCompile(match).FindStringSubmatch(engl)
		if len(parts) == 3 && parts[1]+"ay" != engl {
			return parts[2] + parts[1] + "ay"
		}
	}
	return engl + "ay"
}
