package maxconsecutiveones

func longestOnes(nums []int, k int) int {
	i, j := 0, 0
	for j = 0; j < len(nums); j++ {
		k -= 1 - nums[j]
		if k < 0 {
			k += 1 - nums[i]
			i++
		}
	}
	return j - i
}
