package longestsubstringwithoutrepeatingchars

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]bool)
	maxCounter := 0
	counter := 0
	var ch rune
	var start_index int = 0
	i := start_index
	for i < len(s) {
		ch = rune(s[i])

		if _, ok := m[ch]; !ok {
			m[ch] = true
			counter++
		} else {
			counter = 0
			m = make(map[rune]bool)
			start_index++
			i = start_index
			continue
		}

		if counter > maxCounter {
			maxCounter = counter
		}

		i++
	}

	return maxCounter
}
