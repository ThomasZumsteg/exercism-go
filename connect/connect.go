package connect

//player plays the game of connect
type player struct {
	Name              string
	Char              byte
	GetStartPositions func([]string) []coord
	AtWinningPosition func([]string, coord) bool
}

//coord is a position on a connectBoard
type coord struct {
	Row, Col int
}

//connectBoard is board for the game of connect
type connectBoard []string

/*ResultOf determines which player has won the game of conenct.*/
func ResultOf(board connectBoard) (string, error) {
	black := player{"black", 'X', getFirstCol, atLastCol}
	white := player{"white", 'O', getFirstRow, atLastRow}
	for _, p := range []player{black, white} {
		var pos coord
		seenPos := make(map[coord]bool)
		for queue := p.GetStartPositions(board); len(queue) != 0; {
			pos, queue = queue[0], queue[1:]
			seen, ok := seenPos[pos]
			if (ok && seen) || board[pos.Row][pos.Col] != p.Char {
				continue
			} else if p.AtWinningPosition(board, pos) {
				return p.Name, nil
			}
			seenPos[pos] = true
			queue = append(queue, board.getAdjacent(pos)...)
		}
	}
	return "", nil
}

/*getAdjacent gets all valid adjacent positions.*/
func (b connectBoard) getAdjacent(pos coord) []coord {
	newPos := []coord{
		{pos.Row - 1, pos.Col}, {pos.Row - 1, pos.Col + 1},
		{pos.Row, pos.Col - 1}, {pos.Row, pos.Col + 1},
		{pos.Row + 1, pos.Col - 1}, {pos.Row + 1, pos.Col},
	}
	var adjacent []coord
	for _, p := range newPos {
		validRow := 0 <= p.Row && p.Row < len(b)
		validCol := validRow && 0 <= p.Col && p.Col < len(b)
		if validRow && validCol {
			adjacent = append(adjacent, p)
		}
	}
	return adjacent
}

/*getFirstCol gets all positions in the first column of the board.*/
func getFirstCol(board []string) []coord {
	firstCol := make([]coord, len(board))
	for r := range board {
		firstCol[r] = coord{r, 0}
	}
	return firstCol
}

/*getFirstRow gets all positions in the first row of the board.*/
func getFirstRow(board []string) []coord {
	firstRow := make([]coord, len(board[0]))
	for c := range board[0] {
		firstRow[c] = coord{0, c}
	}
	return firstRow
}

/*atLastCol if a position is at the last column of a board.*/
func atLastCol(board []string, pos coord) bool {
	return pos.Col == len(board[pos.Row])-1
}

/*atLastRow if a position is at the last row of a board.*/
func atLastRow(board []string, pos coord) bool {
	return pos.Row == len(board)-1
}
