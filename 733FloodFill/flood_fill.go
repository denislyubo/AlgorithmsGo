package floodfill

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	numRows := len(image)
	numCols := len(image[0])

	res := make([][]int, numRows)
	visited := make([][]bool, numRows)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, numCols)
		visited[i] = make([]bool, numCols)
	}

	copy(res, image)

	fill(res, image, visited, sr, sc, image[sr][sc], color, numRows, numCols)

	return res
}

func fill(res, image [][]int, visited [][]bool, i int, j int, srColor, color int, nr, nc int) {
	visited[i][j] = true
	if srColor == image[i][j] {
		res[i][j] = color
	} else {
		return
	}

	if i-1 >= 0 && !visited[i-1][j] {
		fill(res, image, visited, i-1, j, srColor, color, nr, nc)
	}
	if i+1 < nr && !visited[i+1][j] {
		fill(res, image, visited, i+1, j, srColor, color, nr, nc)
	}
	if j-1 >= 0 && !visited[i][j-1] {
		fill(res, image, visited, i, j-1, srColor, color, nr, nc)
	}
	if j+1 < nc && !visited[i][j+1] {
		fill(res, image, visited, i, j+1, srColor, color, nr, nc)
	}
}
