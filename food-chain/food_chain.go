package foodchain

import "strings"

// TestVersion is the unit test version that this program is built to pass.
const TestVersion = 1

/*Verse sings a verse of "I know an old lady who swallowed a fly".*/
func Verse(v int) string {
	verse := "I know an old lady who swallowed " + first[v]
	if 8 <= v {
		return verse
	}
	for i := v; 0 <= i; i-- {
		verse += refrain[i]
	}
	return verse
}

/*Verses sings a set of verses of "I know an old lady who swallowed a fly".*/
func Verses(start, stop int) string {
	var verses = []string{}
	for v := start; v <= stop; v++ {
		verses = append(verses, Verse(v))
	}
	return strings.Join(verses, "\n\n")
}

/*Song sings "I know an old lady who swallowed a fly" in it's entirty.*/
func Song() string {
	return Verses(1, 8)
}

// first is the first line for a verse of "I know an old lady who swallowed a fly".
var first = []string{
	"",
	"a fly.\n",
	"a spider.\nIt wriggled and jiggled and tickled inside her.\n",
	"a bird.\nHow absurd to swallow a bird!\n",
	"a cat.\nImagine that, to swallow a cat!\n",
	"a dog.\nWhat a hog, to swallow a dog!\n",
	"a goat.\nJust opened her throat and swallowed a goat!\n",
	"a cow.\nI don't know how she swallowed a cow!\n",
	"a horse.\nShe's dead, of course!",
}

// refrain in the repeating lines in "I know an old lady who swallowed a fly"
var refrain = []string{
	"",
	"I don't know why she swallowed the fly. Perhaps she'll die.",
	"She swallowed the spider to catch the fly.\n",
	"She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.\n",
	"She swallowed the cat to catch the bird.\n",
	"She swallowed the dog to catch the cat.\n",
	"She swallowed the goat to catch the dog.\n",
	"She swallowed the cow to catch the goat.\n",
}
