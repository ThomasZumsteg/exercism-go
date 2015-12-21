package wordcount

import "unicode"

//TestVersion is the unit tests that this passes.
const TestVersion = 1

//Frequency tracks the occurances of strings.
type Frequency map[string]int

/*WordCount finds the frequency of words in a phrase
only letters and numbers are counted, punctuation is ignored*/
func WordCount(phrase string) Frequency {
	counter := make(Frequency)
	for _, word := range splitWords(phrase) {
		counter[word]++
	}
	return counter
}

/*splitWords divides a string into a sequence of words. Continuous
sequences of letters and digits are counted as words*/
func splitWords(phrase string) []string {
	var words []string
	word := ""
	for _, v := range phrase + " " {
		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			word += string(unicode.ToLower(v))
		} else if word != "" {
			words = append(words, word)
			word = ""
		}
	}
	return words
}
