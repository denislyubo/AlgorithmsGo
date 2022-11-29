package combinations

func combine(n int, k int) [][]int {
	var ans [][]int
	var current []int
	helper(1, current, &ans, k, n)
	return ans
}

func helper(idx int, current []int, ans *[][]int, k int, n int) {
	if len(current) == k {
		*ans = append(*ans, current)
		return
	}

	for i := idx; i < n+1; i++ {
		current = append(current, i)
		curCopy := make([]int, len(current), cap(current))
		copy(curCopy, current)
		helper(i+1, curCopy, ans, k, n)
		current = current[0 : len(current)-1]
	}
}
