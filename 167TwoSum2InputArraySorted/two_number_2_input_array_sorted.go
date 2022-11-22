package twosum2inputarraysorted

// N*log(N)
func twoSum(numbers []int, target int) []int {
	if len(numbers) > 0 {
		for b, e := 0, len(numbers)-1; b != e; {
			if numbers[b]+numbers[e] == target {
				return []int{b + 1, e + 1}
			}
			if numbers[b]+numbers[e] < target {
				b++
			} else {
				e--
			}
		}
	}

	return []int{}
}
