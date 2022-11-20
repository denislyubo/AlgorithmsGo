package searchinsertposition

func searchInsert(nums []int, target int) int {
	l := len(nums)
	beg, end := 0, l-1
	for beg <= end {
		mid := beg + (end-beg)/2

		if target < nums[mid] {
			end = mid - 1
		} else if target > nums[mid] {
			beg = mid + 1
		} else {
			return mid
		}
	}

	return beg
}
