package maxsubarray

func Max(items ...int) int {
	max := 0
	for i := range items {
		if items[i] > max {
			max = items[i]
		}
	}

	return max
}

func maxSubArray(a []int) []int {
	ans := a[0]
	ans_l := 0
	ans_r := 0
	sum := 0
	minus_pos := -1
	for r := 0; r < len(a); r++ {
		sum += a[r]

		if sum > ans {
			ans = sum
			ans_l = minus_pos + 1
			ans_r = r
		}

		if sum < 0 {
			sum = 0
			minus_pos = r
		}
	}

	return a[ans_l : ans_r+1]
}
