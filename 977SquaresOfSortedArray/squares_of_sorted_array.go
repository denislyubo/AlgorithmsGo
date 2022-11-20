package squaresofsortedarray

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func sortedSquares(nums []int) []int {
	l := len(nums)
	ret := make([]int, l, l)
	for i, j, k := 0, l-1, l-1; i <= j; {
		if abs(nums[i]) >= abs(nums[j]) {
			ret[k] = nums[i] * nums[i]
			i++
		} else {
			ret[k] = nums[j] * nums[j]
			j--
		}
		k--
	}

	return ret
}
