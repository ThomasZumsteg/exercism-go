package matrix

//Pair is a row and column in a matrix
type Pair struct {
	row, col int
}

/*Saddle finds saddle points in a matrix.*/
func (m *Matrix) Saddle() []Pair {
	var saddlePoints []Pair
	cols := m.Cols()
	for r, row := range m.Rows() {
		for c, elem := range row {
			smallestInRow := compare(row, min) == elem
			largestInCol := compare(cols[c], max) == elem
			if smallestInRow && largestInCol {
				saddlePoints = append(saddlePoints, Pair{r, c})
			}
		}
	}
	return saddlePoints
}

/*compare finds the element in an array that best fits some function.*/
func compare(items []int, comp func(int, int) bool) int {
	best := items[0]
	for _, item := range items {
		if comp(best, item) {
			best = item
		}
	}
	return best
}

/*min compares two elemens for smallness.*/
func min(a, b int) bool {
	return a < b
}

/*max compares two elements for bigness.*/
func max(a, b int) bool {
	return a > b
}
