package minesweeper

import (
	"bytes"
	"fmt"
)

//Board is a board for minesweeper.
type Board [][]byte

/*String converts the minesweeper board to a string.*/
func (b Board) String() string {
	return "\n" + string(bytes.Join(b, []byte{'\n'}))
}

/*Count validates the board and fills in the mine counts.*/
func (b Board) Count() error {
	for r, row := range b {
		for c := range row {
			if !b.validSquare(r, c) {
				return fmt.Errorf("Invalid character at (%d, %d)", r, c)
			}
			b.countMines(r, c)
		}
	}
	return nil
}

/*countMines counts adjacent mines and fills in the result if the element is blank.*/
func (b Board) countMines(row, col int) {
	elem := b[row][col]
	if elem != ' ' {
		return
	}
	mines := byte('0')
	for r := row - 1; r <= row+1 && r < len(b); r++ {
		for c := col - 1; c <= col+1 && c < len(b[r]); c++ {
			if b[r][c] == '*' {
				mines++
			}
		}
	}
	if mines != byte('0') {
		b[row][col] = mines
	}
}

/*validSquare checks if a given location is valid and has a valid character.*/
func (b Board) validSquare(row, col int) bool {
	if row < 0 || len(b) <= row || col < 0 || len(b[row]) <= col {
		return false
	}
	height, width := len(b), len(b[0])
	onTopOrBottom := row == 0 || row == height-1
	onLeftOrRight := col == 0 || col == width-1
	switch b[row][col] {
	case '+':
		return onTopOrBottom && onLeftOrRight
	case '-':
		return onTopOrBottom && !onLeftOrRight
	case '|':
		return !onTopOrBottom && onLeftOrRight
	case '*', ' ':
		return 0 < row && row < height && 0 < col && col < width
	default:
		return false
	}
}
