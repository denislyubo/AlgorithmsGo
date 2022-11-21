package rotate_array

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func rotate(nums []int, k int) {
	reverse(nums)
	k = k % len(nums)
	if k <= len(nums) {
		reverse(nums[:k])
	}
	if k < len(nums) {
		reverse(nums[k:])
	}
}
