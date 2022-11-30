package houserobber

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func rob(nums []int) int {
	rob1, rob2 := 0, 0
	var temp int
	//[rob1, rob2, n, n+1, ...]
	for _, n := range nums {
		temp = max(n+rob1, rob2)
		rob1 = rob2
		rob2 = temp
	}
	return rob2
}
