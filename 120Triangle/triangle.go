package triangle

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func minimumTotal(triangle [][]int) int {
	dp := make([]int, len(triangle)+1)

	for i := len(triangle) - 1; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[j] = triangle[i][j] + min(dp[j], dp[j+1])
		}
	}

	return dp[0]
}
