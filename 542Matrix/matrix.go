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
			bfs(mat, dist, vis, 0, i, j, nRows, nCols)
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

func bfs(mat, dist [][]int, vis [][]bool, prevWeight, i, j, nRows, nCols int) {

}
