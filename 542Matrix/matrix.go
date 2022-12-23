package matrix

import queue "algogo/Queue"

type Pair struct {
	i, j int
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func updateMatrix(mat [][]int) [][]int {
	nRows, nCols := len(mat), len(mat[0])
	res := make([][]int, nRows)
	for i := 0; i < nRows; i++ {
		res[i] = make([]int, nCols)
	}
	vis := make([][]bool, nRows)
	for k := 0; k < nRows; k++ {
		vis[k] = make([]bool, nCols)
	}

	for i := 0; i < nRows; i++ {
		for j := 0; j < nCols; j++ {
			if mat[i][j] == 0 {
				res[i][j] = 0
			}
			if mat[i][j] == 1 {
				bfs(mat, res, vis, i, j, nRows, nCols)
			}
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func bfs(mat, res [][]int, vis [][]bool, i, j, nRows, nCols int) {
	Q := queue.Queue{}

	Q.PushBack(Pair{i, j})
	dist := 0
	for i := 0; i < nRows; i++ {
		for j := 0; j < nCols; j++ {
			vis[i][j] = false
		}
	}
label:
	for !Q.Empty() {
		for lenQ := Q.Size(); lenQ > 0; lenQ-- {
			r := Q.Top().(Pair)
			Q.PopFront()
			ii, jj := r.i, r.j
			vis[i][j] = true
			for _, item := range []Pair{{ii + 1, jj}, {ii - 1, jj}, {ii, jj + 1}, {ii, jj - 1}} {
				x, y := item.i, item.j
				if x >= 0 && x < nRows && y >= 0 && y < nCols {
					if mat[x][y] != 0 && !vis[x][y] {
						Q.PushBack(Pair{x, y})
					} else if !vis[x][y] {
						dist++
						break label
					}
				}
			}
		}
		dist++
	}
	res[i][j] = dist
}
