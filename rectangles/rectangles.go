package rectangles

func Count(grid []string) int {
	total := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			for h := i + 1; h < len(grid); h++ {
				for w := j + 1; w < len(grid[i]); w++ {
					if CheckRectangle(grid, i, j, h, w) {
						total++
					}
				}
			}
		}
	}
	return total
}

func CheckRectangle(grid []string, i, j, h, w int) bool {
	for x := i; x <= h; x++ {
		if (x == i || x == h) && (grid[x][j] != '+' || grid[x][w] != '+') {
			return false
		} else if !(grid[x][j] == '|' || grid[x][j] == '+') ||
			!(grid[x][w] == '|' || grid[x][w] == '+') {
			return false
		}
	}
	for y := j; y <= w; y++ {
		if (y == j || y == w) && (grid[i][y] != '+' || grid[h][y] != '+') {
			return false
		} else if !(grid[i][y] == '-' || grid[i][y] == '+') ||
			!(grid[h][y] == '-' || grid[h][y] == '+') {
			return false
		}
	}
	return true
}
