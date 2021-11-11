package missingnumber

func missingNumber(nums []int) int {
	var count int = 0

	for i := 0; i < len(nums); i++ {
		count += i - nums[i]
	}

	return count + len(nums)
}
