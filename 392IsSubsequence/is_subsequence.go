package issubsequence

func isSubsequence(s string, t string) bool {
	sRune, tRune := []rune(s), []rune(t)
	lenS, lenT := len(sRune), len(tRune)
	if lenS == 0 {
		return true
	}

	for i, j := 0, 0; i < lenS && j < lenT; j++ {
		if sRune[i] == tRune[j] {
			i++
		}
		if i == lenS {
			return true
		}
	}

	return false
}
