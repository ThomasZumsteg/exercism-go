package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

//TestVersion is the unit tests that this program passes
const TestVersion = 1

/*Encode preforms square code encrytion on the alphanumeric parts of some text*/
func Encode(text string) string {
	plainText := cleanText(text)
	l := int(math.Ceil(math.Sqrt(float64(len(plainText)))))
	cipherText := make([]string, l)
	for i := 0; i < l; i++ {
		for j := i; j < len(plainText); j += l {
			cipherText[i] += string(plainText[j])
		}
	}
	return strings.Join(cipherText, " ")
}

/*cleanText returns only the alphanumeric parts of a string*/
func cleanText(text string) string {
	cleanText := ""
	for _, v := range strings.ToLower(text) {
		if unicode.IsLetter(v) || unicode.IsNumber(v) {
			cleanText += string(v)
		}
	}
	return cleanText
}
