package maxareaofisland

func maxAreaOfIsland(grid [][]int) int {
	nRows, nCols := len(grid), len(grid[0])

	vis := make([][]bool, nRows)
	for i := 0; i < nRows; i++ {
		vis[i] = make([]bool, nCols)
	}

	maxArea := 0

	for i := 0; i < nRows; i++ {
		for j := 0; j < nCols; j++ {
			if grid[i][j] == 1 && !vis[i][j] {
				area := f(grid, vis, i, j, nRows, nCols)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}

func f(grid [][]int, vis [][]bool, i, j int, nRows, nCols int) (sum int) {
	if i < 0 || i >= nRows || j < 0 || j >= nCols || vis[i][j] || grid[i][j] == 0 {
		return 0
	}

	vis[i][j] = true

	sum = 1 +
		f(grid, vis, i-1, j, nRows, nCols) +
		f(grid, vis, i+1, j, nRows, nCols) +
		f(grid, vis, i, j-1, nRows, nCols) +
		f(grid, vis, i, j+1, nRows, nCols)

	return sum
}
