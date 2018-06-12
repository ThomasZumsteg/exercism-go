package spiralmatrix

func SpiralMatrix(depth int) [][]int {
	result := [][]int{}
	for row := 0; row < depth; row++ {
		result = append(result, make([]int, depth))
	}
	row, col := 0, 0
	dr, dc := 0, 1
	for i := 1; i <= depth*depth; i++ {
		result[row][col] = i
		if row+dr < 0 || depth <= row+dr || col+dc < 0 ||
			depth <= col+dc || result[row+dr][col+dc] != 0 {
			dr, dc = dc, -dr
		}
		row, col = row+dr, col+dc
	}
	return result
}
