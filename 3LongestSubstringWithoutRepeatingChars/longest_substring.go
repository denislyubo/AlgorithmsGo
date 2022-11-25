package longestsubstringwithoutrepeatingchars

func lengthOfLongestSubstring(s string) int {
	res := 0
	m := make(map[uint8]struct{})
	var l int

	for r := 0; r < len(s); r++ {
		_, exist := m[s[r]]
		for ; exist; _, exist = m[s[r]] {
			delete(m, s[l])
			l++
		}
		m[s[r]] = struct{}{}
		if r-l+1 > res {
			res = r - l + 1
		}
	}

	return res
}
