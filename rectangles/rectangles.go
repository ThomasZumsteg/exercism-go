package rectangles

func Count(grid []string) int {
    total := 0
    for i, row := range grid {
        for j, char := range row {
            if char == '+' {
                for height := 1; height < len(grid) - i; height++ {
                    for width := 1; width < len(grid[i]) - j; width++ {
                        if CheckRectangle(grid, height, width) {
                            total++
                        }
                    }
                }
            }
        }
    }
    return total
}

func CheckRectangle(grid []string, height, width int) bool {
    for i := 0; i <= height; i++ {
        if (i == 0 || i == height) && (grid[i][width] != '+' || grid[i][0] != '+') {
            return false
        } else if !(grid[i][width] == '|' || grid[i][width] == '+') || !(grid[i][0] == '|' || grid[i][0] == '+') {
            return false
        }
    }
    for j := 0; j < width; j++ {
        if (j == 0 || j == height) && (grid[height][j] != '+' || grid[0][j] != '+') {
            return false
        } else if !(grid[width][j] == '-' || grid[width][j] == '+') || !(grid[0][j] == '-' || grid[0][j] == '+') {
            return false
        }
    }
    return true;
}
