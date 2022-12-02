package lettercasepermutation

import "unicode"

func letterCasePermutation(s string) []string {
	var ans []string
	var current []byte

	helper([]byte(s), current, &ans)

	return ans
}

func helper(s, current []byte, ans *[]string) {
	if len(s) == 0 {
		*ans = append(*ans, string(current))
		return
	}

	if s[0] >= '0' && s[0] <= '9' {
		current = append(current, s[0])
		helper(s[1:], current, ans)
	} else {
		for _, b := range []byte{byte(unicode.ToLower(rune(s[0]))), byte(unicode.ToUpper(rune(s[0])))} {
			current = append(current, b)
			curCopy := make([]byte, len(current), cap(current))
			copy(curCopy, current)

			helper(s[1:], curCopy, ans)

			current = current[0 : len(current)-1]
		}
	}
}
