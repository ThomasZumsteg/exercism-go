package anagram

import (
	"reflect"
	"sort"
	"strings"
)

//letters creates an array of characters in a string for sorting
// TODO go all in and create an equal interface
type letters []rune

func (s letters) Len() int           { return len(s) }
func (s letters) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s letters) Less(i, j int) bool { return s[i] < s[j] }

/*Detect finds any anagrams in a list of candidates, case insensitive*/
func Detect(subject string, candidates []string) []string {
	word := strings.ToLower(subject)
	wordLetters := letters(word)
	sort.Sort(wordLetters)
	var anagrams []string
	for _, candidate := range candidates {
		lowerCandidate := strings.ToLower(candidate)
		candidateLetters := letters(lowerCandidate)
		sort.Sort(candidateLetters)
		if reflect.DeepEqual(candidateLetters, wordLetters) &&
			word != lowerCandidate {
			anagrams = append(anagrams, lowerCandidate)
		}

	}
	return anagrams
}
