package scrabble_score

import "strings"

/*Score determines the value of a scrabble word*/
func Score(word string) int {
	score := 0
	letterValues := makeMap()
	for _, l := range strings.ToUpper(word) {
		if v, ok := letterValues[l]; ok {
			score += v
		}
	}
	return score
}

/*makeMap builds a letter score map for a game of scrabble*/
func makeMap() map[rune]int {
	letterMap := map[rune]int{}
	for value, letters := range letterScores {
		for _, letter := range letters {
			letterMap[letter] = value
		}
	}
	return letterMap
}

//letterScores sets the value of letters
var letterScores = []string{
	"",
	"AEIOULNRST",
	"DG",
	"BCMP",
	"FHVWY",
	"K",
	"",
	"",
	"JX",
	"",
	"QZ",
}
