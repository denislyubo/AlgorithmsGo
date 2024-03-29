package binarysearch

// Search is exported version of search for use otside package
func Search(nums []int, target int) int {
	return search(nums, target)
}

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		}
	}
	return -1
}
