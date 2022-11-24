package longestsubstringwithoutrepeatingchars

func lengthOfLongestSubstring(s string) int {
	res := 0
	m := make(map[uint8]struct{})

	for l, r := 0, 0; r < len(s); {
		if _, exist := m[s[r]]; !exist {
			m[s[r]] = struct{}{}
			r++
		} else {
			for l < len(s) && r < len(s) {
				if _, exist := m[s[r]]; !exist {
					break
				}
				delete(m, s[l])
				l++
			}
		}

		if r-l > res {
			res = r - l
		}
	}

	return res
}
