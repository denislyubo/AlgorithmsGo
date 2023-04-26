package rottingoranges

import (
	"algogo/Queue"
)

type pair struct {
	i, j int
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func orangesRotting(grid [][]int) int {
	nRows, nCols := len(grid), len(grid[0])
	Q := queue.Queue{}
	cnt := 0

	for i := 0; i < nRows; i++ {
		for j := 0; j < nCols; j++ {
			if grid[i][j] == 1 {
				cnt++
			}
			if grid[i][j] == 2 {
				Q.PushBack(pair{i, j})
			}
		}
	}
	res := 0

	for !Q.Empty() {
		for lenQ := Q.Size(); lenQ > 0; lenQ-- {
			r := Q.Top().(pair)
			Q.PopFront()
			i, j := r.i, r.j
			for _, item := range []pair{{i + 1, j}, {i - 1, j}, {i, j + 1}, {i, j - 1}} {
				x, y := item.i, item.j
				if x >= 0 && x < nRows && y >= 0 && y < nCols && grid[x][y] == 1 {
					grid[x][y] = 2
					cnt--
					Q.PushBack(pair{x, y})
				}
			}
		}
		res++
	}
	if cnt == 0 {
		return max(0, res-1)
	}

	return -1
}
