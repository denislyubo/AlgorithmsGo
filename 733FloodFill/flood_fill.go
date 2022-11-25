package floodfill

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	numRows := len(image)
	numCols := len(image[0])

	fill(image, sr, sc, image[sr][sc], color, numRows, numCols)

	return image
}

func fill(image [][]int, i, j, srColor, color, nr, nc int) {

	if i < 0 || i >= nr || j < 0 || j >= nc || image[i][j] == color || image[i][j] != srColor {
		return
	}
	image[i][j] = color

	fill(image, i-1, j, srColor, color, nr, nc)
	fill(image, i+1, j, srColor, color, nr, nc)
	fill(image, i, j-1, srColor, color, nr, nc)
	fill(image, i, j+1, srColor, color, nr, nc)
}
