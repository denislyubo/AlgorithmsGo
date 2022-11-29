package permutations

func permute(nums []int) [][]int {
	var ans [][]int
	var current []int

	helper(nums, current, &ans)

	return ans
}

func helper(nums, current []int, ans *[][]int) {
	if 0 == len(nums) {
		*ans = append(*ans, current)
		return
	}
	for i := 0; i < len(nums); i++ {
		current = append(current, nums[i])
		curCopy := make([]int, len(current), cap(current))
		copy(curCopy, current)
		nums[0], nums[i] = nums[i], nums[0]
		helper(nums[1:len(nums)], curCopy, ans)
		nums[0], nums[i] = nums[i], nums[0]
		current = current[0 : len(current)-1]
	}
}
