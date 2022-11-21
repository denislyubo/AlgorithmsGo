package movezeroes

func moveZeroes(nums []int) {
	pos := 0
	for k := 0; k < len(nums); k++ {
		if nums[k] != 0 {
			nums[pos] = nums[k]
			pos++
		}
	}
	for pos < len(nums) {
		nums[pos] = 0
		pos++
	}
}
