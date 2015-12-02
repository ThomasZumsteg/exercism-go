package pascal

/*Triangle generates pascals triangle for a number of rows.*/
func Triangle(rows int) [][]int {
	var triangle = make([][]int, rows)
	for r := 0; r < rows; r++ {
		triangle[r] = make([]int, r+1)
		for c := 0; c <= r; c++ {
			triangle[r][c] = cell(r, c)
		}
	}
	return triangle
}

/*cell calculates the value of a cell in pascals triangle
based on combinations
https://en.wikipedia.org/wiki/Pascal%27s_triangle#Combinations*/
func cell(n, k int) int {
	return factorial(n) / (factorial(k) * factorial(n-k))
}

// fact_cache is a cache to speed up the factorial computation
var factCache = make(map[int]int)

/*factorial computes n * (n-1) * (n-2) * ... * 2 * 1*/
func factorial(n int) int {
	fact, ok := factCache[n]
	if !ok {
		fact = 1
		for i := 1; i <= n; i++ {
			fact *= i
		}
		factCache[n] = fact
	}
	return fact
}
