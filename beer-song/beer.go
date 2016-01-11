package beer

import "fmt"

/*Verse sings a verse of "Bottles of Beer on the Wall".*/
func Verse(v int) (string, error) {
	if v < 0 || 99 < v {
		return "", fmt.Errorf("Verse must be between 0 and 99: %d", v)
	}
	switch v {
	case 0:
		return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n", nil
	case 1:
		return "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", nil
	case 2:
		return "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n", nil
	default:
		return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", v, v, v-1), nil
	}
}

/*Verses sings many verses of "Bottles of Beer on the Wall".*/
func Verses(start, stop int) (string, error) {
	verses := ""
	for i := start; stop <= i; i-- {
		if verse, err := Verse(i); err == nil {
			verses += verse + "\n"
		} else {
			return "", err
		}
	}
	if verses == "" {
		return "", fmt.Errorf("Empty verses from %d to %d", start, stop)
	}
	return verses, nil
}

/*Song sings the entire "Bottles of Beer on the Wall" song.*/
func Song() string {
	if song, ok := Verses(99, 0); ok == nil {
		return song
	}
	return ""
}
