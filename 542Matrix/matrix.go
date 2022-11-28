package matrix

func updateMatrix(mat [][]int) [][]int {
	nRows := len(mat)
	nCols := len(mat[0])
	dist := make([][]int, nRows)
	for i := 0; i < nRows; i++ {
		dist[i] = make([]int, nCols)
	}
	vis := make([][]bool, nRows)
	for i := 0; i < nRows; i++ {
		vis[i] = make([]bool, nCols)
	}

	for i := 0; i < nRows; i++ {
		for j := 0; j < nRows; j++ {
			dist[i][j] = bfs(mat, dist, vis, i, j, nRows, nCols)
		}
	}

	return dist
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func bfs(mat, dist [][]int, vis [][]bool, i, j, nRows, nCols int) int {
	if i < 0 || i >= nRows || j < 0 || j >= nCols {
		return nRows * nCols
	}
	if vis[i][j] {
		return dist[i][j]
	}
	if mat[i][j] == 0 {
		return 0
	} else {
		vis[i][j] = true
		dist[i][j] = 1 + min(
			min(
				bfs(mat, dist, vis, i-1, j, nRows, nCols),
				bfs(mat, dist, vis, i+1, j, nRows, nCols)),
			min(
				bfs(mat, dist, vis, i, j-1, nRows, nCols),
				bfs(mat, dist, vis, i, j+1, nRows, nCols)),
		)
		return dist[i][j]
	}
}
