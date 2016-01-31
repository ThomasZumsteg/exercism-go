package wordsearch

//TestVersion is the verion of the unit test that this will pass
const TestVersion = 1

//slice is a slice of the puzzle and it's starting and stopping position
type slice struct {
	word string
	pos  [2][2]int
}

/*Solve searches the puzzle for instances of words.*/
func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	slices := make(chan slice)
	matches := make(map[string][2][2]int)
	go makeSlices(puzzle, slices)
	for wordSlice := range slices {
		if contains(wordSlice.word, words) {
			matches[wordSlice.word] = wordSlice.pos
		}
	}
	return matches, nil
}

/*contains check the array contains an item.*/
func contains(match string, matches []string) bool {
	for _, m := range matches {
		if m == match {
			return true
		}
	}
	return false
}

/*makeSlices generates all word slices in the puzzle.*/
func makeSlices(puzzle []string, slices chan slice) {
	var sliceList [3]slice
	for rStart, row := range puzzle {
		for cStart := range row {
			start := [2]int{cStart, rStart}
			for l := 2; l+rStart <= len(row) || l+cStart < len(puzzle); l++ {
				rowWord := getRow(puzzle, rStart, cStart, l)
				rowStop := [2]int{cStart + l - 1, rStart}
				sliceList[0] = slice{rowWord, [2][2]int{start, rowStop}}

				colWord := getCol(puzzle, rStart, cStart, l)
				colStop := [2]int{cStart, rStart + l - 1}
				sliceList[1] = slice{colWord, [2][2]int{start, colStop}}

				diaWord := getDia(puzzle, rStart, cStart, l)
				diaStop := [2]int{cStart + l - 1, rStart + l - 1}
				sliceList[2] = slice{diaWord, [2][2]int{start, diaStop}}
				for _, s := range sliceList {
					if s.word != "" {
						slices <- s
						slices <- slice{reverse(s.word), [2][2]int{s.pos[1], s.pos[0]}}
					}
				}
			}
		}
	}
	close(slices)
}

/*getDia gets a diagonal slice.*/
func getDia(puzzle []string, row, col, length int) string {
	if len(puzzle) <= row+length || len(puzzle[row+length]) <= col+length {
		return ""
	}
	var runes []byte
	for i := 0; i < length; i++ {
		runes = append(runes, puzzle[row+i][col+i])
	}
	return string(runes)
}

/*getRow gets a row slice.*/
func getRow(puzzle []string, row, col, length int) string {
	if len(puzzle[row]) < col+length {
		return ""
	}
	return puzzle[row][col : col+length]
}

/*getCol gets a column slice.*/
func getCol(puzzle []string, row, col, length int) string {
	if len(puzzle) < row+length {
		return ""
	}
	var runes []byte
	for i := row; i < row+length; i++ {
		runes = append(runes, puzzle[i][col])
	}
	return string(runes)
}

/*reverse reverses a string.*/
func reverse(word string) string {
	runes := []rune(word)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
