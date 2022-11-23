package longestsubstringwithoutrepeatingchars

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]bool)
	maxCounter := 0
	counter := 0
	var ch rune
	var startIndex int = 0
	i := startIndex
	for i < len(s) {
		ch = rune(s[i])

		if _, ok := m[ch]; !ok {
			m[ch] = true
			counter++
		} else {
			counter = 0
			m = make(map[rune]bool)
			startIndex++
			i = startIndex
			continue
		}

		if counter > maxCounter {
			maxCounter = counter
		}

		i++
	}

	return maxCounter
}
