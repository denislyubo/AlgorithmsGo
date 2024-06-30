package FindCelebrity

func findCelebrity(matrix [][]bool) int {
	numPeople := len(matrix)

	var i, j int
	for i, j = 0, numPeople-1; i != j; {
		if matrix[i][j] {
			i++
		} else {
			j--
		}
	}

	for k := 0; k < numPeople; k++ {
		if k != i && (!matrix[k][i] || matrix[i][k]) {
			return -1
		}
	}

	return i
}
