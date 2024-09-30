package _TwoSum

func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return make([]int, 2)
}

func twoSum1(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		numberToFind := target - nums[i]
		val, ok := m[numberToFind]
		if ok {
			return []int{i, val}
		}

		m[nums[i]] = i
	}

	return make([]int, 2)
}

func twoSum2(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		numberToFind := target - nums[i]
		l, r := i+1, len(nums)-1
		for l <= r {
			m := l + (r-l)/2
			if nums[m] < numberToFind {
				l = m + 1
			} else if nums[m] > numberToFind {
				r = m - 1
			} else {
				return []int{i, m}
			}
		}
	}

	return make([]int, 2)
}

func twoSum3(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		sum := nums[l] + nums[r]
		if sum == target {
			return []int{l, r}
		}
		if sum > target {
			r--
		} else {
			l++
		}
	}

	return make([]int, 2)
}
